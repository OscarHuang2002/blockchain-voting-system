import React, { useState } from 'react';
import { Input, Form, Button, message } from 'antd';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

export default function Register({ walletAddress }) {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate(); 

  const handleRegister = async (values) => {
    const { name, email, password } = values; // 从表单中获取其他字段
    if (!walletAddress) {
      message.error("请先连接 MetaMask");
      return;
    }

    try {
      setLoading(true);
      // 调用后端注册 API
      const response = await axios.post('http://localhost:8080/register', {
        username: name,
        email,
        password,
        address: walletAddress, // 使用从 WalletConnector 获取的地址
      });
      message.success("注册成功，请登录");
      navigate('/login'); // 注册成功后跳转到登录页面
    } catch (error) {
      message.error("注册失败: " + (error.response?.data?.error || error.message));
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ maxWidth: 400, margin: 'auto', padding: '20px' }}>
      <h2>用户注册</h2>
      <Form onFinish={handleRegister}>
        <Form.Item
          name="name"
          rules={[{ required: true, message: '请输入用户名' }]}
        >
          <Input placeholder="用户名" />
        </Form.Item>
        <Form.Item
          name="email"
          rules={[
            { required: true, message: '请输入邮箱' },
            { type: 'email', message: '邮箱格式不正确' },
          ]}
        >
          <Input placeholder="邮箱" />
        </Form.Item>
        <Form.Item
          name="password"
          rules={[{ required: true, message: '请输入密码' }]}
        >
          <Input.Password placeholder="密码" />
        </Form.Item>
        <Form.Item>
          <Input
            placeholder="区块链地址"
            value={walletAddress}
            disabled
          />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" loading={loading}>
            注册
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
}