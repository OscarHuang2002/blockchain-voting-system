package services

import (
	"fmt"
	"log"
	"math/big"

	"backend/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractService struct {
	client       *ethclient.Client
	contract     *contracts.Contracts
	auth         *bind.TransactOpts
	contractAddr common.Address
}

// 初始化合约服务
func NewContractService(clientURL string, privateKey string, contractAddress string) (*ContractService, error) {
	// 连接以太坊客户端
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	// 将私钥字符串解析为 *ecdsa.PrivateKey
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	// 获取链 ID（Hardhat 默认链 ID 为 31337）
	chainID := big.NewInt(31337)

	// 创建交易选项
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction signer: %v", err)
	}

	// 初始化合约实例
	contractAddr := common.HexToAddress(contractAddress)
	contract, err := contracts.NewContracts(contractAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to load contract: %v", err)
	}

	return &ContractService{
		client:       client,
		contract:     contract,
		auth:         auth,
		contractAddr: contractAddr,
	}, nil
}

// 调用 registerVoter 方法
func (cs *ContractService) RegisterVoter(voterAddress string) error {
	// 将选民地址转换为以太坊地址格式
	voter := common.HexToAddress(voterAddress)

	// 调用智能合约的 registerVoter 方法
	tx, err := cs.contract.RegisterVoter(cs.auth, voter)
	if err != nil {
		return fmt.Errorf("failed to register voter: %v", err)
	}

	// 打印交易哈希
	log.Printf("Transaction sent: %s", tx.Hash().Hex())
	return nil
}
