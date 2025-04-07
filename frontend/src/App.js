import React, { useEffect, useState } from "react";
import { Routes, Route, Link } from "react-router-dom";
import { ethers } from "ethers";
import { Layout, Menu, Spin, message, ConfigProvider } from "antd";
import {
  WalletOutlined,
  UserOutlined,
  PieChartOutlined,
  LockOutlined,
  HistoryOutlined, // 添加历史图标
  SafetyOutlined,
} from "@ant-design/icons";
import WalletConnector from "./components/WalletConnector";
import VoterPanel from "./components/VoterPanel";
import AdminPanel from "./components/AdminPanel";
import LiveResults from "./components/LiveResults";
import VoteHistory from "./components/VoteHistory"; // 导入新组件
import Login from "./pages/Login";
import Register from "./pages/Register";
import contractABI from "./contracts/VotingContract.json";
import "./App.css";
import VoteVerifier from "./components/VoteVerifier";

const { Header, Content, Sider } = Layout;
const CONTRACT_ADDRESS = "0x5fbdb2315678afecb367f032d93f642f64180aa3"; // 替换为你的合约地址

message.config({
  top: 100,
  duration: 2,
  maxCount: 1,
});

function App() {
  const [provider, setProvider] = useState(null);
  const [contract, setContract] = useState(null);
  const [account, setAccount] = useState("");
  const [isAdmin, setIsAdmin] = useState(false);
  const [loading, setLoading] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [selectedProjectId, setSelectedProjectId] = useState(null);

  const handleProjectChange = (value) => {
    setSelectedProjectId(value); // 更新选中的 projectId
  };

  useEffect(() => {
    const savedAccount = localStorage.getItem("walletAccount");
    if (savedAccount) {
      setAccount(savedAccount);
      setIsLoggedIn(true);
    }
  }, []);

  // App.js 中添加 useEffect 监听账户变更
  useEffect(() => {
    if (window.ethereum) {
      const handleAccountsChanged = (accounts) => {
        if (accounts.length === 0) {
          message.warning("请连接 MetaMask");
          disconnectWallet();
        } else if (accounts[0] !== account) {
          message.warning("检测到账户变更，请重新登录");
          disconnectWallet();
        }
      };

      window.ethereum.on("accountsChanged", handleAccountsChanged);
      return () => {
        window.ethereum.removeListener(
          "accountsChanged",
          handleAccountsChanged
        );
      };
    }
  }, [account]);

  const connectWallet = async () => {
    if (window.ethereum) {
      try {
        setLoading(true);
        await window.ethereum.request({ method: "eth_requestAccounts" });
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
        localStorage.setItem("walletAccount", address); // 保存钱包地址到本地存储

        return address; // 确保返回钱包地址
      } catch (error) {
        console.error("Connection failed:", error);
        message.error("Wallet connection failed");
        throw error; // 抛出错误以便 WalletConnector 捕获
      } finally {
        setLoading(false);
      }
    } else {
      message.error("Please install MetaMask!");
      throw new Error("MetaMask is not installed");
    }
  };

  const disconnectWallet = () => {
    setProvider(null);
    setContract(null);
    setAccount("");
    setIsAdmin(false);
    setIsLoggedIn(false);
    localStorage.removeItem("walletAccount"); // 清除本地存储的钱包地址
  };

  return (
    <ConfigProvider>
      <Layout>
        <Header className="header">
          <div className="logo">区块链选举投票系统</div>
          <WalletConnector
            account={account}
            onConnect={connectWallet}
            onDisconnect={disconnectWallet}
            loading={loading}
          />
        </Header>

        <Content>
          <Layout>
            {isLoggedIn && (
              <Sider width={200} theme="light">
                <Menu
                  mode="inline"
                  items={[
                    {
                      key: "vote",
                      icon: <UserOutlined />,
                      label: <Link to="/VoterPanel">投票界面</Link>,
                    },
                    {
                      key: "results",
                      icon: <PieChartOutlined />,
                      label: <Link to="/results">实时投票统计</Link>,
                    },
                    {
                      key: "voteHistory",
                      icon: <HistoryOutlined />,
                      label: <Link to="/voteHistory">我的投票记录</Link>,
                    },
                    {
                      key: "verify",
                      icon: <SafetyOutlined />,
                      label: <Link to="/verify">投票验证</Link>,
                    },
                    isAdmin && {
                      key: "admin",
                      icon: <LockOutlined />,
                      label: <Link to="/AdminPanel">管理员界面</Link>,
                    },
                  ].filter(Boolean)}
                />
              </Sider>
            )}
            <Content className="content">
              <Routes>
                {!isLoggedIn ? (
                  <>
                    <Route
                      path="/login"
                      element={
                        <Login
                          walletAddress={account}
                          setToken={() => setIsLoggedIn(true)}
                          setIsAdmin={setIsAdmin} // 传递 setIsAdmin
                        />
                      }
                    />
                    <Route
                      path="/register"
                      element={<Register walletAddress={account} />}
                    />
                    <Route
                      path="*"
                      element={
                        <Login
                          walletAddress={account}
                          setToken={() => setIsLoggedIn(true)}
                          setIsAdmin={setIsAdmin} // 传递 setIsAdmin
                        />
                      }
                    />
                  </>
                ) : (
                  <>
                    <Route
                      path="/VoterPanel"
                      element={
                        <VoterPanel contract={contract} account={account} />
                      }
                    />
                    <Route
                      path="/results"
                      element={<LiveResults contract={contract} />}
                    />
                    <Route
                      path="/voteHistory"
                      element={<VoteHistory account={account} />}
                    />
                    <Route
                      path="/verify"
                      element={
                        <VoteVerifier account={account} contract={contract} />
                      }
                    />
                    {isAdmin && (
                      <Route
                        path="/AdminPanel"
                        element={<AdminPanel contract={contract} />}
                      />
                    )}
                  </>
                )}
              </Routes>
            </Content>
          </Layout>
        </Content>
      </Layout>
    </ConfigProvider>
  );
}

export default App;
