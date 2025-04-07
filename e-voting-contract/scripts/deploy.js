const hre = require("hardhat");

async function main() {
  // 获取部署者的地址
  const [deployer] = await hre.ethers.getSigners();

  console.log("Deploying contract with the account:", deployer.address);

  // 获取 VotingContract 合约工厂
  const VotingContractFactory = await hre.ethers.getContractFactory("VotingContract");

  // 部署合约
  const votingContract = await VotingContractFactory.deploy();

  // 等待部署完成
  await votingContract.waitForDeployment();

  // 输出合约地址
  console.log("VotingContract deployed to:", await votingContract.getAddress());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });