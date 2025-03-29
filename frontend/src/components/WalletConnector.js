import React, { useState } from 'react';
import { Button, message, Spin } from 'antd';
import { WalletOutlined, LogoutOutlined } from '@ant-design/icons';

export default function WalletConnector({ account, onConnect, onDisconnect, loading }) {
  const [isConnecting, setIsConnecting] = useState(false); // 防止重复调用

  const connectWallet = async () => {
    if (isConnecting) return; // 如果正在连接，直接返回
    setIsConnecting(true);

    if (!window.ethereum) {
      
      message.error("MetaMask is not installed");
      setIsConnecting(false);
      return;
    }

    try {
      await onConnect(); // 调用父组件传递的连接逻辑
      message.success("Wallet connected successfully");
    } catch (error) {
      message.error("Wallet connection failed: " + error.message);
    } finally {
      setIsConnecting(false);
    }
  };

  const disconnectWallet = async () => {
    if (!window.ethereum) {
      message.error("MetaMask is not installed");
      return;
    }

    try {
      onDisconnect(); // 调用父组件传递的断开逻辑
      message.success("You have logged out of the current account");
    } catch (error) {
      message.error("Log out failed: " + error.message);
    }
  };

  return (
    <div>
      {account ? (
        <div>
          <Button type="dashed" icon={<WalletOutlined />}>
            {account.slice(0, 6)}...{account.slice(-4)}
          </Button>
          <Button
            type="primary"
            icon={<LogoutOutlined />}
            onClick={disconnectWallet}
            style={{ marginLeft: 8 }}
          >
            Log out
          </Button>
        </div>
      ) : (
        <Button
          type="primary"
          icon={<WalletOutlined />}
          onClick={connectWallet}
          loading={loading || isConnecting}
        >
          {loading || isConnecting ? (
            <span>
              Connecting...
              <Spin size="small" style={{ marginLeft: 8 }} />
            </span>
          ) : (
            "Connect Wallet"
          )}
        </Button>
      )}
    </div>
  );
}