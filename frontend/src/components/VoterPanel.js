import React, { useState, useEffect } from "react";
import { Table, Button, Modal, message, Tag, Select, Spin } from "antd";
import { CheckCircleOutlined, CloseCircleOutlined } from "@ant-design/icons";
import axios from "axios";

const { Option } = Select;

export default function VoterPanel({ contract, account }) {
  const [candidates, setCandidates] = useState([]);
  const [projects, setProjects] = useState([]);
  const [selectedProject, setSelectedProject] = useState(1);
  const [loading, setLoading] = useState(false);
  const [hasVoted, setHasVoted] = useState(false);
  const [votingStatus, setVotingStatus] = useState(false);
  const [voteModalVisible, setVoteModalVisible] = useState(false);
  const [selectedCandidate, setSelectedCandidate] = useState(null);

  // 添加一个状态来跟踪页面是否初始化
  const [initialized, setInitialized] = useState(false);

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
      setProjects(projectList);

      // 如果有项目并且没有选择项目，则选择第一个
      if (
        projectList.length > 0 &&
        (!selectedProject || !projectList.find((p) => p.id === selectedProject))
      ) {
        setSelectedProject(projectList[0].id);
        // 立即加载这个项目的候选人
        await loadCandidates(projectList[0].id);
      } else if (selectedProject) {
        // 如果已经选择了项目，刷新这个项目的数据
        await loadCandidates(selectedProject);
      }

      return projectList;
    } catch (error) {
      console.error("获取项目失败:", error);
      message.error("获取项目失败: " + error.message);
      return [];
    }
  };

  // 加载指定项目的候选人 - 更新状态逻辑
  const loadCandidates = async (projectId) => {
    if (!projectId) return;

    try {
      setLoading(true);
      console.log("加载项目ID=", projectId, "的候选人");

      // 加载候选人
      const response = await axios.get("http://localhost:8080/candidates", {
        params: { projectId },
      });
      setCandidates(response.data.candidates || []);

      // 更新项目状态
      const projectInfo = projects.find((p) => p.id === projectId);
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

  // 显示投票确认模态框
  const showVoteModal = (candidate) => {
    setSelectedCandidate(candidate);
    setVoteModalVisible(true);
  };

  // 项目选择变更处理
  const handleProjectChange = (value) => {
    setSelectedProject(value);
    loadCandidates(value);
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
      await fetchProjects();
      await loadCandidates(selectedProject);
    } catch (error) {
      console.error("投票错误:", error);
      message.error("投票失败: " + (error.reason || error.message));
    } finally {
      setLoading(false);
    }
  };

  // // 添加自动刷新功能
  // useEffect(() => {
  //   // 只在初始化后开始轮询
  //   if (!initialized) return;

  //   const refreshInterval = setInterval(() => {
  //     fetchProjects();
  //     if (selectedProject) {
  //       loadCandidates(selectedProject);
  //     }
  //   }, 15000); // 每15秒刷新一次

  //   return () => clearInterval(refreshInterval);
  // }, [initialized, selectedProject, account]);

  // 添加页面可见性变化检测
  useEffect(() => {
    const handleVisibilityChange = () => {
      if (document.visibilityState === "visible" && initialized) {
        fetchProjects();
        if (selectedProject) {
          loadCandidates(selectedProject);
        }
      }
    };

    document.addEventListener("visibilitychange", handleVisibilityChange);
    return () => {
      document.removeEventListener("visibilitychange", handleVisibilityChange);
    };
  }, [initialized, selectedProject]);

  // 显示候选人详情
  const showCandidateDetails = (candidate) => {
    Modal.info({
      title: "候选人详情",
      content: (
        <div>
          <p>姓名: {candidate.name}</p>
          <p>简介: {candidate.description}</p>
          <p>票数: {candidate.votes}</p>
          {/* <p>状态: {candidate.isActive ? "活跃" : "非活跃"}</p> */}
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
        <>
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
        </>
      ),
    },
  ];

  return (
    <div>
      <div style={{ marginBottom: 16 }}>
        <span style={{ marginRight: 8 }}>选择项目:</span>
        <Select
          value={selectedProject}
          style={{ width: 300 }}
          onChange={handleProjectChange}
        >
          {projects.map((project) => (
            <Option key={project.id} value={project.id}>
              {project.description.substring(0, 30)}...
            </Option>
          ))}
        </Select>
      </div>

      <Table
        columns={columns}
        dataSource={candidates}
        loading={loading}
        rowKey="id"
      />

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
