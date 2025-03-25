const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("VotingContract", function () {
  let VotingContract, votingContract, admin, voter1, voter2;

  beforeEach(async function () {
    // 部署合约
    [admin, voter1, voter2] = await ethers.getSigners();
    VotingContract = await ethers.getContractFactory("VotingContract");
    votingContract = await VotingContract.deploy(); // 不需要调用 deployed()
  });

  it("should register a voter successfully", async function () {
    // 调用 registerVoter 方法
    await expect(votingContract.connect(admin).registerVoter(voter1.address))
      .to.emit(votingContract, "VoterRegistered")
      .withArgs(voter1.address);

    // 验证选民是否已注册
    const voter = await votingContract.voters(voter1.address);
    expect(voter.isRegistered).to.be.true;
    expect(voter.hasVoted).to.be.false;
  });

  it("should allow any user to register voters", async function () {
    // 非管理员尝试注册选民
    await expect(
      votingContract.connect(voter1).registerVoter(voter2.address)
    ).to.emit(votingContract, "VoterRegistered")
      .withArgs(voter2.address);
  
    // 验证选民是否已注册
    const voter = await votingContract.voters(voter2.address);
    expect(voter.isRegistered).to.be.true;
    expect(voter.hasVoted).to.be.false;
  });

  it("should not allow duplicate voter registration", async function () {
    // 第一次注册
    await votingContract.connect(admin).registerVoter(voter1.address);

    // 尝试重复注册
    await expect(
      votingContract.connect(admin).registerVoter(voter1.address)
    ).to.be.revertedWith("Voter already registered");
  });
});