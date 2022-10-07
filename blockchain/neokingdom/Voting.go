// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package neokingdom

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
)

// VotingMetaData contains all meta data concerning the Voting contract.
var VotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentDelegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingPower\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingPower\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"afterAddContributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"afterTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"beforeRemoveContributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"canVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"canVoteAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDelegate\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSnapshotId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"getDelegateAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalVotingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"getTotalVotingPowerAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getVotingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"getVotingPowerAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIShareholderRegistry\",\"name\":\"shareholderRegistry\",\"type\":\"address\"}],\"name\":\"setShareholderRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VotingABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingMetaData.ABI instead.
var VotingABI = VotingMetaData.ABI

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Voting *VotingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Voting *VotingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Voting.Contract.DEFAULTADMINROLE(&_Voting.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Voting *VotingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Voting.Contract.DEFAULTADMINROLE(&_Voting.CallOpts)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address account) view returns(bool)
func (_Voting *VotingCaller) CanVote(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "canVote", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address account) view returns(bool)
func (_Voting *VotingSession) CanVote(account common.Address) (bool, error) {
	return _Voting.Contract.CanVote(&_Voting.CallOpts, account)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote(address account) view returns(bool)
func (_Voting *VotingCallerSession) CanVote(account common.Address) (bool, error) {
	return _Voting.Contract.CanVote(&_Voting.CallOpts, account)
}

// CanVoteAt is a free data retrieval call binding the contract method 0xb652d57b.
//
// Solidity: function canVoteAt(address account, uint256 snapshotId) view returns(bool)
func (_Voting *VotingCaller) CanVoteAt(opts *bind.CallOpts, account common.Address, snapshotId *big.Int) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "canVoteAt", account, snapshotId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanVoteAt is a free data retrieval call binding the contract method 0xb652d57b.
//
// Solidity: function canVoteAt(address account, uint256 snapshotId) view returns(bool)
func (_Voting *VotingSession) CanVoteAt(account common.Address, snapshotId *big.Int) (bool, error) {
	return _Voting.Contract.CanVoteAt(&_Voting.CallOpts, account, snapshotId)
}

// CanVoteAt is a free data retrieval call binding the contract method 0xb652d57b.
//
// Solidity: function canVoteAt(address account, uint256 snapshotId) view returns(bool)
func (_Voting *VotingCallerSession) CanVoteAt(account common.Address, snapshotId *big.Int) (bool, error) {
	return _Voting.Contract.CanVoteAt(&_Voting.CallOpts, account, snapshotId)
}

// GetCurrentSnapshotId is a free data retrieval call binding the contract method 0x5439ad86.
//
// Solidity: function getCurrentSnapshotId() view returns(uint256)
func (_Voting *VotingCaller) GetCurrentSnapshotId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getCurrentSnapshotId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentSnapshotId is a free data retrieval call binding the contract method 0x5439ad86.
//
// Solidity: function getCurrentSnapshotId() view returns(uint256)
func (_Voting *VotingSession) GetCurrentSnapshotId() (*big.Int, error) {
	return _Voting.Contract.GetCurrentSnapshotId(&_Voting.CallOpts)
}

// GetCurrentSnapshotId is a free data retrieval call binding the contract method 0x5439ad86.
//
// Solidity: function getCurrentSnapshotId() view returns(uint256)
func (_Voting *VotingCallerSession) GetCurrentSnapshotId() (*big.Int, error) {
	return _Voting.Contract.GetCurrentSnapshotId(&_Voting.CallOpts)
}

// GetDelegate is a free data retrieval call binding the contract method 0x544d8564.
//
// Solidity: function getDelegate(address account) view returns(address)
func (_Voting *VotingCaller) GetDelegate(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getDelegate", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDelegate is a free data retrieval call binding the contract method 0x544d8564.
//
// Solidity: function getDelegate(address account) view returns(address)
func (_Voting *VotingSession) GetDelegate(account common.Address) (common.Address, error) {
	return _Voting.Contract.GetDelegate(&_Voting.CallOpts, account)
}

// GetDelegate is a free data retrieval call binding the contract method 0x544d8564.
//
// Solidity: function getDelegate(address account) view returns(address)
func (_Voting *VotingCallerSession) GetDelegate(account common.Address) (common.Address, error) {
	return _Voting.Contract.GetDelegate(&_Voting.CallOpts, account)
}

// GetDelegateAt is a free data retrieval call binding the contract method 0x1f992ff2.
//
// Solidity: function getDelegateAt(address account, uint256 snapshotId) view returns(address)
func (_Voting *VotingCaller) GetDelegateAt(opts *bind.CallOpts, account common.Address, snapshotId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getDelegateAt", account, snapshotId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDelegateAt is a free data retrieval call binding the contract method 0x1f992ff2.
//
// Solidity: function getDelegateAt(address account, uint256 snapshotId) view returns(address)
func (_Voting *VotingSession) GetDelegateAt(account common.Address, snapshotId *big.Int) (common.Address, error) {
	return _Voting.Contract.GetDelegateAt(&_Voting.CallOpts, account, snapshotId)
}

// GetDelegateAt is a free data retrieval call binding the contract method 0x1f992ff2.
//
// Solidity: function getDelegateAt(address account, uint256 snapshotId) view returns(address)
func (_Voting *VotingCallerSession) GetDelegateAt(account common.Address, snapshotId *big.Int) (common.Address, error) {
	return _Voting.Contract.GetDelegateAt(&_Voting.CallOpts, account, snapshotId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Voting *VotingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Voting *VotingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Voting.Contract.GetRoleAdmin(&_Voting.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Voting *VotingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Voting.Contract.GetRoleAdmin(&_Voting.CallOpts, role)
}

// GetTotalVotingPower is a free data retrieval call binding the contract method 0x11acc1a7.
//
// Solidity: function getTotalVotingPower() view returns(uint256)
func (_Voting *VotingCaller) GetTotalVotingPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getTotalVotingPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalVotingPower is a free data retrieval call binding the contract method 0x11acc1a7.
//
// Solidity: function getTotalVotingPower() view returns(uint256)
func (_Voting *VotingSession) GetTotalVotingPower() (*big.Int, error) {
	return _Voting.Contract.GetTotalVotingPower(&_Voting.CallOpts)
}

// GetTotalVotingPower is a free data retrieval call binding the contract method 0x11acc1a7.
//
// Solidity: function getTotalVotingPower() view returns(uint256)
func (_Voting *VotingCallerSession) GetTotalVotingPower() (*big.Int, error) {
	return _Voting.Contract.GetTotalVotingPower(&_Voting.CallOpts)
}

// GetTotalVotingPowerAt is a free data retrieval call binding the contract method 0x3fadf73b.
//
// Solidity: function getTotalVotingPowerAt(uint256 snapshotId) view returns(uint256)
func (_Voting *VotingCaller) GetTotalVotingPowerAt(opts *bind.CallOpts, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getTotalVotingPowerAt", snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalVotingPowerAt is a free data retrieval call binding the contract method 0x3fadf73b.
//
// Solidity: function getTotalVotingPowerAt(uint256 snapshotId) view returns(uint256)
func (_Voting *VotingSession) GetTotalVotingPowerAt(snapshotId *big.Int) (*big.Int, error) {
	return _Voting.Contract.GetTotalVotingPowerAt(&_Voting.CallOpts, snapshotId)
}

// GetTotalVotingPowerAt is a free data retrieval call binding the contract method 0x3fadf73b.
//
// Solidity: function getTotalVotingPowerAt(uint256 snapshotId) view returns(uint256)
func (_Voting *VotingCallerSession) GetTotalVotingPowerAt(snapshotId *big.Int) (*big.Int, error) {
	return _Voting.Contract.GetTotalVotingPowerAt(&_Voting.CallOpts, snapshotId)
}

// GetVotingPower is a free data retrieval call binding the contract method 0xbb4d4436.
//
// Solidity: function getVotingPower(address account) view returns(uint256)
func (_Voting *VotingCaller) GetVotingPower(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getVotingPower", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotingPower is a free data retrieval call binding the contract method 0xbb4d4436.
//
// Solidity: function getVotingPower(address account) view returns(uint256)
func (_Voting *VotingSession) GetVotingPower(account common.Address) (*big.Int, error) {
	return _Voting.Contract.GetVotingPower(&_Voting.CallOpts, account)
}

// GetVotingPower is a free data retrieval call binding the contract method 0xbb4d4436.
//
// Solidity: function getVotingPower(address account) view returns(uint256)
func (_Voting *VotingCallerSession) GetVotingPower(account common.Address) (*big.Int, error) {
	return _Voting.Contract.GetVotingPower(&_Voting.CallOpts, account)
}

// GetVotingPowerAt is a free data retrieval call binding the contract method 0xeaeded5f.
//
// Solidity: function getVotingPowerAt(address account, uint256 snapshotId) view returns(uint256)
func (_Voting *VotingCaller) GetVotingPowerAt(opts *bind.CallOpts, account common.Address, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getVotingPowerAt", account, snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotingPowerAt is a free data retrieval call binding the contract method 0xeaeded5f.
//
// Solidity: function getVotingPowerAt(address account, uint256 snapshotId) view returns(uint256)
func (_Voting *VotingSession) GetVotingPowerAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _Voting.Contract.GetVotingPowerAt(&_Voting.CallOpts, account, snapshotId)
}

// GetVotingPowerAt is a free data retrieval call binding the contract method 0xeaeded5f.
//
// Solidity: function getVotingPowerAt(address account, uint256 snapshotId) view returns(uint256)
func (_Voting *VotingCallerSession) GetVotingPowerAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _Voting.Contract.GetVotingPowerAt(&_Voting.CallOpts, account, snapshotId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Voting *VotingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Voting *VotingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Voting.Contract.HasRole(&_Voting.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Voting *VotingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Voting.Contract.HasRole(&_Voting.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Voting *VotingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Voting *VotingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Voting.Contract.SupportsInterface(&_Voting.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Voting *VotingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Voting.Contract.SupportsInterface(&_Voting.CallOpts, interfaceId)
}

// AfterAddContributor is a paid mutator transaction binding the contract method 0x4693feab.
//
// Solidity: function afterAddContributor(address account) returns()
func (_Voting *VotingTransactor) AfterAddContributor(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "afterAddContributor", account)
}

// AfterAddContributor is a paid mutator transaction binding the contract method 0x4693feab.
//
// Solidity: function afterAddContributor(address account) returns()
func (_Voting *VotingSession) AfterAddContributor(account common.Address) (*types.Transaction, error) {
	return _Voting.Contract.AfterAddContributor(&_Voting.TransactOpts, account)
}

// AfterAddContributor is a paid mutator transaction binding the contract method 0x4693feab.
//
// Solidity: function afterAddContributor(address account) returns()
func (_Voting *VotingTransactorSession) AfterAddContributor(account common.Address) (*types.Transaction, error) {
	return _Voting.Contract.AfterAddContributor(&_Voting.TransactOpts, account)
}

// AfterTokenTransfer is a paid mutator transaction binding the contract method 0xdf6258be.
//
// Solidity: function afterTokenTransfer(address from, address to, uint256 amount) returns()
func (_Voting *VotingTransactor) AfterTokenTransfer(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "afterTokenTransfer", from, to, amount)
}

// AfterTokenTransfer is a paid mutator transaction binding the contract method 0xdf6258be.
//
// Solidity: function afterTokenTransfer(address from, address to, uint256 amount) returns()
func (_Voting *VotingSession) AfterTokenTransfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.AfterTokenTransfer(&_Voting.TransactOpts, from, to, amount)
}

// AfterTokenTransfer is a paid mutator transaction binding the contract method 0xdf6258be.
//
// Solidity: function afterTokenTransfer(address from, address to, uint256 amount) returns()
func (_Voting *VotingTransactorSession) AfterTokenTransfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.AfterTokenTransfer(&_Voting.TransactOpts, from, to, amount)
}

// BeforeRemoveContributor is a paid mutator transaction binding the contract method 0xd86f5a07.
//
// Solidity: function beforeRemoveContributor(address account) returns()
func (_Voting *VotingTransactor) BeforeRemoveContributor(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "beforeRemoveContributor", account)
}

// BeforeRemoveContributor is a paid mutator transaction binding the contract method 0xd86f5a07.
//
// Solidity: function beforeRemoveContributor(address account) returns()
func (_Voting *VotingSession) BeforeRemoveContributor(account common.Address) (*types.Transaction, error) {
	return _Voting.Contract.BeforeRemoveContributor(&_Voting.TransactOpts, account)
}

// BeforeRemoveContributor is a paid mutator transaction binding the contract method 0xd86f5a07.
//
// Solidity: function beforeRemoveContributor(address account) returns()
func (_Voting *VotingTransactorSession) BeforeRemoveContributor(account common.Address) (*types.Transaction, error) {
	return _Voting.Contract.BeforeRemoveContributor(&_Voting.TransactOpts, account)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address newDelegate) returns()
func (_Voting *VotingTransactor) Delegate(opts *bind.TransactOpts, newDelegate common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "delegate", newDelegate)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address newDelegate) returns()
func (_Voting *VotingSession) Delegate(newDelegate common.Address) (*types.Transaction, error) {
	return _Voting.Contract.Delegate(&_Voting.TransactOpts, newDelegate)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address newDelegate) returns()
func (_Voting *VotingTransactorSession) Delegate(newDelegate common.Address) (*types.Transaction, error) {
	return _Voting.Contract.Delegate(&_Voting.TransactOpts, newDelegate)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Voting *VotingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Voting *VotingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Voting.Contract.GrantRole(&_Voting.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Voting *VotingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Voting.Contract.GrantRole(&_Voting.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Voting *VotingTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Voting *VotingSession) Initialize() (*types.Transaction, error) {
	return _Voting.Contract.Initialize(&_Voting.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Voting *VotingTransactorSession) Initialize() (*types.Transaction, error) {
	return _Voting.Contract.Initialize(&_Voting.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Voting *VotingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Voting *VotingSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Voting.Contract.RenounceRole(&_Voting.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Voting *VotingTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Voting.Contract.RenounceRole(&_Voting.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Voting *VotingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Voting *VotingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Voting.Contract.RevokeRole(&_Voting.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Voting *VotingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Voting.Contract.RevokeRole(&_Voting.TransactOpts, role, account)
}

// SetShareholderRegistry is a paid mutator transaction binding the contract method 0xa260b102.
//
// Solidity: function setShareholderRegistry(address shareholderRegistry) returns()
func (_Voting *VotingTransactor) SetShareholderRegistry(opts *bind.TransactOpts, shareholderRegistry common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "setShareholderRegistry", shareholderRegistry)
}

// SetShareholderRegistry is a paid mutator transaction binding the contract method 0xa260b102.
//
// Solidity: function setShareholderRegistry(address shareholderRegistry) returns()
func (_Voting *VotingSession) SetShareholderRegistry(shareholderRegistry common.Address) (*types.Transaction, error) {
	return _Voting.Contract.SetShareholderRegistry(&_Voting.TransactOpts, shareholderRegistry)
}

// SetShareholderRegistry is a paid mutator transaction binding the contract method 0xa260b102.
//
// Solidity: function setShareholderRegistry(address shareholderRegistry) returns()
func (_Voting *VotingTransactorSession) SetShareholderRegistry(shareholderRegistry common.Address) (*types.Transaction, error) {
	return _Voting.Contract.SetShareholderRegistry(&_Voting.TransactOpts, shareholderRegistry)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address token) returns()
func (_Voting *VotingTransactor) SetToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "setToken", token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address token) returns()
func (_Voting *VotingSession) SetToken(token common.Address) (*types.Transaction, error) {
	return _Voting.Contract.SetToken(&_Voting.TransactOpts, token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address token) returns()
func (_Voting *VotingTransactorSession) SetToken(token common.Address) (*types.Transaction, error) {
	return _Voting.Contract.SetToken(&_Voting.TransactOpts, token)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256)
func (_Voting *VotingTransactor) Snapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "snapshot")
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256)
func (_Voting *VotingSession) Snapshot() (*types.Transaction, error) {
	return _Voting.Contract.Snapshot(&_Voting.TransactOpts)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256)
func (_Voting *VotingTransactorSession) Snapshot() (*types.Transaction, error) {
	return _Voting.Contract.Snapshot(&_Voting.TransactOpts)
}

// VotingDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the Voting contract.
type VotingDelegateChangedIterator struct {
	Event *VotingDelegateChanged // Event containing the contract specifics and raw log

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
func (it *VotingDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingDelegateChanged)
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
		it.Event = new(VotingDelegateChanged)
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
func (it *VotingDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingDelegateChanged represents a DelegateChanged event raised by the Voting contract.
type VotingDelegateChanged struct {
	Delegator       common.Address
	CurrentDelegate common.Address
	NewDelegate     common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address currentDelegate, address newDelegate)
func (_Voting *VotingFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address) (*VotingDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "DelegateChanged", delegatorRule)
	if err != nil {
		return nil, err
	}
	return &VotingDelegateChangedIterator{contract: _Voting.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address currentDelegate, address newDelegate)
func (_Voting *VotingFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *VotingDelegateChanged, delegator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "DelegateChanged", delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingDelegateChanged)
				if err := _Voting.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address currentDelegate, address newDelegate)
func (_Voting *VotingFilterer) ParseDelegateChanged(log types.Log) (*VotingDelegateChanged, error) {
	event := new(VotingDelegateChanged)
	if err := _Voting.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the Voting contract.
type VotingDelegateVotesChangedIterator struct {
	Event *VotingDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *VotingDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingDelegateVotesChanged)
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
		it.Event = new(VotingDelegateVotesChanged)
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
func (it *VotingDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingDelegateVotesChanged represents a DelegateVotesChanged event raised by the Voting contract.
type VotingDelegateVotesChanged struct {
	Account        common.Address
	OldVotingPower *big.Int
	NewVotingPower *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed account, uint256 oldVotingPower, uint256 newVotingPower)
func (_Voting *VotingFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, account []common.Address) (*VotingDelegateVotesChangedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "DelegateVotesChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return &VotingDelegateVotesChangedIterator{contract: _Voting.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed account, uint256 oldVotingPower, uint256 newVotingPower)
func (_Voting *VotingFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *VotingDelegateVotesChanged, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "DelegateVotesChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingDelegateVotesChanged)
				if err := _Voting.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed account, uint256 oldVotingPower, uint256 newVotingPower)
func (_Voting *VotingFilterer) ParseDelegateVotesChanged(log types.Log) (*VotingDelegateVotesChanged, error) {
	event := new(VotingDelegateVotesChanged)
	if err := _Voting.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Voting contract.
type VotingInitializedIterator struct {
	Event *VotingInitialized // Event containing the contract specifics and raw log

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
func (it *VotingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingInitialized)
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
		it.Event = new(VotingInitialized)
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
func (it *VotingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingInitialized represents a Initialized event raised by the Voting contract.
type VotingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voting *VotingFilterer) FilterInitialized(opts *bind.FilterOpts) (*VotingInitializedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VotingInitializedIterator{contract: _Voting.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voting *VotingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VotingInitialized) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingInitialized)
				if err := _Voting.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voting *VotingFilterer) ParseInitialized(log types.Log) (*VotingInitialized, error) {
	event := new(VotingInitialized)
	if err := _Voting.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Voting contract.
type VotingRoleAdminChangedIterator struct {
	Event *VotingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VotingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingRoleAdminChanged)
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
		it.Event = new(VotingRoleAdminChanged)
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
func (it *VotingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingRoleAdminChanged represents a RoleAdminChanged event raised by the Voting contract.
type VotingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Voting *VotingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VotingRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VotingRoleAdminChangedIterator{contract: _Voting.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Voting *VotingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VotingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingRoleAdminChanged)
				if err := _Voting.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Voting *VotingFilterer) ParseRoleAdminChanged(log types.Log) (*VotingRoleAdminChanged, error) {
	event := new(VotingRoleAdminChanged)
	if err := _Voting.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Voting contract.
type VotingRoleGrantedIterator struct {
	Event *VotingRoleGranted // Event containing the contract specifics and raw log

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
func (it *VotingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingRoleGranted)
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
		it.Event = new(VotingRoleGranted)
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
func (it *VotingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingRoleGranted represents a RoleGranted event raised by the Voting contract.
type VotingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Voting *VotingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VotingRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VotingRoleGrantedIterator{contract: _Voting.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Voting *VotingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VotingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingRoleGranted)
				if err := _Voting.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Voting *VotingFilterer) ParseRoleGranted(log types.Log) (*VotingRoleGranted, error) {
	event := new(VotingRoleGranted)
	if err := _Voting.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Voting contract.
type VotingRoleRevokedIterator struct {
	Event *VotingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VotingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingRoleRevoked)
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
		it.Event = new(VotingRoleRevoked)
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
func (it *VotingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingRoleRevoked represents a RoleRevoked event raised by the Voting contract.
type VotingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Voting *VotingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VotingRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VotingRoleRevokedIterator{contract: _Voting.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Voting *VotingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VotingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingRoleRevoked)
				if err := _Voting.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Voting *VotingFilterer) ParseRoleRevoked(log types.Log) (*VotingRoleRevoked, error) {
	event := new(VotingRoleRevoked)
	if err := _Voting.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingSnapshotIterator is returned from FilterSnapshot and is used to iterate over the raw logs and unpacked data for Snapshot events raised by the Voting contract.
type VotingSnapshotIterator struct {
	Event *VotingSnapshot // Event containing the contract specifics and raw log

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
func (it *VotingSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingSnapshot)
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
		it.Event = new(VotingSnapshot)
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
func (it *VotingSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingSnapshot represents a Snapshot event raised by the Voting contract.
type VotingSnapshot struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSnapshot is a free log retrieval operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_Voting *VotingFilterer) FilterSnapshot(opts *bind.FilterOpts) (*VotingSnapshotIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return &VotingSnapshotIterator{contract: _Voting.contract, event: "Snapshot", logs: logs, sub: sub}, nil
}

// WatchSnapshot is a free log subscription operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_Voting *VotingFilterer) WatchSnapshot(opts *bind.WatchOpts, sink chan<- *VotingSnapshot) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingSnapshot)
				if err := _Voting.contract.UnpackLog(event, "Snapshot", log); err != nil {
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

// ParseSnapshot is a log parse operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_Voting *VotingFilterer) ParseSnapshot(log types.Log) (*VotingSnapshot, error) {
	event := new(VotingSnapshot)
	if err := _Voting.contract.UnpackLog(event, "Snapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
