import React, { useState, useEffect } from "react";
import {
  Table,
  Button,
  Modal,
  message,
  Tag,
  Select,
  Spin,
  Input,
  Space,
  Card,
  Row,
  Col,
  Typography,
  Empty,
  Divider,
} from "antd";
import {
  CheckCircleOutlined,
  CloseCircleOutlined,
  SearchOutlined,
  FilterOutlined,
  RightOutlined,
  CalendarOutlined,
  TeamOutlined,
} from "@ant-design/icons";
import axios from "axios";

const { Option } = Select;
const { Search } = Input;
const { Title, Paragraph, Text } = Typography;

export default function VoterPanel({ contract, account }) {
  const [candidates, setCandidates] = useState([]);
  const [allProjects, setAllProjects] = useState([]);
  const [filteredProjects, setFilteredProjects] = useState([]);
  const [selectedProject, setSelectedProject] = useState(null);
  const [loading, setLoading] = useState(false);
  const [hasVoted, setHasVoted] = useState(false);
  const [votingStatus, setVotingStatus] = useState(false);
  const [voteModalVisible, setVoteModalVisible] = useState(false);
  const [selectedCandidate, setSelectedCandidate] = useState(null);
  const [projectSearchText, setProjectSearchText] = useState("");
  const [projectStatusFilter, setProjectStatusFilter] = useState("all");
  const [initialized, setInitialized] = useState(false);
  const [showCandidates, setShowCandidates] = useState(false);

  // 进入页面时立即刷新所有数据
  useEffect(() => {
    const initializeData = async () => {
      setLoading(true);
      try {
        // 获取所有项目
        await fetchProjects();
        setInitialized(true);
      } catch (error) {
        console.error("初始化数据失败:", error);
        message.error("加载数据失败，请刷新页面重试");
      } finally {
        setLoading(false);
      }
    };

    initializeData();
  }, []);

  // 获取所有项目并初始化选择
  const fetchProjects = async () => {
    try {
      const response = await axios.get("http://localhost:8080/projects");
      const projectList = response.data.projects || [];
      setAllProjects(projectList);
      applyProjectFilters(projectList, projectSearchText, projectStatusFilter);
      return projectList;
    } catch (error) {
      console.error("获取项目失败:", error);
      message.error("获取项目失败: " + error.message);
      return [];
    }
  };

  // 应用项目过滤器
  const applyProjectFilters = (projects, searchText, statusFilter) => {
    let filtered = [...projects];

    // 搜索过滤
    if (searchText) {
      const lowercaseSearch = searchText.toLowerCase();
      filtered = filtered.filter(
        (project) =>
          project.id.toString().includes(searchText) ||
          project.description.toLowerCase().includes(lowercaseSearch)
      );
    }

    // 状态过滤
    if (statusFilter !== "all") {
      const isActive = statusFilter === "active";
      filtered = filtered.filter((project) => project.isActive === isActive);
    }

    setFilteredProjects(filtered);
  };

  // 加载指定项目的候选人 - 更新状态逻辑
  const loadCandidates = async (projectId) => {
    if (!projectId) return;

    try {
      setLoading(true);
      setSelectedProject(projectId);
      setShowCandidates(true);
      console.log("加载项目ID=", projectId, "的候选人");

      // 加载候选人
      const response = await axios.get("http://localhost:8080/candidates", {
        params: { projectId },
      });
      setCandidates(response.data.candidates || []);

      // 更新项目状态
      const projectInfo = allProjects.find((p) => p.id === projectId);
      setVotingStatus(projectInfo ? projectInfo.isActive : false);

      // 检查用户投票状态
      await checkVoteStatus(projectId);
    } catch (error) {
      console.error("加载候选人失败:", error);
      message.error("加载候选人失败，请稍后重试");
      setCandidates([]);
    } finally {
      setLoading(false);
    }
  };

  // 单独检查投票状态的函数
  const checkVoteStatus = async (projectId) => {
    if (!account) return;

    try {
      const response = await axios.get("http://localhost:8080/userVoteDetail", {
        params: {
          project_id: projectId,
          user_address: account,
        },
      });

      setHasVoted(response.data.hasVoted || false);
    } catch (error) {
      console.error("检查投票状态错误:", error);
      setHasVoted(false);
    }
  };

  // 当搜索文本或状态过滤器改变时，重新应用过滤
  useEffect(() => {
    applyProjectFilters(allProjects, projectSearchText, projectStatusFilter);
  }, [projectSearchText, projectStatusFilter, allProjects]);

  // 显示投票确认模态框
  const showVoteModal = (candidate) => {
    setSelectedCandidate(candidate);
    setVoteModalVisible(true);
  };

  // 投票处理函数
  const handleVote = async () => {
    try {
      setLoading(true);

      // 使用用户自己的钱包签名
      const tx = await contract.vote(selectedProject, selectedCandidate.id);

      // 等待交易确认
      await tx.wait();

      message.success("投票成功！");
      setVoteModalVisible(false);

      // 刷新数据
      await loadCandidates(selectedProject);
    } catch (error) {
      console.error("投票错误:", error);
      message.error("投票失败: " + (error.reason || error.message));
    } finally {
      setLoading(false);
    }
  };

  // 添加页面可见性变化检测
  useEffect(() => {
    const handleVisibilityChange = () => {
      if (document.visibilityState === "visible" && initialized) {
        fetchProjects();
        if (selectedProject && showCandidates) {
          loadCandidates(selectedProject);
        }
      }
    };

    document.addEventListener("visibilitychange", handleVisibilityChange);
    return () => {
      document.removeEventListener("visibilitychange", handleVisibilityChange);
    };
  }, [initialized, selectedProject, showCandidates]);

  // 显示候选人详情
  const showCandidateDetails = (candidate) => {
    Modal.info({
      title: "候选人详情",
      content: (
        <div>
          <p>姓名: {candidate.name}</p>
          <p>简介: {candidate.description}</p>
          <p>票数: {candidate.votes}</p>
          <img
            src={candidate.imageUrl}
            alt="候选人图片"
            style={{ width: "100%", marginTop: 16 }}
          />
        </div>
      ),
      onOk() {},
    });
  };

  // 返回项目选择界面
  const backToProjects = () => {
    setShowCandidates(false);
    setSelectedProject(null);
  };

  const columns = [
    {
      title: "候选人名称",
      dataIndex: "name",
      key: "name",
      render: (text) => <b>{text}</b>,
    },
    {
      title: "当前票数",
      dataIndex: "votes",
      key: "votes",
      render: (votes) => <Tag color="blue">{votes}</Tag>,
    },
    {
      title: "操作",
      key: "action",
      render: (_, record) => (
        <Space>
          <Button type="link" onClick={() => showCandidateDetails(record)}>
            查看详情
          </Button>
          <Button
            type="primary"
            icon={hasVoted ? <CheckCircleOutlined /> : <CloseCircleOutlined />}
            disabled={hasVoted || !votingStatus}
            onClick={() => showVoteModal(record)}
          >
            {hasVoted ? "已投票" : votingStatus ? "投票" : "投票未开放"}
          </Button>
        </Space>
      ),
    },
  ];

  // 渲染项目卡片
  const renderProjectCard = (project) => {
    return (
      <Col
        xs={24}
        sm={12}
        md={8}
        lg={6}
        key={project.id}
        style={{ marginBottom: "16px" }}
      >
        <Card
          hoverable
          style={{ height: "100%" }}
          actions={[
            <Button
              type="primary"
              onClick={() => loadCandidates(project.id)}
              icon={<TeamOutlined />}
            >
              查看候选人
            </Button>,
          ]}
        >
          <Card.Meta
            title={<span>项目 #{project.id}</span>}
            description={
              <>
                <Paragraph ellipsis={{ rows: 2 }}>
                  {project.description}
                </Paragraph>
                <div style={{ marginTop: "8px" }}>
                  <Tag color={project.isActive ? "green" : "default"}>
                    {project.isActive ? "进行中" : "未开始/已结束"}
                  </Tag>
                </div>
              </>
            }
          />
        </Card>
      </Col>
    );
  };

  // 渲染项目列表页面
  const renderProjectsPage = () => (
    <div>
      <div
        style={{
          marginBottom: 16,
          display: "flex",
          justifyContent: "space-between",
        }}
      >
        <Title level={4}>投票项目列表</Title>

        <div style={{ display: "flex", gap: 8 }}>
          <Search
            placeholder="搜索项目ID或描述"
            allowClear
            style={{ width: 250 }}
            value={projectSearchText}
            onChange={(e) => setProjectSearchText(e.target.value)}
          />

          <Select
            style={{ width: 150 }}
            value={projectStatusFilter}
            onChange={setProjectStatusFilter}
          >
            <Option value="all">所有状态</Option>
            <Option value="active">进行中</Option>
            <Option value="inactive">未开始/已结束</Option>
          </Select>
        </div>
      </div>

      <Spin spinning={loading}>
        {filteredProjects.length > 0 ? (
          <Row gutter={16}>{filteredProjects.map(renderProjectCard)}</Row>
        ) : (
          <Empty description="没有找到符合条件的项目" />
        )}
      </Spin>
    </div>
  );

  // 渲染候选人页面
  const renderCandidatesPage = () => {
    const project = allProjects.find((p) => p.id === selectedProject);

    return (
      <div>
        <div style={{ marginBottom: 16 }}>
          <Button type="link" onClick={backToProjects} style={{ padding: 0 }}>
            <RightOutlined style={{ transform: "rotate(180deg)" }} />{" "}
            返回项目列表
          </Button>
          <Title level={4} style={{ marginTop: 16 }}>
            项目 #{selectedProject}: {project?.description.substring(0, 30)}...
          </Title>
          <Tag
            color={project?.isActive ? "green" : "default"}
            style={{ marginBottom: 16 }}
          >
            {project?.isActive ? "进行中" : "未开始/已结束"}
          </Tag>

          {hasVoted && (
            <Tag color="blue" style={{ marginLeft: 8 }}>
              <CheckCircleOutlined /> 您已经在此项目中投票
            </Tag>
          )}
        </div>

        <Table
          columns={columns}
          dataSource={candidates}
          loading={loading}
          rowKey="id"
          pagination={{ pageSize: 10 }}
        />
      </div>
    );
  };

  return (
    <div>
      {showCandidates ? renderCandidatesPage() : renderProjectsPage()}

      <Modal
        title="确认投票"
        open={voteModalVisible}
        onOk={handleVote}
        onCancel={() => setVoteModalVisible(false)}
        confirmLoading={loading}
      >
        <p>
          您确定要投票给 <b>{selectedCandidate?.name}</b> 吗？
        </p>
        <p>此操作不可撤销，请谨慎操作。</p>
      </Modal>
    </div>
  );
}
