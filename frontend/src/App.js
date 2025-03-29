import React, { useState } from 'react';
import { Routes, Route, Link } from 'react-router-dom';
import { ethers } from 'ethers';
import { Layout, Menu, Spin } from 'antd';
import { WalletOutlined, UserOutlined, PieChartOutlined, LockOutlined } from '@ant-design/icons';
import WalletConnector from './components/WalletConnector';
import VoterPanel from './components/VoterPanel';
import AdminPanel from './components/AdminPanel';
import LiveResults from './components/LiveResults';
import Login from './pages/Login';
import Register from './pages/Register';
import contractABI from './contracts/VotingContract.json';
import './App.css';

const { Header, Content, Sider } = Layout;
const CONTRACT_ADDRESS = "0x5fbdb2315678afecb367f032d93f642f64180aa3"; // 替换为你的合约地址

function App() {
  const [provider, setProvider] = useState(null);
  const [contract, setContract] = useState(null);
  const [account, setAccount] = useState('');
  const [isAdmin, setIsAdmin] = useState(false);
  const [loading, setLoading] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const connectWallet = async () => {
    if (window.ethereum) {
      try {
        setLoading(true);
        await window.ethereum.request({ method: 'eth_requestAccounts' });
        const web3Provider = new ethers.providers.Web3Provider(window.ethereum);
        const signer = web3Provider.getSigner();
        const address = await signer.getAddress();

        const votingContract = new ethers.Contract(CONTRACT_ADDRESS, contractABI, signer);

        setProvider(web3Provider);
        setContract(votingContract);
        setAccount(address);
        console.log("Wallet address set in App.js:", address); 
      } catch (error) {
        console.error("Connection failed:", error);
        message.error("Wallet connection failed");
      } finally {
        setLoading(false);
      }
    } else {
      message.error("Please install MetaMask!");
    }
  };

  const disconnectWallet = () => {
    setProvider(null);
    setContract(null);
    setAccount('');
    setIsAdmin(false);
    setIsLoggedIn(false);
  };

  return (
    <Layout>
      <Header className="header">
        <div className="logo">Blockchain E-Voting System</div>
        <WalletConnector
          account={account}
          onConnect={connectWallet}
          onDisconnect={disconnectWallet}
          loading={loading}
        />
      </Header>

      <Layout>
        <Content className="content">
          <Routes>
            {!isLoggedIn ? (
              <>
                <Route path="/login" element={<Login walletAddress={account} setToken={() => setIsLoggedIn(true)} />}
                  key={account} // 动态 key，确保组件重新渲染    
                />
                <Route path="/register" element={<Register walletAddress={account} />} />
                <Route path="*" element={<Login walletAddress={account} setToken={() => setIsLoggedIn(true)} />} />
              </>
            ) : (
              <>
                <Route path="/" element={<VoterPanel contract={contract} />} />
                <Route path="/results" element={<LiveResults contract={contract} />} />
                {isAdmin && <Route path="/admin" element={<AdminPanel contract={contract} />} />}
              </>
            )}
          </Routes>
        </Content>
      </Layout>

  
        {isLoggedIn && (
          <Sider width={200} theme="light">
            <Menu mode="inline">
              <Menu.Item key="vote" icon={<UserOutlined />}>
                <Link to="/">VoterPanel</Link>
              </Menu.Item>
              <Menu.Item key="results" icon={<PieChartOutlined />}>
                <Link to="/results">LiveResults</Link>
              </Menu.Item>
              {isAdmin && (
                <Menu.Item key="admin" icon={<LockOutlined />}>
                  <Link to="/admin">AdminPanel</Link>
                </Menu.Item>
              )}
            </Menu>
          </Sider>
        )}    
    </Layout>
  );
}

export default App;