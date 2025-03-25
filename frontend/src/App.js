import React, { useState, useEffect } from 'react';
import { 
  Routes,
  Route,
  Link,
  useNavigate
} from 'react-router-dom';
import { ethers } from 'ethers';
import { Layout, Menu, Spin, Alert } from 'antd';
import { 
  WalletOutlined, 
  UserOutlined, 
  PieChartOutlined,
  LockOutlined 
} from '@ant-design/icons';
import WalletConnector from './components/WalletConnector';
import VoterPanel from './components/VoterPanel';
import AdminPanel from './components/AdminPanel';
import LiveResults from './components/LiveResults';
import Login from './pages/Login';
import Register from './pages/Register';
import contractABI from './contracts/VotingContract.json';
import './App.css';

const { Header, Content, Sider } = Layout;
const CONTRACT_ADDRESS = "0xa131AD247055FD2e2aA8b156A11bdEc81b9eAD95"; // Replace with your contract address

function App() {
  const [provider, setProvider] = useState(null);
  const [contract, setContract] = useState(null);
  const [account, setAccount] = useState('');
  const [isAdmin, setIsAdmin] = useState(false);
  const [loading, setLoading] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const navigate = useNavigate();

  const handleLogin = (address, isAdmin) => {
    setAccount(address);
    setIsAdmin(isAdmin);
    setIsLoggedIn(true);
    navigate('/');
  };

  // connectWallet
  const connectWallet = async () => {
    if (window.ethereum) {
      try {
        setLoading(true);
        await window.ethereum.request({ method: 'eth_requestAccounts' });
        const web3Provider = new ethers.providers.Web3Provider(window.ethereum);
        const signer = web3Provider.getSigner();
        const address = await signer.getAddress();
        
        const votingContract = new ethers.Contract(
          CONTRACT_ADDRESS, 
          contractABI, 
          signer
        );
        
        setProvider(web3Provider);
        setContract(votingContract);
        setAccount(address);
      } catch (error) {
        console.error("Connected failed:", error);
        Alert.error("Wallet connected Failed");
      } finally {
        setLoading(false);
      }
    } else {
      Alert.error("Please install MetaMask!");
    }
  };

  // disconnectWallet
  const disconnectWallet = () => {
    setProvider(null);
    setContract(null);
    setAccount('');
    setIsAdmin(false);
    setIsLoggedIn(false);
  };

  // initBlockchain
  const initBlockchain = async () => {
    if (window.ethereum) {
      try {
        await window.ethereum.request({ method: 'eth_requestAccounts' });
        const web3Provider = new ethers.providers.Web3Provider(window.ethereum);
        const signer = web3Provider.getSigner();
        const address = await signer.getAddress();
        
        const votingContract = new ethers.Contract(
          CONTRACT_ADDRESS, 
          contractABI, 
          signer
        );

        // Verify administrator identity
        const isAdmin = await votingContract.isAdmin(address);
        
        setProvider(web3Provider);
        setContract(votingContract);
        setAccount(address);
        setIsAdmin(isAdmin);
      } catch (error) {
        console.error("Initialization failed:", error);
      }
    }
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

          <Content className="content">
            {loading ? (
              <Spin tip="Blockchain transaction processing..." size="large" />
            ) : (
              <Routes>
                {!isLoggedIn ? (
                  <>
                    <Route path="/login" element={<Login onLogin={handleLogin} />} />
                    <Route path="/register" element={<Register />} />
                    <Route path="*" element={<Login onLogin={handleLogin} />} />
                  </>
                ) : (
                  <>
                    <Route path="/" element={<VoterPanel contract={contract} />} />
                    <Route path="/results" element={<LiveResults contract={contract} />} />
                    {isAdmin && <Route path="/admin" element={<AdminPanel contract={contract} />} />}
                  </>
                )}
              </Routes>
            )}
          </Content>
        </Layout>
      </Layout>
  );
}

export default App;