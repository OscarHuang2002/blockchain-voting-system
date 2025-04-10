import React, { useState, useEffect } from "react";
import {
  Table,
  Card,
  Typography,
  Spin,
  Empty,
  Tag,
  Image,
  Input,
  Space,
} from "antd";
import { CheckCircleOutlined, SearchOutlined } from "@ant-design/icons";
import axios from "axios";

const { Title, Text } = Typography;
const { Search } = Input;

const VoteHistory = ({ account }) => {
  const [loading, setLoading] = useState(true);
  const [voteHistory, setVoteHistory] = useState([]);
  const [filteredVoteHistory, setFilteredVoteHistory] = useState([]);
  const [searchText, setSearchText] = useState("");

  useEffect(() => {
    if (account) {
      fetchVoteHistory();
    }
  }, [account]);

  // 当原始投票历史或搜索文本变化时，过滤数据
  useEffect(() => {
    filterVoteHistory();
  }, [voteHistory, searchText]);

  const fetchVoteHistory = async () => {
    try {
      setLoading(true);
      const response = await axios.get(
        `http://localhost:8080/userVoteHistory?address=${account}`
      );
      if (response.data && response.data.votes) {
        setVoteHistory(response.data.votes);
        setFilteredVoteHistory(response.data.votes);
      }
    } catch (error) {
      console.error("获取投票历史失败:", error);
    } finally {
      setLoading(false);
    }
  };

  // 处理搜索功能
  const handleSearch = (value) => {
    setSearchText(value);
  };

  // 过滤投票历史
  const filterVoteHistory = () => {
    if (!searchText) {
      setFilteredVoteHistory(voteHistory);
      return;
    }

    const lowerCaseSearch = searchText.toLowerCase();
    const filtered = voteHistory.filter(
      (vote) =>
        vote.projectId.toString().includes(lowerCaseSearch) ||
        vote.projectName.toLowerCase().includes(lowerCaseSearch) ||
        vote.candidateName.toLowerCase().includes(lowerCaseSearch)
    );

    setFilteredVoteHistory(filtered);
  };

  const columns = [
    {
      title: "项目ID",
      dataIndex: "projectId",
      key: "projectId",
      render: (text) => <Text code>{text}</Text>,
    },
    {
      title: "项目名称",
      dataIndex: "projectName",
      key: "projectName",
    },
    {
      title: "候选人",
      dataIndex: "candidateName",
      key: "candidateName",
      render: (text, record) => (
        <div style={{ display: "flex", alignItems: "center" }}>
          <Image
            src={record.candidateImage}
            alt={text}
            width={40}
            height={40}
            style={{ borderRadius: "50%", marginRight: 10 }}
            fallback="data:image/png;base64,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"
          />
          {text}
        </div>
      ),
    },
    {
      title: "区块哈希",
      dataIndex: "blockHash",
      key: "blockHash",
      ellipsis: true,
      render: (text) => (
        <Text copyable style={{ width: 180 }} ellipsis>
          {text}
        </Text>
      ),
    },
    {
      title: "状态",
      key: "status",
      render: () => (
        <Tag icon={<CheckCircleOutlined />} color="success">
          投票成功
        </Tag>
      ),
    },
  ];

  return (
    <Card className="vote-history-card">
      <Title level={3}>我的投票历史</Title>

      <div style={{ marginBottom: 16 }}>
        <Search
          placeholder="搜索项目ID、项目名称或候选人"
          allowClear
          enterButton={<SearchOutlined />}
          onSearch={handleSearch}
          onChange={(e) => handleSearch(e.target.value)}
          style={{ width: 300 }}
        />
      </div>

      <div className="vote-history-content">
        <Spin spinning={loading}>
          {filteredVoteHistory.length > 0 ? (
            <Table
              dataSource={filteredVoteHistory}
              columns={columns}
              rowKey={(record) => `${record.projectId}-${record.candidateId}`}
              pagination={
                filteredVoteHistory.length > 10 ? { pageSize: 10 } : false
              }
            />
          ) : (
            <Empty
              description={
                searchText ? "没有找到匹配的投票记录" : "您还没有参与任何投票"
              }
              style={{ marginTop: 40 }}
            />
          )}
        </Spin>
      </div>
    </Card>
  );
};

export default VoteHistory;
