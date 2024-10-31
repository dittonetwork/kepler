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

// IValidatorDataValidatorData is an auto generated low-level Go binding around an user-defined struct.
type IValidatorDataValidatorData struct {
	Operator common.Address
	Stake    *big.Int
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addCollateralOracle\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addToWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorStake\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorSet\",\"inputs\":[],\"outputs\":[{\"name\":\"validatorsData\",\"type\":\"tuple[]\",\"internalType\":\"structIValidatorData.ValidatorData[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isActiveOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedVault\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorIsVaultRegistry\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeCollateralOracle\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeFromWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinStake\",\"inputs\":[{\"name\":\"minStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSubnetworks\",\"inputs\":[{\"name\":\"subnetworks\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unRegisterOperator\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CollateralOracleAdded\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralOracleRemoved\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinStakeChanged\",\"inputs\":[{\"name\":\"oldMinstake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newMinstake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUnregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubnetworksChanged\",\"inputs\":[{\"name\":\"oldSubnetworks\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newSubnetworks\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultRegistered\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultUnregistered\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CollateralAlreadyUseForVaults\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CollateralAlreadyUsedOtherOracle\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CollateralWithOracleAlreadyUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyAddressCollateralOrOracle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSubnetworks\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSupportedCollateral\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotVault\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegisterVault\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotOptedIn\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotOptedInVault\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorStakeIsSmall\",\"inputs\":[{\"name\":\"minStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
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

// GetOperatorStake is a free data retrieval call binding the contract method 0xe4e88de8.
//
// Solidity: function getOperatorStake(address operator) view returns(uint256)
func (_Contracts *ContractsCaller) GetOperatorStake(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getOperatorStake", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorStake is a free data retrieval call binding the contract method 0xe4e88de8.
//
// Solidity: function getOperatorStake(address operator) view returns(uint256)
func (_Contracts *ContractsSession) GetOperatorStake(operator common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetOperatorStake(&_Contracts.CallOpts, operator)
}

// GetOperatorStake is a free data retrieval call binding the contract method 0xe4e88de8.
//
// Solidity: function getOperatorStake(address operator) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetOperatorStake(operator common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetOperatorStake(&_Contracts.CallOpts, operator)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0xcf331250.
//
// Solidity: function getValidatorSet() view returns((address,uint256)[] validatorsData)
func (_Contracts *ContractsCaller) GetValidatorSet(opts *bind.CallOpts) ([]IValidatorDataValidatorData, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getValidatorSet")

	if err != nil {
		return *new([]IValidatorDataValidatorData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IValidatorDataValidatorData)).(*[]IValidatorDataValidatorData)

	return out0, err

}

// GetValidatorSet is a free data retrieval call binding the contract method 0xcf331250.
//
// Solidity: function getValidatorSet() view returns((address,uint256)[] validatorsData)
func (_Contracts *ContractsSession) GetValidatorSet() ([]IValidatorDataValidatorData, error) {
	return _Contracts.Contract.GetValidatorSet(&_Contracts.CallOpts)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0xcf331250.
//
// Solidity: function getValidatorSet() view returns((address,uint256)[] validatorsData)
func (_Contracts *ContractsCallerSession) GetValidatorSet() ([]IValidatorDataValidatorData, error) {
	return _Contracts.Contract.GetValidatorSet(&_Contracts.CallOpts)
}

// IsActiveOperator is a free data retrieval call binding the contract method 0x3367cca5.
//
// Solidity: function isActiveOperator(address operator) view returns(bool)
func (_Contracts *ContractsCaller) IsActiveOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isActiveOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveOperator is a free data retrieval call binding the contract method 0x3367cca5.
//
// Solidity: function isActiveOperator(address operator) view returns(bool)
func (_Contracts *ContractsSession) IsActiveOperator(operator common.Address) (bool, error) {
	return _Contracts.Contract.IsActiveOperator(&_Contracts.CallOpts, operator)
}

// IsActiveOperator is a free data retrieval call binding the contract method 0x3367cca5.
//
// Solidity: function isActiveOperator(address operator) view returns(bool)
func (_Contracts *ContractsCallerSession) IsActiveOperator(operator common.Address) (bool, error) {
	return _Contracts.Contract.IsActiveOperator(&_Contracts.CallOpts, operator)
}

// IsSupportedVault is a free data retrieval call binding the contract method 0xb2b4ec48.
//
// Solidity: function isSupportedVault(address vault) view returns(bool)
func (_Contracts *ContractsCaller) IsSupportedVault(opts *bind.CallOpts, vault common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isSupportedVault", vault)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedVault is a free data retrieval call binding the contract method 0xb2b4ec48.
//
// Solidity: function isSupportedVault(address vault) view returns(bool)
func (_Contracts *ContractsSession) IsSupportedVault(vault common.Address) (bool, error) {
	return _Contracts.Contract.IsSupportedVault(&_Contracts.CallOpts, vault)
}

// IsSupportedVault is a free data retrieval call binding the contract method 0xb2b4ec48.
//
// Solidity: function isSupportedVault(address vault) view returns(bool)
func (_Contracts *ContractsCallerSession) IsSupportedVault(vault common.Address) (bool, error) {
	return _Contracts.Contract.IsSupportedVault(&_Contracts.CallOpts, vault)
}

// OperatorIsVaultRegistry is a free data retrieval call binding the contract method 0x2a42a230.
//
// Solidity: function operatorIsVaultRegistry(address operator, address vault) view returns(bool)
func (_Contracts *ContractsCaller) OperatorIsVaultRegistry(opts *bind.CallOpts, operator common.Address, vault common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "operatorIsVaultRegistry", operator, vault)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorIsVaultRegistry is a free data retrieval call binding the contract method 0x2a42a230.
//
// Solidity: function operatorIsVaultRegistry(address operator, address vault) view returns(bool)
func (_Contracts *ContractsSession) OperatorIsVaultRegistry(operator common.Address, vault common.Address) (bool, error) {
	return _Contracts.Contract.OperatorIsVaultRegistry(&_Contracts.CallOpts, operator, vault)
}

// OperatorIsVaultRegistry is a free data retrieval call binding the contract method 0x2a42a230.
//
// Solidity: function operatorIsVaultRegistry(address operator, address vault) view returns(bool)
func (_Contracts *ContractsCallerSession) OperatorIsVaultRegistry(operator common.Address, vault common.Address) (bool, error) {
	return _Contracts.Contract.OperatorIsVaultRegistry(&_Contracts.CallOpts, operator, vault)
}

// AddCollateralOracle is a paid mutator transaction binding the contract method 0x43d92408.
//
// Solidity: function addCollateralOracle(address collateral, address oracle) returns()
func (_Contracts *ContractsTransactor) AddCollateralOracle(opts *bind.TransactOpts, collateral common.Address, oracle common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addCollateralOracle", collateral, oracle)
}

// AddCollateralOracle is a paid mutator transaction binding the contract method 0x43d92408.
//
// Solidity: function addCollateralOracle(address collateral, address oracle) returns()
func (_Contracts *ContractsSession) AddCollateralOracle(collateral common.Address, oracle common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddCollateralOracle(&_Contracts.TransactOpts, collateral, oracle)
}

// AddCollateralOracle is a paid mutator transaction binding the contract method 0x43d92408.
//
// Solidity: function addCollateralOracle(address collateral, address oracle) returns()
func (_Contracts *ContractsTransactorSession) AddCollateralOracle(collateral common.Address, oracle common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddCollateralOracle(&_Contracts.TransactOpts, collateral, oracle)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address operator) returns()
func (_Contracts *ContractsTransactor) AddToWhitelist(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addToWhitelist", operator)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address operator) returns()
func (_Contracts *ContractsSession) AddToWhitelist(operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddToWhitelist(&_Contracts.TransactOpts, operator)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address operator) returns()
func (_Contracts *ContractsTransactorSession) AddToWhitelist(operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddToWhitelist(&_Contracts.TransactOpts, operator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address vault) returns()
func (_Contracts *ContractsTransactor) RegisterOperator(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerOperator", vault)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address vault) returns()
func (_Contracts *ContractsSession) RegisterOperator(vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterOperator(&_Contracts.TransactOpts, vault)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address vault) returns()
func (_Contracts *ContractsTransactorSession) RegisterOperator(vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterOperator(&_Contracts.TransactOpts, vault)
}

// RemoveCollateralOracle is a paid mutator transaction binding the contract method 0xaf3049fa.
//
// Solidity: function removeCollateralOracle(address collateral) returns()
func (_Contracts *ContractsTransactor) RemoveCollateralOracle(opts *bind.TransactOpts, collateral common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removeCollateralOracle", collateral)
}

// RemoveCollateralOracle is a paid mutator transaction binding the contract method 0xaf3049fa.
//
// Solidity: function removeCollateralOracle(address collateral) returns()
func (_Contracts *ContractsSession) RemoveCollateralOracle(collateral common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveCollateralOracle(&_Contracts.TransactOpts, collateral)
}

// RemoveCollateralOracle is a paid mutator transaction binding the contract method 0xaf3049fa.
//
// Solidity: function removeCollateralOracle(address collateral) returns()
func (_Contracts *ContractsTransactorSession) RemoveCollateralOracle(collateral common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveCollateralOracle(&_Contracts.TransactOpts, collateral)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address operator) returns()
func (_Contracts *ContractsTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removeFromWhitelist", operator)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address operator) returns()
func (_Contracts *ContractsSession) RemoveFromWhitelist(operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveFromWhitelist(&_Contracts.TransactOpts, operator)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address operator) returns()
func (_Contracts *ContractsTransactorSession) RemoveFromWhitelist(operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveFromWhitelist(&_Contracts.TransactOpts, operator)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 minStake) returns()
func (_Contracts *ContractsTransactor) SetMinStake(opts *bind.TransactOpts, minStake *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMinStake", minStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 minStake) returns()
func (_Contracts *ContractsSession) SetMinStake(minStake *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMinStake(&_Contracts.TransactOpts, minStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 minStake) returns()
func (_Contracts *ContractsTransactorSession) SetMinStake(minStake *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMinStake(&_Contracts.TransactOpts, minStake)
}

// SetSubnetworks is a paid mutator transaction binding the contract method 0x1bc70105.
//
// Solidity: function setSubnetworks(uint256 subnetworks) returns()
func (_Contracts *ContractsTransactor) SetSubnetworks(opts *bind.TransactOpts, subnetworks *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setSubnetworks", subnetworks)
}

// SetSubnetworks is a paid mutator transaction binding the contract method 0x1bc70105.
//
// Solidity: function setSubnetworks(uint256 subnetworks) returns()
func (_Contracts *ContractsSession) SetSubnetworks(subnetworks *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetSubnetworks(&_Contracts.TransactOpts, subnetworks)
}

// SetSubnetworks is a paid mutator transaction binding the contract method 0x1bc70105.
//
// Solidity: function setSubnetworks(uint256 subnetworks) returns()
func (_Contracts *ContractsTransactorSession) SetSubnetworks(subnetworks *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetSubnetworks(&_Contracts.TransactOpts, subnetworks)
}

// UnRegisterOperator is a paid mutator transaction binding the contract method 0xdaf5c2e5.
//
// Solidity: function unRegisterOperator(address vault) returns()
func (_Contracts *ContractsTransactor) UnRegisterOperator(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unRegisterOperator", vault)
}

// UnRegisterOperator is a paid mutator transaction binding the contract method 0xdaf5c2e5.
//
// Solidity: function unRegisterOperator(address vault) returns()
func (_Contracts *ContractsSession) UnRegisterOperator(vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnRegisterOperator(&_Contracts.TransactOpts, vault)
}

// UnRegisterOperator is a paid mutator transaction binding the contract method 0xdaf5c2e5.
//
// Solidity: function unRegisterOperator(address vault) returns()
func (_Contracts *ContractsTransactorSession) UnRegisterOperator(vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnRegisterOperator(&_Contracts.TransactOpts, vault)
}

// ContractsCollateralOracleAddedIterator is returned from FilterCollateralOracleAdded and is used to iterate over the raw logs and unpacked data for CollateralOracleAdded events raised by the Contracts contract.
type ContractsCollateralOracleAddedIterator struct {
	Event *ContractsCollateralOracleAdded // Event containing the contract specifics and raw log

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
func (it *ContractsCollateralOracleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCollateralOracleAdded)
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
		it.Event = new(ContractsCollateralOracleAdded)
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
func (it *ContractsCollateralOracleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCollateralOracleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCollateralOracleAdded represents a CollateralOracleAdded event raised by the Contracts contract.
type ContractsCollateralOracleAdded struct {
	Collateral common.Address
	Oracle     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollateralOracleAdded is a free log retrieval operation binding the contract event 0xf9f707881258111d41a59e99aa6ea48067b66fa1da9f06c9e2f0fd42ea1376db.
//
// Solidity: event CollateralOracleAdded(address indexed collateral, address indexed oracle)
func (_Contracts *ContractsFilterer) FilterCollateralOracleAdded(opts *bind.FilterOpts, collateral []common.Address, oracle []common.Address) (*ContractsCollateralOracleAddedIterator, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CollateralOracleAdded", collateralRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &ContractsCollateralOracleAddedIterator{contract: _Contracts.contract, event: "CollateralOracleAdded", logs: logs, sub: sub}, nil
}

// WatchCollateralOracleAdded is a free log subscription operation binding the contract event 0xf9f707881258111d41a59e99aa6ea48067b66fa1da9f06c9e2f0fd42ea1376db.
//
// Solidity: event CollateralOracleAdded(address indexed collateral, address indexed oracle)
func (_Contracts *ContractsFilterer) WatchCollateralOracleAdded(opts *bind.WatchOpts, sink chan<- *ContractsCollateralOracleAdded, collateral []common.Address, oracle []common.Address) (event.Subscription, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CollateralOracleAdded", collateralRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCollateralOracleAdded)
				if err := _Contracts.contract.UnpackLog(event, "CollateralOracleAdded", log); err != nil {
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

// ParseCollateralOracleAdded is a log parse operation binding the contract event 0xf9f707881258111d41a59e99aa6ea48067b66fa1da9f06c9e2f0fd42ea1376db.
//
// Solidity: event CollateralOracleAdded(address indexed collateral, address indexed oracle)
func (_Contracts *ContractsFilterer) ParseCollateralOracleAdded(log types.Log) (*ContractsCollateralOracleAdded, error) {
	event := new(ContractsCollateralOracleAdded)
	if err := _Contracts.contract.UnpackLog(event, "CollateralOracleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCollateralOracleRemovedIterator is returned from FilterCollateralOracleRemoved and is used to iterate over the raw logs and unpacked data for CollateralOracleRemoved events raised by the Contracts contract.
type ContractsCollateralOracleRemovedIterator struct {
	Event *ContractsCollateralOracleRemoved // Event containing the contract specifics and raw log

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
func (it *ContractsCollateralOracleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCollateralOracleRemoved)
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
		it.Event = new(ContractsCollateralOracleRemoved)
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
func (it *ContractsCollateralOracleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCollateralOracleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCollateralOracleRemoved represents a CollateralOracleRemoved event raised by the Contracts contract.
type ContractsCollateralOracleRemoved struct {
	Collateral common.Address
	Oracle     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollateralOracleRemoved is a free log retrieval operation binding the contract event 0x55f3a531fb6e334353b181d6c237c4df30a3cf04c971a3593c11d49955d2301c.
//
// Solidity: event CollateralOracleRemoved(address indexed collateral, address indexed oracle)
func (_Contracts *ContractsFilterer) FilterCollateralOracleRemoved(opts *bind.FilterOpts, collateral []common.Address, oracle []common.Address) (*ContractsCollateralOracleRemovedIterator, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CollateralOracleRemoved", collateralRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &ContractsCollateralOracleRemovedIterator{contract: _Contracts.contract, event: "CollateralOracleRemoved", logs: logs, sub: sub}, nil
}

// WatchCollateralOracleRemoved is a free log subscription operation binding the contract event 0x55f3a531fb6e334353b181d6c237c4df30a3cf04c971a3593c11d49955d2301c.
//
// Solidity: event CollateralOracleRemoved(address indexed collateral, address indexed oracle)
func (_Contracts *ContractsFilterer) WatchCollateralOracleRemoved(opts *bind.WatchOpts, sink chan<- *ContractsCollateralOracleRemoved, collateral []common.Address, oracle []common.Address) (event.Subscription, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CollateralOracleRemoved", collateralRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCollateralOracleRemoved)
				if err := _Contracts.contract.UnpackLog(event, "CollateralOracleRemoved", log); err != nil {
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

// ParseCollateralOracleRemoved is a log parse operation binding the contract event 0x55f3a531fb6e334353b181d6c237c4df30a3cf04c971a3593c11d49955d2301c.
//
// Solidity: event CollateralOracleRemoved(address indexed collateral, address indexed oracle)
func (_Contracts *ContractsFilterer) ParseCollateralOracleRemoved(log types.Log) (*ContractsCollateralOracleRemoved, error) {
	event := new(ContractsCollateralOracleRemoved)
	if err := _Contracts.contract.UnpackLog(event, "CollateralOracleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMinStakeChangedIterator is returned from FilterMinStakeChanged and is used to iterate over the raw logs and unpacked data for MinStakeChanged events raised by the Contracts contract.
type ContractsMinStakeChangedIterator struct {
	Event *ContractsMinStakeChanged // Event containing the contract specifics and raw log

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
func (it *ContractsMinStakeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMinStakeChanged)
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
		it.Event = new(ContractsMinStakeChanged)
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
func (it *ContractsMinStakeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMinStakeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMinStakeChanged represents a MinStakeChanged event raised by the Contracts contract.
type ContractsMinStakeChanged struct {
	OldMinstake *big.Int
	NewMinstake *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinStakeChanged is a free log retrieval operation binding the contract event 0xca11c8a4c461b60c9f485404c272650c2aaae260b2067d72e9924abb68556593.
//
// Solidity: event MinStakeChanged(uint256 oldMinstake, uint256 newMinstake)
func (_Contracts *ContractsFilterer) FilterMinStakeChanged(opts *bind.FilterOpts) (*ContractsMinStakeChangedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MinStakeChanged")
	if err != nil {
		return nil, err
	}
	return &ContractsMinStakeChangedIterator{contract: _Contracts.contract, event: "MinStakeChanged", logs: logs, sub: sub}, nil
}

// WatchMinStakeChanged is a free log subscription operation binding the contract event 0xca11c8a4c461b60c9f485404c272650c2aaae260b2067d72e9924abb68556593.
//
// Solidity: event MinStakeChanged(uint256 oldMinstake, uint256 newMinstake)
func (_Contracts *ContractsFilterer) WatchMinStakeChanged(opts *bind.WatchOpts, sink chan<- *ContractsMinStakeChanged) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MinStakeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMinStakeChanged)
				if err := _Contracts.contract.UnpackLog(event, "MinStakeChanged", log); err != nil {
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

// ParseMinStakeChanged is a log parse operation binding the contract event 0xca11c8a4c461b60c9f485404c272650c2aaae260b2067d72e9924abb68556593.
//
// Solidity: event MinStakeChanged(uint256 oldMinstake, uint256 newMinstake)
func (_Contracts *ContractsFilterer) ParseMinStakeChanged(log types.Log) (*ContractsMinStakeChanged, error) {
	event := new(ContractsMinStakeChanged)
	if err := _Contracts.contract.UnpackLog(event, "MinStakeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOperatorAddedToWhitelistIterator is returned from FilterOperatorAddedToWhitelist and is used to iterate over the raw logs and unpacked data for OperatorAddedToWhitelist events raised by the Contracts contract.
type ContractsOperatorAddedToWhitelistIterator struct {
	Event *ContractsOperatorAddedToWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractsOperatorAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOperatorAddedToWhitelist)
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
		it.Event = new(ContractsOperatorAddedToWhitelist)
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
func (it *ContractsOperatorAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOperatorAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOperatorAddedToWhitelist represents a OperatorAddedToWhitelist event raised by the Contracts contract.
type ContractsOperatorAddedToWhitelist struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToWhitelist is a free log retrieval operation binding the contract event 0x697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f.
//
// Solidity: event OperatorAddedToWhitelist(address indexed operator)
func (_Contracts *ContractsFilterer) FilterOperatorAddedToWhitelist(opts *bind.FilterOpts, operator []common.Address) (*ContractsOperatorAddedToWhitelistIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OperatorAddedToWhitelist", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOperatorAddedToWhitelistIterator{contract: _Contracts.contract, event: "OperatorAddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToWhitelist is a free log subscription operation binding the contract event 0x697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f.
//
// Solidity: event OperatorAddedToWhitelist(address indexed operator)
func (_Contracts *ContractsFilterer) WatchOperatorAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *ContractsOperatorAddedToWhitelist, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OperatorAddedToWhitelist", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOperatorAddedToWhitelist)
				if err := _Contracts.contract.UnpackLog(event, "OperatorAddedToWhitelist", log); err != nil {
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

// ParseOperatorAddedToWhitelist is a log parse operation binding the contract event 0x697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f.
//
// Solidity: event OperatorAddedToWhitelist(address indexed operator)
func (_Contracts *ContractsFilterer) ParseOperatorAddedToWhitelist(log types.Log) (*ContractsOperatorAddedToWhitelist, error) {
	event := new(ContractsOperatorAddedToWhitelist)
	if err := _Contracts.contract.UnpackLog(event, "OperatorAddedToWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the Contracts contract.
type ContractsOperatorRegisteredIterator struct {
	Event *ContractsOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractsOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOperatorRegistered)
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
		it.Event = new(ContractsOperatorRegistered)
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
func (it *ContractsOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOperatorRegistered represents a OperatorRegistered event raised by the Contracts contract.
type ContractsOperatorRegistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_Contracts *ContractsFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractsOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOperatorRegisteredIterator{contract: _Contracts.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_Contracts *ContractsFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractsOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOperatorRegistered)
				if err := _Contracts.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_Contracts *ContractsFilterer) ParseOperatorRegistered(log types.Log) (*ContractsOperatorRegistered, error) {
	event := new(ContractsOperatorRegistered)
	if err := _Contracts.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOperatorRemovedFromWhitelistIterator is returned from FilterOperatorRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromWhitelist events raised by the Contracts contract.
type ContractsOperatorRemovedFromWhitelistIterator struct {
	Event *ContractsOperatorRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractsOperatorRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOperatorRemovedFromWhitelist)
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
		it.Event = new(ContractsOperatorRemovedFromWhitelist)
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
func (it *ContractsOperatorRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOperatorRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOperatorRemovedFromWhitelist represents a OperatorRemovedFromWhitelist event raised by the Contracts contract.
type ContractsOperatorRemovedFromWhitelist struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x90ff506c206ebc7071c8746ca959ea01c75ddbd1464c1263398da71530aaf23a.
//
// Solidity: event OperatorRemovedFromWhitelist(address indexed operator)
func (_Contracts *ContractsFilterer) FilterOperatorRemovedFromWhitelist(opts *bind.FilterOpts, operator []common.Address) (*ContractsOperatorRemovedFromWhitelistIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OperatorRemovedFromWhitelist", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOperatorRemovedFromWhitelistIterator{contract: _Contracts.contract, event: "OperatorRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromWhitelist is a free log subscription operation binding the contract event 0x90ff506c206ebc7071c8746ca959ea01c75ddbd1464c1263398da71530aaf23a.
//
// Solidity: event OperatorRemovedFromWhitelist(address indexed operator)
func (_Contracts *ContractsFilterer) WatchOperatorRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *ContractsOperatorRemovedFromWhitelist, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OperatorRemovedFromWhitelist", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOperatorRemovedFromWhitelist)
				if err := _Contracts.contract.UnpackLog(event, "OperatorRemovedFromWhitelist", log); err != nil {
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

// ParseOperatorRemovedFromWhitelist is a log parse operation binding the contract event 0x90ff506c206ebc7071c8746ca959ea01c75ddbd1464c1263398da71530aaf23a.
//
// Solidity: event OperatorRemovedFromWhitelist(address indexed operator)
func (_Contracts *ContractsFilterer) ParseOperatorRemovedFromWhitelist(log types.Log) (*ContractsOperatorRemovedFromWhitelist, error) {
	event := new(ContractsOperatorRemovedFromWhitelist)
	if err := _Contracts.contract.UnpackLog(event, "OperatorRemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOperatorUnregisteredIterator is returned from FilterOperatorUnregistered and is used to iterate over the raw logs and unpacked data for OperatorUnregistered events raised by the Contracts contract.
type ContractsOperatorUnregisteredIterator struct {
	Event *ContractsOperatorUnregistered // Event containing the contract specifics and raw log

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
func (it *ContractsOperatorUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOperatorUnregistered)
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
		it.Event = new(ContractsOperatorUnregistered)
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
func (it *ContractsOperatorUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOperatorUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOperatorUnregistered represents a OperatorUnregistered event raised by the Contracts contract.
type ContractsOperatorUnregistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnregistered is a free log retrieval operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address indexed operator)
func (_Contracts *ContractsFilterer) FilterOperatorUnregistered(opts *bind.FilterOpts, operator []common.Address) (*ContractsOperatorUnregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OperatorUnregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOperatorUnregisteredIterator{contract: _Contracts.contract, event: "OperatorUnregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorUnregistered is a free log subscription operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address indexed operator)
func (_Contracts *ContractsFilterer) WatchOperatorUnregistered(opts *bind.WatchOpts, sink chan<- *ContractsOperatorUnregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OperatorUnregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOperatorUnregistered)
				if err := _Contracts.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
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

// ParseOperatorUnregistered is a log parse operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address indexed operator)
func (_Contracts *ContractsFilterer) ParseOperatorUnregistered(log types.Log) (*ContractsOperatorUnregistered, error) {
	event := new(ContractsOperatorUnregistered)
	if err := _Contracts.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSubnetworksChangedIterator is returned from FilterSubnetworksChanged and is used to iterate over the raw logs and unpacked data for SubnetworksChanged events raised by the Contracts contract.
type ContractsSubnetworksChangedIterator struct {
	Event *ContractsSubnetworksChanged // Event containing the contract specifics and raw log

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
func (it *ContractsSubnetworksChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSubnetworksChanged)
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
		it.Event = new(ContractsSubnetworksChanged)
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
func (it *ContractsSubnetworksChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSubnetworksChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSubnetworksChanged represents a SubnetworksChanged event raised by the Contracts contract.
type ContractsSubnetworksChanged struct {
	OldSubnetworks *big.Int
	NewSubnetworks *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubnetworksChanged is a free log retrieval operation binding the contract event 0x8b6b192217e39ecdda36ddf273727104d74ade70c0e44dfd59d6efcf54dc8381.
//
// Solidity: event SubnetworksChanged(uint256 oldSubnetworks, uint256 newSubnetworks)
func (_Contracts *ContractsFilterer) FilterSubnetworksChanged(opts *bind.FilterOpts) (*ContractsSubnetworksChangedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SubnetworksChanged")
	if err != nil {
		return nil, err
	}
	return &ContractsSubnetworksChangedIterator{contract: _Contracts.contract, event: "SubnetworksChanged", logs: logs, sub: sub}, nil
}

// WatchSubnetworksChanged is a free log subscription operation binding the contract event 0x8b6b192217e39ecdda36ddf273727104d74ade70c0e44dfd59d6efcf54dc8381.
//
// Solidity: event SubnetworksChanged(uint256 oldSubnetworks, uint256 newSubnetworks)
func (_Contracts *ContractsFilterer) WatchSubnetworksChanged(opts *bind.WatchOpts, sink chan<- *ContractsSubnetworksChanged) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SubnetworksChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSubnetworksChanged)
				if err := _Contracts.contract.UnpackLog(event, "SubnetworksChanged", log); err != nil {
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

// ParseSubnetworksChanged is a log parse operation binding the contract event 0x8b6b192217e39ecdda36ddf273727104d74ade70c0e44dfd59d6efcf54dc8381.
//
// Solidity: event SubnetworksChanged(uint256 oldSubnetworks, uint256 newSubnetworks)
func (_Contracts *ContractsFilterer) ParseSubnetworksChanged(log types.Log) (*ContractsSubnetworksChanged, error) {
	event := new(ContractsSubnetworksChanged)
	if err := _Contracts.contract.UnpackLog(event, "SubnetworksChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsVaultRegisteredIterator is returned from FilterVaultRegistered and is used to iterate over the raw logs and unpacked data for VaultRegistered events raised by the Contracts contract.
type ContractsVaultRegisteredIterator struct {
	Event *ContractsVaultRegistered // Event containing the contract specifics and raw log

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
func (it *ContractsVaultRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsVaultRegistered)
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
		it.Event = new(ContractsVaultRegistered)
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
func (it *ContractsVaultRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsVaultRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsVaultRegistered represents a VaultRegistered event raised by the Contracts contract.
type ContractsVaultRegistered struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVaultRegistered is a free log retrieval operation binding the contract event 0x8e0930709528779f1112249aac8fcca15dbb9c595db31092c7bc7f954b567933.
//
// Solidity: event VaultRegistered(address indexed vault)
func (_Contracts *ContractsFilterer) FilterVaultRegistered(opts *bind.FilterOpts, vault []common.Address) (*ContractsVaultRegisteredIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "VaultRegistered", vaultRule)
	if err != nil {
		return nil, err
	}
	return &ContractsVaultRegisteredIterator{contract: _Contracts.contract, event: "VaultRegistered", logs: logs, sub: sub}, nil
}

// WatchVaultRegistered is a free log subscription operation binding the contract event 0x8e0930709528779f1112249aac8fcca15dbb9c595db31092c7bc7f954b567933.
//
// Solidity: event VaultRegistered(address indexed vault)
func (_Contracts *ContractsFilterer) WatchVaultRegistered(opts *bind.WatchOpts, sink chan<- *ContractsVaultRegistered, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "VaultRegistered", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsVaultRegistered)
				if err := _Contracts.contract.UnpackLog(event, "VaultRegistered", log); err != nil {
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

// ParseVaultRegistered is a log parse operation binding the contract event 0x8e0930709528779f1112249aac8fcca15dbb9c595db31092c7bc7f954b567933.
//
// Solidity: event VaultRegistered(address indexed vault)
func (_Contracts *ContractsFilterer) ParseVaultRegistered(log types.Log) (*ContractsVaultRegistered, error) {
	event := new(ContractsVaultRegistered)
	if err := _Contracts.contract.UnpackLog(event, "VaultRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsVaultUnregisteredIterator is returned from FilterVaultUnregistered and is used to iterate over the raw logs and unpacked data for VaultUnregistered events raised by the Contracts contract.
type ContractsVaultUnregisteredIterator struct {
	Event *ContractsVaultUnregistered // Event containing the contract specifics and raw log

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
func (it *ContractsVaultUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsVaultUnregistered)
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
		it.Event = new(ContractsVaultUnregistered)
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
func (it *ContractsVaultUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsVaultUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsVaultUnregistered represents a VaultUnregistered event raised by the Contracts contract.
type ContractsVaultUnregistered struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVaultUnregistered is a free log retrieval operation binding the contract event 0x6fd85269676191488efa05d4b8cef674e502d3db3864638c729a0394048423cf.
//
// Solidity: event VaultUnregistered(address indexed vault)
func (_Contracts *ContractsFilterer) FilterVaultUnregistered(opts *bind.FilterOpts, vault []common.Address) (*ContractsVaultUnregisteredIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "VaultUnregistered", vaultRule)
	if err != nil {
		return nil, err
	}
	return &ContractsVaultUnregisteredIterator{contract: _Contracts.contract, event: "VaultUnregistered", logs: logs, sub: sub}, nil
}

// WatchVaultUnregistered is a free log subscription operation binding the contract event 0x6fd85269676191488efa05d4b8cef674e502d3db3864638c729a0394048423cf.
//
// Solidity: event VaultUnregistered(address indexed vault)
func (_Contracts *ContractsFilterer) WatchVaultUnregistered(opts *bind.WatchOpts, sink chan<- *ContractsVaultUnregistered, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "VaultUnregistered", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsVaultUnregistered)
				if err := _Contracts.contract.UnpackLog(event, "VaultUnregistered", log); err != nil {
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

// ParseVaultUnregistered is a log parse operation binding the contract event 0x6fd85269676191488efa05d4b8cef674e502d3db3864638c729a0394048423cf.
//
// Solidity: event VaultUnregistered(address indexed vault)
func (_Contracts *ContractsFilterer) ParseVaultUnregistered(log types.Log) (*ContractsVaultUnregistered, error) {
	event := new(ContractsVaultUnregistered)
	if err := _Contracts.contract.UnpackLog(event, "VaultUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
