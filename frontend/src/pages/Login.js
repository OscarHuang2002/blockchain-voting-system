import React, { useState, useEffect } from 'react';
import { Form, Input, Button, message } from 'antd';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

export default function Login({ walletAddress, setToken }) {
  useEffect(() => {
    console.log("Wallet address updated in Login.js:", walletAddress);
  }, [walletAddress]);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  // 使用钱包地址和密码登录
  const handleLogin = async (values) => {
    const { password } = values;
    if (!walletAddress) {
      message.error("无法获取钱包地址，请检查 MetaMask 是否已连接");
      return;
    }

    try {
      setLoading(true);
      const response = await axios.post('http://localhost:8080/login', {
        address: walletAddress,
        passwd: password,
      });
      message.success(`登录成功，欢迎 ${response.data.user}`);
      setToken(response.data.token); // 保存 Token
      navigate('/voter'); 
    } catch (error) {
      if (error.response?.data?.error === "用户未注册，请先注册") {
        message.error("该地址未注册，请先注册");
      } else {
        message.error("登录失败: " + (error.response?.data?.error || error.message));
      }
      navigate('/login');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ maxWidth: 400, margin: 'auto', padding: '20px' }}>
      <h2>用户登录</h2>
      <Form onFinish={handleLogin}>
      
        <Form.Item>
          <Input
            placeholder="钱包地址"
            value={walletAddress || "未连接"}
            disabled
          />
        </Form.Item>
        <Form.Item
          name="password"
          rules={[{ required: true, message: '请输入密码' }]}
        >
          <Input.Password placeholder="密码" />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" loading={loading}>
            登录
          </Button>
        </Form.Item>
        <Form.Item>
          <Button type="link" onClick={() => navigate('/register')}>
            没有账户？点击注册
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
}