// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package network

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

// ITypesSlashingReport is an auto generated low-level Go binding around an user-defined struct.
type ITypesSlashingReport struct {
	ValidatorAddress common.Address
	SlashAmount      *big.Int
	Reason           string
}

// ITypesValidator is an auto generated low-level Go binding around an user-defined struct.
type ITypesValidator struct {
	Operator    common.Address
	VotingPower *big.Int
	PublicKey   [32]byte
	IsEmergency bool
	Status      uint8
	Protocol    uint8
}

// INetworkMetaData contains all meta data concerning the INetwork contract.
var INetworkMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidCommitteeAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommitteeSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMiddlewareAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperatorAddress\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"publicKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isEmergency\",\"type\":\"bool\"},{\"internalType\":\"enumITypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumITypes.Protocol\",\"name\":\"protocol\",\"type\":\"uint8\"}],\"internalType\":\"structITypes.Validator\",\"name\":\"validator\",\"type\":\"tuple\"}],\"name\":\"addDittoValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorSet\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"publicKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isEmergency\",\"type\":\"bool\"},{\"internalType\":\"enumITypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumITypes.Protocol\",\"name\":\"protocol\",\"type\":\"uint8\"}],\"internalType\":\"structITypes.Validator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"markValidatorEmergency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"internalType\":\"structITypes.SlashingReport[]\",\"name\":\"reports\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"sendSlashingReports\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// INetworkABI is the input ABI used to generate the binding from.
// Deprecated: Use INetworkMetaData.ABI instead.
var INetworkABI = INetworkMetaData.ABI

// INetwork is an auto generated Go binding around an Ethereum contract.
type INetwork struct {
	INetworkCaller     // Read-only binding to the contract
	INetworkTransactor // Write-only binding to the contract
	INetworkFilterer   // Log filterer for contract events
}

// INetworkCaller is an auto generated read-only Go binding around an Ethereum contract.
type INetworkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INetworkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INetworkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INetworkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INetworkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INetworkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INetworkSession struct {
	Contract     *INetwork         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INetworkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INetworkCallerSession struct {
	Contract *INetworkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// INetworkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INetworkTransactorSession struct {
	Contract     *INetworkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// INetworkRaw is an auto generated low-level Go binding around an Ethereum contract.
type INetworkRaw struct {
	Contract *INetwork // Generic contract binding to access the raw methods on
}

// INetworkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INetworkCallerRaw struct {
	Contract *INetworkCaller // Generic read-only contract binding to access the raw methods on
}

// INetworkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INetworkTransactorRaw struct {
	Contract *INetworkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINetwork creates a new instance of INetwork, bound to a specific deployed contract.
func NewINetwork(address common.Address, backend bind.ContractBackend) (*INetwork, error) {
	contract, err := bindINetwork(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INetwork{INetworkCaller: INetworkCaller{contract: contract}, INetworkTransactor: INetworkTransactor{contract: contract}, INetworkFilterer: INetworkFilterer{contract: contract}}, nil
}

// NewINetworkCaller creates a new read-only instance of INetwork, bound to a specific deployed contract.
func NewINetworkCaller(address common.Address, caller bind.ContractCaller) (*INetworkCaller, error) {
	contract, err := bindINetwork(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INetworkCaller{contract: contract}, nil
}

// NewINetworkTransactor creates a new write-only instance of INetwork, bound to a specific deployed contract.
func NewINetworkTransactor(address common.Address, transactor bind.ContractTransactor) (*INetworkTransactor, error) {
	contract, err := bindINetwork(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INetworkTransactor{contract: contract}, nil
}

// NewINetworkFilterer creates a new log filterer instance of INetwork, bound to a specific deployed contract.
func NewINetworkFilterer(address common.Address, filterer bind.ContractFilterer) (*INetworkFilterer, error) {
	contract, err := bindINetwork(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INetworkFilterer{contract: contract}, nil
}

// bindINetwork binds a generic wrapper to an already deployed contract.
func bindINetwork(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := INetworkMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INetwork *INetworkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INetwork.Contract.INetworkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INetwork *INetworkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INetwork.Contract.INetworkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INetwork *INetworkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INetwork.Contract.INetworkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INetwork *INetworkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INetwork.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INetwork *INetworkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INetwork.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INetwork *INetworkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INetwork.Contract.contract.Transact(opts, method, params...)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0xcf331250.
//
// Solidity: function getValidatorSet() view returns((address,uint256,bytes32,bool,uint8,uint8)[] validators)
func (_INetwork *INetworkCaller) GetValidatorSet(opts *bind.CallOpts) ([]ITypesValidator, error) {
	var out []interface{}
	err := _INetwork.contract.Call(opts, &out, "getValidatorSet")

	if err != nil {
		return *new([]ITypesValidator), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITypesValidator)).(*[]ITypesValidator)

	return out0, err

}

// GetValidatorSet is a free data retrieval call binding the contract method 0xcf331250.
//
// Solidity: function getValidatorSet() view returns((address,uint256,bytes32,bool,uint8,uint8)[] validators)
func (_INetwork *INetworkSession) GetValidatorSet() ([]ITypesValidator, error) {
	return _INetwork.Contract.GetValidatorSet(&_INetwork.CallOpts)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0xcf331250.
//
// Solidity: function getValidatorSet() view returns((address,uint256,bytes32,bool,uint8,uint8)[] validators)
func (_INetwork *INetworkCallerSession) GetValidatorSet() ([]ITypesValidator, error) {
	return _INetwork.Contract.GetValidatorSet(&_INetwork.CallOpts)
}

// AddDittoValidator is a paid mutator transaction binding the contract method 0x87797dca.
//
// Solidity: function addDittoValidator((address,uint256,bytes32,bool,uint8,uint8) validator) returns()
func (_INetwork *INetworkTransactor) AddDittoValidator(opts *bind.TransactOpts, validator ITypesValidator) (*types.Transaction, error) {
	return _INetwork.contract.Transact(opts, "addDittoValidator", validator)
}

// AddDittoValidator is a paid mutator transaction binding the contract method 0x87797dca.
//
// Solidity: function addDittoValidator((address,uint256,bytes32,bool,uint8,uint8) validator) returns()
func (_INetwork *INetworkSession) AddDittoValidator(validator ITypesValidator) (*types.Transaction, error) {
	return _INetwork.Contract.AddDittoValidator(&_INetwork.TransactOpts, validator)
}

// AddDittoValidator is a paid mutator transaction binding the contract method 0x87797dca.
//
// Solidity: function addDittoValidator((address,uint256,bytes32,bool,uint8,uint8) validator) returns()
func (_INetwork *INetworkTransactorSession) AddDittoValidator(validator ITypesValidator) (*types.Transaction, error) {
	return _INetwork.Contract.AddDittoValidator(&_INetwork.TransactOpts, validator)
}

// MarkValidatorEmergency is a paid mutator transaction binding the contract method 0x9f2110e4.
//
// Solidity: function markValidatorEmergency(address validatorAddress) returns()
func (_INetwork *INetworkTransactor) MarkValidatorEmergency(opts *bind.TransactOpts, validatorAddress common.Address) (*types.Transaction, error) {
	return _INetwork.contract.Transact(opts, "markValidatorEmergency", validatorAddress)
}

// MarkValidatorEmergency is a paid mutator transaction binding the contract method 0x9f2110e4.
//
// Solidity: function markValidatorEmergency(address validatorAddress) returns()
func (_INetwork *INetworkSession) MarkValidatorEmergency(validatorAddress common.Address) (*types.Transaction, error) {
	return _INetwork.Contract.MarkValidatorEmergency(&_INetwork.TransactOpts, validatorAddress)
}

// MarkValidatorEmergency is a paid mutator transaction binding the contract method 0x9f2110e4.
//
// Solidity: function markValidatorEmergency(address validatorAddress) returns()
func (_INetwork *INetworkTransactorSession) MarkValidatorEmergency(validatorAddress common.Address) (*types.Transaction, error) {
	return _INetwork.Contract.MarkValidatorEmergency(&_INetwork.TransactOpts, validatorAddress)
}

// SendSlashingReports is a paid mutator transaction binding the contract method 0x420052f3.
//
// Solidity: function sendSlashingReports((address,uint256,string)[] reports, bytes[] signatures) returns()
func (_INetwork *INetworkTransactor) SendSlashingReports(opts *bind.TransactOpts, reports []ITypesSlashingReport, signatures [][]byte) (*types.Transaction, error) {
	return _INetwork.contract.Transact(opts, "sendSlashingReports", reports, signatures)
}

// SendSlashingReports is a paid mutator transaction binding the contract method 0x420052f3.
//
// Solidity: function sendSlashingReports((address,uint256,string)[] reports, bytes[] signatures) returns()
func (_INetwork *INetworkSession) SendSlashingReports(reports []ITypesSlashingReport, signatures [][]byte) (*types.Transaction, error) {
	return _INetwork.Contract.SendSlashingReports(&_INetwork.TransactOpts, reports, signatures)
}

// SendSlashingReports is a paid mutator transaction binding the contract method 0x420052f3.
//
// Solidity: function sendSlashingReports((address,uint256,string)[] reports, bytes[] signatures) returns()
func (_INetwork *INetworkTransactorSession) SendSlashingReports(reports []ITypesSlashingReport, signatures [][]byte) (*types.Transaction, error) {
	return _INetwork.Contract.SendSlashingReports(&_INetwork.TransactOpts, reports, signatures)
}
