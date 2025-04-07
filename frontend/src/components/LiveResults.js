import React, { useState, useEffect } from "react";
import {
  PieChart,
  Pie,
  Cell,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";
import {
  Spin,
  message,
  Select,
  Empty,
  Card,
  Statistic,
  Row,
  Col,
  Divider,
  Carousel,
  Typography,
  Input,
  Tag,
  Button,
  Space,
} from "antd";
import { ArrowLeftOutlined } from "@ant-design/icons";
import axios from "axios";

const { Option } = Select;
const { Title, Paragraph } = Typography;
const { Search } = Input;

// 饼图颜色
const COLORS = [
  "#0088FE",
  "#00C49F",
  "#FFBB28",
  "#FF8042",
  "#8884D8",
  "#FF6B6B",
  "#6FCF97",
  "#9B51E0",
];

export default function LiveResults({ account }) {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [projects, setProjects] = useState([]);
  const [filteredProjects, setFilteredProjects] = useState([]);
  const [selectedProject, setSelectedProject] = useState(null);
  const [totalVotes, setTotalVotes] = useState(0);
  const [initialized, setInitialized] = useState(false);
  const [projectSearchText, setProjectSearchText] = useState("");
  const [hotProjects, setHotProjects] = useState([]);
  const [hotProjectsData, setHotProjectsData] = useState([]);
  const [loadingHotProjects, setLoadingHotProjects] = useState(true);

  // 初始化：获取项目列表
  useEffect(() => {
    const initializeData = async () => {
      try {
        setLoading(true);
        setLoadingHotProjects(true);
        const projectsList = await fetchProjects();

        // 获取热门项目
        await fetchHotProjects(projectsList);

        setInitialized(true);
      } catch (error) {
        console.error("初始化数据失败:", error);
        message.error("加载项目数据失败，请刷新页面重试");
      } finally {
        setLoading(false);
      }
    };

    initializeData();
  }, []);

  // 获取所有项目
  const fetchProjects = async () => {
    try {
      const response = await axios.get("http://localhost:8080/projects");
      const projectList = response.data.projects || [];
      setProjects(projectList);
      setFilteredProjects(projectList);
      return projectList;
    } catch (error) {
      console.error("获取项目失败:", error);
      message.error("获取项目失败: " + error.message);
      return [];
    }
  };

  // 搜索项目
  const searchProjects = (text) => {
    setProjectSearchText(text);

    if (!text) {
      setFilteredProjects(projects);
      return;
    }

    const lowercaseSearch = text.toLowerCase();
    const filtered = projects.filter(
      (project) =>
        project.id.toString().includes(text) ||
        project.description.toLowerCase().includes(lowercaseSearch)
    );

    setFilteredProjects(filtered);
  };

  // 获取热门项目（按投票数排序）
  const fetchHotProjects = async (projectList) => {
    setLoadingHotProjects(true);
    try {
      const projectVotesPromises = projectList.map(async (project) => {
        const response = await axios.get("http://localhost:8080/candidates", {
          params: { projectId: project.id },
        });
        const candidates = response.data.candidates || [];
        const totalVotes = candidates.reduce((sum, c) => sum + c.votes, 0);
        return {
          ...project,
          totalVotes,
          candidates,
        };
      });

      const projectsWithVotes = await Promise.all(projectVotesPromises);

      // 根据总票数排序，得到热门项目
      const sorted = [...projectsWithVotes].sort(
        (a, b) => b.totalVotes - a.totalVotes
      );
      const top3Projects = sorted.slice(0, 3);

      setHotProjects(top3Projects);

      // 为每个热门项目准备投票数据
      const hotProjectsChartData = top3Projects.map((project) => {
        // 准备饼图数据
        const chartData = project.candidates.map((candidate) => ({
          name: candidate.name,
          value: candidate.votes,
        }));

        return {
          project,
          chartData,
        };
      });

      setHotProjectsData(hotProjectsChartData);
    } catch (error) {
      console.error("获取热门项目失败:", error);
      message.error("获取热门项目失败，请稍后重试");
    } finally {
      setLoadingHotProjects(false);
    }
  };

  // 加载投票结果
  const loadVotingResults = async (projectId) => {
    if (!projectId) return;

    try {
      setLoading(true);
      console.log("加载项目ID=", projectId, "的投票结果");

      // 从后端获取候选人列表和投票数据
      const response = await axios.get("http://localhost:8080/candidates", {
        params: { projectId },
      });

      const candidates = response.data.candidates || [];

      // 转换数据格式用于饼图显示
      const chartData = candidates.map((candidate) => ({
        name: candidate.name,
        value: candidate.votes,
      }));

      setData(chartData);

      // 计算总票数
      const total = candidates.reduce(
        (sum, candidate) => sum + candidate.votes,
        0
      );
      setTotalVotes(total);
    } catch (error) {
      console.error("加载投票结果失败:", error);
      message.error("获取投票结果失败，请稍后重试");
      setData([]);
      setTotalVotes(0);
    } finally {
      setLoading(false);
    }
  };

  // 处理项目选择变更
  const handleProjectChange = (value) => {
    setSelectedProject(value);
    loadVotingResults(value);
  };

  // 返回热门项目列表
  const backToHotProjects = () => {
    setSelectedProject(null);
    setData([]);
  };

  // 定时刷新数据（每30秒）
  useEffect(() => {
    if (!initialized) return;

    const refreshInterval = setInterval(() => {
      if (selectedProject) {
        loadVotingResults(selectedProject);
      } else if (projects.length > 0) {
        // 只有在未选择特定项目时刷新热门项目
        fetchHotProjects(projects);
      }
    }, 30000);

    return () => clearInterval(refreshInterval);
  }, [initialized, selectedProject, projects]);

  // 添加页面可见性变化检测，切换回页面时刷新数据
  useEffect(() => {
    const handleVisibilityChange = () => {
      if (document.visibilityState === "visible" && initialized) {
        if (selectedProject) {
          loadVotingResults(selectedProject);
        } else if (projects.length > 0) {
          fetchHotProjects(projects);
        }
      }
    };

    document.addEventListener("visibilitychange", handleVisibilityChange);
    return () => {
      document.removeEventListener("visibilitychange", handleVisibilityChange);
    };
  }, [initialized, selectedProject, projects]);

  const getProjectName = (projectId) => {
    const project = projects.find((p) => p.id === projectId);
    return project ? project.description : "未选择项目";
  };

  // 渲染热门项目轮播图
  const renderHotProjectsCarousel = () => {
    if (loadingHotProjects) {
      return (
        <div style={{ textAlign: "center", padding: "50px 0" }}>
          <Spin size="large" tip="加载热门项目中..." />
        </div>
      );
    }

    if (hotProjectsData.length === 0) {
      return (
        <Empty
          description="暂无投票数据"
          image={Empty.PRESENTED_IMAGE_SIMPLE}
          style={{ marginTop: 50 }}
        />
      );
    }

    return (
      <Carousel
        autoplay
        autoplaySpeed={5000}
        effect="fade"
        style={{ marginTop: 24, marginBottom: 50 }}
      >
        {hotProjectsData.map((item, index) => (
          <div key={item.project.id}>
            <Card
              title={
                <div
                  style={{
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                  }}
                >
                  <span>
                    热门项目 #{index + 1}:{" "}
                    {item.project.description.substring(0, 30)}
                    {item.project.description.length > 30 ? "..." : ""}
                  </span>
                  <Tag color={item.project.isActive ? "green" : "default"}>
                    {item.project.isActive ? "进行中" : "未开始/已结束"}
                  </Tag>
                </div>
              }
              extra={
                <Button
                  type="primary"
                  onClick={() => handleProjectChange(item.project.id)}
                >
                  查看详情
                </Button>
              }
            >
              <Row gutter={16}>
                <Col xs={24} md={8}>
                  <Card>
                    <Statistic title="项目ID" value={item.project.id} />
                  </Card>
                </Col>
                <Col xs={24} md={8}>
                  <Card>
                    <Statistic
                      title="总投票数"
                      value={item.project.totalVotes}
                      valueStyle={{ color: "#3f8600" }}
                    />
                  </Card>
                </Col>
                <Col xs={24} md={8}>
                  <Card>
                    <Statistic
                      title="候选人数量"
                      value={item.chartData.length}
                    />
                  </Card>
                </Col>
              </Row>

              <div style={{ height: 300, marginTop: 24 }}>
                <ResponsiveContainer width="100%" height="100%">
                  <PieChart>
                    <Pie
                      data={item.chartData}
                      dataKey="value"
                      nameKey="name"
                      cx="50%"
                      cy="50%"
                      outerRadius={100}
                      label={({ name, value, percent }) =>
                        `${name}: ${value} (${(percent * 100).toFixed(1)}%)`
                      }
                    >
                      {item.chartData.map((entry, i) => (
                        <Cell
                          key={`cell-${i}`}
                          fill={COLORS[i % COLORS.length]}
                        />
                      ))}
                    </Pie>
                    <Tooltip formatter={(value) => `${value} 票`} />
                    <Legend />
                  </PieChart>
                </ResponsiveContainer>
              </div>
            </Card>
          </div>
        ))}
      </Carousel>
    );
  };

  // 渲染项目详情
  const renderProjectDetail = () => {
    return (
      <div style={{ marginTop: 24 }}>
        <div style={{ marginBottom: 16 }}>
          <Button
            type="link"
            icon={<ArrowLeftOutlined />}
            onClick={backToHotProjects}
            style={{ paddingLeft: 0 }}
          >
            返回热门项目
          </Button>
        </div>

        {loading ? (
          <div style={{ textAlign: "center", padding: "50px 0" }}>
            <Spin size="large" tip="加载投票数据中..." />
          </div>
        ) : (
          <>
            <div style={{ marginBottom: 24 }}>
              <Row gutter={16}>
                <Col span={24}>
                  <Card>
                    <Statistic
                      title="当前项目"
                      value={getProjectName(selectedProject)}
                      valueStyle={{ fontSize: 16 }}
                    />
                  </Card>
                </Col>
              </Row>
              <Divider />
              <Row gutter={16}>
                <Col span={8}>
                  <Card>
                    <Statistic
                      title="总投票数"
                      value={totalVotes}
                      valueStyle={{ color: "#3f8600" }}
                    />
                  </Card>
                </Col>
                <Col span={8}>
                  <Card>
                    <Statistic title="候选人数量" value={data.length} />
                  </Card>
                </Col>
                <Col span={8}>
                  <Card>
                    <Statistic
                      title="最新刷新时间"
                      value={new Date().toLocaleTimeString()}
                      suffix="自动每30秒刷新"
                    />
                  </Card>
                </Col>
              </Row>
            </div>

            {data.length > 0 ? (
              <div style={{ width: "100%", height: 400 }}>
                <ResponsiveContainer>
                  <PieChart>
                    <Pie
                      data={data}
                      dataKey="value"
                      nameKey="name"
                      cx="50%"
                      cy="50%"
                      outerRadius={150}
                      label={({ name, value, percent }) =>
                        `${name}: ${value} (${(percent * 100).toFixed(2)}%)`
                      }
                      labelLine={true}
                    >
                      {data.map((entry, index) => (
                        <Cell
                          key={`cell-${index}`}
                          fill={COLORS[index % COLORS.length]}
                        />
                      ))}
                    </Pie>
                    <Tooltip formatter={(value) => `${value} 票`} />
                    <Legend />
                  </PieChart>
                </ResponsiveContainer>
              </div>
            ) : (
              <Empty
                description="暂无投票数据"
                image={Empty.PRESENTED_IMAGE_SIMPLE}
                style={{ marginTop: 50 }}
              />
            )}
          </>
        )}
      </div>
    );
  };

  return (
    <div className="live-results-container">
      <Card title="实时投票统计" bordered={false}>
        <Title level={4}>项目选择</Title>
        <div
          style={{
            marginBottom: 16,
            display: "flex",
            justifyContent: "space-between",
          }}
        >
          <Space>
            <span>选择项目:</span>
            <Select
              value={selectedProject}
              style={{ width: 300 }}
              onChange={handleProjectChange}
              loading={loading && projects.length === 0}
              disabled={loading && projects.length === 0}
              allowClear
              placeholder="请选择项目"
              dropdownRender={(menu) => (
                <>
                  <div style={{ padding: "8px" }}>
                    <Search
                      placeholder="搜索项目ID或描述"
                      allowClear
                      value={projectSearchText}
                      onChange={(e) => searchProjects(e.target.value)}
                      style={{ width: "100%" }}
                    />
                  </div>
                  <Divider style={{ margin: "4px 0" }} />
                  {menu}
                </>
              )}
            >
              {filteredProjects.map((project) => (
                <Option key={project.id} value={project.id}>
                  项目#{project.id}: {project.description.substring(0, 30)}...
                </Option>
              ))}
            </Select>
          </Space>
        </div>

        <Divider />

        {selectedProject ? (
          // 如果选择了项目，显示项目详情
          renderProjectDetail()
        ) : (
          // 否则显示热门项目轮播图
          <>
            <Title level={4}>热门项目</Title>
            {renderHotProjectsCarousel()}
          </>
        )}
      </Card>
    </div>
  );
}
