import React, { useState } from 'react';
import { Button, Input, Form, message } from 'antd';
import axios from 'axios';

export default function Register() {
  const [loading, setLoading] = useState(false);

  const handleRegister = async (values) => {
    const { name, email, password, address } = values; // Include address field
    // console.log(values);
    
    try {
      setLoading(true);
      // Call the backend registration API
      const response = await axios.post('http://localhost:8080/register', {
        username: name,
        email,
        password,
        address, // Send the address field to the backend
      });
      message.success("注册成功，请登录");
    } catch (error) {
      message.error("注册失败: " + (error.response?.data?.error || error));
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
            { type: 'email', message: '邮箱格式不正确' }
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
        <Form.Item
          name="address"
          rules={[{ required: true, message: '请输入区块链地址' }]} // Add validation for address
        >
          <Input placeholder="区块链地址" />
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