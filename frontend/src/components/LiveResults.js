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
} from "antd";
import axios from "axios";

const { Option } = Select;

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
  const [selectedProject, setSelectedProject] = useState(null);
  const [totalVotes, setTotalVotes] = useState(0);
  const [initialized, setInitialized] = useState(false);

  // 初始化：获取项目列表
  useEffect(() => {
    const initializeData = async () => {
      try {
        setLoading(true);
        await fetchProjects();
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

      // 如果有项目，选择第一个并加载其投票结果
      if (projectList.length > 0) {
        setSelectedProject(projectList[0].id);
        await loadVotingResults(projectList[0].id);
      }

      return projectList;
    } catch (error) {
      console.error("获取项目失败:", error);
      message.error("获取项目失败: " + error.message);
      return [];
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

  // 定时刷新数据（每30秒）
  useEffect(() => {
    if (!initialized) return;

    const refreshInterval = setInterval(() => {
      if (selectedProject) {
        loadVotingResults(selectedProject);
      }
    }, 30000);

    return () => clearInterval(refreshInterval);
  }, [initialized, selectedProject]);

  // 添加页面可见性变化检测，切换回页面时刷新数据
  useEffect(() => {
    const handleVisibilityChange = () => {
      if (
        document.visibilityState === "visible" &&
        initialized &&
        selectedProject
      ) {
        loadVotingResults(selectedProject);
      }
    };

    document.addEventListener("visibilitychange", handleVisibilityChange);
    return () => {
      document.removeEventListener("visibilitychange", handleVisibilityChange);
    };
  }, [initialized, selectedProject]);

  const getProjectName = (projectId) => {
    const project = projects.find((p) => p.id === projectId);
    return project ? project.description : "未选择项目";
  };

  return (
    <div className="live-results-container">
      <Card title="实时投票统计" bordered={false}>
        <div style={{ marginBottom: 16 }}>
          <span style={{ marginRight: 8 }}>选择项目:</span>
          <Select
            value={selectedProject}
            style={{ width: 300 }}
            onChange={handleProjectChange}
            loading={loading && projects.length === 0}
            disabled={loading && projects.length === 0}
          >
            {projects.map((project) => (
              <Option key={project.id} value={project.id}>
                {project.description.substring(0, 30)}...
              </Option>
            ))}
          </Select>
        </div>

        {loading ? (
          <div style={{ textAlign: "center", marginTop: 50, marginBottom: 50 }}>
            <Spin size="large" tip="加载投票数据中..." />
          </div>
        ) : (
          <>
            {selectedProject && (
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
            )}

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
      </Card>
    </div>
  );
}
