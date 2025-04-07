import React, { useState } from "react";
import { Button, message, Spin } from "antd";
import { WalletOutlined, LogoutOutlined } from "@ant-design/icons";

export default function WalletConnector({
  account,
  onConnect,
  onDisconnect,
  loading,
}) {
  const [isConnecting, setIsConnecting] = useState(false);

  const connectWallet = async () => {
    if (isConnecting) return;
    setIsConnecting(true);

    if (!window.ethereum) {
      message.error("MetaMask没有安装");
      setIsConnecting(false);
      return;
    }

    try {
      await onConnect();
      message.success("钱包连接成功");
    } catch (error) {
      message.error("钱包连接失败: " + error.message);
    } finally {
      setIsConnecting(false);
    }
  };

  const disconnectWallet = async () => {
    if (!window.ethereum) {
      message.error("MetaMask没有安装");
      return;
    }

    try {
      onDisconnect();
      message.success("你已登出当前账户");
    } catch (error) {
      message.error("登出失败: " + error.message);
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
            登出
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
              连接中...
              <Spin size="small" style={{ marginLeft: 8 }} />
            </span>
          ) : (
            "连接钱包"
          )}
        </Button>
      )}
    </div>
  );
}
