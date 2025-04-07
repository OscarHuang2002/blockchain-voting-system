// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;

contract VotingContract {
    address public admin;

    // 全局注册的选民（如果需要全局注册，可提前注册后在各项目中使用）
    struct Voter {
        bool isRegistered;
    }
    mapping(address => Voter) public voters;
    
    // 候选人结构体（不再包含地址，新增图片与简介字段）
    struct Candidate {
        uint256 id;
        string name;
        string imageUrl;             // 候选人图片链接
        string candidateDescription; // 候选人简介
        uint256 voteCount;
        bool isActive;
    }
    
    // 投票记录结构体：记录候选人ID与投票时的区块哈希
    struct VoteDetail {
        uint256 candidateId;
        bytes32 blockHash;
    }
    
    // 投票项目结构体，包含项目简介、投票状态、候选人列表、投票记录（每个地址对应一次投票）
    struct VotingProject {
        uint256 id;
        string description;     // 项目简介
        bool votingOpen;        // 投票状态
        Candidate[] candidates; // 候选人列表
        mapping(address => bool) hasVoted; // 记录用户是否在本项目中已投票
    }
    
    // 使用映射存储投票项目，并通过 projectCount 记录项目数量
    mapping(uint256 => VotingProject) private projects;
    uint256 public projectCount;
    
    // 投票记录映射，按项目和选民记录投票详情
    mapping(uint256 => mapping(address => VoteDetail)) public voteDetails;

    // 事件定义
    event VoterRegistered(address voter);
    event VotingProjectCreated(uint256 projectId, string description);
    event VotingProjectStarted(uint256 projectId);
    event VotingProjectEnded(uint256 projectId);
    event VoteCast(address indexed voter, uint256 projectId, uint256 candidateId, bytes32 blockHash);
    event CandidateAdded(uint256 projectId, uint256 candidateId, string name, string imageUrl, string candidateDescription);
    event CandidateUpdated(uint256 projectId, uint256 candidateId, string newName, string newImageUrl, string newCandidateDescription);
    event CandidateDeleted(uint256 projectId, uint256 candidateId);

    // 仅管理员访问
    modifier onlyAdmin() {
        require(msg.sender == admin, "Only admin can perform this action");
        _;
    }

    // 构造函数：初始化管理员
    constructor() {
        admin = msg.sender;
    }

    // 注册选民（全局）
    function registerVoter(address _voter) external {
        require(!voters[_voter].isRegistered, "Voter already registered");
        voters[_voter] = Voter(true);
        emit VoterRegistered(_voter);
    }
    
    // 创建投票项目，设置项目简介
    function createVotingProject(string memory _description) external onlyAdmin {
        projectCount++;
        VotingProject storage project = projects[projectCount];
        project.id = projectCount;
        project.description = _description;
        project.votingOpen = false; // 默认关闭
        emit VotingProjectCreated(project.id, _description);
    }
    
    // 开始某个项目的投票
    function startVotingForProject(uint256 _projectId) external onlyAdmin {
        VotingProject storage project = projects[_projectId];
        require(!project.votingOpen, "Voting already started for this project");
        project.votingOpen = true;
        emit VotingProjectStarted(_projectId);
    }
    
    // 结束某个项目的投票
    function endVotingForProject(uint256 _projectId) external onlyAdmin {
        VotingProject storage project = projects[_projectId];
        require(project.votingOpen, "Voting not started for this project");
        project.votingOpen = false;
        emit VotingProjectEnded(_projectId);
    }
    
    // 添加候选人至指定项目
    function addCandidate(
        uint256 _projectId,
        string memory _name,
        string memory _imageUrl,
        string memory _candidateDescription
    ) external onlyAdmin {
        VotingProject storage project = projects[_projectId];
        uint256 candidateId = project.candidates.length;
        project.candidates.push(Candidate({
            id: candidateId,
            name: _name,
            imageUrl: _imageUrl,
            candidateDescription: _candidateDescription,
            voteCount: 0,
            isActive: true
        }));
        emit CandidateAdded(_projectId, candidateId, _name, _imageUrl, _candidateDescription);
    }
    
    // 修改候选人信息
    function updateCandidate(
        uint256 _projectId,
        uint256 _candidateId,
        string memory _newName,
        string memory _newImageUrl,
        string memory _newCandidateDescription
    ) external onlyAdmin {
        VotingProject storage project = projects[_projectId];
        require(_candidateId < project.candidates.length, "Invalid candidate ID");
        Candidate storage candidate = project.candidates[_candidateId];
        require(candidate.isActive, "Candidate is not active");
        candidate.name = _newName;
        candidate.imageUrl = _newImageUrl;
        candidate.candidateDescription = _newCandidateDescription;
        emit CandidateUpdated(_projectId, _candidateId, _newName, _newImageUrl, _newCandidateDescription);
    }
    
    // 删除候选人（逻辑删除）
    function deleteCandidate(uint256 _projectId, uint256 _candidateId) external onlyAdmin {
        VotingProject storage project = projects[_projectId];
        require(_candidateId < project.candidates.length, "Invalid candidate ID");
        Candidate storage candidate = project.candidates[_candidateId];
        require(candidate.isActive, "Candidate is already inactive");
        candidate.isActive = false;
        emit CandidateDeleted(_projectId, _candidateId);
    }
    
    // 投票：针对指定项目投票，每个选民在每个项目中仅可投票一次
    function vote(uint256 _projectId, uint256 _candidateId) external {
        require(voters[msg.sender].isRegistered, "Voter not registered");
        VotingProject storage project = projects[_projectId];
        require(project.votingOpen, "Voting is not open for this project");
        require(!project.hasVoted[msg.sender], "Already voted in this project");
        require(_candidateId < project.candidates.length, "Invalid candidate");
        Candidate storage candidate = project.candidates[_candidateId];
        require(candidate.isActive, "Candidate is not active");

        candidate.voteCount++;
        project.hasVoted[msg.sender] = true;
        // 记录投票详情：由于当前区块的哈希不可获取，因此保存前一个区块的哈希
        bytes32 voteBlockHash = blockhash(block.number - 1);
        voteDetails[_projectId][msg.sender] = VoteDetail(_candidateId, voteBlockHash);
        emit VoteCast(msg.sender, _projectId, _candidateId, voteBlockHash);
    }
    
    // 获取指定项目候选人数量
    function getCandidateCount(uint256 _projectId) external view returns (uint256) {
        return projects[_projectId].candidates.length;
    }
    
    // 获取指定项目中单个候选人信息
    function getCandidate(uint256 _projectId, uint256 _candidateId) external view returns (
        uint256, string memory, string memory, string memory, uint256, bool
    ) {
        VotingProject storage project = projects[_projectId];
        require(_candidateId < project.candidates.length, "Invalid candidate ID");
        Candidate storage candidate = project.candidates[_candidateId];
        return (
            candidate.id,
            candidate.name,
            candidate.imageUrl,
            candidate.candidateDescription,
            candidate.voteCount,
            candidate.isActive
        );
    }
    
    // 获取指定项目所有候选人信息（注意：不能直接返回包含映射的结构体，因此只返回候选人数组）
    function getAllCandidates(uint256 _projectId) external view returns (Candidate[] memory) {
        return projects[_projectId].candidates;
    }
    
    // 获取指定项目的投票状态与简介
    function getProjectInfo(uint256 _projectId) external view returns (uint256, string memory, bool) {
        VotingProject storage project = projects[_projectId];
        return (project.id, project.description, project.votingOpen);
    }
    
    // 检查用户是否在指定项目中已经投票
    function hasUserVoted(uint256 _projectId, address _user) external view returns (bool) {
        return projects[_projectId].hasVoted[_user];
    }
    
    // 获取当前用户在指定项目中的投票详情（候选人ID与记录的区块哈希）
    function getMyVoteDetail(uint256 _projectId) external view returns (uint256 candidateId, bytes32 blockHash) {
        require(projects[_projectId].hasVoted[msg.sender], "User has not voted in this project");
        VoteDetail storage detail = voteDetails[_projectId][msg.sender];
        return (detail.candidateId, detail.blockHash);
    }
}
