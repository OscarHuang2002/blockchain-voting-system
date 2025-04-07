import React, { useState, useEffect } from "react";
import {
  Form,
  Input,
  Button,
  Modal,
  Table,
  message,
  Select,
  Tabs,
  Space,
  Upload,
  Input as AntInput,
  Tag,
} from "antd";
import {
  SearchOutlined,
  ExclamationCircleOutlined,
  DeleteOutlined,
  EditOutlined,
  TeamOutlined,
  PlusOutlined,
  LoadingOutlined,
} from "@ant-design/icons";
import axios from "axios";
import { handleError } from "../utils/errorHandler";

const { TabPane } = Tabs;
const { Option } = Select;
const { Search } = AntInput;
const { confirm } = Modal;

export default function AdminPanel({ contract }) {
  const [form] = Form.useForm();
  const [projectForm] = Form.useForm();
  const [visible, setVisible] = useState(false);
  const [projectVisible, setProjectVisible] = useState(false);
  const [candidates, setCandidates] = useState([]);
  const [allProjects, setAllProjects] = useState([]);
  const [filteredProjects, setFilteredProjects] = useState([]);
  const [selectedProject, setSelectedProject] = useState(1);
  const [isEditing, setIsEditing] = useState(false);
  const [editingId, setEditingId] = useState(null);
  const [projectSearchText, setProjectSearchText] = useState("");
  const [projectStatusFilter, setProjectStatusFilter] = useState("all");
  const [activeTab, setActiveTab] = useState("1");
  const [imageUrl, setImageUrl] = useState(""); // 存储上传图片的URL
  const [uploadLoading, setUploadLoading] = useState(false); // 上传状态

  // 获取所有项目
  const fetchProjects = async () => {
    try {
      const response = await axios.get("http://localhost:8080/projects");
      const projects = response.data.projects || [];
      setAllProjects(projects);
      applyProjectFilters(projects, projectSearchText, projectStatusFilter);

      // 如果有项目数据且selectedProject未设置，自动选择第一个项目
      if (projects.length > 0 && !selectedProject) {
        setSelectedProject(projects[0].id);
      }
    } catch (error) {
      message.error("获取项目失败: " + error.message);
      setAllProjects([]);
      setFilteredProjects([]);
    }
  };

  // 图片上传前检查
  const beforeUpload = (file) => {
    const isJpgOrPng = file.type === "image/jpeg" || file.type === "image/png";
    if (!isJpgOrPng) {
      message.error("只能上传JPG/PNG格式的图片!");
    }
    const isLt2M = file.size / 1024 / 1024 < 2;
    if (!isLt2M) {
      message.error("图片必须小于2MB!");
    }
    return isJpgOrPng && isLt2M;
  };

  // 图片上传状态变化处理
  const handleImageChange = (info) => {
    if (info.file.status === "uploading") {
      setUploadLoading(true);
      return;
    }

    if (info.file.status === "done") {
      setUploadLoading(false);
      // 获取上传后的图片URL
      const url = info.file.response.url;
      setImageUrl(url);
      // 自动更新表单中的imageUrl字段
      form.setFieldsValue({ imageUrl: url });
    } else if (info.file.status === "error") {
      setUploadLoading(false);
      message.error("图片上传失败");
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

  // 获取候选人列表
  const fetchCandidates = async (projectId) => {
    try {
      const response = await axios.get("http://localhost:8080/candidates", {
        params: { projectId },
      });
      setCandidates(response.data.candidates || []);
    } catch (error) {
      message.error("获取候选人失败: " + error.message);
    }
  };

  useEffect(() => {
    fetchProjects();
  }, [contract]);

  useEffect(() => {
    if (selectedProject) {
      fetchCandidates(selectedProject);
    }
  }, [selectedProject]);

  // 当搜索文本或状态过滤器改变时，重新应用过滤
  useEffect(() => {
    applyProjectFilters(allProjects, projectSearchText, projectStatusFilter);
  }, [projectSearchText, projectStatusFilter, allProjects]);

  // 创建项目
  const createProject = async (values) => {
    try {
      await axios.post("http://localhost:8080/AdminPanel/createProject", {
        description: values.description,
      });
      message.success("项目创建成功");
      setProjectVisible(false);
      projectForm.resetFields();
      fetchProjects();
    } catch (error) {
      message.error("创建项目失败: " + error.message);
    }
  };

  // 删除项目
  const deleteProject = (projectId) => {
    confirm({
      title: "确认删除项目",
      icon: <ExclamationCircleOutlined />,
      content: "删除后项目将不再显示，确定要删除此项目吗？",
      onOk: async () => {
        try {
          await axios.delete(
            `http://localhost:8080/AdminPanel/deleteProject/${projectId}`
          );
          message.success("项目删除成功");
          fetchProjects();
        } catch (error) {
          message.error("删除项目失败: " + error.message);
        }
      },
    });
  };

  // 重置表单和上传状态
  const resetForm = () => {
    form.resetFields();
    setImageUrl("");
    setIsEditing(false);
    setEditingId(null);
  };

  // 添加候选人
  const addCandidate = async (values) => {
    try {
      await axios.post("http://localhost:8080/AdminPanel/addCandidate", {
        project_id: selectedProject,
        name: values.name,
        image_url: values.imageUrl,
        description: values.description,
      });
      message.success("候选人添加成功");
      setVisible(false);
      resetForm();
      fetchCandidates(selectedProject);
    } catch (error) {
      message.error("添加候选人失败: " + error.message);
    }
  };

  // 编辑候选人
  const editCandidate = (record) => {
    form.setFieldsValue({
      name: record.name,
      imageUrl: record.imageUrl,
      description: record.description,
    });
    setImageUrl(record.imageUrl); // 设置图片URL
    setIsEditing(true);
    setEditingId(record.id);
    setVisible(true);
  };

  // 更新候选人
  const updateCandidate = async (values) => {
    try {
      await axios.put("http://localhost:8080/AdminPanel/updateCandidate", {
        project_id: selectedProject,
        candidate_id: editingId,
        name: values.name,
        image_url: values.imageUrl,
        description: values.description,
      });
      message.success("候选人修改成功");
      setVisible(false);
      resetForm();
      fetchCandidates(selectedProject);
    } catch (error) {
      message.error("修改候选人失败: " + error.message);
    }
  };

  // 删除候选人
  const deleteCandidate = async (id) => {
    try {
      await axios.delete(
        `http://localhost:8080/AdminPanel/deleteCandidate/${id}?project_id=${selectedProject}`
      );
      message.success("候选人删除成功");
      fetchCandidates(selectedProject);
    } catch (error) {
      message.error("删除候选人失败: " + error.message);
    }
  };

  // 提交表单
  const handleFinish = async (values) => {
    if (isEditing) {
      await updateCandidate(values);
    } else {
      await addCandidate(values);
    }
  };

  // 开始投票
  const startVoting = async () => {
    try {
      await axios.post("http://localhost:8080/AdminPanel/startVoting", {
        project_id: selectedProject,
      });
      message.success("投票已开始");
      fetchProjects();
      fetchCandidates(selectedProject);
    } catch (error) {
      handleError(error);
    }
  };

  // 结束投票
  const endVoting = async () => {
    try {
      await axios.post("http://localhost:8080/AdminPanel/endVoting", {
        project_id: selectedProject,
      });
      message.success("投票已结束");
      fetchProjects();
      fetchCandidates(selectedProject);
    } catch (error) {
      handleError(error);
    }
  };

  // 转到候选人管理
  const goToCandidateManagement = (projectId) => {
    setSelectedProject(projectId);
    setActiveTab("2"); // 切换到候选人管理选项卡
  };

  // 上传按钮
  const uploadButton = (
    <div>
      {uploadLoading ? <LoadingOutlined /> : <PlusOutlined />}
      <div style={{ marginTop: 8 }}>上传图片</div>
    </div>
  );

  const candidateColumns = [
    { title: "ID", dataIndex: "id", key: "id" },
    { title: "姓名", dataIndex: "name", key: "name" },
    {
      title: "图片",
      dataIndex: "imageUrl",
      key: "imageUrl",
      render: (imageUrl) => (
        <img
          src={imageUrl}
          alt="候选人图片"
          style={{ width: 50, height: 50, objectFit: "cover" }}
        />
      ),
    },
    { title: "票数", dataIndex: "votes", key: "votes" },
    {
      title: "操作",
      key: "action",
      render: (_, record) => (
        <Space>
          <Button icon={<EditOutlined />} onClick={() => editCandidate(record)}>
            编辑
          </Button>
          <Button
            danger
            icon={<DeleteOutlined />}
            onClick={() => deleteCandidate(record.id)}
          >
            删除
          </Button>
        </Space>
      ),
    },
  ];

  const projectColumns = [
    { title: "ID", dataIndex: "id", key: "id" },
    {
      title: "描述",
      dataIndex: "description",
      key: "description",
      ellipsis: true,
    },
    {
      title: "状态",
      dataIndex: "isActive",
      key: "isActive",
      render: (isActive) => (
        <Tag color={isActive ? "green" : "default"}>
          {isActive ? "进行中" : "未开始/已结束"}
        </Tag>
      ),
    },
    {
      title: "操作",
      key: "action",
      render: (_, record) => (
        <Space>
          <Button
            type="primary"
            icon={<TeamOutlined />}
            onClick={() => goToCandidateManagement(record.id)}
          >
            管理候选人
          </Button>
          <Button
            danger
            icon={<DeleteOutlined />}
            onClick={() => deleteProject(record.id)}
          >
            删除
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <div>
      <Tabs activeKey={activeTab} onChange={setActiveTab}>
        <TabPane tab="项目管理" key="1">
          <div
            style={{
              marginBottom: 16,
              display: "flex",
              justifyContent: "space-between",
            }}
          >
            <Button type="primary" onClick={() => setProjectVisible(true)}>
              创建新项目
            </Button>

            <div style={{ display: "flex", gap: 8 }}>
              <Search
                placeholder="搜索项目ID或描述"
                allowClear
                style={{ width: 250 }}
                value={projectSearchText}
                onChange={(e) => setProjectSearchText(e.target.value)}
              />

              <Select
                defaultValue="all"
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

          <Table
            dataSource={filteredProjects}
            columns={projectColumns}
            rowKey="id"
            pagination={{ pageSize: 10 }}
          />

          <Modal
            title="创建新项目"
            open={projectVisible}
            onCancel={() => setProjectVisible(false)}
            footer={null}
          >
            <Form form={projectForm} onFinish={createProject}>
              <Form.Item
                name="description"
                label="项目描述"
                rules={[{ required: true, message: "请输入项目描述" }]}
              >
                <Input.TextArea rows={4} />
              </Form.Item>
              <Form.Item>
                <Button type="primary" htmlType="submit">
                  创建
                </Button>
              </Form.Item>
            </Form>
          </Modal>
        </TabPane>

        <TabPane tab="候选人管理" key="2">
          <div style={{ marginBottom: 16 }}>
            <div
              style={{
                display: "flex",
                justifyContent: "space-between",
                marginBottom: 16,
              }}
            >
              <div>
                <span style={{ marginRight: 8 }}>选择项目:</span>
                <Select
                  value={selectedProject}
                  onChange={setSelectedProject}
                  style={{ width: 300 }}
                >
                  {allProjects.map((project) => (
                    <Option key={project.id} value={project.id}>
                      {project.description.substring(0, 30)}...
                    </Option>
                  ))}
                </Select>
              </div>

              <div>
                <Button
                  type="primary"
                  onClick={() => {
                    setVisible(true);
                    resetForm();
                  }}
                  style={{ marginRight: 8 }}
                >
                  添加候选人
                </Button>

                <Button
                  type="default"
                  onClick={startVoting}
                  style={{ marginRight: 8 }}
                >
                  开始投票
                </Button>

                <Button type="default" onClick={endVoting}>
                  结束投票
                </Button>
              </div>
            </div>
          </div>

          <Table
            dataSource={candidates}
            columns={candidateColumns}
            rowKey="id"
            pagination={{ pageSize: 10 }}
          />

          <Modal
            title={isEditing ? "编辑候选人" : "添加候选人"}
            open={visible}
            onCancel={() => {
              resetForm();
              setVisible(false);
            }}
            footer={null}
          >
            <Form form={form} onFinish={handleFinish}>
              <Form.Item
                name="name"
                label="候选人姓名"
                rules={[{ required: true, message: "请输入候选人姓名" }]}
              >
                <Input />
              </Form.Item>

              <Form.Item
                name="imageUrl"
                label="候选人图片"
                rules={[{ required: true, message: "请上传候选人图片" }]}
                style={{ display: "none" }} // 隐藏实际存储URL的表单项
              >
                <Input />
              </Form.Item>

              <Form.Item
                label="上传图片"
                extra="支持 JPG/PNG 格式图片，最大 2MB"
              >
                <Upload
                  name="image"
                  listType="picture-card"
                  className="avatar-uploader"
                  showUploadList={false}
                  action="http://localhost:8080/upload/image"
                  beforeUpload={beforeUpload}
                  onChange={handleImageChange}
                >
                  {imageUrl ? (
                    <img
                      src={imageUrl}
                      alt="avatar"
                      style={{
                        width: "100%",
                        height: "100%",
                        objectFit: "cover",
                      }}
                    />
                  ) : (
                    uploadButton
                  )}
                </Upload>
              </Form.Item>

              <Form.Item
                name="description"
                label="候选人描述"
                rules={[{ required: true, message: "请输入候选人描述" }]}
              >
                <Input.TextArea rows={3} />
              </Form.Item>
              <Form.Item>
                <Button type="primary" htmlType="submit">
                  提交
                </Button>
              </Form.Item>
            </Form>
          </Modal>
        </TabPane>
      </Tabs>
    </div>
  );
}
