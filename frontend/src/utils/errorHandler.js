import { message } from "antd";

export const handleError = (error) => {
  const errorMessage =
    error?.response?.data?.error || // 后端返回的错误信息
    error?.reason || // ethers.js 合约调用的错误原因
    error?.data?.message || // ethers.js 内部错误
    error?.message || // 通用错误信息
    "发生未知错误";

  // 根据错误内容显示友好的提示
  if (errorMessage.includes("Voting is not open")) {
    message.error("投票尚未开始，请稍后再试！");
  } else if (errorMessage.includes("Voting not started")) {
    message.error("投票尚未开始，无法结束投票！");
  } else if (errorMessage.includes("Already voted")) {
    message.error("您已经投过票了！");
  } else if (errorMessage.includes("Candidate is not active")) {
    message.error("该候选人已被删除，无法投票！");
  } else if (errorMessage.includes("Only admin can perform this action")) {
    message.error("只有管理员可以执行此操作！");
  } else {
    message.error(`操作失败: ${errorMessage}`);
  }
};
