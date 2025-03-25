const hre = require("hardhat");

async function main() {
  const VotingContractFactory = await hre.ethers.getContractFactory("VotingContract");
  const votingContract = await VotingContractFactory.deploy();
  await votingContract.waitForDeployment(); // 改用 waitForDeployment()
  console.log("VotingContract deployed to:", await votingContract.getAddress()); // 使用 getAddress()
}

// 处理错误并运行主函数
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });