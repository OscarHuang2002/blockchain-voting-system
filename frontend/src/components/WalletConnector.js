import React from 'react';
import { Button, message, Spin } from 'antd';
import { WalletOutlined, LogoutOutlined} from '@ant-design/icons';

export default function WalletConnector({ account, onConnect, onDisconnect, loading }) {
  const connectWallet = async () => {
    try {
      await onConnect();
      message.success("Wallet connected successfully");
    } catch (error) {
      message.error("Wallet connected failed: " + error.message);
    }
  };

  const disconnectWallet = async () => {
    try {
      await window.ethereum.request({
        method: 'wallet_requestPermissions',
        params: [{
          eth_accounts: {}
        }]
      });
      onDisconnect();
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
          <Button type="primary" icon={<LogoutOutlined />} onClick={disconnectWallet} style={{ marginLeft: 8 }}>
            Log out
          </Button>
        </div>
      ) : (
        <Button 
          type="primary" 
          icon={<WalletOutlined />}
          onClick={connectWallet}
          loading={loading}
        >
          {loading ? (
            <span>
              Connecting...
              <Spin size="small" style={{ marginLeft: 8 }} />
            </span>
          ) : "ConnectWallet"}
        </Button>
      )}
    </div>
  );
}
