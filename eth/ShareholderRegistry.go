// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// ShareholderRegistryMetaData contains all meta data concerning the ShareholderRegistry contract.
var ShareholderRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previous\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"current\",\"type\":\"bytes32\"}],\"name\":\"StatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONTRIBUTOR_STATUS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INVESTOR_STATUS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGING_BOARD_STATUS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHAREHOLDER_STATUS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"}],\"name\":\"batchTransferFromDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"getBalanceAndStatusAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSnapshotId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"getStatusAt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"status\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"status\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAtLeast\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"status\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"isAtLeastAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"status\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVoting\",\"name\":\"voting\",\"type\":\"address\"}],\"name\":\"setVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ShareholderRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ShareholderRegistryMetaData.ABI instead.
var ShareholderRegistryABI = ShareholderRegistryMetaData.ABI

// ShareholderRegistry is an auto generated Go binding around an Ethereum contract.
type ShareholderRegistry struct {
	ShareholderRegistryCaller     // Read-only binding to the contract
	ShareholderRegistryTransactor // Write-only binding to the contract
	ShareholderRegistryFilterer   // Log filterer for contract events
}

// ShareholderRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShareholderRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareholderRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShareholderRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareholderRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShareholderRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareholderRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShareholderRegistrySession struct {
	Contract     *ShareholderRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ShareholderRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShareholderRegistryCallerSession struct {
	Contract *ShareholderRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ShareholderRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShareholderRegistryTransactorSession struct {
	Contract     *ShareholderRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ShareholderRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShareholderRegistryRaw struct {
	Contract *ShareholderRegistry // Generic contract binding to access the raw methods on
}

// ShareholderRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShareholderRegistryCallerRaw struct {
	Contract *ShareholderRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ShareholderRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShareholderRegistryTransactorRaw struct {
	Contract *ShareholderRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShareholderRegistry creates a new instance of ShareholderRegistry, bound to a specific deployed contract.
func NewShareholderRegistry(address common.Address, backend bind.ContractBackend) (*ShareholderRegistry, error) {
	contract, err := bindShareholderRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShareholderRegistry{ShareholderRegistryCaller: ShareholderRegistryCaller{contract: contract}, ShareholderRegistryTransactor: ShareholderRegistryTransactor{contract: contract}, ShareholderRegistryFilterer: ShareholderRegistryFilterer{contract: contract}}, nil
}

// NewShareholderRegistryCaller creates a new read-only instance of ShareholderRegistry, bound to a specific deployed contract.
func NewShareholderRegistryCaller(address common.Address, caller bind.ContractCaller) (*ShareholderRegistryCaller, error) {
	contract, err := bindShareholderRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShareholderRegistryCaller{contract: contract}, nil
}

// NewShareholderRegistryTransactor creates a new write-only instance of ShareholderRegistry, bound to a specific deployed contract.
func NewShareholderRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ShareholderRegistryTransactor, error) {
	contract, err := bindShareholderRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShareholderRegistryTransactor{contract: contract}, nil
}

// NewShareholderRegistryFilterer creates a new log filterer instance of ShareholderRegistry, bound to a specific deployed contract.
func NewShareholderRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ShareholderRegistryFilterer, error) {
	contract, err := bindShareholderRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShareholderRegistryFilterer{contract: contract}, nil
}

// bindShareholderRegistry binds a generic wrapper to an already deployed contract.
func bindShareholderRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShareholderRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShareholderRegistry *ShareholderRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShareholderRegistry.Contract.ShareholderRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShareholderRegistry *ShareholderRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.ShareholderRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShareholderRegistry *ShareholderRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.ShareholderRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShareholderRegistry *ShareholderRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShareholderRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShareholderRegistry *ShareholderRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShareholderRegistry *ShareholderRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.contract.Transact(opts, method, params...)
}

// CONTRIBUTORSTATUS is a free data retrieval call binding the contract method 0xefbc89bc.
//
// Solidity: function CONTRIBUTOR_STATUS() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCaller) CONTRIBUTORSTATUS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "CONTRIBUTOR_STATUS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTRIBUTORSTATUS is a free data retrieval call binding the contract method 0xefbc89bc.
//
// Solidity: function CONTRIBUTOR_STATUS() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistrySession) CONTRIBUTORSTATUS() ([32]byte, error) {
	return _ShareholderRegistry.Contract.CONTRIBUTORSTATUS(&_ShareholderRegistry.CallOpts)
}

// CONTRIBUTORSTATUS is a free data retrieval call binding the contract method 0xefbc89bc.
//
// Solidity: function CONTRIBUTOR_STATUS() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) CONTRIBUTORSTATUS() ([32]byte, error) {
	return _ShareholderRegistry.Contract.CONTRIBUTORSTATUS(&_ShareholderRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ShareholderRegistry.Contract.DEFAULTADMINROLE(&_ShareholderRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ShareholderRegistry.Contract.DEFAULTADMINROLE(&_ShareholderRegistry.CallOpts)
}

// INVESTORSTATUS is a free data retrieval call binding the contract method 0xfe4c7da8.
//
// Solidity: function INVESTOR_STATUS() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCaller) INVESTORSTATUS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "INVESTOR_STATUS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INVESTORSTATUS is a free data retrieval call binding the contract method 0xfe4c7da8.
//
// Solidity: function INVESTOR_STATUS() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistrySession) INVESTORSTATUS() ([32]byte, error) {
	return _ShareholderRegistry.Contract.INVESTORSTATUS(&_ShareholderRegistry.CallOpts)
}

// INVESTORSTATUS is a free data retrieval call binding the contract method 0xfe4c7da8.
//
// Solidity: function INVESTOR_STATUS() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) INVESTORSTATUS() ([32]byte, error) {
	return _ShareholderRegistry.Contract.INVESTORSTATUS(&_ShareholderRegistry.CallOpts)
}

// MANAGINGBOARDSTATUS is a free data retrieval call binding the contract method 0x06c1c03b.
//
// Solidity: function MANAGING_BOARD_STATUS() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCaller) MANAGINGBOARDSTATUS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "MANAGING_BOARD_STATUS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGINGBOARDSTATUS is a free data retrieval call binding the contract method 0x06c1c03b.
//
// Solidity: function MANAGING_BOARD_STATUS() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistrySession) MANAGINGBOARDSTATUS() ([32]byte, error) {
	return _ShareholderRegistry.Contract.MANAGINGBOARDSTATUS(&_ShareholderRegistry.CallOpts)
}

// MANAGINGBOARDSTATUS is a free data retrieval call binding the contract method 0x06c1c03b.
//
// Solidity: function MANAGING_BOARD_STATUS() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) MANAGINGBOARDSTATUS() ([32]byte, error) {
	return _ShareholderRegistry.Contract.MANAGINGBOARDSTATUS(&_ShareholderRegistry.CallOpts)
}

// SHAREHOLDERSTATUS is a free data retrieval call binding the contract method 0xed76fff3.
//
// Solidity: function SHAREHOLDER_STATUS() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCaller) SHAREHOLDERSTATUS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "SHAREHOLDER_STATUS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SHAREHOLDERSTATUS is a free data retrieval call binding the contract method 0xed76fff3.
//
// Solidity: function SHAREHOLDER_STATUS() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistrySession) SHAREHOLDERSTATUS() ([32]byte, error) {
	return _ShareholderRegistry.Contract.SHAREHOLDERSTATUS(&_ShareholderRegistry.CallOpts)
}

// SHAREHOLDERSTATUS is a free data retrieval call binding the contract method 0xed76fff3.
//
// Solidity: function SHAREHOLDER_STATUS() view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) SHAREHOLDERSTATUS() ([32]byte, error) {
	return _ShareholderRegistry.Contract.SHAREHOLDERSTATUS(&_ShareholderRegistry.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistryCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistrySession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ShareholderRegistry.Contract.Allowance(&_ShareholderRegistry.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ShareholderRegistry.Contract.Allowance(&_ShareholderRegistry.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistryCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistrySession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ShareholderRegistry.Contract.BalanceOf(&_ShareholderRegistry.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ShareholderRegistry.Contract.BalanceOf(&_ShareholderRegistry.CallOpts, account)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256 balance)
func (_ShareholderRegistry *ShareholderRegistryCaller) BalanceOfAt(opts *bind.CallOpts, account common.Address, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "balanceOfAt", account, snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256 balance)
func (_ShareholderRegistry *ShareholderRegistrySession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _ShareholderRegistry.Contract.BalanceOfAt(&_ShareholderRegistry.CallOpts, account, snapshotId)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256 balance)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _ShareholderRegistry.Contract.BalanceOfAt(&_ShareholderRegistry.CallOpts, account, snapshotId)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShareholderRegistry *ShareholderRegistryCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShareholderRegistry *ShareholderRegistrySession) Decimals() (uint8, error) {
	return _ShareholderRegistry.Contract.Decimals(&_ShareholderRegistry.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) Decimals() (uint8, error) {
	return _ShareholderRegistry.Contract.Decimals(&_ShareholderRegistry.CallOpts)
}

// GetBalanceAndStatusAt is a free data retrieval call binding the contract method 0x4baa52cc.
//
// Solidity: function getBalanceAndStatusAt(address account, uint256 snapshotId) view returns(uint256, bytes32)
func (_ShareholderRegistry *ShareholderRegistryCaller) GetBalanceAndStatusAt(opts *bind.CallOpts, account common.Address, snapshotId *big.Int) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "getBalanceAndStatusAt", account, snapshotId)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetBalanceAndStatusAt is a free data retrieval call binding the contract method 0x4baa52cc.
//
// Solidity: function getBalanceAndStatusAt(address account, uint256 snapshotId) view returns(uint256, bytes32)
func (_ShareholderRegistry *ShareholderRegistrySession) GetBalanceAndStatusAt(account common.Address, snapshotId *big.Int) (*big.Int, [32]byte, error) {
	return _ShareholderRegistry.Contract.GetBalanceAndStatusAt(&_ShareholderRegistry.CallOpts, account, snapshotId)
}

// GetBalanceAndStatusAt is a free data retrieval call binding the contract method 0x4baa52cc.
//
// Solidity: function getBalanceAndStatusAt(address account, uint256 snapshotId) view returns(uint256, bytes32)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) GetBalanceAndStatusAt(account common.Address, snapshotId *big.Int) (*big.Int, [32]byte, error) {
	return _ShareholderRegistry.Contract.GetBalanceAndStatusAt(&_ShareholderRegistry.CallOpts, account, snapshotId)
}

// GetCurrentSnapshotId is a free data retrieval call binding the contract method 0x5439ad86.
//
// Solidity: function getCurrentSnapshotId() view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistryCaller) GetCurrentSnapshotId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "getCurrentSnapshotId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentSnapshotId is a free data retrieval call binding the contract method 0x5439ad86.
//
// Solidity: function getCurrentSnapshotId() view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistrySession) GetCurrentSnapshotId() (*big.Int, error) {
	return _ShareholderRegistry.Contract.GetCurrentSnapshotId(&_ShareholderRegistry.CallOpts)
}

// GetCurrentSnapshotId is a free data retrieval call binding the contract method 0x5439ad86.
//
// Solidity: function getCurrentSnapshotId() view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) GetCurrentSnapshotId() (*big.Int, error) {
	return _ShareholderRegistry.Contract.GetCurrentSnapshotId(&_ShareholderRegistry.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ShareholderRegistry.Contract.GetRoleAdmin(&_ShareholderRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ShareholderRegistry.Contract.GetRoleAdmin(&_ShareholderRegistry.CallOpts, role)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address account) view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCaller) GetStatus(opts *bind.CallOpts, account common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "getStatus", account)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address account) view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistrySession) GetStatus(account common.Address) ([32]byte, error) {
	return _ShareholderRegistry.Contract.GetStatus(&_ShareholderRegistry.CallOpts, account)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address account) view returns(bytes32)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) GetStatus(account common.Address) ([32]byte, error) {
	return _ShareholderRegistry.Contract.GetStatus(&_ShareholderRegistry.CallOpts, account)
}

// GetStatusAt is a free data retrieval call binding the contract method 0x3397cc23.
//
// Solidity: function getStatusAt(address account, uint256 snapshotId) view returns(bytes32 status)
func (_ShareholderRegistry *ShareholderRegistryCaller) GetStatusAt(opts *bind.CallOpts, account common.Address, snapshotId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "getStatusAt", account, snapshotId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStatusAt is a free data retrieval call binding the contract method 0x3397cc23.
//
// Solidity: function getStatusAt(address account, uint256 snapshotId) view returns(bytes32 status)
func (_ShareholderRegistry *ShareholderRegistrySession) GetStatusAt(account common.Address, snapshotId *big.Int) ([32]byte, error) {
	return _ShareholderRegistry.Contract.GetStatusAt(&_ShareholderRegistry.CallOpts, account, snapshotId)
}

// GetStatusAt is a free data retrieval call binding the contract method 0x3397cc23.
//
// Solidity: function getStatusAt(address account, uint256 snapshotId) view returns(bytes32 status)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) GetStatusAt(account common.Address, snapshotId *big.Int) ([32]byte, error) {
	return _ShareholderRegistry.Contract.GetStatusAt(&_ShareholderRegistry.CallOpts, account, snapshotId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ShareholderRegistry *ShareholderRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ShareholderRegistry *ShareholderRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ShareholderRegistry.Contract.HasRole(&_ShareholderRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ShareholderRegistry.Contract.HasRole(&_ShareholderRegistry.CallOpts, role, account)
}

// IsAtLeast is a free data retrieval call binding the contract method 0x0ea018f7.
//
// Solidity: function isAtLeast(bytes32 status, address account) view returns(bool)
func (_ShareholderRegistry *ShareholderRegistryCaller) IsAtLeast(opts *bind.CallOpts, status [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "isAtLeast", status, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAtLeast is a free data retrieval call binding the contract method 0x0ea018f7.
//
// Solidity: function isAtLeast(bytes32 status, address account) view returns(bool)
func (_ShareholderRegistry *ShareholderRegistrySession) IsAtLeast(status [32]byte, account common.Address) (bool, error) {
	return _ShareholderRegistry.Contract.IsAtLeast(&_ShareholderRegistry.CallOpts, status, account)
}

// IsAtLeast is a free data retrieval call binding the contract method 0x0ea018f7.
//
// Solidity: function isAtLeast(bytes32 status, address account) view returns(bool)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) IsAtLeast(status [32]byte, account common.Address) (bool, error) {
	return _ShareholderRegistry.Contract.IsAtLeast(&_ShareholderRegistry.CallOpts, status, account)
}

// IsAtLeastAt is a free data retrieval call binding the contract method 0x4b4b35f6.
//
// Solidity: function isAtLeastAt(bytes32 status, address account, uint256 snapshotId) view returns(bool)
func (_ShareholderRegistry *ShareholderRegistryCaller) IsAtLeastAt(opts *bind.CallOpts, status [32]byte, account common.Address, snapshotId *big.Int) (bool, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "isAtLeastAt", status, account, snapshotId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAtLeastAt is a free data retrieval call binding the contract method 0x4b4b35f6.
//
// Solidity: function isAtLeastAt(bytes32 status, address account, uint256 snapshotId) view returns(bool)
func (_ShareholderRegistry *ShareholderRegistrySession) IsAtLeastAt(status [32]byte, account common.Address, snapshotId *big.Int) (bool, error) {
	return _ShareholderRegistry.Contract.IsAtLeastAt(&_ShareholderRegistry.CallOpts, status, account, snapshotId)
}

// IsAtLeastAt is a free data retrieval call binding the contract method 0x4b4b35f6.
//
// Solidity: function isAtLeastAt(bytes32 status, address account, uint256 snapshotId) view returns(bool)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) IsAtLeastAt(status [32]byte, account common.Address, snapshotId *big.Int) (bool, error) {
	return _ShareholderRegistry.Contract.IsAtLeastAt(&_ShareholderRegistry.CallOpts, status, account, snapshotId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShareholderRegistry *ShareholderRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShareholderRegistry *ShareholderRegistrySession) Name() (string, error) {
	return _ShareholderRegistry.Contract.Name(&_ShareholderRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) Name() (string, error) {
	return _ShareholderRegistry.Contract.Name(&_ShareholderRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShareholderRegistry *ShareholderRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShareholderRegistry *ShareholderRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ShareholderRegistry.Contract.SupportsInterface(&_ShareholderRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ShareholderRegistry.Contract.SupportsInterface(&_ShareholderRegistry.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShareholderRegistry *ShareholderRegistryCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShareholderRegistry *ShareholderRegistrySession) Symbol() (string, error) {
	return _ShareholderRegistry.Contract.Symbol(&_ShareholderRegistry.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) Symbol() (string, error) {
	return _ShareholderRegistry.Contract.Symbol(&_ShareholderRegistry.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistryCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistrySession) TotalSupply() (*big.Int, error) {
	return _ShareholderRegistry.Contract.TotalSupply(&_ShareholderRegistry.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) TotalSupply() (*big.Int, error) {
	return _ShareholderRegistry.Contract.TotalSupply(&_ShareholderRegistry.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistryCaller) TotalSupplyAt(opts *bind.CallOpts, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShareholderRegistry.contract.Call(opts, &out, "totalSupplyAt", snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistrySession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _ShareholderRegistry.Contract.TotalSupplyAt(&_ShareholderRegistry.CallOpts, snapshotId)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_ShareholderRegistry *ShareholderRegistryCallerSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _ShareholderRegistry.Contract.TotalSupplyAt(&_ShareholderRegistry.CallOpts, snapshotId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ShareholderRegistry *ShareholderRegistryTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ShareholderRegistry *ShareholderRegistrySession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.Approve(&_ShareholderRegistry.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.Approve(&_ShareholderRegistry.TransactOpts, spender, amount)
}

// BatchTransferFromDAO is a paid mutator transaction binding the contract method 0x35650b5b.
//
// Solidity: function batchTransferFromDAO(address[] recipients) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactor) BatchTransferFromDAO(opts *bind.TransactOpts, recipients []common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "batchTransferFromDAO", recipients)
}

// BatchTransferFromDAO is a paid mutator transaction binding the contract method 0x35650b5b.
//
// Solidity: function batchTransferFromDAO(address[] recipients) returns()
func (_ShareholderRegistry *ShareholderRegistrySession) BatchTransferFromDAO(recipients []common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.BatchTransferFromDAO(&_ShareholderRegistry.TransactOpts, recipients)
}

// BatchTransferFromDAO is a paid mutator transaction binding the contract method 0x35650b5b.
//
// Solidity: function batchTransferFromDAO(address[] recipients) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) BatchTransferFromDAO(recipients []common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.BatchTransferFromDAO(&_ShareholderRegistry.TransactOpts, recipients)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_ShareholderRegistry *ShareholderRegistrySession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.Burn(&_ShareholderRegistry.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.Burn(&_ShareholderRegistry.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ShareholderRegistry *ShareholderRegistryTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ShareholderRegistry *ShareholderRegistrySession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.DecreaseAllowance(&_ShareholderRegistry.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.DecreaseAllowance(&_ShareholderRegistry.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ShareholderRegistry *ShareholderRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.GrantRole(&_ShareholderRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.GrantRole(&_ShareholderRegistry.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ShareholderRegistry *ShareholderRegistryTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ShareholderRegistry *ShareholderRegistrySession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.IncreaseAllowance(&_ShareholderRegistry.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.IncreaseAllowance(&_ShareholderRegistry.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name, string symbol) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "initialize", name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name, string symbol) returns()
func (_ShareholderRegistry *ShareholderRegistrySession) Initialize(name string, symbol string) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.Initialize(&_ShareholderRegistry.TransactOpts, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name, string symbol) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) Initialize(name string, symbol string) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.Initialize(&_ShareholderRegistry.TransactOpts, name, symbol)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ShareholderRegistry *ShareholderRegistrySession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.Mint(&_ShareholderRegistry.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.Mint(&_ShareholderRegistry.TransactOpts, account, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ShareholderRegistry *ShareholderRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.RenounceRole(&_ShareholderRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.RenounceRole(&_ShareholderRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ShareholderRegistry *ShareholderRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.RevokeRole(&_ShareholderRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.RevokeRole(&_ShareholderRegistry.TransactOpts, role, account)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8bdbbb75.
//
// Solidity: function setStatus(bytes32 status, address account) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactor) SetStatus(opts *bind.TransactOpts, status [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "setStatus", status, account)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8bdbbb75.
//
// Solidity: function setStatus(bytes32 status, address account) returns()
func (_ShareholderRegistry *ShareholderRegistrySession) SetStatus(status [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.SetStatus(&_ShareholderRegistry.TransactOpts, status, account)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8bdbbb75.
//
// Solidity: function setStatus(bytes32 status, address account) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) SetStatus(status [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.SetStatus(&_ShareholderRegistry.TransactOpts, status, account)
}

// SetVoting is a paid mutator transaction binding the contract method 0xc4d07951.
//
// Solidity: function setVoting(address voting) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactor) SetVoting(opts *bind.TransactOpts, voting common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "setVoting", voting)
}

// SetVoting is a paid mutator transaction binding the contract method 0xc4d07951.
//
// Solidity: function setVoting(address voting) returns()
func (_ShareholderRegistry *ShareholderRegistrySession) SetVoting(voting common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.SetVoting(&_ShareholderRegistry.TransactOpts, voting)
}

// SetVoting is a paid mutator transaction binding the contract method 0xc4d07951.
//
// Solidity: function setVoting(address voting) returns()
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) SetVoting(voting common.Address) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.SetVoting(&_ShareholderRegistry.TransactOpts, voting)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256)
func (_ShareholderRegistry *ShareholderRegistryTransactor) Snapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "snapshot")
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256)
func (_ShareholderRegistry *ShareholderRegistrySession) Snapshot() (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.Snapshot(&_ShareholderRegistry.TransactOpts)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256)
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) Snapshot() (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.Snapshot(&_ShareholderRegistry.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ShareholderRegistry *ShareholderRegistryTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ShareholderRegistry *ShareholderRegistrySession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.Transfer(&_ShareholderRegistry.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.Transfer(&_ShareholderRegistry.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ShareholderRegistry *ShareholderRegistryTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ShareholderRegistry *ShareholderRegistrySession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.TransferFrom(&_ShareholderRegistry.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ShareholderRegistry *ShareholderRegistryTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShareholderRegistry.Contract.TransferFrom(&_ShareholderRegistry.TransactOpts, from, to, amount)
}

// ShareholderRegistryApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ShareholderRegistry contract.
type ShareholderRegistryApprovalIterator struct {
	Event *ShareholderRegistryApproval // Event containing the contract specifics and raw log

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
func (it *ShareholderRegistryApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareholderRegistryApproval)
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
		it.Event = new(ShareholderRegistryApproval)
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
func (it *ShareholderRegistryApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareholderRegistryApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareholderRegistryApproval represents a Approval event raised by the ShareholderRegistry contract.
type ShareholderRegistryApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ShareholderRegistry *ShareholderRegistryFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ShareholderRegistryApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ShareholderRegistry.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ShareholderRegistryApprovalIterator{contract: _ShareholderRegistry.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ShareholderRegistry *ShareholderRegistryFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ShareholderRegistryApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ShareholderRegistry.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareholderRegistryApproval)
				if err := _ShareholderRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ShareholderRegistry *ShareholderRegistryFilterer) ParseApproval(log types.Log) (*ShareholderRegistryApproval, error) {
	event := new(ShareholderRegistryApproval)
	if err := _ShareholderRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShareholderRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ShareholderRegistry contract.
type ShareholderRegistryInitializedIterator struct {
	Event *ShareholderRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ShareholderRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareholderRegistryInitialized)
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
		it.Event = new(ShareholderRegistryInitialized)
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
func (it *ShareholderRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareholderRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareholderRegistryInitialized represents a Initialized event raised by the ShareholderRegistry contract.
type ShareholderRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ShareholderRegistry *ShareholderRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ShareholderRegistryInitializedIterator, error) {

	logs, sub, err := _ShareholderRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ShareholderRegistryInitializedIterator{contract: _ShareholderRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ShareholderRegistry *ShareholderRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ShareholderRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ShareholderRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareholderRegistryInitialized)
				if err := _ShareholderRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ShareholderRegistry *ShareholderRegistryFilterer) ParseInitialized(log types.Log) (*ShareholderRegistryInitialized, error) {
	event := new(ShareholderRegistryInitialized)
	if err := _ShareholderRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShareholderRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ShareholderRegistry contract.
type ShareholderRegistryRoleAdminChangedIterator struct {
	Event *ShareholderRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ShareholderRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareholderRegistryRoleAdminChanged)
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
		it.Event = new(ShareholderRegistryRoleAdminChanged)
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
func (it *ShareholderRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareholderRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareholderRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the ShareholderRegistry contract.
type ShareholderRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ShareholderRegistry *ShareholderRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ShareholderRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ShareholderRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ShareholderRegistryRoleAdminChangedIterator{contract: _ShareholderRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ShareholderRegistry *ShareholderRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ShareholderRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ShareholderRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareholderRegistryRoleAdminChanged)
				if err := _ShareholderRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ShareholderRegistry *ShareholderRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*ShareholderRegistryRoleAdminChanged, error) {
	event := new(ShareholderRegistryRoleAdminChanged)
	if err := _ShareholderRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShareholderRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ShareholderRegistry contract.
type ShareholderRegistryRoleGrantedIterator struct {
	Event *ShareholderRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *ShareholderRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareholderRegistryRoleGranted)
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
		it.Event = new(ShareholderRegistryRoleGranted)
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
func (it *ShareholderRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareholderRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareholderRegistryRoleGranted represents a RoleGranted event raised by the ShareholderRegistry contract.
type ShareholderRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShareholderRegistry *ShareholderRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ShareholderRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _ShareholderRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ShareholderRegistryRoleGrantedIterator{contract: _ShareholderRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShareholderRegistry *ShareholderRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ShareholderRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ShareholderRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareholderRegistryRoleGranted)
				if err := _ShareholderRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ShareholderRegistry *ShareholderRegistryFilterer) ParseRoleGranted(log types.Log) (*ShareholderRegistryRoleGranted, error) {
	event := new(ShareholderRegistryRoleGranted)
	if err := _ShareholderRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShareholderRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ShareholderRegistry contract.
type ShareholderRegistryRoleRevokedIterator struct {
	Event *ShareholderRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ShareholderRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareholderRegistryRoleRevoked)
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
		it.Event = new(ShareholderRegistryRoleRevoked)
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
func (it *ShareholderRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareholderRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareholderRegistryRoleRevoked represents a RoleRevoked event raised by the ShareholderRegistry contract.
type ShareholderRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShareholderRegistry *ShareholderRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ShareholderRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _ShareholderRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ShareholderRegistryRoleRevokedIterator{contract: _ShareholderRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShareholderRegistry *ShareholderRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ShareholderRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ShareholderRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareholderRegistryRoleRevoked)
				if err := _ShareholderRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ShareholderRegistry *ShareholderRegistryFilterer) ParseRoleRevoked(log types.Log) (*ShareholderRegistryRoleRevoked, error) {
	event := new(ShareholderRegistryRoleRevoked)
	if err := _ShareholderRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShareholderRegistrySnapshotIterator is returned from FilterSnapshot and is used to iterate over the raw logs and unpacked data for Snapshot events raised by the ShareholderRegistry contract.
type ShareholderRegistrySnapshotIterator struct {
	Event *ShareholderRegistrySnapshot // Event containing the contract specifics and raw log

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
func (it *ShareholderRegistrySnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareholderRegistrySnapshot)
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
		it.Event = new(ShareholderRegistrySnapshot)
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
func (it *ShareholderRegistrySnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareholderRegistrySnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareholderRegistrySnapshot represents a Snapshot event raised by the ShareholderRegistry contract.
type ShareholderRegistrySnapshot struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSnapshot is a free log retrieval operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_ShareholderRegistry *ShareholderRegistryFilterer) FilterSnapshot(opts *bind.FilterOpts) (*ShareholderRegistrySnapshotIterator, error) {

	logs, sub, err := _ShareholderRegistry.contract.FilterLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return &ShareholderRegistrySnapshotIterator{contract: _ShareholderRegistry.contract, event: "Snapshot", logs: logs, sub: sub}, nil
}

// WatchSnapshot is a free log subscription operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_ShareholderRegistry *ShareholderRegistryFilterer) WatchSnapshot(opts *bind.WatchOpts, sink chan<- *ShareholderRegistrySnapshot) (event.Subscription, error) {

	logs, sub, err := _ShareholderRegistry.contract.WatchLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareholderRegistrySnapshot)
				if err := _ShareholderRegistry.contract.UnpackLog(event, "Snapshot", log); err != nil {
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
func (_ShareholderRegistry *ShareholderRegistryFilterer) ParseSnapshot(log types.Log) (*ShareholderRegistrySnapshot, error) {
	event := new(ShareholderRegistrySnapshot)
	if err := _ShareholderRegistry.contract.UnpackLog(event, "Snapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShareholderRegistryStatusChangedIterator is returned from FilterStatusChanged and is used to iterate over the raw logs and unpacked data for StatusChanged events raised by the ShareholderRegistry contract.
type ShareholderRegistryStatusChangedIterator struct {
	Event *ShareholderRegistryStatusChanged // Event containing the contract specifics and raw log

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
func (it *ShareholderRegistryStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareholderRegistryStatusChanged)
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
		it.Event = new(ShareholderRegistryStatusChanged)
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
func (it *ShareholderRegistryStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareholderRegistryStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareholderRegistryStatusChanged represents a StatusChanged event raised by the ShareholderRegistry contract.
type ShareholderRegistryStatusChanged struct {
	Account  common.Address
	Previous [32]byte
	Current  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStatusChanged is a free log retrieval operation binding the contract event 0x63d1f3ab699618c7d6d972a3c28f0d08fa090bebc96a34c3a30e5fcda6397ab2.
//
// Solidity: event StatusChanged(address indexed account, bytes32 previous, bytes32 current)
func (_ShareholderRegistry *ShareholderRegistryFilterer) FilterStatusChanged(opts *bind.FilterOpts, account []common.Address) (*ShareholderRegistryStatusChangedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ShareholderRegistry.contract.FilterLogs(opts, "StatusChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return &ShareholderRegistryStatusChangedIterator{contract: _ShareholderRegistry.contract, event: "StatusChanged", logs: logs, sub: sub}, nil
}

// WatchStatusChanged is a free log subscription operation binding the contract event 0x63d1f3ab699618c7d6d972a3c28f0d08fa090bebc96a34c3a30e5fcda6397ab2.
//
// Solidity: event StatusChanged(address indexed account, bytes32 previous, bytes32 current)
func (_ShareholderRegistry *ShareholderRegistryFilterer) WatchStatusChanged(opts *bind.WatchOpts, sink chan<- *ShareholderRegistryStatusChanged, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ShareholderRegistry.contract.WatchLogs(opts, "StatusChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareholderRegistryStatusChanged)
				if err := _ShareholderRegistry.contract.UnpackLog(event, "StatusChanged", log); err != nil {
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

// ParseStatusChanged is a log parse operation binding the contract event 0x63d1f3ab699618c7d6d972a3c28f0d08fa090bebc96a34c3a30e5fcda6397ab2.
//
// Solidity: event StatusChanged(address indexed account, bytes32 previous, bytes32 current)
func (_ShareholderRegistry *ShareholderRegistryFilterer) ParseStatusChanged(log types.Log) (*ShareholderRegistryStatusChanged, error) {
	event := new(ShareholderRegistryStatusChanged)
	if err := _ShareholderRegistry.contract.UnpackLog(event, "StatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShareholderRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ShareholderRegistry contract.
type ShareholderRegistryTransferIterator struct {
	Event *ShareholderRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *ShareholderRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareholderRegistryTransfer)
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
		it.Event = new(ShareholderRegistryTransfer)
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
func (it *ShareholderRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareholderRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareholderRegistryTransfer represents a Transfer event raised by the ShareholderRegistry contract.
type ShareholderRegistryTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ShareholderRegistry *ShareholderRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ShareholderRegistryTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShareholderRegistry.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShareholderRegistryTransferIterator{contract: _ShareholderRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ShareholderRegistry *ShareholderRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ShareholderRegistryTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShareholderRegistry.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareholderRegistryTransfer)
				if err := _ShareholderRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ShareholderRegistry *ShareholderRegistryFilterer) ParseTransfer(log types.Log) (*ShareholderRegistryTransfer, error) {
	event := new(ShareholderRegistryTransfer)
	if err := _ShareholderRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
