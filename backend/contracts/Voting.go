// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VotingContractCandidate is an auto generated low-level Go binding around an user-defined struct.
type VotingContractCandidate struct {
	Id                   *big.Int
	Name                 string
	ImageUrl             string
	CandidateDescription string
	VoteCount            *big.Int
	IsActive             bool
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"imageUrl\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"candidateDescription\",\"type\":\"string\"}],\"name\":\"CandidateAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"}],\"name\":\"CandidateDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newImageUrl\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newCandidateDescription\",\"type\":\"string\"}],\"name\":\"CandidateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VoterRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"VotingProjectCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"}],\"name\":\"VotingProjectEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"}],\"name\":\"VotingProjectStarted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_imageUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_candidateDescription\",\"type\":\"string\"}],\"name\":\"addCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"createVotingProject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_candidateId\",\"type\":\"uint256\"}],\"name\":\"deleteCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"endVotingForProject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"getAllCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"candidateDescription\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structVotingContract.Candidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_candidateId\",\"type\":\"uint256\"}],\"name\":\"getCandidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"getCandidateCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"getMyVoteDetail\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"getProjectInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"hasUserVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"projectCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"registerVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"startVotingForProject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_candidateId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_newName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_newImageUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_newCandidateDescription\",\"type\":\"string\"}],\"name\":\"updateCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_candidateId\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voteDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contracts *ContractsCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contracts *ContractsSession) Admin() (common.Address, error) {
	return _Contracts.Contract.Admin(&_Contracts.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contracts *ContractsCallerSession) Admin() (common.Address, error) {
	return _Contracts.Contract.Admin(&_Contracts.CallOpts)
}

// GetAllCandidates is a free data retrieval call binding the contract method 0xf9ded646.
//
// Solidity: function getAllCandidates(uint256 _projectId) view returns((uint256,string,string,string,uint256,bool)[])
func (_Contracts *ContractsCaller) GetAllCandidates(opts *bind.CallOpts, _projectId *big.Int) ([]VotingContractCandidate, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAllCandidates", _projectId)

	if err != nil {
		return *new([]VotingContractCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new([]VotingContractCandidate)).(*[]VotingContractCandidate)

	return out0, err

}

// GetAllCandidates is a free data retrieval call binding the contract method 0xf9ded646.
//
// Solidity: function getAllCandidates(uint256 _projectId) view returns((uint256,string,string,string,uint256,bool)[])
func (_Contracts *ContractsSession) GetAllCandidates(_projectId *big.Int) ([]VotingContractCandidate, error) {
	return _Contracts.Contract.GetAllCandidates(&_Contracts.CallOpts, _projectId)
}

// GetAllCandidates is a free data retrieval call binding the contract method 0xf9ded646.
//
// Solidity: function getAllCandidates(uint256 _projectId) view returns((uint256,string,string,string,uint256,bool)[])
func (_Contracts *ContractsCallerSession) GetAllCandidates(_projectId *big.Int) ([]VotingContractCandidate, error) {
	return _Contracts.Contract.GetAllCandidates(&_Contracts.CallOpts, _projectId)
}

// GetCandidate is a free data retrieval call binding the contract method 0x4bd46448.
//
// Solidity: function getCandidate(uint256 _projectId, uint256 _candidateId) view returns(uint256, string, string, string, uint256, bool)
func (_Contracts *ContractsCaller) GetCandidate(opts *bind.CallOpts, _projectId *big.Int, _candidateId *big.Int) (*big.Int, string, string, string, *big.Int, bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCandidate", _projectId, _candidateId)

	if err != nil {
		return *new(*big.Int), *new(string), *new(string), *new(string), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, err

}

// GetCandidate is a free data retrieval call binding the contract method 0x4bd46448.
//
// Solidity: function getCandidate(uint256 _projectId, uint256 _candidateId) view returns(uint256, string, string, string, uint256, bool)
func (_Contracts *ContractsSession) GetCandidate(_projectId *big.Int, _candidateId *big.Int) (*big.Int, string, string, string, *big.Int, bool, error) {
	return _Contracts.Contract.GetCandidate(&_Contracts.CallOpts, _projectId, _candidateId)
}

// GetCandidate is a free data retrieval call binding the contract method 0x4bd46448.
//
// Solidity: function getCandidate(uint256 _projectId, uint256 _candidateId) view returns(uint256, string, string, string, uint256, bool)
func (_Contracts *ContractsCallerSession) GetCandidate(_projectId *big.Int, _candidateId *big.Int) (*big.Int, string, string, string, *big.Int, bool, error) {
	return _Contracts.Contract.GetCandidate(&_Contracts.CallOpts, _projectId, _candidateId)
}

// GetCandidateCount is a free data retrieval call binding the contract method 0x26f6a2aa.
//
// Solidity: function getCandidateCount(uint256 _projectId) view returns(uint256)
func (_Contracts *ContractsCaller) GetCandidateCount(opts *bind.CallOpts, _projectId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCandidateCount", _projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCandidateCount is a free data retrieval call binding the contract method 0x26f6a2aa.
//
// Solidity: function getCandidateCount(uint256 _projectId) view returns(uint256)
func (_Contracts *ContractsSession) GetCandidateCount(_projectId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetCandidateCount(&_Contracts.CallOpts, _projectId)
}

// GetCandidateCount is a free data retrieval call binding the contract method 0x26f6a2aa.
//
// Solidity: function getCandidateCount(uint256 _projectId) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetCandidateCount(_projectId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetCandidateCount(&_Contracts.CallOpts, _projectId)
}

// GetMyVoteDetail is a free data retrieval call binding the contract method 0x35ff2223.
//
// Solidity: function getMyVoteDetail(uint256 _projectId) view returns(uint256 candidateId, bytes32 blockHash)
func (_Contracts *ContractsCaller) GetMyVoteDetail(opts *bind.CallOpts, _projectId *big.Int) (struct {
	CandidateId *big.Int
	BlockHash   [32]byte
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMyVoteDetail", _projectId)

	outstruct := new(struct {
		CandidateId *big.Int
		BlockHash   [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CandidateId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetMyVoteDetail is a free data retrieval call binding the contract method 0x35ff2223.
//
// Solidity: function getMyVoteDetail(uint256 _projectId) view returns(uint256 candidateId, bytes32 blockHash)
func (_Contracts *ContractsSession) GetMyVoteDetail(_projectId *big.Int) (struct {
	CandidateId *big.Int
	BlockHash   [32]byte
}, error) {
	return _Contracts.Contract.GetMyVoteDetail(&_Contracts.CallOpts, _projectId)
}

// GetMyVoteDetail is a free data retrieval call binding the contract method 0x35ff2223.
//
// Solidity: function getMyVoteDetail(uint256 _projectId) view returns(uint256 candidateId, bytes32 blockHash)
func (_Contracts *ContractsCallerSession) GetMyVoteDetail(_projectId *big.Int) (struct {
	CandidateId *big.Int
	BlockHash   [32]byte
}, error) {
	return _Contracts.Contract.GetMyVoteDetail(&_Contracts.CallOpts, _projectId)
}

// GetProjectInfo is a free data retrieval call binding the contract method 0xfabf5968.
//
// Solidity: function getProjectInfo(uint256 _projectId) view returns(uint256, string, bool)
func (_Contracts *ContractsCaller) GetProjectInfo(opts *bind.CallOpts, _projectId *big.Int) (*big.Int, string, bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getProjectInfo", _projectId)

	if err != nil {
		return *new(*big.Int), *new(string), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetProjectInfo is a free data retrieval call binding the contract method 0xfabf5968.
//
// Solidity: function getProjectInfo(uint256 _projectId) view returns(uint256, string, bool)
func (_Contracts *ContractsSession) GetProjectInfo(_projectId *big.Int) (*big.Int, string, bool, error) {
	return _Contracts.Contract.GetProjectInfo(&_Contracts.CallOpts, _projectId)
}

// GetProjectInfo is a free data retrieval call binding the contract method 0xfabf5968.
//
// Solidity: function getProjectInfo(uint256 _projectId) view returns(uint256, string, bool)
func (_Contracts *ContractsCallerSession) GetProjectInfo(_projectId *big.Int) (*big.Int, string, bool, error) {
	return _Contracts.Contract.GetProjectInfo(&_Contracts.CallOpts, _projectId)
}

// HasUserVoted is a free data retrieval call binding the contract method 0xdc296ae1.
//
// Solidity: function hasUserVoted(uint256 _projectId, address _user) view returns(bool)
func (_Contracts *ContractsCaller) HasUserVoted(opts *bind.CallOpts, _projectId *big.Int, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "hasUserVoted", _projectId, _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserVoted is a free data retrieval call binding the contract method 0xdc296ae1.
//
// Solidity: function hasUserVoted(uint256 _projectId, address _user) view returns(bool)
func (_Contracts *ContractsSession) HasUserVoted(_projectId *big.Int, _user common.Address) (bool, error) {
	return _Contracts.Contract.HasUserVoted(&_Contracts.CallOpts, _projectId, _user)
}

// HasUserVoted is a free data retrieval call binding the contract method 0xdc296ae1.
//
// Solidity: function hasUserVoted(uint256 _projectId, address _user) view returns(bool)
func (_Contracts *ContractsCallerSession) HasUserVoted(_projectId *big.Int, _user common.Address) (bool, error) {
	return _Contracts.Contract.HasUserVoted(&_Contracts.CallOpts, _projectId, _user)
}

// ProjectCount is a free data retrieval call binding the contract method 0x36fbad26.
//
// Solidity: function projectCount() view returns(uint256)
func (_Contracts *ContractsCaller) ProjectCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "projectCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectCount is a free data retrieval call binding the contract method 0x36fbad26.
//
// Solidity: function projectCount() view returns(uint256)
func (_Contracts *ContractsSession) ProjectCount() (*big.Int, error) {
	return _Contracts.Contract.ProjectCount(&_Contracts.CallOpts)
}

// ProjectCount is a free data retrieval call binding the contract method 0x36fbad26.
//
// Solidity: function projectCount() view returns(uint256)
func (_Contracts *ContractsCallerSession) ProjectCount() (*big.Int, error) {
	return _Contracts.Contract.ProjectCount(&_Contracts.CallOpts)
}

// VoteDetails is a free data retrieval call binding the contract method 0x59191ba1.
//
// Solidity: function voteDetails(uint256 , address ) view returns(uint256 candidateId, bytes32 blockHash)
func (_Contracts *ContractsCaller) VoteDetails(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	CandidateId *big.Int
	BlockHash   [32]byte
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "voteDetails", arg0, arg1)

	outstruct := new(struct {
		CandidateId *big.Int
		BlockHash   [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CandidateId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// VoteDetails is a free data retrieval call binding the contract method 0x59191ba1.
//
// Solidity: function voteDetails(uint256 , address ) view returns(uint256 candidateId, bytes32 blockHash)
func (_Contracts *ContractsSession) VoteDetails(arg0 *big.Int, arg1 common.Address) (struct {
	CandidateId *big.Int
	BlockHash   [32]byte
}, error) {
	return _Contracts.Contract.VoteDetails(&_Contracts.CallOpts, arg0, arg1)
}

// VoteDetails is a free data retrieval call binding the contract method 0x59191ba1.
//
// Solidity: function voteDetails(uint256 , address ) view returns(uint256 candidateId, bytes32 blockHash)
func (_Contracts *ContractsCallerSession) VoteDetails(arg0 *big.Int, arg1 common.Address) (struct {
	CandidateId *big.Int
	BlockHash   [32]byte
}, error) {
	return _Contracts.Contract.VoteDetails(&_Contracts.CallOpts, arg0, arg1)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool isRegistered)
func (_Contracts *ContractsCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "voters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool isRegistered)
func (_Contracts *ContractsSession) Voters(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Voters(&_Contracts.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool isRegistered)
func (_Contracts *ContractsCallerSession) Voters(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Voters(&_Contracts.CallOpts, arg0)
}

// AddCandidate is a paid mutator transaction binding the contract method 0xce4a2743.
//
// Solidity: function addCandidate(uint256 _projectId, string _name, string _imageUrl, string _candidateDescription) returns()
func (_Contracts *ContractsTransactor) AddCandidate(opts *bind.TransactOpts, _projectId *big.Int, _name string, _imageUrl string, _candidateDescription string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addCandidate", _projectId, _name, _imageUrl, _candidateDescription)
}

// AddCandidate is a paid mutator transaction binding the contract method 0xce4a2743.
//
// Solidity: function addCandidate(uint256 _projectId, string _name, string _imageUrl, string _candidateDescription) returns()
func (_Contracts *ContractsSession) AddCandidate(_projectId *big.Int, _name string, _imageUrl string, _candidateDescription string) (*types.Transaction, error) {
	return _Contracts.Contract.AddCandidate(&_Contracts.TransactOpts, _projectId, _name, _imageUrl, _candidateDescription)
}

// AddCandidate is a paid mutator transaction binding the contract method 0xce4a2743.
//
// Solidity: function addCandidate(uint256 _projectId, string _name, string _imageUrl, string _candidateDescription) returns()
func (_Contracts *ContractsTransactorSession) AddCandidate(_projectId *big.Int, _name string, _imageUrl string, _candidateDescription string) (*types.Transaction, error) {
	return _Contracts.Contract.AddCandidate(&_Contracts.TransactOpts, _projectId, _name, _imageUrl, _candidateDescription)
}

// CreateVotingProject is a paid mutator transaction binding the contract method 0xcde5fdca.
//
// Solidity: function createVotingProject(string _description) returns()
func (_Contracts *ContractsTransactor) CreateVotingProject(opts *bind.TransactOpts, _description string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createVotingProject", _description)
}

// CreateVotingProject is a paid mutator transaction binding the contract method 0xcde5fdca.
//
// Solidity: function createVotingProject(string _description) returns()
func (_Contracts *ContractsSession) CreateVotingProject(_description string) (*types.Transaction, error) {
	return _Contracts.Contract.CreateVotingProject(&_Contracts.TransactOpts, _description)
}

// CreateVotingProject is a paid mutator transaction binding the contract method 0xcde5fdca.
//
// Solidity: function createVotingProject(string _description) returns()
func (_Contracts *ContractsTransactorSession) CreateVotingProject(_description string) (*types.Transaction, error) {
	return _Contracts.Contract.CreateVotingProject(&_Contracts.TransactOpts, _description)
}

// DeleteCandidate is a paid mutator transaction binding the contract method 0xcf4a6060.
//
// Solidity: function deleteCandidate(uint256 _projectId, uint256 _candidateId) returns()
func (_Contracts *ContractsTransactor) DeleteCandidate(opts *bind.TransactOpts, _projectId *big.Int, _candidateId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deleteCandidate", _projectId, _candidateId)
}

// DeleteCandidate is a paid mutator transaction binding the contract method 0xcf4a6060.
//
// Solidity: function deleteCandidate(uint256 _projectId, uint256 _candidateId) returns()
func (_Contracts *ContractsSession) DeleteCandidate(_projectId *big.Int, _candidateId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DeleteCandidate(&_Contracts.TransactOpts, _projectId, _candidateId)
}

// DeleteCandidate is a paid mutator transaction binding the contract method 0xcf4a6060.
//
// Solidity: function deleteCandidate(uint256 _projectId, uint256 _candidateId) returns()
func (_Contracts *ContractsTransactorSession) DeleteCandidate(_projectId *big.Int, _candidateId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DeleteCandidate(&_Contracts.TransactOpts, _projectId, _candidateId)
}

// EndVotingForProject is a paid mutator transaction binding the contract method 0xa8ec9f15.
//
// Solidity: function endVotingForProject(uint256 _projectId) returns()
func (_Contracts *ContractsTransactor) EndVotingForProject(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "endVotingForProject", _projectId)
}

// EndVotingForProject is a paid mutator transaction binding the contract method 0xa8ec9f15.
//
// Solidity: function endVotingForProject(uint256 _projectId) returns()
func (_Contracts *ContractsSession) EndVotingForProject(_projectId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.EndVotingForProject(&_Contracts.TransactOpts, _projectId)
}

// EndVotingForProject is a paid mutator transaction binding the contract method 0xa8ec9f15.
//
// Solidity: function endVotingForProject(uint256 _projectId) returns()
func (_Contracts *ContractsTransactorSession) EndVotingForProject(_projectId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.EndVotingForProject(&_Contracts.TransactOpts, _projectId)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x38db6dd3.
//
// Solidity: function registerVoter(address _voter) returns()
func (_Contracts *ContractsTransactor) RegisterVoter(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerVoter", _voter)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x38db6dd3.
//
// Solidity: function registerVoter(address _voter) returns()
func (_Contracts *ContractsSession) RegisterVoter(_voter common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterVoter(&_Contracts.TransactOpts, _voter)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x38db6dd3.
//
// Solidity: function registerVoter(address _voter) returns()
func (_Contracts *ContractsTransactorSession) RegisterVoter(_voter common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterVoter(&_Contracts.TransactOpts, _voter)
}

// StartVotingForProject is a paid mutator transaction binding the contract method 0x4f91c2a4.
//
// Solidity: function startVotingForProject(uint256 _projectId) returns()
func (_Contracts *ContractsTransactor) StartVotingForProject(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "startVotingForProject", _projectId)
}

// StartVotingForProject is a paid mutator transaction binding the contract method 0x4f91c2a4.
//
// Solidity: function startVotingForProject(uint256 _projectId) returns()
func (_Contracts *ContractsSession) StartVotingForProject(_projectId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.StartVotingForProject(&_Contracts.TransactOpts, _projectId)
}

// StartVotingForProject is a paid mutator transaction binding the contract method 0x4f91c2a4.
//
// Solidity: function startVotingForProject(uint256 _projectId) returns()
func (_Contracts *ContractsTransactorSession) StartVotingForProject(_projectId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.StartVotingForProject(&_Contracts.TransactOpts, _projectId)
}

// UpdateCandidate is a paid mutator transaction binding the contract method 0x1c4e8b76.
//
// Solidity: function updateCandidate(uint256 _projectId, uint256 _candidateId, string _newName, string _newImageUrl, string _newCandidateDescription) returns()
func (_Contracts *ContractsTransactor) UpdateCandidate(opts *bind.TransactOpts, _projectId *big.Int, _candidateId *big.Int, _newName string, _newImageUrl string, _newCandidateDescription string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateCandidate", _projectId, _candidateId, _newName, _newImageUrl, _newCandidateDescription)
}

// UpdateCandidate is a paid mutator transaction binding the contract method 0x1c4e8b76.
//
// Solidity: function updateCandidate(uint256 _projectId, uint256 _candidateId, string _newName, string _newImageUrl, string _newCandidateDescription) returns()
func (_Contracts *ContractsSession) UpdateCandidate(_projectId *big.Int, _candidateId *big.Int, _newName string, _newImageUrl string, _newCandidateDescription string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateCandidate(&_Contracts.TransactOpts, _projectId, _candidateId, _newName, _newImageUrl, _newCandidateDescription)
}

// UpdateCandidate is a paid mutator transaction binding the contract method 0x1c4e8b76.
//
// Solidity: function updateCandidate(uint256 _projectId, uint256 _candidateId, string _newName, string _newImageUrl, string _newCandidateDescription) returns()
func (_Contracts *ContractsTransactorSession) UpdateCandidate(_projectId *big.Int, _candidateId *big.Int, _newName string, _newImageUrl string, _newCandidateDescription string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateCandidate(&_Contracts.TransactOpts, _projectId, _candidateId, _newName, _newImageUrl, _newCandidateDescription)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 _projectId, uint256 _candidateId) returns()
func (_Contracts *ContractsTransactor) Vote(opts *bind.TransactOpts, _projectId *big.Int, _candidateId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "vote", _projectId, _candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 _projectId, uint256 _candidateId) returns()
func (_Contracts *ContractsSession) Vote(_projectId *big.Int, _candidateId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Vote(&_Contracts.TransactOpts, _projectId, _candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 _projectId, uint256 _candidateId) returns()
func (_Contracts *ContractsTransactorSession) Vote(_projectId *big.Int, _candidateId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Vote(&_Contracts.TransactOpts, _projectId, _candidateId)
}

// ContractsCandidateAddedIterator is returned from FilterCandidateAdded and is used to iterate over the raw logs and unpacked data for CandidateAdded events raised by the Contracts contract.
type ContractsCandidateAddedIterator struct {
	Event *ContractsCandidateAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsCandidateAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCandidateAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsCandidateAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsCandidateAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCandidateAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCandidateAdded represents a CandidateAdded event raised by the Contracts contract.
type ContractsCandidateAdded struct {
	ProjectId            *big.Int
	CandidateId          *big.Int
	Name                 string
	ImageUrl             string
	CandidateDescription string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCandidateAdded is a free log retrieval operation binding the contract event 0x5311339e5c8f602abcad009c6ae5047e68f731a4f05ebeff0a4496784128ecd3.
//
// Solidity: event CandidateAdded(uint256 projectId, uint256 candidateId, string name, string imageUrl, string candidateDescription)
func (_Contracts *ContractsFilterer) FilterCandidateAdded(opts *bind.FilterOpts) (*ContractsCandidateAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CandidateAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsCandidateAddedIterator{contract: _Contracts.contract, event: "CandidateAdded", logs: logs, sub: sub}, nil
}

// WatchCandidateAdded is a free log subscription operation binding the contract event 0x5311339e5c8f602abcad009c6ae5047e68f731a4f05ebeff0a4496784128ecd3.
//
// Solidity: event CandidateAdded(uint256 projectId, uint256 candidateId, string name, string imageUrl, string candidateDescription)
func (_Contracts *ContractsFilterer) WatchCandidateAdded(opts *bind.WatchOpts, sink chan<- *ContractsCandidateAdded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CandidateAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCandidateAdded)
				if err := _Contracts.contract.UnpackLog(event, "CandidateAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCandidateAdded is a log parse operation binding the contract event 0x5311339e5c8f602abcad009c6ae5047e68f731a4f05ebeff0a4496784128ecd3.
//
// Solidity: event CandidateAdded(uint256 projectId, uint256 candidateId, string name, string imageUrl, string candidateDescription)
func (_Contracts *ContractsFilterer) ParseCandidateAdded(log types.Log) (*ContractsCandidateAdded, error) {
	event := new(ContractsCandidateAdded)
	if err := _Contracts.contract.UnpackLog(event, "CandidateAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCandidateDeletedIterator is returned from FilterCandidateDeleted and is used to iterate over the raw logs and unpacked data for CandidateDeleted events raised by the Contracts contract.
type ContractsCandidateDeletedIterator struct {
	Event *ContractsCandidateDeleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsCandidateDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCandidateDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsCandidateDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsCandidateDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCandidateDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCandidateDeleted represents a CandidateDeleted event raised by the Contracts contract.
type ContractsCandidateDeleted struct {
	ProjectId   *big.Int
	CandidateId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCandidateDeleted is a free log retrieval operation binding the contract event 0x48bf3ef4944e20b74a7ae7804d664a8e9ba28141c5fc26163c87f7fbc9e90e2f.
//
// Solidity: event CandidateDeleted(uint256 projectId, uint256 candidateId)
func (_Contracts *ContractsFilterer) FilterCandidateDeleted(opts *bind.FilterOpts) (*ContractsCandidateDeletedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CandidateDeleted")
	if err != nil {
		return nil, err
	}
	return &ContractsCandidateDeletedIterator{contract: _Contracts.contract, event: "CandidateDeleted", logs: logs, sub: sub}, nil
}

// WatchCandidateDeleted is a free log subscription operation binding the contract event 0x48bf3ef4944e20b74a7ae7804d664a8e9ba28141c5fc26163c87f7fbc9e90e2f.
//
// Solidity: event CandidateDeleted(uint256 projectId, uint256 candidateId)
func (_Contracts *ContractsFilterer) WatchCandidateDeleted(opts *bind.WatchOpts, sink chan<- *ContractsCandidateDeleted) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CandidateDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCandidateDeleted)
				if err := _Contracts.contract.UnpackLog(event, "CandidateDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCandidateDeleted is a log parse operation binding the contract event 0x48bf3ef4944e20b74a7ae7804d664a8e9ba28141c5fc26163c87f7fbc9e90e2f.
//
// Solidity: event CandidateDeleted(uint256 projectId, uint256 candidateId)
func (_Contracts *ContractsFilterer) ParseCandidateDeleted(log types.Log) (*ContractsCandidateDeleted, error) {
	event := new(ContractsCandidateDeleted)
	if err := _Contracts.contract.UnpackLog(event, "CandidateDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCandidateUpdatedIterator is returned from FilterCandidateUpdated and is used to iterate over the raw logs and unpacked data for CandidateUpdated events raised by the Contracts contract.
type ContractsCandidateUpdatedIterator struct {
	Event *ContractsCandidateUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsCandidateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCandidateUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsCandidateUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsCandidateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCandidateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCandidateUpdated represents a CandidateUpdated event raised by the Contracts contract.
type ContractsCandidateUpdated struct {
	ProjectId               *big.Int
	CandidateId             *big.Int
	NewName                 string
	NewImageUrl             string
	NewCandidateDescription string
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterCandidateUpdated is a free log retrieval operation binding the contract event 0x438b97487e2ce2ab75b5fc681e52fc9d2ed2a63de9588b6b970afb661dbdc6f0.
//
// Solidity: event CandidateUpdated(uint256 projectId, uint256 candidateId, string newName, string newImageUrl, string newCandidateDescription)
func (_Contracts *ContractsFilterer) FilterCandidateUpdated(opts *bind.FilterOpts) (*ContractsCandidateUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CandidateUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsCandidateUpdatedIterator{contract: _Contracts.contract, event: "CandidateUpdated", logs: logs, sub: sub}, nil
}

// WatchCandidateUpdated is a free log subscription operation binding the contract event 0x438b97487e2ce2ab75b5fc681e52fc9d2ed2a63de9588b6b970afb661dbdc6f0.
//
// Solidity: event CandidateUpdated(uint256 projectId, uint256 candidateId, string newName, string newImageUrl, string newCandidateDescription)
func (_Contracts *ContractsFilterer) WatchCandidateUpdated(opts *bind.WatchOpts, sink chan<- *ContractsCandidateUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CandidateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCandidateUpdated)
				if err := _Contracts.contract.UnpackLog(event, "CandidateUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCandidateUpdated is a log parse operation binding the contract event 0x438b97487e2ce2ab75b5fc681e52fc9d2ed2a63de9588b6b970afb661dbdc6f0.
//
// Solidity: event CandidateUpdated(uint256 projectId, uint256 candidateId, string newName, string newImageUrl, string newCandidateDescription)
func (_Contracts *ContractsFilterer) ParseCandidateUpdated(log types.Log) (*ContractsCandidateUpdated, error) {
	event := new(ContractsCandidateUpdated)
	if err := _Contracts.contract.UnpackLog(event, "CandidateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the Contracts contract.
type ContractsVoteCastIterator struct {
	Event *ContractsVoteCast // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsVoteCast)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsVoteCast)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsVoteCast represents a VoteCast event raised by the Contracts contract.
type ContractsVoteCast struct {
	Voter       common.Address
	ProjectId   *big.Int
	CandidateId *big.Int
	BlockHash   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0x07948c9515930cad4eac944a15a78a009b2a6db5bbe06a7560e7a7348825bb72.
//
// Solidity: event VoteCast(address indexed voter, uint256 projectId, uint256 candidateId, bytes32 blockHash)
func (_Contracts *ContractsFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*ContractsVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &ContractsVoteCastIterator{contract: _Contracts.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0x07948c9515930cad4eac944a15a78a009b2a6db5bbe06a7560e7a7348825bb72.
//
// Solidity: event VoteCast(address indexed voter, uint256 projectId, uint256 candidateId, bytes32 blockHash)
func (_Contracts *ContractsFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *ContractsVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsVoteCast)
				if err := _Contracts.contract.UnpackLog(event, "VoteCast", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCast is a log parse operation binding the contract event 0x07948c9515930cad4eac944a15a78a009b2a6db5bbe06a7560e7a7348825bb72.
//
// Solidity: event VoteCast(address indexed voter, uint256 projectId, uint256 candidateId, bytes32 blockHash)
func (_Contracts *ContractsFilterer) ParseVoteCast(log types.Log) (*ContractsVoteCast, error) {
	event := new(ContractsVoteCast)
	if err := _Contracts.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsVoterRegisteredIterator is returned from FilterVoterRegistered and is used to iterate over the raw logs and unpacked data for VoterRegistered events raised by the Contracts contract.
type ContractsVoterRegisteredIterator struct {
	Event *ContractsVoterRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsVoterRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsVoterRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsVoterRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsVoterRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsVoterRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsVoterRegistered represents a VoterRegistered event raised by the Contracts contract.
type ContractsVoterRegistered struct {
	Voter common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoterRegistered is a free log retrieval operation binding the contract event 0xb6be2187d059cc2a55fe29e0e503b566e1e0f8c8780096e185429350acffd3dd.
//
// Solidity: event VoterRegistered(address voter)
func (_Contracts *ContractsFilterer) FilterVoterRegistered(opts *bind.FilterOpts) (*ContractsVoterRegisteredIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "VoterRegistered")
	if err != nil {
		return nil, err
	}
	return &ContractsVoterRegisteredIterator{contract: _Contracts.contract, event: "VoterRegistered", logs: logs, sub: sub}, nil
}

// WatchVoterRegistered is a free log subscription operation binding the contract event 0xb6be2187d059cc2a55fe29e0e503b566e1e0f8c8780096e185429350acffd3dd.
//
// Solidity: event VoterRegistered(address voter)
func (_Contracts *ContractsFilterer) WatchVoterRegistered(opts *bind.WatchOpts, sink chan<- *ContractsVoterRegistered) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "VoterRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsVoterRegistered)
				if err := _Contracts.contract.UnpackLog(event, "VoterRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoterRegistered is a log parse operation binding the contract event 0xb6be2187d059cc2a55fe29e0e503b566e1e0f8c8780096e185429350acffd3dd.
//
// Solidity: event VoterRegistered(address voter)
func (_Contracts *ContractsFilterer) ParseVoterRegistered(log types.Log) (*ContractsVoterRegistered, error) {
	event := new(ContractsVoterRegistered)
	if err := _Contracts.contract.UnpackLog(event, "VoterRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsVotingProjectCreatedIterator is returned from FilterVotingProjectCreated and is used to iterate over the raw logs and unpacked data for VotingProjectCreated events raised by the Contracts contract.
type ContractsVotingProjectCreatedIterator struct {
	Event *ContractsVotingProjectCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsVotingProjectCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsVotingProjectCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsVotingProjectCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsVotingProjectCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsVotingProjectCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsVotingProjectCreated represents a VotingProjectCreated event raised by the Contracts contract.
type ContractsVotingProjectCreated struct {
	ProjectId   *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVotingProjectCreated is a free log retrieval operation binding the contract event 0x837f3d36cc844e134d5c0d833ddd4313bb3c42964532da67acb125af4c9ef02c.
//
// Solidity: event VotingProjectCreated(uint256 projectId, string description)
func (_Contracts *ContractsFilterer) FilterVotingProjectCreated(opts *bind.FilterOpts) (*ContractsVotingProjectCreatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "VotingProjectCreated")
	if err != nil {
		return nil, err
	}
	return &ContractsVotingProjectCreatedIterator{contract: _Contracts.contract, event: "VotingProjectCreated", logs: logs, sub: sub}, nil
}

// WatchVotingProjectCreated is a free log subscription operation binding the contract event 0x837f3d36cc844e134d5c0d833ddd4313bb3c42964532da67acb125af4c9ef02c.
//
// Solidity: event VotingProjectCreated(uint256 projectId, string description)
func (_Contracts *ContractsFilterer) WatchVotingProjectCreated(opts *bind.WatchOpts, sink chan<- *ContractsVotingProjectCreated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "VotingProjectCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsVotingProjectCreated)
				if err := _Contracts.contract.UnpackLog(event, "VotingProjectCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVotingProjectCreated is a log parse operation binding the contract event 0x837f3d36cc844e134d5c0d833ddd4313bb3c42964532da67acb125af4c9ef02c.
//
// Solidity: event VotingProjectCreated(uint256 projectId, string description)
func (_Contracts *ContractsFilterer) ParseVotingProjectCreated(log types.Log) (*ContractsVotingProjectCreated, error) {
	event := new(ContractsVotingProjectCreated)
	if err := _Contracts.contract.UnpackLog(event, "VotingProjectCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsVotingProjectEndedIterator is returned from FilterVotingProjectEnded and is used to iterate over the raw logs and unpacked data for VotingProjectEnded events raised by the Contracts contract.
type ContractsVotingProjectEndedIterator struct {
	Event *ContractsVotingProjectEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsVotingProjectEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsVotingProjectEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsVotingProjectEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsVotingProjectEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsVotingProjectEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsVotingProjectEnded represents a VotingProjectEnded event raised by the Contracts contract.
type ContractsVotingProjectEnded struct {
	ProjectId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVotingProjectEnded is a free log retrieval operation binding the contract event 0x1c32f86c7e5d094a801c70a05662bcdfd69b6ce9a3fff7b210b8a7b899eb63f5.
//
// Solidity: event VotingProjectEnded(uint256 projectId)
func (_Contracts *ContractsFilterer) FilterVotingProjectEnded(opts *bind.FilterOpts) (*ContractsVotingProjectEndedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "VotingProjectEnded")
	if err != nil {
		return nil, err
	}
	return &ContractsVotingProjectEndedIterator{contract: _Contracts.contract, event: "VotingProjectEnded", logs: logs, sub: sub}, nil
}

// WatchVotingProjectEnded is a free log subscription operation binding the contract event 0x1c32f86c7e5d094a801c70a05662bcdfd69b6ce9a3fff7b210b8a7b899eb63f5.
//
// Solidity: event VotingProjectEnded(uint256 projectId)
func (_Contracts *ContractsFilterer) WatchVotingProjectEnded(opts *bind.WatchOpts, sink chan<- *ContractsVotingProjectEnded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "VotingProjectEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsVotingProjectEnded)
				if err := _Contracts.contract.UnpackLog(event, "VotingProjectEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVotingProjectEnded is a log parse operation binding the contract event 0x1c32f86c7e5d094a801c70a05662bcdfd69b6ce9a3fff7b210b8a7b899eb63f5.
//
// Solidity: event VotingProjectEnded(uint256 projectId)
func (_Contracts *ContractsFilterer) ParseVotingProjectEnded(log types.Log) (*ContractsVotingProjectEnded, error) {
	event := new(ContractsVotingProjectEnded)
	if err := _Contracts.contract.UnpackLog(event, "VotingProjectEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsVotingProjectStartedIterator is returned from FilterVotingProjectStarted and is used to iterate over the raw logs and unpacked data for VotingProjectStarted events raised by the Contracts contract.
type ContractsVotingProjectStartedIterator struct {
	Event *ContractsVotingProjectStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsVotingProjectStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsVotingProjectStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsVotingProjectStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsVotingProjectStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsVotingProjectStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsVotingProjectStarted represents a VotingProjectStarted event raised by the Contracts contract.
type ContractsVotingProjectStarted struct {
	ProjectId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVotingProjectStarted is a free log retrieval operation binding the contract event 0x3a132177e16ff48392058b78d04b6898d7a776195658e94ab38c384562d73049.
//
// Solidity: event VotingProjectStarted(uint256 projectId)
func (_Contracts *ContractsFilterer) FilterVotingProjectStarted(opts *bind.FilterOpts) (*ContractsVotingProjectStartedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "VotingProjectStarted")
	if err != nil {
		return nil, err
	}
	return &ContractsVotingProjectStartedIterator{contract: _Contracts.contract, event: "VotingProjectStarted", logs: logs, sub: sub}, nil
}

// WatchVotingProjectStarted is a free log subscription operation binding the contract event 0x3a132177e16ff48392058b78d04b6898d7a776195658e94ab38c384562d73049.
//
// Solidity: event VotingProjectStarted(uint256 projectId)
func (_Contracts *ContractsFilterer) WatchVotingProjectStarted(opts *bind.WatchOpts, sink chan<- *ContractsVotingProjectStarted) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "VotingProjectStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsVotingProjectStarted)
				if err := _Contracts.contract.UnpackLog(event, "VotingProjectStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVotingProjectStarted is a log parse operation binding the contract event 0x3a132177e16ff48392058b78d04b6898d7a776195658e94ab38c384562d73049.
//
// Solidity: event VotingProjectStarted(uint256 projectId)
func (_Contracts *ContractsFilterer) ParseVotingProjectStarted(log types.Log) (*ContractsVotingProjectStarted, error) {
	event := new(ContractsVotingProjectStarted)
	if err := _Contracts.contract.UnpackLog(event, "VotingProjectStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
