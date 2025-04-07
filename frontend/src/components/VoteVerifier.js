import React, { useState, useEffect } from "react";
import {
  Card,
  Typography,
  Input,
  Button,
  Form,
  Result,
  Spin,
  Divider,
  Alert,
  Descriptions,
  Tag,
  Image,
  Row,
  Col,
  Collapse,
} from "antd";
import {
  CheckCircleOutlined,
  CloseCircleOutlined,
  SearchOutlined,
  LinkOutlined,
  DownOutlined,
  BlockOutlined,
  FieldTimeOutlined,
  FileTextOutlined,
  NumberOutlined,
} from "@ant-design/icons";
import axios from "axios";
import { ethers } from "ethers";

const { Title, Text, Paragraph } = Typography;
const { Panel } = Collapse;

const VoteVerifier = ({ account, contract }) => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);
  const [verificationResult, setVerificationResult] = useState(null);
  const [errorMessage, setErrorMessage] = useState("");
  const [blockDetails, setBlockDetails] = useState(null);
  const [networkInfo, setNetworkInfo] = useState({
    name: "未知网络",
    chainId: 0,
    isLocalNetwork: false,
  });

  // 获取当前网络信息
  useEffect(() => {
    const getNetworkInfo = async () => {
      if (window.ethereum) {
        try {
          const provider = new ethers.providers.Web3Provider(window.ethereum);
          const network = await provider.getNetwork();

          // Hardhat本地网络的chainId通常是31337
          const isLocalNetwork = network.chainId === 31337;

          setNetworkInfo({
            name: isLocalNetwork ? "Hardhat本地网络" : network.name,
            chainId: network.chainId,
            isLocalNetwork,
          });
        } catch (error) {
          console.error("获取网络信息失败:", error);
        }
      }
    };

    getNetworkInfo();
  }, []);

  const verifyVote = async (values) => {
    try {
      setLoading(true);
      setErrorMessage("");
      setVerificationResult(null);
      setBlockDetails(null);

      const { projectId, blockHash } = values;

      // 验证区块哈希是否存在于区块链上
      const provider = new ethers.providers.Web3Provider(window.ethereum);
      const blockByHash = await provider.getBlock(blockHash).catch(() => null);

      if (!blockByHash) {
        throw new Error("提供的区块哈希在区块链上不存在");
      }

      // 保存区块详情以便显示
      setBlockDetails(blockByHash);

      // 从后端验证投票记录
      const response = await axios.get(
        `http://localhost:8080/verifyVote?address=${account}&projectId=${projectId}&blockHash=${blockHash}`
      );

      setVerificationResult(response.data);
    } catch (error) {
      console.error("验证失败:", error);
      setErrorMessage(error.response?.data?.error || error.message);
    } finally {
      setLoading(false);
    }
  };

  // 将Unix时间戳转换为可读时间
  const formatTimestamp = (timestamp) => {
    if (!timestamp) return "未知";
    const date = new Date(timestamp * 1000);
    return date.toLocaleString("zh-CN");
  };

  // 格式化gas使用数据
  const formatGas = (gas) => {
    if (!gas) return "未知";
    return ethers.utils.commify(gas.toString());
  };

  return (
    <Card className="vote-verifier-card">
      <Title level={3}>投票验证</Title>
      <Paragraph>
        通过区块哈希验证您的投票是否被准确记录在区块链上，验证投票的真实性和不可篡改性。
      </Paragraph>

      <Alert
        message={`当前网络: ${networkInfo.name} (ChainID: ${networkInfo.chainId})`}
        type={networkInfo.isLocalNetwork ? "info" : "warning"}
        description={
          networkInfo.isLocalNetwork
            ? "您正在使用Hardhat本地网络，可以查看完整的区块信息"
            : "注意：非本地网络的区块详情可能无法完全显示"
        }
        showIcon
        style={{ marginBottom: 16 }}
      />

      <Divider />

      <Form form={form} onFinish={verifyVote} layout="vertical">
        <Row gutter={16}>
          <Col xs={24} md={12}>
            <Form.Item
              name="projectId"
              label="项目 ID"
              rules={[{ required: true, message: "请输入项目 ID" }]}
            >
              <Input placeholder="输入您投票的项目 ID" />
            </Form.Item>
          </Col>
          <Col xs={24} md={12}>
            <Form.Item
              name="blockHash"
              label="区块哈希"
              rules={[{ required: true, message: "请输入区块哈希" }]}
            >
              <Input placeholder="输入您投票记录中的区块哈希值" allowClear />
            </Form.Item>
          </Col>
        </Row>

        <Form.Item>
          <Button
            type="primary"
            htmlType="submit"
            icon={<SearchOutlined />}
            loading={loading}
          >
            验证投票
          </Button>
        </Form.Item>
      </Form>

      {errorMessage && (
        <Alert
          message="验证失败"
          description={errorMessage}
          type="error"
          showIcon
          style={{ marginBottom: 16 }}
        />
      )}

      {verificationResult && (
        <div className="verification-result">
          <Divider />
          <Result
            status={verificationResult.verified ? "success" : "error"}
            title={verificationResult.verified ? "验证成功" : "验证失败"}
            subTitle={
              verificationResult.verified
                ? "您的投票已成功验证，区块链上的记录与提供的信息一致"
                : "验证失败，区块链上的记录与提供的信息不匹配"
            }
          />

          {verificationResult.verified && (
            <>
              <Descriptions
                bordered
                column={{ xs: 1, sm: 2 }}
                style={{ marginBottom: 16 }}
              >
                <Descriptions.Item label="投票者地址">
                  <Text copyable ellipsis style={{ maxWidth: 250 }}>
                    {verificationResult.voterAddress}
                  </Text>
                </Descriptions.Item>
                <Descriptions.Item label="项目 ID">
                  {verificationResult.projectId}
                </Descriptions.Item>
                <Descriptions.Item label="项目名称">
                  {verificationResult.projectName}
                </Descriptions.Item>
                <Descriptions.Item label="候选人">
                  <div style={{ display: "flex", alignItems: "center" }}>
                    {verificationResult.candidateImage && (
                      <Image
                        src={verificationResult.candidateImage}
                        alt={verificationResult.candidateName}
                        width={40}
                        height={40}
                        style={{ borderRadius: "50%", marginRight: 10 }}
                        fallback="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMIAAADDCAYAAADQvc6UAAABRWlDQ1BJQ0MgUHJvZmlsZQAAKJFjYGASSSwoyGFhYGDIzSspCnJ3UoiIjFJgf8LAwSDCIMogwMCcmFxc4BgQ4ANUwgCjUcG3awyMIPqyLsis7PPOq3QdDFcvjV3jOD1boQVTPQrgSkktTgbSf4A4LbmgqISBgTEFyFYuLykAsTuAbJEioKOA7DkgdjqEvQHEToKwj4DVhAQ5A9k3gGyB5IxEoBmML4BsnSQk8XQkNtReEOBxcfXxUQg1Mjc0dyHgXNJBSWpFCYh2zi+oLMpMzyhRcASGUqqCZ16yno6CkYGRAQMDKMwhqj/fAIcloxgHQqxAjIHBEugw5sUIsSQpBobtQPdLciLEVJYzMPBHMDBsayhILEqEO4DxG0txmrERhM29nYGBddr//5/DGRjYNRkY/l7////39v///y4Dmn+LgeHANwDrkl1AuO+pmgAAADhlWElmTU0AKgAAAAgAAYdpAAQAAAABAAAAGgAAAAAAAqACAAQAAAABAAAAwqADAAQAAAABAAAAwwAAAAD9b/HnAAAHlklEQVR4Ae3dP3PTWBSGcbGzM6GCKqlIBRV0dHRJFarQ0eUT8LH4BnRU0NHR0UEFVdIlFRV7TzRksomPY8uykTk/zewQfKw/9znv4yvJynLv4uLiV2dBoDiBf4qP3/ARuCRABEFAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghgg0Aj8i0JO4OzsrPv69Wv+hi2qPHr0qNvf39+iI97soRIh4f3z58/u7du3SXX7Xt7Z2enevHmzfQe+oSN2apSAPj09TSrb+XKI/f379+08+A0cNRE2ANkupk+ACNPvkSPcAAEibACyXUyfABGm3yNHuAECRNgAZLuYPgEirKlHu7u7XdyytGwHAd8jjNyng4OD7vnz51dbPT8/7z58+NB9+/bt6jU/TI+AGWHEnrx48eJ/EsSmHzx40L18+fLyzxF3ZVMjEyDCiEDjMYZZS5wiPXnyZFbJaxMhQIQRGzHvWR7XCyOCXsOmiDAi1HmPMMQjDpbpEiDCiL358eNHurW/5SnWdIBbXiDCiA38/Pnzrce2YyZ4//59F3ePLNMl4PbpiL2J0L979+7yDtHDhw8vtzzvdGnEXdvUigSIsCLAWavHp/+qM0BcXMd/q25n1vF57TYBp0a3mUzilePj4+7k5KSLb6gt6ydAhPUzXnoPR0dHl79WGTNCfBnn1uvSCJdegQhLI1vvCk+fPu2ePXt2tZOYEV6/fn31dz+shwAR1sP1cqvLntbEN9MxA9xcYjsxS1jWR4AIa2Ibzx0tc44fYX/16lV6NDFLXH+YL32jwiACRBiEbf5KcXoTIsQSpzXx4N28Ja4BQoK7rgXiydbHjx/P25TaQAJEGAguWy0+2Q8PD6/Ki4R8EVl+bzBOnZY95fq9rj9zAkTI2SxdidBHqG9+skdw43borCXO/ZcJdraPWdv22uIEiLA4q7nvvCug8WTqzQveOH26fodo7g6uFe/a17W3+nFBAkRYENRdb1vkkz1CH9cPsVy/jrhr27PqMYvENYNlHAIesRiBYwRy0V+8iXP8+/fvX11Mr7L7ECueb/r48eMqm7FuI2BGWDEG8cm+7G3NEOfmdcTQw4h9/55lhm7DekRYKQPZF2ArbXTAyu4kDYB2YxUzwg0gi/41ztHnfQG26HbGel/crVrm7tNY+/1btkOEAZ2M05r4FB7r9GbAIdxaZYrHdOsgJ/wCEQY0J74TmOKnbxxT9n3FgGGWWsVdowHtjt9Nnvf7yQM2aZU/TIAIAxrw6dOnAWtZZcoEnBpNuTuObWMEiLAx1HY0ZQJEmHJ3HNvGCBBhY6jtaMoEiJB0Z29vL6ls58vxPcO8/zfrdo5qvKO+d3Fx8Wu8zf1dW4p/cPzLly/dtv9Ts/EbcvGAHhHyfBIhZ6NSiIBTo0LNNtScABFyNiqFCBChULMNNSdAhJyNSiECRCjUbEPNCRAhZ6NSiAARCjXbUHMCRMjZqBQiQIRCzTbUnAARcjYqhQgQoVCzDTUnQIScjUohAkQo1GxDzQkQIWejUogAEQo121BzAkTI2agUIkCEQs021JwAEXI2KoUIEKFQsw01J0CEnI1KIQJEKNRsQ80JECFno1KIABEKNdtQcwJEyNmoFCJAhELNNtScABFyNiqFCBChULMNNSdAhJyNSiECRCjUbEPNCRAhZ6NSiAARCjXbUHMCRMjZqBQiQIRCzTbUnAARcjYqhQgQoVCzDTUnQIScjUohAkQo1GxDzQkQIWejUogAEQo121BzAkTI2agUIkCEQs021JwAEXI2KoUIEKFQsw01J0CEnI1KIQJEKNRsQ80JECFno1KIABEKNdtQcwJEyNmoFCJAhELNNtScABFyNiqFCBChULMNNSdAhJyNSiECRCjUbEPNCRAhZ6NSiAARCjXbUHMCRMjZqBQiQIRCzTbUnAARcjYqhQgQoVCzDTUnQIScjUohAkQo1GxDzQkQIWejUogAEQo121BzAkTI2agUIkCEQs021JwAEXI2KoUIEKFQsw01J0CEnI1KIQJEKNRsQ80JECFno1KIABEKNdtQcwJEyNmoFCJAhELNNtScABFyNiqFCBChULMNNSdAhJyNSiEC/wGgKKC4YMA4TAAAAABJRU5ErkJggg=="
                      />
                    )}
                    {verificationResult.candidateName}
                  </div>
                </Descriptions.Item>
              </Descriptions>

              {/* 区块信息部分 */}
              {blockDetails && (
                <Collapse defaultActiveKey={["1"]} style={{ marginBottom: 20 }}>
                  <Panel
                    header={
                      <span>
                        <BlockOutlined /> 区块详细信息 (高度:{" "}
                        {blockDetails.number})
                      </span>
                    }
                    key="1"
                  >
                    <Descriptions bordered column={{ xs: 1, sm: 2 }}>
                      <Descriptions.Item
                        label={
                          <>
                            <NumberOutlined /> 区块高度
                          </>
                        }
                      >
                        {blockDetails.number}
                      </Descriptions.Item>
                      <Descriptions.Item
                        label={
                          <>
                            <FieldTimeOutlined /> 时间戳
                          </>
                        }
                      >
                        {formatTimestamp(blockDetails.timestamp)}
                      </Descriptions.Item>
                      <Descriptions.Item
                        label={
                          <>
                            <BlockOutlined /> 区块哈希
                          </>
                        }
                      >
                        <Text copyable ellipsis style={{ maxWidth: 250 }}>
                          {blockDetails.hash}
                        </Text>
                      </Descriptions.Item>
                      <Descriptions.Item label="父区块哈希">
                        <Text copyable ellipsis style={{ maxWidth: 250 }}>
                          {blockDetails.parentHash}
                        </Text>
                      </Descriptions.Item>
                      <Descriptions.Item label="交易数量">
                        {blockDetails.transactions
                          ? blockDetails.transactions.length
                          : 0}
                      </Descriptions.Item>
                      <Descriptions.Item label="Gas 限制">
                        {formatGas(blockDetails.gasLimit)}
                      </Descriptions.Item>
                      <Descriptions.Item label="Gas 使用量">
                        {formatGas(blockDetails.gasUsed)}
                      </Descriptions.Item>
                      <Descriptions.Item label="矿工/验证人">
                        <Text copyable>{blockDetails.miner}</Text>
                      </Descriptions.Item>
                      {blockDetails.baseFeePerGas && (
                        <Descriptions.Item label="基础 Gas 费">
                          {ethers.utils.formatUnits(
                            blockDetails.baseFeePerGas,
                            "gwei"
                          )}{" "}
                          Gwei
                        </Descriptions.Item>
                      )}
                      <Descriptions.Item label="区块大小">
                        {blockDetails.size
                          ? `${blockDetails.size} 字节`
                          : "未知"}
                      </Descriptions.Item>
                    </Descriptions>

                    {blockDetails.transactions &&
                      blockDetails.transactions.length > 0 && (
                        <>
                          <Divider orientation="left">
                            <FileTextOutlined /> 交易列表
                          </Divider>
                          <ul
                            style={{
                              maxHeight: "200px",
                              overflowY: "auto",
                              padding: "0 20px",
                            }}
                          >
                            {blockDetails.transactions.map((tx, index) => (
                              <li key={index}>
                                <Text
                                  copyable
                                  ellipsis
                                  style={{ maxWidth: "100%" }}
                                >
                                  {tx}
                                </Text>
                              </li>
                            ))}
                          </ul>
                        </>
                      )}
                  </Panel>
                </Collapse>
              )}

              {/* 提供查看区块的链接（仅在非本地网络时显示） */}
              {!networkInfo.isLocalNetwork && (
                <Descriptions bordered>
                  <Descriptions.Item label="区块浏览器">
                    <Button
                      type="link"
                      icon={<LinkOutlined />}
                      href={`https://goerli.etherscan.io/block/${verificationResult.blockHash}`}
                      target="_blank"
                    >
                      在 Etherscan 中查看区块
                    </Button>
                  </Descriptions.Item>
                </Descriptions>
              )}
            </>
          )}
        </div>
      )}
    </Card>
  );
};

export default VoteVerifier;
