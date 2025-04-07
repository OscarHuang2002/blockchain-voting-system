import React, { useState, useEffect } from "react";
import { Form, Input, Button, Modal, Table, message, Select, Tabs } from "antd";
import axios from "axios";
import { handleError } from "../utils/errorHandler";

const { TabPane } = Tabs;
const { Option } = Select;

export default function AdminPanel({ contract }) {
  const [form] = Form.useForm();
  const [projectForm] = Form.useForm();
  const [visible, setVisible] = useState(false);
  const [projectVisible, setProjectVisible] = useState(false);
  const [candidates, setCandidates] = useState([]);
  const [projects, setProjects] = useState([]);
  const [selectedProject, setSelectedProject] = useState(1);
  const [isEditing, setIsEditing] = useState(false);
  const [editingId, setEditingId] = useState(null);

  // 获取所有项目
  const fetchProjects = async () => {
    try {
      const response = await axios.get("http://localhost:8080/projects");
      // 确保设置的是数组，即使API返回null或undefined
      setProjects(response.data.projects || []);

      // 如果有项目数据且selectedProject未设置，自动选择第一个项目
      if (
        response.data.projects &&
        response.data.projects.length > 0 &&
        !selectedProject
      ) {
        setSelectedProject(response.data.projects[0].id);
      }
    } catch (error) {
      message.error("获取项目失败: " + error.message);
      setProjects([]); // 确保错误时也设置为空数组
    }
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
      form.resetFields();
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
      address: record.address,
    });
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
      form.resetFields();
      setIsEditing(false);
      setEditingId(null);
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

      // 添加刷新项目和候选人列表的代码
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

      // 添加刷新项目和候选人列表的代码
      fetchProjects();
      fetchCandidates(selectedProject);
    } catch (error) {
      handleError(error);
    }
  };

  const candidateColumns = [
    { title: "ID", dataIndex: "id", key: "id" },
    { title: "姓名", dataIndex: "name", key: "name" },
    { title: "票数", dataIndex: "votes", key: "votes" },
    {
      title: "操作",
      key: "action",
      render: (_, record) => (
        <>
          <Button
            onClick={() => editCandidate(record)}
            style={{ marginRight: 8 }}
          >
            编辑
          </Button>
          <Button onClick={() => deleteCandidate(record.id)} danger>
            删除
          </Button>
        </>
      ),
    },
  ];

  const projectColumns = [
    { title: "ID", dataIndex: "id", key: "id" },
    { title: "描述", dataIndex: "description", key: "description" },
    {
      title: "状态",
      dataIndex: "isActive",
      key: "isActive",
      render: (isActive) => (isActive ? "进行中" : "未开始/已结束"),
    },
    {
      title: "操作",
      key: "action",
      render: (_, record) => (
        <Button onClick={() => setSelectedProject(record.id)} type="primary">
          管理候选人
        </Button>
      ),
    },
  ];

  return (
    <div>
      <Tabs defaultActiveKey="1">
        <TabPane tab="项目管理" key="1">
          <Button
            type="primary"
            onClick={() => setProjectVisible(true)}
            style={{ marginBottom: 16 }}
          >
            创建新项目
          </Button>

          <Table dataSource={projects} columns={projectColumns} rowKey="id" />

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
            <span style={{ marginRight: 8 }}>选择项目:</span>
            <Select
              value={selectedProject}
              onChange={setSelectedProject}
              style={{ width: 200 }}
            >
              {projects.map((project) => (
                <Option key={project.id} value={project.id}>
                  {project.description.substring(0, 20)}...
                </Option>
              ))}
            </Select>

            <Button
              type="primary"
              onClick={() => {
                setVisible(true);
                setIsEditing(false);
                form.resetFields();
              }}
              style={{ marginLeft: 16 }}
            >
              添加候选人
            </Button>

            <Button
              type="default"
              onClick={startVoting}
              style={{ marginLeft: 8 }}
            >
              开始投票
            </Button>

            <Button
              type="default"
              onClick={endVoting}
              style={{ marginLeft: 8 }}
            >
              结束投票
            </Button>
          </div>

          <Table
            dataSource={candidates}
            columns={candidateColumns}
            rowKey="id"
          />

          <Modal
            title={isEditing ? "编辑候选人" : "添加候选人"}
            open={visible}
            onCancel={() => setVisible(false)}
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
                label="图片URL"
                rules={[{ required: true, message: "请输入图片URL" }]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                name="description"
                label="候选人描述"
                rules={[{ required: true, message: "请输入候选人描述" }]}
              >
                <Input.TextArea rows={3} />
              </Form.Item>
              {!isEditing && (
                <Form.Item
                  name="address"
                  label="候选人地址"
                  rules={[{ required: true, message: "请输入候选人地址" }]}
                >
                  <Input />
                </Form.Item>
              )}
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
