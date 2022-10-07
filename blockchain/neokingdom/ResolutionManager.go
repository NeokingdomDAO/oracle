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

// ResolutionManagerMetaData contains all meta data concerning the ResolutionManager contract.
var ResolutionManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegateLostVotingPower\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"}],\"name\":\"ResolutionApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"}],\"name\":\"ResolutionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"}],\"name\":\"ResolutionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"}],\"name\":\"ResolutionRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"typeIndex\",\"type\":\"uint256\"}],\"name\":\"ResolutionTypeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"}],\"name\":\"ResolutionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isYes\",\"type\":\"bool\"}],\"name\":\"ResolutionVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noticePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canBeNegative\",\"type\":\"bool\"}],\"name\":\"addResolutionType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"}],\"name\":\"approveResolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dataURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"resolutionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNegative\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"executionTo\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"executionData\",\"type\":\"bytes[]\"}],\"name\":\"createResolution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"}],\"name\":\"executeResolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"}],\"name\":\"getExecutionDetails\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"}],\"name\":\"getResolutionResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getVoterVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isYes\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIShareholderRegistry\",\"name\":\"shareholderRegistry\",\"type\":\"address\"},{\"internalType\":\"contractITelediskoToken\",\"name\":\"telediskoToken\",\"type\":\"address\"},{\"internalType\":\"contractIVoting\",\"name\":\"voting\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"}],\"name\":\"rejectResolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"resolutionTypes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noticePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canBeNegative\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"resolutions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"dataURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"resolutionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"approveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"yesVotesTotal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNegative\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"rejectionTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIShareholderRegistry\",\"name\":\"shareholderRegistry\",\"type\":\"address\"}],\"name\":\"setShareholderRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITelediskoToken\",\"name\":\"telediskoToken\",\"type\":\"address\"}],\"name\":\"setTelediskoToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVoting\",\"name\":\"voting\",\"type\":\"address\"}],\"name\":\"setVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dataURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"resolutionTypeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNegative\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"executionTo\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"executionData\",\"type\":\"bytes[]\"}],\"name\":\"updateResolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resolutionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isYes\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ResolutionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ResolutionManagerMetaData.ABI instead.
var ResolutionManagerABI = ResolutionManagerMetaData.ABI

// ResolutionManager is an auto generated Go binding around an Ethereum contract.
type ResolutionManager struct {
	ResolutionManagerCaller     // Read-only binding to the contract
	ResolutionManagerTransactor // Write-only binding to the contract
	ResolutionManagerFilterer   // Log filterer for contract events
}

// ResolutionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolutionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolutionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolutionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolutionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolutionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolutionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolutionManagerSession struct {
	Contract     *ResolutionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ResolutionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolutionManagerCallerSession struct {
	Contract *ResolutionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ResolutionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolutionManagerTransactorSession struct {
	Contract     *ResolutionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ResolutionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolutionManagerRaw struct {
	Contract *ResolutionManager // Generic contract binding to access the raw methods on
}

// ResolutionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolutionManagerCallerRaw struct {
	Contract *ResolutionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ResolutionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolutionManagerTransactorRaw struct {
	Contract *ResolutionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolutionManager creates a new instance of ResolutionManager, bound to a specific deployed contract.
func NewResolutionManager(address common.Address, backend bind.ContractBackend) (*ResolutionManager, error) {
	contract, err := bindResolutionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResolutionManager{ResolutionManagerCaller: ResolutionManagerCaller{contract: contract}, ResolutionManagerTransactor: ResolutionManagerTransactor{contract: contract}, ResolutionManagerFilterer: ResolutionManagerFilterer{contract: contract}}, nil
}

// NewResolutionManagerCaller creates a new read-only instance of ResolutionManager, bound to a specific deployed contract.
func NewResolutionManagerCaller(address common.Address, caller bind.ContractCaller) (*ResolutionManagerCaller, error) {
	contract, err := bindResolutionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerCaller{contract: contract}, nil
}

// NewResolutionManagerTransactor creates a new write-only instance of ResolutionManager, bound to a specific deployed contract.
func NewResolutionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolutionManagerTransactor, error) {
	contract, err := bindResolutionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerTransactor{contract: contract}, nil
}

// NewResolutionManagerFilterer creates a new log filterer instance of ResolutionManager, bound to a specific deployed contract.
func NewResolutionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ResolutionManagerFilterer, error) {
	contract, err := bindResolutionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerFilterer{contract: contract}, nil
}

// bindResolutionManager binds a generic wrapper to an already deployed contract.
func bindResolutionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolutionManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolutionManager *ResolutionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolutionManager.Contract.ResolutionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolutionManager *ResolutionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolutionManager.Contract.ResolutionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolutionManager *ResolutionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolutionManager.Contract.ResolutionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolutionManager *ResolutionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolutionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolutionManager *ResolutionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolutionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolutionManager *ResolutionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolutionManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ResolutionManager *ResolutionManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ResolutionManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ResolutionManager *ResolutionManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ResolutionManager.Contract.DEFAULTADMINROLE(&_ResolutionManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ResolutionManager *ResolutionManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ResolutionManager.Contract.DEFAULTADMINROLE(&_ResolutionManager.CallOpts)
}

// GetExecutionDetails is a free data retrieval call binding the contract method 0x245b9a32.
//
// Solidity: function getExecutionDetails(uint256 resolutionId) view returns(address[], bytes[])
func (_ResolutionManager *ResolutionManagerCaller) GetExecutionDetails(opts *bind.CallOpts, resolutionId *big.Int) ([]common.Address, [][]byte, error) {
	var out []interface{}
	err := _ResolutionManager.contract.Call(opts, &out, "getExecutionDetails", resolutionId)

	if err != nil {
		return *new([]common.Address), *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)

	return out0, out1, err

}

// GetExecutionDetails is a free data retrieval call binding the contract method 0x245b9a32.
//
// Solidity: function getExecutionDetails(uint256 resolutionId) view returns(address[], bytes[])
func (_ResolutionManager *ResolutionManagerSession) GetExecutionDetails(resolutionId *big.Int) ([]common.Address, [][]byte, error) {
	return _ResolutionManager.Contract.GetExecutionDetails(&_ResolutionManager.CallOpts, resolutionId)
}

// GetExecutionDetails is a free data retrieval call binding the contract method 0x245b9a32.
//
// Solidity: function getExecutionDetails(uint256 resolutionId) view returns(address[], bytes[])
func (_ResolutionManager *ResolutionManagerCallerSession) GetExecutionDetails(resolutionId *big.Int) ([]common.Address, [][]byte, error) {
	return _ResolutionManager.Contract.GetExecutionDetails(&_ResolutionManager.CallOpts, resolutionId)
}

// GetResolutionResult is a free data retrieval call binding the contract method 0x607bbbf4.
//
// Solidity: function getResolutionResult(uint256 resolutionId) view returns(bool)
func (_ResolutionManager *ResolutionManagerCaller) GetResolutionResult(opts *bind.CallOpts, resolutionId *big.Int) (bool, error) {
	var out []interface{}
	err := _ResolutionManager.contract.Call(opts, &out, "getResolutionResult", resolutionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetResolutionResult is a free data retrieval call binding the contract method 0x607bbbf4.
//
// Solidity: function getResolutionResult(uint256 resolutionId) view returns(bool)
func (_ResolutionManager *ResolutionManagerSession) GetResolutionResult(resolutionId *big.Int) (bool, error) {
	return _ResolutionManager.Contract.GetResolutionResult(&_ResolutionManager.CallOpts, resolutionId)
}

// GetResolutionResult is a free data retrieval call binding the contract method 0x607bbbf4.
//
// Solidity: function getResolutionResult(uint256 resolutionId) view returns(bool)
func (_ResolutionManager *ResolutionManagerCallerSession) GetResolutionResult(resolutionId *big.Int) (bool, error) {
	return _ResolutionManager.Contract.GetResolutionResult(&_ResolutionManager.CallOpts, resolutionId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ResolutionManager *ResolutionManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ResolutionManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ResolutionManager *ResolutionManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ResolutionManager.Contract.GetRoleAdmin(&_ResolutionManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ResolutionManager *ResolutionManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ResolutionManager.Contract.GetRoleAdmin(&_ResolutionManager.CallOpts, role)
}

// GetVoterVote is a free data retrieval call binding the contract method 0x3ff973d9.
//
// Solidity: function getVoterVote(uint256 resolutionId, address voter) view returns(bool isYes, bool hasVoted, uint256 votingPower)
func (_ResolutionManager *ResolutionManagerCaller) GetVoterVote(opts *bind.CallOpts, resolutionId *big.Int, voter common.Address) (struct {
	IsYes       bool
	HasVoted    bool
	VotingPower *big.Int
}, error) {
	var out []interface{}
	err := _ResolutionManager.contract.Call(opts, &out, "getVoterVote", resolutionId, voter)

	outstruct := new(struct {
		IsYes       bool
		HasVoted    bool
		VotingPower *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsYes = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.HasVoted = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.VotingPower = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetVoterVote is a free data retrieval call binding the contract method 0x3ff973d9.
//
// Solidity: function getVoterVote(uint256 resolutionId, address voter) view returns(bool isYes, bool hasVoted, uint256 votingPower)
func (_ResolutionManager *ResolutionManagerSession) GetVoterVote(resolutionId *big.Int, voter common.Address) (struct {
	IsYes       bool
	HasVoted    bool
	VotingPower *big.Int
}, error) {
	return _ResolutionManager.Contract.GetVoterVote(&_ResolutionManager.CallOpts, resolutionId, voter)
}

// GetVoterVote is a free data retrieval call binding the contract method 0x3ff973d9.
//
// Solidity: function getVoterVote(uint256 resolutionId, address voter) view returns(bool isYes, bool hasVoted, uint256 votingPower)
func (_ResolutionManager *ResolutionManagerCallerSession) GetVoterVote(resolutionId *big.Int, voter common.Address) (struct {
	IsYes       bool
	HasVoted    bool
	VotingPower *big.Int
}, error) {
	return _ResolutionManager.Contract.GetVoterVote(&_ResolutionManager.CallOpts, resolutionId, voter)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ResolutionManager *ResolutionManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ResolutionManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ResolutionManager *ResolutionManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ResolutionManager.Contract.HasRole(&_ResolutionManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ResolutionManager *ResolutionManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ResolutionManager.Contract.HasRole(&_ResolutionManager.CallOpts, role, account)
}

// ResolutionTypes is a free data retrieval call binding the contract method 0x96648206.
//
// Solidity: function resolutionTypes(uint256 ) view returns(string name, uint256 quorum, uint256 noticePeriod, uint256 votingPeriod, bool canBeNegative)
func (_ResolutionManager *ResolutionManagerCaller) ResolutionTypes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name          string
	Quorum        *big.Int
	NoticePeriod  *big.Int
	VotingPeriod  *big.Int
	CanBeNegative bool
}, error) {
	var out []interface{}
	err := _ResolutionManager.contract.Call(opts, &out, "resolutionTypes", arg0)

	outstruct := new(struct {
		Name          string
		Quorum        *big.Int
		NoticePeriod  *big.Int
		VotingPeriod  *big.Int
		CanBeNegative bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Quorum = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NoticePeriod = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.VotingPeriod = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CanBeNegative = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// ResolutionTypes is a free data retrieval call binding the contract method 0x96648206.
//
// Solidity: function resolutionTypes(uint256 ) view returns(string name, uint256 quorum, uint256 noticePeriod, uint256 votingPeriod, bool canBeNegative)
func (_ResolutionManager *ResolutionManagerSession) ResolutionTypes(arg0 *big.Int) (struct {
	Name          string
	Quorum        *big.Int
	NoticePeriod  *big.Int
	VotingPeriod  *big.Int
	CanBeNegative bool
}, error) {
	return _ResolutionManager.Contract.ResolutionTypes(&_ResolutionManager.CallOpts, arg0)
}

// ResolutionTypes is a free data retrieval call binding the contract method 0x96648206.
//
// Solidity: function resolutionTypes(uint256 ) view returns(string name, uint256 quorum, uint256 noticePeriod, uint256 votingPeriod, bool canBeNegative)
func (_ResolutionManager *ResolutionManagerCallerSession) ResolutionTypes(arg0 *big.Int) (struct {
	Name          string
	Quorum        *big.Int
	NoticePeriod  *big.Int
	VotingPeriod  *big.Int
	CanBeNegative bool
}, error) {
	return _ResolutionManager.Contract.ResolutionTypes(&_ResolutionManager.CallOpts, arg0)
}

// Resolutions is a free data retrieval call binding the contract method 0xa4b7f5ce.
//
// Solidity: function resolutions(uint256 ) view returns(string dataURI, uint256 resolutionTypeId, uint256 approveTimestamp, uint256 snapshotId, uint256 yesVotesTotal, bool isNegative, uint256 rejectionTimestamp, uint256 executionTimestamp)
func (_ResolutionManager *ResolutionManagerCaller) Resolutions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DataURI            string
	ResolutionTypeId   *big.Int
	ApproveTimestamp   *big.Int
	SnapshotId         *big.Int
	YesVotesTotal      *big.Int
	IsNegative         bool
	RejectionTimestamp *big.Int
	ExecutionTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _ResolutionManager.contract.Call(opts, &out, "resolutions", arg0)

	outstruct := new(struct {
		DataURI            string
		ResolutionTypeId   *big.Int
		ApproveTimestamp   *big.Int
		SnapshotId         *big.Int
		YesVotesTotal      *big.Int
		IsNegative         bool
		RejectionTimestamp *big.Int
		ExecutionTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DataURI = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ResolutionTypeId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ApproveTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SnapshotId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.YesVotesTotal = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsNegative = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.RejectionTimestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.ExecutionTimestamp = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Resolutions is a free data retrieval call binding the contract method 0xa4b7f5ce.
//
// Solidity: function resolutions(uint256 ) view returns(string dataURI, uint256 resolutionTypeId, uint256 approveTimestamp, uint256 snapshotId, uint256 yesVotesTotal, bool isNegative, uint256 rejectionTimestamp, uint256 executionTimestamp)
func (_ResolutionManager *ResolutionManagerSession) Resolutions(arg0 *big.Int) (struct {
	DataURI            string
	ResolutionTypeId   *big.Int
	ApproveTimestamp   *big.Int
	SnapshotId         *big.Int
	YesVotesTotal      *big.Int
	IsNegative         bool
	RejectionTimestamp *big.Int
	ExecutionTimestamp *big.Int
}, error) {
	return _ResolutionManager.Contract.Resolutions(&_ResolutionManager.CallOpts, arg0)
}

// Resolutions is a free data retrieval call binding the contract method 0xa4b7f5ce.
//
// Solidity: function resolutions(uint256 ) view returns(string dataURI, uint256 resolutionTypeId, uint256 approveTimestamp, uint256 snapshotId, uint256 yesVotesTotal, bool isNegative, uint256 rejectionTimestamp, uint256 executionTimestamp)
func (_ResolutionManager *ResolutionManagerCallerSession) Resolutions(arg0 *big.Int) (struct {
	DataURI            string
	ResolutionTypeId   *big.Int
	ApproveTimestamp   *big.Int
	SnapshotId         *big.Int
	YesVotesTotal      *big.Int
	IsNegative         bool
	RejectionTimestamp *big.Int
	ExecutionTimestamp *big.Int
}, error) {
	return _ResolutionManager.Contract.Resolutions(&_ResolutionManager.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ResolutionManager *ResolutionManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ResolutionManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ResolutionManager *ResolutionManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ResolutionManager.Contract.SupportsInterface(&_ResolutionManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ResolutionManager *ResolutionManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ResolutionManager.Contract.SupportsInterface(&_ResolutionManager.CallOpts, interfaceId)
}

// AddResolutionType is a paid mutator transaction binding the contract method 0x25bd38ea.
//
// Solidity: function addResolutionType(string name, uint256 quorum, uint256 noticePeriod, uint256 votingPeriod, bool canBeNegative) returns()
func (_ResolutionManager *ResolutionManagerTransactor) AddResolutionType(opts *bind.TransactOpts, name string, quorum *big.Int, noticePeriod *big.Int, votingPeriod *big.Int, canBeNegative bool) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "addResolutionType", name, quorum, noticePeriod, votingPeriod, canBeNegative)
}

// AddResolutionType is a paid mutator transaction binding the contract method 0x25bd38ea.
//
// Solidity: function addResolutionType(string name, uint256 quorum, uint256 noticePeriod, uint256 votingPeriod, bool canBeNegative) returns()
func (_ResolutionManager *ResolutionManagerSession) AddResolutionType(name string, quorum *big.Int, noticePeriod *big.Int, votingPeriod *big.Int, canBeNegative bool) (*types.Transaction, error) {
	return _ResolutionManager.Contract.AddResolutionType(&_ResolutionManager.TransactOpts, name, quorum, noticePeriod, votingPeriod, canBeNegative)
}

// AddResolutionType is a paid mutator transaction binding the contract method 0x25bd38ea.
//
// Solidity: function addResolutionType(string name, uint256 quorum, uint256 noticePeriod, uint256 votingPeriod, bool canBeNegative) returns()
func (_ResolutionManager *ResolutionManagerTransactorSession) AddResolutionType(name string, quorum *big.Int, noticePeriod *big.Int, votingPeriod *big.Int, canBeNegative bool) (*types.Transaction, error) {
	return _ResolutionManager.Contract.AddResolutionType(&_ResolutionManager.TransactOpts, name, quorum, noticePeriod, votingPeriod, canBeNegative)
}

// ApproveResolution is a paid mutator transaction binding the contract method 0x61053090.
//
// Solidity: function approveResolution(uint256 resolutionId) returns()
func (_ResolutionManager *ResolutionManagerTransactor) ApproveResolution(opts *bind.TransactOpts, resolutionId *big.Int) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "approveResolution", resolutionId)
}

// ApproveResolution is a paid mutator transaction binding the contract method 0x61053090.
//
// Solidity: function approveResolution(uint256 resolutionId) returns()
func (_ResolutionManager *ResolutionManagerSession) ApproveResolution(resolutionId *big.Int) (*types.Transaction, error) {
	return _ResolutionManager.Contract.ApproveResolution(&_ResolutionManager.TransactOpts, resolutionId)
}

// ApproveResolution is a paid mutator transaction binding the contract method 0x61053090.
//
// Solidity: function approveResolution(uint256 resolutionId) returns()
func (_ResolutionManager *ResolutionManagerTransactorSession) ApproveResolution(resolutionId *big.Int) (*types.Transaction, error) {
	return _ResolutionManager.Contract.ApproveResolution(&_ResolutionManager.TransactOpts, resolutionId)
}

// CreateResolution is a paid mutator transaction binding the contract method 0x325ae67a.
//
// Solidity: function createResolution(string dataURI, uint256 resolutionTypeId, bool isNegative, address[] executionTo, bytes[] executionData) returns(uint256)
func (_ResolutionManager *ResolutionManagerTransactor) CreateResolution(opts *bind.TransactOpts, dataURI string, resolutionTypeId *big.Int, isNegative bool, executionTo []common.Address, executionData [][]byte) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "createResolution", dataURI, resolutionTypeId, isNegative, executionTo, executionData)
}

// CreateResolution is a paid mutator transaction binding the contract method 0x325ae67a.
//
// Solidity: function createResolution(string dataURI, uint256 resolutionTypeId, bool isNegative, address[] executionTo, bytes[] executionData) returns(uint256)
func (_ResolutionManager *ResolutionManagerSession) CreateResolution(dataURI string, resolutionTypeId *big.Int, isNegative bool, executionTo []common.Address, executionData [][]byte) (*types.Transaction, error) {
	return _ResolutionManager.Contract.CreateResolution(&_ResolutionManager.TransactOpts, dataURI, resolutionTypeId, isNegative, executionTo, executionData)
}

// CreateResolution is a paid mutator transaction binding the contract method 0x325ae67a.
//
// Solidity: function createResolution(string dataURI, uint256 resolutionTypeId, bool isNegative, address[] executionTo, bytes[] executionData) returns(uint256)
func (_ResolutionManager *ResolutionManagerTransactorSession) CreateResolution(dataURI string, resolutionTypeId *big.Int, isNegative bool, executionTo []common.Address, executionData [][]byte) (*types.Transaction, error) {
	return _ResolutionManager.Contract.CreateResolution(&_ResolutionManager.TransactOpts, dataURI, resolutionTypeId, isNegative, executionTo, executionData)
}

// ExecuteResolution is a paid mutator transaction binding the contract method 0x7cbf3b6b.
//
// Solidity: function executeResolution(uint256 resolutionId) returns()
func (_ResolutionManager *ResolutionManagerTransactor) ExecuteResolution(opts *bind.TransactOpts, resolutionId *big.Int) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "executeResolution", resolutionId)
}

// ExecuteResolution is a paid mutator transaction binding the contract method 0x7cbf3b6b.
//
// Solidity: function executeResolution(uint256 resolutionId) returns()
func (_ResolutionManager *ResolutionManagerSession) ExecuteResolution(resolutionId *big.Int) (*types.Transaction, error) {
	return _ResolutionManager.Contract.ExecuteResolution(&_ResolutionManager.TransactOpts, resolutionId)
}

// ExecuteResolution is a paid mutator transaction binding the contract method 0x7cbf3b6b.
//
// Solidity: function executeResolution(uint256 resolutionId) returns()
func (_ResolutionManager *ResolutionManagerTransactorSession) ExecuteResolution(resolutionId *big.Int) (*types.Transaction, error) {
	return _ResolutionManager.Contract.ExecuteResolution(&_ResolutionManager.TransactOpts, resolutionId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ResolutionManager *ResolutionManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ResolutionManager *ResolutionManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.GrantRole(&_ResolutionManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ResolutionManager *ResolutionManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.GrantRole(&_ResolutionManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address shareholderRegistry, address telediskoToken, address voting) returns()
func (_ResolutionManager *ResolutionManagerTransactor) Initialize(opts *bind.TransactOpts, shareholderRegistry common.Address, telediskoToken common.Address, voting common.Address) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "initialize", shareholderRegistry, telediskoToken, voting)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address shareholderRegistry, address telediskoToken, address voting) returns()
func (_ResolutionManager *ResolutionManagerSession) Initialize(shareholderRegistry common.Address, telediskoToken common.Address, voting common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.Initialize(&_ResolutionManager.TransactOpts, shareholderRegistry, telediskoToken, voting)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address shareholderRegistry, address telediskoToken, address voting) returns()
func (_ResolutionManager *ResolutionManagerTransactorSession) Initialize(shareholderRegistry common.Address, telediskoToken common.Address, voting common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.Initialize(&_ResolutionManager.TransactOpts, shareholderRegistry, telediskoToken, voting)
}

// RejectResolution is a paid mutator transaction binding the contract method 0xfd2c9719.
//
// Solidity: function rejectResolution(uint256 resolutionId) returns()
func (_ResolutionManager *ResolutionManagerTransactor) RejectResolution(opts *bind.TransactOpts, resolutionId *big.Int) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "rejectResolution", resolutionId)
}

// RejectResolution is a paid mutator transaction binding the contract method 0xfd2c9719.
//
// Solidity: function rejectResolution(uint256 resolutionId) returns()
func (_ResolutionManager *ResolutionManagerSession) RejectResolution(resolutionId *big.Int) (*types.Transaction, error) {
	return _ResolutionManager.Contract.RejectResolution(&_ResolutionManager.TransactOpts, resolutionId)
}

// RejectResolution is a paid mutator transaction binding the contract method 0xfd2c9719.
//
// Solidity: function rejectResolution(uint256 resolutionId) returns()
func (_ResolutionManager *ResolutionManagerTransactorSession) RejectResolution(resolutionId *big.Int) (*types.Transaction, error) {
	return _ResolutionManager.Contract.RejectResolution(&_ResolutionManager.TransactOpts, resolutionId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ResolutionManager *ResolutionManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ResolutionManager *ResolutionManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.RenounceRole(&_ResolutionManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ResolutionManager *ResolutionManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.RenounceRole(&_ResolutionManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ResolutionManager *ResolutionManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ResolutionManager *ResolutionManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.RevokeRole(&_ResolutionManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ResolutionManager *ResolutionManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.RevokeRole(&_ResolutionManager.TransactOpts, role, account)
}

// SetShareholderRegistry is a paid mutator transaction binding the contract method 0xa260b102.
//
// Solidity: function setShareholderRegistry(address shareholderRegistry) returns()
func (_ResolutionManager *ResolutionManagerTransactor) SetShareholderRegistry(opts *bind.TransactOpts, shareholderRegistry common.Address) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "setShareholderRegistry", shareholderRegistry)
}

// SetShareholderRegistry is a paid mutator transaction binding the contract method 0xa260b102.
//
// Solidity: function setShareholderRegistry(address shareholderRegistry) returns()
func (_ResolutionManager *ResolutionManagerSession) SetShareholderRegistry(shareholderRegistry common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.SetShareholderRegistry(&_ResolutionManager.TransactOpts, shareholderRegistry)
}

// SetShareholderRegistry is a paid mutator transaction binding the contract method 0xa260b102.
//
// Solidity: function setShareholderRegistry(address shareholderRegistry) returns()
func (_ResolutionManager *ResolutionManagerTransactorSession) SetShareholderRegistry(shareholderRegistry common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.SetShareholderRegistry(&_ResolutionManager.TransactOpts, shareholderRegistry)
}

// SetTelediskoToken is a paid mutator transaction binding the contract method 0xa0aae32f.
//
// Solidity: function setTelediskoToken(address telediskoToken) returns()
func (_ResolutionManager *ResolutionManagerTransactor) SetTelediskoToken(opts *bind.TransactOpts, telediskoToken common.Address) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "setTelediskoToken", telediskoToken)
}

// SetTelediskoToken is a paid mutator transaction binding the contract method 0xa0aae32f.
//
// Solidity: function setTelediskoToken(address telediskoToken) returns()
func (_ResolutionManager *ResolutionManagerSession) SetTelediskoToken(telediskoToken common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.SetTelediskoToken(&_ResolutionManager.TransactOpts, telediskoToken)
}

// SetTelediskoToken is a paid mutator transaction binding the contract method 0xa0aae32f.
//
// Solidity: function setTelediskoToken(address telediskoToken) returns()
func (_ResolutionManager *ResolutionManagerTransactorSession) SetTelediskoToken(telediskoToken common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.SetTelediskoToken(&_ResolutionManager.TransactOpts, telediskoToken)
}

// SetVoting is a paid mutator transaction binding the contract method 0xc4d07951.
//
// Solidity: function setVoting(address voting) returns()
func (_ResolutionManager *ResolutionManagerTransactor) SetVoting(opts *bind.TransactOpts, voting common.Address) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "setVoting", voting)
}

// SetVoting is a paid mutator transaction binding the contract method 0xc4d07951.
//
// Solidity: function setVoting(address voting) returns()
func (_ResolutionManager *ResolutionManagerSession) SetVoting(voting common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.SetVoting(&_ResolutionManager.TransactOpts, voting)
}

// SetVoting is a paid mutator transaction binding the contract method 0xc4d07951.
//
// Solidity: function setVoting(address voting) returns()
func (_ResolutionManager *ResolutionManagerTransactorSession) SetVoting(voting common.Address) (*types.Transaction, error) {
	return _ResolutionManager.Contract.SetVoting(&_ResolutionManager.TransactOpts, voting)
}

// UpdateResolution is a paid mutator transaction binding the contract method 0x444be179.
//
// Solidity: function updateResolution(uint256 resolutionId, string dataURI, uint256 resolutionTypeId, bool isNegative, address[] executionTo, bytes[] executionData) returns()
func (_ResolutionManager *ResolutionManagerTransactor) UpdateResolution(opts *bind.TransactOpts, resolutionId *big.Int, dataURI string, resolutionTypeId *big.Int, isNegative bool, executionTo []common.Address, executionData [][]byte) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "updateResolution", resolutionId, dataURI, resolutionTypeId, isNegative, executionTo, executionData)
}

// UpdateResolution is a paid mutator transaction binding the contract method 0x444be179.
//
// Solidity: function updateResolution(uint256 resolutionId, string dataURI, uint256 resolutionTypeId, bool isNegative, address[] executionTo, bytes[] executionData) returns()
func (_ResolutionManager *ResolutionManagerSession) UpdateResolution(resolutionId *big.Int, dataURI string, resolutionTypeId *big.Int, isNegative bool, executionTo []common.Address, executionData [][]byte) (*types.Transaction, error) {
	return _ResolutionManager.Contract.UpdateResolution(&_ResolutionManager.TransactOpts, resolutionId, dataURI, resolutionTypeId, isNegative, executionTo, executionData)
}

// UpdateResolution is a paid mutator transaction binding the contract method 0x444be179.
//
// Solidity: function updateResolution(uint256 resolutionId, string dataURI, uint256 resolutionTypeId, bool isNegative, address[] executionTo, bytes[] executionData) returns()
func (_ResolutionManager *ResolutionManagerTransactorSession) UpdateResolution(resolutionId *big.Int, dataURI string, resolutionTypeId *big.Int, isNegative bool, executionTo []common.Address, executionData [][]byte) (*types.Transaction, error) {
	return _ResolutionManager.Contract.UpdateResolution(&_ResolutionManager.TransactOpts, resolutionId, dataURI, resolutionTypeId, isNegative, executionTo, executionData)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 resolutionId, bool isYes) returns()
func (_ResolutionManager *ResolutionManagerTransactor) Vote(opts *bind.TransactOpts, resolutionId *big.Int, isYes bool) (*types.Transaction, error) {
	return _ResolutionManager.contract.Transact(opts, "vote", resolutionId, isYes)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 resolutionId, bool isYes) returns()
func (_ResolutionManager *ResolutionManagerSession) Vote(resolutionId *big.Int, isYes bool) (*types.Transaction, error) {
	return _ResolutionManager.Contract.Vote(&_ResolutionManager.TransactOpts, resolutionId, isYes)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 resolutionId, bool isYes) returns()
func (_ResolutionManager *ResolutionManagerTransactorSession) Vote(resolutionId *big.Int, isYes bool) (*types.Transaction, error) {
	return _ResolutionManager.Contract.Vote(&_ResolutionManager.TransactOpts, resolutionId, isYes)
}

// ResolutionManagerDelegateLostVotingPowerIterator is returned from FilterDelegateLostVotingPower and is used to iterate over the raw logs and unpacked data for DelegateLostVotingPower events raised by the ResolutionManager contract.
type ResolutionManagerDelegateLostVotingPowerIterator struct {
	Event *ResolutionManagerDelegateLostVotingPower // Event containing the contract specifics and raw log

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
func (it *ResolutionManagerDelegateLostVotingPowerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionManagerDelegateLostVotingPower)
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
		it.Event = new(ResolutionManagerDelegateLostVotingPower)
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
func (it *ResolutionManagerDelegateLostVotingPowerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionManagerDelegateLostVotingPowerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionManagerDelegateLostVotingPower represents a DelegateLostVotingPower event raised by the ResolutionManager contract.
type ResolutionManagerDelegateLostVotingPower struct {
	From         common.Address
	ResolutionId *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateLostVotingPower is a free log retrieval operation binding the contract event 0x2a753d47dc7411744b037903d0f3d1273bd5046e19f509d13d0a896bb6dca896.
//
// Solidity: event DelegateLostVotingPower(address indexed from, uint256 indexed resolutionId, uint256 amount)
func (_ResolutionManager *ResolutionManagerFilterer) FilterDelegateLostVotingPower(opts *bind.FilterOpts, from []common.Address, resolutionId []*big.Int) (*ResolutionManagerDelegateLostVotingPowerIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.FilterLogs(opts, "DelegateLostVotingPower", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerDelegateLostVotingPowerIterator{contract: _ResolutionManager.contract, event: "DelegateLostVotingPower", logs: logs, sub: sub}, nil
}

// WatchDelegateLostVotingPower is a free log subscription operation binding the contract event 0x2a753d47dc7411744b037903d0f3d1273bd5046e19f509d13d0a896bb6dca896.
//
// Solidity: event DelegateLostVotingPower(address indexed from, uint256 indexed resolutionId, uint256 amount)
func (_ResolutionManager *ResolutionManagerFilterer) WatchDelegateLostVotingPower(opts *bind.WatchOpts, sink chan<- *ResolutionManagerDelegateLostVotingPower, from []common.Address, resolutionId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.WatchLogs(opts, "DelegateLostVotingPower", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionManagerDelegateLostVotingPower)
				if err := _ResolutionManager.contract.UnpackLog(event, "DelegateLostVotingPower", log); err != nil {
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

// ParseDelegateLostVotingPower is a log parse operation binding the contract event 0x2a753d47dc7411744b037903d0f3d1273bd5046e19f509d13d0a896bb6dca896.
//
// Solidity: event DelegateLostVotingPower(address indexed from, uint256 indexed resolutionId, uint256 amount)
func (_ResolutionManager *ResolutionManagerFilterer) ParseDelegateLostVotingPower(log types.Log) (*ResolutionManagerDelegateLostVotingPower, error) {
	event := new(ResolutionManagerDelegateLostVotingPower)
	if err := _ResolutionManager.contract.UnpackLog(event, "DelegateLostVotingPower", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ResolutionManager contract.
type ResolutionManagerInitializedIterator struct {
	Event *ResolutionManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ResolutionManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionManagerInitialized)
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
		it.Event = new(ResolutionManagerInitialized)
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
func (it *ResolutionManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionManagerInitialized represents a Initialized event raised by the ResolutionManager contract.
type ResolutionManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ResolutionManager *ResolutionManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ResolutionManagerInitializedIterator, error) {

	logs, sub, err := _ResolutionManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerInitializedIterator{contract: _ResolutionManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ResolutionManager *ResolutionManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ResolutionManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ResolutionManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionManagerInitialized)
				if err := _ResolutionManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ResolutionManager *ResolutionManagerFilterer) ParseInitialized(log types.Log) (*ResolutionManagerInitialized, error) {
	event := new(ResolutionManagerInitialized)
	if err := _ResolutionManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionManagerResolutionApprovedIterator is returned from FilterResolutionApproved and is used to iterate over the raw logs and unpacked data for ResolutionApproved events raised by the ResolutionManager contract.
type ResolutionManagerResolutionApprovedIterator struct {
	Event *ResolutionManagerResolutionApproved // Event containing the contract specifics and raw log

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
func (it *ResolutionManagerResolutionApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionManagerResolutionApproved)
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
		it.Event = new(ResolutionManagerResolutionApproved)
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
func (it *ResolutionManagerResolutionApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionManagerResolutionApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionManagerResolutionApproved represents a ResolutionApproved event raised by the ResolutionManager contract.
type ResolutionManagerResolutionApproved struct {
	From         common.Address
	ResolutionId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterResolutionApproved is a free log retrieval operation binding the contract event 0x6f6c6ecde2adc206ce4f038800c59db57260d8a236c7a599438a553d7e4b3aa6.
//
// Solidity: event ResolutionApproved(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) FilterResolutionApproved(opts *bind.FilterOpts, from []common.Address, resolutionId []*big.Int) (*ResolutionManagerResolutionApprovedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.FilterLogs(opts, "ResolutionApproved", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerResolutionApprovedIterator{contract: _ResolutionManager.contract, event: "ResolutionApproved", logs: logs, sub: sub}, nil
}

// WatchResolutionApproved is a free log subscription operation binding the contract event 0x6f6c6ecde2adc206ce4f038800c59db57260d8a236c7a599438a553d7e4b3aa6.
//
// Solidity: event ResolutionApproved(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) WatchResolutionApproved(opts *bind.WatchOpts, sink chan<- *ResolutionManagerResolutionApproved, from []common.Address, resolutionId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.WatchLogs(opts, "ResolutionApproved", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionManagerResolutionApproved)
				if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionApproved", log); err != nil {
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

// ParseResolutionApproved is a log parse operation binding the contract event 0x6f6c6ecde2adc206ce4f038800c59db57260d8a236c7a599438a553d7e4b3aa6.
//
// Solidity: event ResolutionApproved(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) ParseResolutionApproved(log types.Log) (*ResolutionManagerResolutionApproved, error) {
	event := new(ResolutionManagerResolutionApproved)
	if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionManagerResolutionCreatedIterator is returned from FilterResolutionCreated and is used to iterate over the raw logs and unpacked data for ResolutionCreated events raised by the ResolutionManager contract.
type ResolutionManagerResolutionCreatedIterator struct {
	Event *ResolutionManagerResolutionCreated // Event containing the contract specifics and raw log

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
func (it *ResolutionManagerResolutionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionManagerResolutionCreated)
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
		it.Event = new(ResolutionManagerResolutionCreated)
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
func (it *ResolutionManagerResolutionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionManagerResolutionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionManagerResolutionCreated represents a ResolutionCreated event raised by the ResolutionManager contract.
type ResolutionManagerResolutionCreated struct {
	From         common.Address
	ResolutionId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterResolutionCreated is a free log retrieval operation binding the contract event 0x5cb605df30db4389dcc26eb7b6ac40cee45e3da41204619ed102f777bf814f8b.
//
// Solidity: event ResolutionCreated(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) FilterResolutionCreated(opts *bind.FilterOpts, from []common.Address, resolutionId []*big.Int) (*ResolutionManagerResolutionCreatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.FilterLogs(opts, "ResolutionCreated", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerResolutionCreatedIterator{contract: _ResolutionManager.contract, event: "ResolutionCreated", logs: logs, sub: sub}, nil
}

// WatchResolutionCreated is a free log subscription operation binding the contract event 0x5cb605df30db4389dcc26eb7b6ac40cee45e3da41204619ed102f777bf814f8b.
//
// Solidity: event ResolutionCreated(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) WatchResolutionCreated(opts *bind.WatchOpts, sink chan<- *ResolutionManagerResolutionCreated, from []common.Address, resolutionId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.WatchLogs(opts, "ResolutionCreated", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionManagerResolutionCreated)
				if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionCreated", log); err != nil {
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

// ParseResolutionCreated is a log parse operation binding the contract event 0x5cb605df30db4389dcc26eb7b6ac40cee45e3da41204619ed102f777bf814f8b.
//
// Solidity: event ResolutionCreated(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) ParseResolutionCreated(log types.Log) (*ResolutionManagerResolutionCreated, error) {
	event := new(ResolutionManagerResolutionCreated)
	if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionManagerResolutionExecutedIterator is returned from FilterResolutionExecuted and is used to iterate over the raw logs and unpacked data for ResolutionExecuted events raised by the ResolutionManager contract.
type ResolutionManagerResolutionExecutedIterator struct {
	Event *ResolutionManagerResolutionExecuted // Event containing the contract specifics and raw log

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
func (it *ResolutionManagerResolutionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionManagerResolutionExecuted)
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
		it.Event = new(ResolutionManagerResolutionExecuted)
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
func (it *ResolutionManagerResolutionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionManagerResolutionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionManagerResolutionExecuted represents a ResolutionExecuted event raised by the ResolutionManager contract.
type ResolutionManagerResolutionExecuted struct {
	From         common.Address
	ResolutionId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterResolutionExecuted is a free log retrieval operation binding the contract event 0x14a68c83cab032d52f3c5eb75dcda4cfdc0efbe72a39c36bda85edd44918df09.
//
// Solidity: event ResolutionExecuted(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) FilterResolutionExecuted(opts *bind.FilterOpts, from []common.Address, resolutionId []*big.Int) (*ResolutionManagerResolutionExecutedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.FilterLogs(opts, "ResolutionExecuted", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerResolutionExecutedIterator{contract: _ResolutionManager.contract, event: "ResolutionExecuted", logs: logs, sub: sub}, nil
}

// WatchResolutionExecuted is a free log subscription operation binding the contract event 0x14a68c83cab032d52f3c5eb75dcda4cfdc0efbe72a39c36bda85edd44918df09.
//
// Solidity: event ResolutionExecuted(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) WatchResolutionExecuted(opts *bind.WatchOpts, sink chan<- *ResolutionManagerResolutionExecuted, from []common.Address, resolutionId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.WatchLogs(opts, "ResolutionExecuted", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionManagerResolutionExecuted)
				if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionExecuted", log); err != nil {
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

// ParseResolutionExecuted is a log parse operation binding the contract event 0x14a68c83cab032d52f3c5eb75dcda4cfdc0efbe72a39c36bda85edd44918df09.
//
// Solidity: event ResolutionExecuted(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) ParseResolutionExecuted(log types.Log) (*ResolutionManagerResolutionExecuted, error) {
	event := new(ResolutionManagerResolutionExecuted)
	if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionManagerResolutionRejectedIterator is returned from FilterResolutionRejected and is used to iterate over the raw logs and unpacked data for ResolutionRejected events raised by the ResolutionManager contract.
type ResolutionManagerResolutionRejectedIterator struct {
	Event *ResolutionManagerResolutionRejected // Event containing the contract specifics and raw log

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
func (it *ResolutionManagerResolutionRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionManagerResolutionRejected)
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
		it.Event = new(ResolutionManagerResolutionRejected)
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
func (it *ResolutionManagerResolutionRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionManagerResolutionRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionManagerResolutionRejected represents a ResolutionRejected event raised by the ResolutionManager contract.
type ResolutionManagerResolutionRejected struct {
	From         common.Address
	ResolutionId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterResolutionRejected is a free log retrieval operation binding the contract event 0x206580a2d87d2bd5906ada5f3ca523aefb5f6b61f702ea347d5399778b647dae.
//
// Solidity: event ResolutionRejected(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) FilterResolutionRejected(opts *bind.FilterOpts, from []common.Address, resolutionId []*big.Int) (*ResolutionManagerResolutionRejectedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.FilterLogs(opts, "ResolutionRejected", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerResolutionRejectedIterator{contract: _ResolutionManager.contract, event: "ResolutionRejected", logs: logs, sub: sub}, nil
}

// WatchResolutionRejected is a free log subscription operation binding the contract event 0x206580a2d87d2bd5906ada5f3ca523aefb5f6b61f702ea347d5399778b647dae.
//
// Solidity: event ResolutionRejected(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) WatchResolutionRejected(opts *bind.WatchOpts, sink chan<- *ResolutionManagerResolutionRejected, from []common.Address, resolutionId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.WatchLogs(opts, "ResolutionRejected", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionManagerResolutionRejected)
				if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionRejected", log); err != nil {
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

// ParseResolutionRejected is a log parse operation binding the contract event 0x206580a2d87d2bd5906ada5f3ca523aefb5f6b61f702ea347d5399778b647dae.
//
// Solidity: event ResolutionRejected(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) ParseResolutionRejected(log types.Log) (*ResolutionManagerResolutionRejected, error) {
	event := new(ResolutionManagerResolutionRejected)
	if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionManagerResolutionTypeCreatedIterator is returned from FilterResolutionTypeCreated and is used to iterate over the raw logs and unpacked data for ResolutionTypeCreated events raised by the ResolutionManager contract.
type ResolutionManagerResolutionTypeCreatedIterator struct {
	Event *ResolutionManagerResolutionTypeCreated // Event containing the contract specifics and raw log

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
func (it *ResolutionManagerResolutionTypeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionManagerResolutionTypeCreated)
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
		it.Event = new(ResolutionManagerResolutionTypeCreated)
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
func (it *ResolutionManagerResolutionTypeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionManagerResolutionTypeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionManagerResolutionTypeCreated represents a ResolutionTypeCreated event raised by the ResolutionManager contract.
type ResolutionManagerResolutionTypeCreated struct {
	From      common.Address
	TypeIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResolutionTypeCreated is a free log retrieval operation binding the contract event 0xc9b91babd5124906e4cca98d9386ddac639dfee98b2bd33a27dbb5c048bd855a.
//
// Solidity: event ResolutionTypeCreated(address indexed from, uint256 indexed typeIndex)
func (_ResolutionManager *ResolutionManagerFilterer) FilterResolutionTypeCreated(opts *bind.FilterOpts, from []common.Address, typeIndex []*big.Int) (*ResolutionManagerResolutionTypeCreatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var typeIndexRule []interface{}
	for _, typeIndexItem := range typeIndex {
		typeIndexRule = append(typeIndexRule, typeIndexItem)
	}

	logs, sub, err := _ResolutionManager.contract.FilterLogs(opts, "ResolutionTypeCreated", fromRule, typeIndexRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerResolutionTypeCreatedIterator{contract: _ResolutionManager.contract, event: "ResolutionTypeCreated", logs: logs, sub: sub}, nil
}

// WatchResolutionTypeCreated is a free log subscription operation binding the contract event 0xc9b91babd5124906e4cca98d9386ddac639dfee98b2bd33a27dbb5c048bd855a.
//
// Solidity: event ResolutionTypeCreated(address indexed from, uint256 indexed typeIndex)
func (_ResolutionManager *ResolutionManagerFilterer) WatchResolutionTypeCreated(opts *bind.WatchOpts, sink chan<- *ResolutionManagerResolutionTypeCreated, from []common.Address, typeIndex []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var typeIndexRule []interface{}
	for _, typeIndexItem := range typeIndex {
		typeIndexRule = append(typeIndexRule, typeIndexItem)
	}

	logs, sub, err := _ResolutionManager.contract.WatchLogs(opts, "ResolutionTypeCreated", fromRule, typeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionManagerResolutionTypeCreated)
				if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionTypeCreated", log); err != nil {
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

// ParseResolutionTypeCreated is a log parse operation binding the contract event 0xc9b91babd5124906e4cca98d9386ddac639dfee98b2bd33a27dbb5c048bd855a.
//
// Solidity: event ResolutionTypeCreated(address indexed from, uint256 indexed typeIndex)
func (_ResolutionManager *ResolutionManagerFilterer) ParseResolutionTypeCreated(log types.Log) (*ResolutionManagerResolutionTypeCreated, error) {
	event := new(ResolutionManagerResolutionTypeCreated)
	if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionTypeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionManagerResolutionUpdatedIterator is returned from FilterResolutionUpdated and is used to iterate over the raw logs and unpacked data for ResolutionUpdated events raised by the ResolutionManager contract.
type ResolutionManagerResolutionUpdatedIterator struct {
	Event *ResolutionManagerResolutionUpdated // Event containing the contract specifics and raw log

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
func (it *ResolutionManagerResolutionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionManagerResolutionUpdated)
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
		it.Event = new(ResolutionManagerResolutionUpdated)
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
func (it *ResolutionManagerResolutionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionManagerResolutionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionManagerResolutionUpdated represents a ResolutionUpdated event raised by the ResolutionManager contract.
type ResolutionManagerResolutionUpdated struct {
	From         common.Address
	ResolutionId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterResolutionUpdated is a free log retrieval operation binding the contract event 0xe9fbf855f59ea63a795f1f64e447af0f3316384d1f74ac2685c9cf4231d15ef8.
//
// Solidity: event ResolutionUpdated(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) FilterResolutionUpdated(opts *bind.FilterOpts, from []common.Address, resolutionId []*big.Int) (*ResolutionManagerResolutionUpdatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.FilterLogs(opts, "ResolutionUpdated", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerResolutionUpdatedIterator{contract: _ResolutionManager.contract, event: "ResolutionUpdated", logs: logs, sub: sub}, nil
}

// WatchResolutionUpdated is a free log subscription operation binding the contract event 0xe9fbf855f59ea63a795f1f64e447af0f3316384d1f74ac2685c9cf4231d15ef8.
//
// Solidity: event ResolutionUpdated(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) WatchResolutionUpdated(opts *bind.WatchOpts, sink chan<- *ResolutionManagerResolutionUpdated, from []common.Address, resolutionId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.WatchLogs(opts, "ResolutionUpdated", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionManagerResolutionUpdated)
				if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionUpdated", log); err != nil {
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

// ParseResolutionUpdated is a log parse operation binding the contract event 0xe9fbf855f59ea63a795f1f64e447af0f3316384d1f74ac2685c9cf4231d15ef8.
//
// Solidity: event ResolutionUpdated(address indexed from, uint256 indexed resolutionId)
func (_ResolutionManager *ResolutionManagerFilterer) ParseResolutionUpdated(log types.Log) (*ResolutionManagerResolutionUpdated, error) {
	event := new(ResolutionManagerResolutionUpdated)
	if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionManagerResolutionVotedIterator is returned from FilterResolutionVoted and is used to iterate over the raw logs and unpacked data for ResolutionVoted events raised by the ResolutionManager contract.
type ResolutionManagerResolutionVotedIterator struct {
	Event *ResolutionManagerResolutionVoted // Event containing the contract specifics and raw log

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
func (it *ResolutionManagerResolutionVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionManagerResolutionVoted)
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
		it.Event = new(ResolutionManagerResolutionVoted)
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
func (it *ResolutionManagerResolutionVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionManagerResolutionVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionManagerResolutionVoted represents a ResolutionVoted event raised by the ResolutionManager contract.
type ResolutionManagerResolutionVoted struct {
	From         common.Address
	ResolutionId *big.Int
	VotingPower  *big.Int
	IsYes        bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterResolutionVoted is a free log retrieval operation binding the contract event 0xe38278e9a4991080fd5c68b1594173a60b3db30bd509a1fe3d0b68f9fcb2e0ab.
//
// Solidity: event ResolutionVoted(address indexed from, uint256 indexed resolutionId, uint256 votingPower, bool isYes)
func (_ResolutionManager *ResolutionManagerFilterer) FilterResolutionVoted(opts *bind.FilterOpts, from []common.Address, resolutionId []*big.Int) (*ResolutionManagerResolutionVotedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.FilterLogs(opts, "ResolutionVoted", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerResolutionVotedIterator{contract: _ResolutionManager.contract, event: "ResolutionVoted", logs: logs, sub: sub}, nil
}

// WatchResolutionVoted is a free log subscription operation binding the contract event 0xe38278e9a4991080fd5c68b1594173a60b3db30bd509a1fe3d0b68f9fcb2e0ab.
//
// Solidity: event ResolutionVoted(address indexed from, uint256 indexed resolutionId, uint256 votingPower, bool isYes)
func (_ResolutionManager *ResolutionManagerFilterer) WatchResolutionVoted(opts *bind.WatchOpts, sink chan<- *ResolutionManagerResolutionVoted, from []common.Address, resolutionId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var resolutionIdRule []interface{}
	for _, resolutionIdItem := range resolutionId {
		resolutionIdRule = append(resolutionIdRule, resolutionIdItem)
	}

	logs, sub, err := _ResolutionManager.contract.WatchLogs(opts, "ResolutionVoted", fromRule, resolutionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionManagerResolutionVoted)
				if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionVoted", log); err != nil {
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

// ParseResolutionVoted is a log parse operation binding the contract event 0xe38278e9a4991080fd5c68b1594173a60b3db30bd509a1fe3d0b68f9fcb2e0ab.
//
// Solidity: event ResolutionVoted(address indexed from, uint256 indexed resolutionId, uint256 votingPower, bool isYes)
func (_ResolutionManager *ResolutionManagerFilterer) ParseResolutionVoted(log types.Log) (*ResolutionManagerResolutionVoted, error) {
	event := new(ResolutionManagerResolutionVoted)
	if err := _ResolutionManager.contract.UnpackLog(event, "ResolutionVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ResolutionManager contract.
type ResolutionManagerRoleAdminChangedIterator struct {
	Event *ResolutionManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ResolutionManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionManagerRoleAdminChanged)
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
		it.Event = new(ResolutionManagerRoleAdminChanged)
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
func (it *ResolutionManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionManagerRoleAdminChanged represents a RoleAdminChanged event raised by the ResolutionManager contract.
type ResolutionManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ResolutionManager *ResolutionManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ResolutionManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ResolutionManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerRoleAdminChangedIterator{contract: _ResolutionManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ResolutionManager *ResolutionManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ResolutionManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ResolutionManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionManagerRoleAdminChanged)
				if err := _ResolutionManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ResolutionManager *ResolutionManagerFilterer) ParseRoleAdminChanged(log types.Log) (*ResolutionManagerRoleAdminChanged, error) {
	event := new(ResolutionManagerRoleAdminChanged)
	if err := _ResolutionManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ResolutionManager contract.
type ResolutionManagerRoleGrantedIterator struct {
	Event *ResolutionManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *ResolutionManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionManagerRoleGranted)
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
		it.Event = new(ResolutionManagerRoleGranted)
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
func (it *ResolutionManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionManagerRoleGranted represents a RoleGranted event raised by the ResolutionManager contract.
type ResolutionManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ResolutionManager *ResolutionManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ResolutionManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _ResolutionManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerRoleGrantedIterator{contract: _ResolutionManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ResolutionManager *ResolutionManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ResolutionManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ResolutionManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionManagerRoleGranted)
				if err := _ResolutionManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ResolutionManager *ResolutionManagerFilterer) ParseRoleGranted(log types.Log) (*ResolutionManagerRoleGranted, error) {
	event := new(ResolutionManagerRoleGranted)
	if err := _ResolutionManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ResolutionManager contract.
type ResolutionManagerRoleRevokedIterator struct {
	Event *ResolutionManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ResolutionManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionManagerRoleRevoked)
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
		it.Event = new(ResolutionManagerRoleRevoked)
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
func (it *ResolutionManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionManagerRoleRevoked represents a RoleRevoked event raised by the ResolutionManager contract.
type ResolutionManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ResolutionManager *ResolutionManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ResolutionManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _ResolutionManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionManagerRoleRevokedIterator{contract: _ResolutionManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ResolutionManager *ResolutionManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ResolutionManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ResolutionManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionManagerRoleRevoked)
				if err := _ResolutionManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ResolutionManager *ResolutionManagerFilterer) ParseRoleRevoked(log types.Log) (*ResolutionManagerRoleRevoked, error) {
	event := new(ResolutionManagerRoleRevoked)
	if err := _ResolutionManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
