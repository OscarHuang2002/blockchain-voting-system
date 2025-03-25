import React, { useState } from 'react';
import { Button, Input, Form, message } from 'antd';
import { useNavigate } from 'react-router-dom';
import { ethers } from 'ethers';
import contractABI from '../contracts/VotingContract.json';

const CONTRACT_ADDRESS = "0xa131AD247055FD2e2aA8b156A11bdEc81b9eAD95"; // Replace with your contract address

export default function Login({ onLogin }) {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleLogin = async (values) => {
    const { address } = values;
    if (window.ethereum) {
      try {
        setLoading(true);
        const web3Provider = new ethers.providers.Web3Provider(window.ethereum);
        const signer = web3Provider.getSigner();
        const votingContract = new ethers.Contract(CONTRACT_ADDRESS, contractABI, signer);

        const isAdmin = await votingContract.isAdmin(address);
        const isRegistered = await votingContract.voters(address);

        if (isRegistered.isRegistered) {
          onLogin(address, isAdmin);
          message.success("Login successful");
        } else {
          message.error("Account not found, please register first");
        }
      } catch (error) {
        console.error("Login failed:", error);
        message.error("Login failed: " + error.message);
      } finally {
        setLoading(false);
      }
    } else {
      message.error("Please install MetaMask!");
    }
  };

  const handleNavigateToRegister = () => {
    navigate('/register');
  };

  return (
    <div style={{ maxWidth: 400, margin: 'auto', padding: '20px' }}>
      <h2>Login</h2>
      <Form onFinish={handleLogin}>
        <Form.Item
          name="address"
          rules={[{ required: true, message: 'Please enter your account' }]}
        >
          <Input placeholder="Email address" />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" loading={loading}>
            Login
          </Button>
        </Form.Item>
      </Form>
      <Button type="primary" onClick={handleNavigateToRegister}>
        Register
      </Button>
    </div>
  );
}