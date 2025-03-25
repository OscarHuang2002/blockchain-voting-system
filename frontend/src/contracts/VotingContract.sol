// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;

contract BlockchainVotingSystem {
    // 管理员地址
    address public admin;
    
    // 投票状态
    bool public votingOpen = false;

    // 选民结构
    struct Voter {
        bool isRegistered;
        bool hasVoted;
    }

    // 候选人结构
    struct Candidate {
        uint256 id;
        string name;
        uint256 voteCount;
    }

    // 数据结构存储
    mapping(address => Voter) public voters;
    Candidate[] public candidates;
    
    // 事件定义
    event VoterRegistered(address voter);
    event VotingStarted();
    event VotingEnded();
    event VoteCast(address indexed voter, uint256 candidateId);
    event CandidateAdded(uint256 candidateId, string name);

    // 权限修饰符
    modifier onlyAdmin() {
        require(msg.sender == admin, "Only admin can perform this action");
        _;
    }

    modifier votingIsOpen() {
        require(votingOpen, "Voting is not open");
        _;
    }

    constructor() {
        admin = msg.sender;
    }

    // 注册选民
    function registerVoter(address _voter) external {
        require(!voters[_voter].isRegistered, "Voter already registered");
        voters[_voter] = Voter(true, false);
        emit VoterRegistered(_voter);
    }

    // 添加候选人（管理员操作）
    function addCandidate(string memory _name) external onlyAdmin {
        candidates.push(Candidate({
            id: candidates.length,
            name: _name,
            voteCount: 0
        }));
        emit CandidateAdded(candidates.length - 1, _name);
    }

    // 开始投票（管理员操作）
    function startVoting() external onlyAdmin {
        require(!votingOpen, "Voting already started");
        votingOpen = true;
        emit VotingStarted();
    }

    // 结束投票（管理员操作）
    function endVoting() external onlyAdmin {
        require(votingOpen, "Voting not started");
        votingOpen = false;
        emit VotingEnded();
    }

    // 执行投票
    function vote(uint256 _candidateId) external votingIsOpen {
        require(voters[msg.sender].isRegistered, "Voter not registered");
        require(!voters[msg.sender].hasVoted, "Already voted");
        require(_candidateId < candidates.length, "Invalid candidate");

        candidates[_candidateId].voteCount++;
        voters[msg.sender].hasVoted = true;
        emit VoteCast(msg.sender, _candidateId);
    }

    // 获取候选人数量
    function getCandidateCount() external view returns (uint256) {
        return candidates.length;
    }

    // 获取投票状态
    function getVotingStatus() external view returns (bool) {
        return votingOpen;
    }
    // 验证管理员身份
    function isAdmin(address _user) external view returns (bool) {
        return _user == admin;    
    }
}