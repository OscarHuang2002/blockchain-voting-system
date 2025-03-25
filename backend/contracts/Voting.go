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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CandidateAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VoterRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"VotingEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"VotingStarted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"addCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCandidateCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotingStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"registerVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_candidateId\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Contracts *ContractsCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "candidates", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Name      string
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Contracts *ContractsSession) Candidates(arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	return _Contracts.Contract.Candidates(&_Contracts.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Contracts *ContractsCallerSession) Candidates(arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	return _Contracts.Contract.Candidates(&_Contracts.CallOpts, arg0)
}

// GetCandidateCount is a free data retrieval call binding the contract method 0x30a56347.
//
// Solidity: function getCandidateCount() view returns(uint256)
func (_Contracts *ContractsCaller) GetCandidateCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCandidateCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCandidateCount is a free data retrieval call binding the contract method 0x30a56347.
//
// Solidity: function getCandidateCount() view returns(uint256)
func (_Contracts *ContractsSession) GetCandidateCount() (*big.Int, error) {
	return _Contracts.Contract.GetCandidateCount(&_Contracts.CallOpts)
}

// GetCandidateCount is a free data retrieval call binding the contract method 0x30a56347.
//
// Solidity: function getCandidateCount() view returns(uint256)
func (_Contracts *ContractsCallerSession) GetCandidateCount() (*big.Int, error) {
	return _Contracts.Contract.GetCandidateCount(&_Contracts.CallOpts)
}

// GetVotingStatus is a free data retrieval call binding the contract method 0x581c281c.
//
// Solidity: function getVotingStatus() view returns(bool)
func (_Contracts *ContractsCaller) GetVotingStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getVotingStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetVotingStatus is a free data retrieval call binding the contract method 0x581c281c.
//
// Solidity: function getVotingStatus() view returns(bool)
func (_Contracts *ContractsSession) GetVotingStatus() (bool, error) {
	return _Contracts.Contract.GetVotingStatus(&_Contracts.CallOpts)
}

// GetVotingStatus is a free data retrieval call binding the contract method 0x581c281c.
//
// Solidity: function getVotingStatus() view returns(bool)
func (_Contracts *ContractsCallerSession) GetVotingStatus() (bool, error) {
	return _Contracts.Contract.GetVotingStatus(&_Contracts.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _user) view returns(bool)
func (_Contracts *ContractsCaller) IsAdmin(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isAdmin", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _user) view returns(bool)
func (_Contracts *ContractsSession) IsAdmin(_user common.Address) (bool, error) {
	return _Contracts.Contract.IsAdmin(&_Contracts.CallOpts, _user)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _user) view returns(bool)
func (_Contracts *ContractsCallerSession) IsAdmin(_user common.Address) (bool, error) {
	return _Contracts.Contract.IsAdmin(&_Contracts.CallOpts, _user)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool isRegistered, bool hasVoted)
func (_Contracts *ContractsCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsRegistered bool
	HasVoted     bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		IsRegistered bool
		HasVoted     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRegistered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.HasVoted = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool isRegistered, bool hasVoted)
func (_Contracts *ContractsSession) Voters(arg0 common.Address) (struct {
	IsRegistered bool
	HasVoted     bool
}, error) {
	return _Contracts.Contract.Voters(&_Contracts.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool isRegistered, bool hasVoted)
func (_Contracts *ContractsCallerSession) Voters(arg0 common.Address) (struct {
	IsRegistered bool
	HasVoted     bool
}, error) {
	return _Contracts.Contract.Voters(&_Contracts.CallOpts, arg0)
}

// VotingOpen is a free data retrieval call binding the contract method 0xa95824b4.
//
// Solidity: function votingOpen() view returns(bool)
func (_Contracts *ContractsCaller) VotingOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "votingOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VotingOpen is a free data retrieval call binding the contract method 0xa95824b4.
//
// Solidity: function votingOpen() view returns(bool)
func (_Contracts *ContractsSession) VotingOpen() (bool, error) {
	return _Contracts.Contract.VotingOpen(&_Contracts.CallOpts)
}

// VotingOpen is a free data retrieval call binding the contract method 0xa95824b4.
//
// Solidity: function votingOpen() view returns(bool)
func (_Contracts *ContractsCallerSession) VotingOpen() (bool, error) {
	return _Contracts.Contract.VotingOpen(&_Contracts.CallOpts)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Contracts *ContractsTransactor) AddCandidate(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addCandidate", _name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Contracts *ContractsSession) AddCandidate(_name string) (*types.Transaction, error) {
	return _Contracts.Contract.AddCandidate(&_Contracts.TransactOpts, _name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Contracts *ContractsTransactorSession) AddCandidate(_name string) (*types.Transaction, error) {
	return _Contracts.Contract.AddCandidate(&_Contracts.TransactOpts, _name)
}

// EndVoting is a paid mutator transaction binding the contract method 0xc3403ddf.
//
// Solidity: function endVoting() returns()
func (_Contracts *ContractsTransactor) EndVoting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "endVoting")
}

// EndVoting is a paid mutator transaction binding the contract method 0xc3403ddf.
//
// Solidity: function endVoting() returns()
func (_Contracts *ContractsSession) EndVoting() (*types.Transaction, error) {
	return _Contracts.Contract.EndVoting(&_Contracts.TransactOpts)
}

// EndVoting is a paid mutator transaction binding the contract method 0xc3403ddf.
//
// Solidity: function endVoting() returns()
func (_Contracts *ContractsTransactorSession) EndVoting() (*types.Transaction, error) {
	return _Contracts.Contract.EndVoting(&_Contracts.TransactOpts)
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

// StartVoting is a paid mutator transaction binding the contract method 0x1ec6b60a.
//
// Solidity: function startVoting() returns()
func (_Contracts *ContractsTransactor) StartVoting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "startVoting")
}

// StartVoting is a paid mutator transaction binding the contract method 0x1ec6b60a.
//
// Solidity: function startVoting() returns()
func (_Contracts *ContractsSession) StartVoting() (*types.Transaction, error) {
	return _Contracts.Contract.StartVoting(&_Contracts.TransactOpts)
}

// StartVoting is a paid mutator transaction binding the contract method 0x1ec6b60a.
//
// Solidity: function startVoting() returns()
func (_Contracts *ContractsTransactorSession) StartVoting() (*types.Transaction, error) {
	return _Contracts.Contract.StartVoting(&_Contracts.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateId) returns()
func (_Contracts *ContractsTransactor) Vote(opts *bind.TransactOpts, _candidateId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "vote", _candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateId) returns()
func (_Contracts *ContractsSession) Vote(_candidateId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Vote(&_Contracts.TransactOpts, _candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateId) returns()
func (_Contracts *ContractsTransactorSession) Vote(_candidateId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Vote(&_Contracts.TransactOpts, _candidateId)
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
	CandidateId *big.Int
	Name        string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCandidateAdded is a free log retrieval operation binding the contract event 0xe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a3.
//
// Solidity: event CandidateAdded(uint256 candidateId, string name)
func (_Contracts *ContractsFilterer) FilterCandidateAdded(opts *bind.FilterOpts) (*ContractsCandidateAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CandidateAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsCandidateAddedIterator{contract: _Contracts.contract, event: "CandidateAdded", logs: logs, sub: sub}, nil
}

// WatchCandidateAdded is a free log subscription operation binding the contract event 0xe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a3.
//
// Solidity: event CandidateAdded(uint256 candidateId, string name)
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

// ParseCandidateAdded is a log parse operation binding the contract event 0xe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a3.
//
// Solidity: event CandidateAdded(uint256 candidateId, string name)
func (_Contracts *ContractsFilterer) ParseCandidateAdded(log types.Log) (*ContractsCandidateAdded, error) {
	event := new(ContractsCandidateAdded)
	if err := _Contracts.contract.UnpackLog(event, "CandidateAdded", log); err != nil {
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
	CandidateId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b.
//
// Solidity: event VoteCast(address indexed voter, uint256 candidateId)
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

// WatchVoteCast is a free log subscription operation binding the contract event 0xa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b.
//
// Solidity: event VoteCast(address indexed voter, uint256 candidateId)
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

// ParseVoteCast is a log parse operation binding the contract event 0xa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b.
//
// Solidity: event VoteCast(address indexed voter, uint256 candidateId)
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

// ContractsVotingEndedIterator is returned from FilterVotingEnded and is used to iterate over the raw logs and unpacked data for VotingEnded events raised by the Contracts contract.
type ContractsVotingEndedIterator struct {
	Event *ContractsVotingEnded // Event containing the contract specifics and raw log

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
func (it *ContractsVotingEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsVotingEnded)
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
		it.Event = new(ContractsVotingEnded)
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
func (it *ContractsVotingEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsVotingEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsVotingEnded represents a VotingEnded event raised by the Contracts contract.
type ContractsVotingEnded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVotingEnded is a free log retrieval operation binding the contract event 0x7a19ed057db79e3c2fa0b97a54b43bef4fce74b31bb6c01af514b9a18a7f70ab.
//
// Solidity: event VotingEnded()
func (_Contracts *ContractsFilterer) FilterVotingEnded(opts *bind.FilterOpts) (*ContractsVotingEndedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "VotingEnded")
	if err != nil {
		return nil, err
	}
	return &ContractsVotingEndedIterator{contract: _Contracts.contract, event: "VotingEnded", logs: logs, sub: sub}, nil
}

// WatchVotingEnded is a free log subscription operation binding the contract event 0x7a19ed057db79e3c2fa0b97a54b43bef4fce74b31bb6c01af514b9a18a7f70ab.
//
// Solidity: event VotingEnded()
func (_Contracts *ContractsFilterer) WatchVotingEnded(opts *bind.WatchOpts, sink chan<- *ContractsVotingEnded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "VotingEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsVotingEnded)
				if err := _Contracts.contract.UnpackLog(event, "VotingEnded", log); err != nil {
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

// ParseVotingEnded is a log parse operation binding the contract event 0x7a19ed057db79e3c2fa0b97a54b43bef4fce74b31bb6c01af514b9a18a7f70ab.
//
// Solidity: event VotingEnded()
func (_Contracts *ContractsFilterer) ParseVotingEnded(log types.Log) (*ContractsVotingEnded, error) {
	event := new(ContractsVotingEnded)
	if err := _Contracts.contract.UnpackLog(event, "VotingEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsVotingStartedIterator is returned from FilterVotingStarted and is used to iterate over the raw logs and unpacked data for VotingStarted events raised by the Contracts contract.
type ContractsVotingStartedIterator struct {
	Event *ContractsVotingStarted // Event containing the contract specifics and raw log

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
func (it *ContractsVotingStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsVotingStarted)
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
		it.Event = new(ContractsVotingStarted)
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
func (it *ContractsVotingStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsVotingStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsVotingStarted represents a VotingStarted event raised by the Contracts contract.
type ContractsVotingStarted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVotingStarted is a free log retrieval operation binding the contract event 0x877e2548498f42b7975a186b94ef1d32c86d420b7b806dd2be2bea293b895904.
//
// Solidity: event VotingStarted()
func (_Contracts *ContractsFilterer) FilterVotingStarted(opts *bind.FilterOpts) (*ContractsVotingStartedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "VotingStarted")
	if err != nil {
		return nil, err
	}
	return &ContractsVotingStartedIterator{contract: _Contracts.contract, event: "VotingStarted", logs: logs, sub: sub}, nil
}

// WatchVotingStarted is a free log subscription operation binding the contract event 0x877e2548498f42b7975a186b94ef1d32c86d420b7b806dd2be2bea293b895904.
//
// Solidity: event VotingStarted()
func (_Contracts *ContractsFilterer) WatchVotingStarted(opts *bind.WatchOpts, sink chan<- *ContractsVotingStarted) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "VotingStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsVotingStarted)
				if err := _Contracts.contract.UnpackLog(event, "VotingStarted", log); err != nil {
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

// ParseVotingStarted is a log parse operation binding the contract event 0x877e2548498f42b7975a186b94ef1d32c86d420b7b806dd2be2bea293b895904.
//
// Solidity: event VotingStarted()
func (_Contracts *ContractsFilterer) ParseVotingStarted(log types.Log) (*ContractsVotingStarted, error) {
	event := new(ContractsVotingStarted)
	if err := _Contracts.contract.UnpackLog(event, "VotingStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
