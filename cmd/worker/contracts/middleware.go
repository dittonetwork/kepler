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
	Operator  common.Address
	Key       [32]byte
	VaultData []IValidatorDataVaultDataSet
}

// IValidatorDataValidatorDataShort is an auto generated low-level Go binding around an user-defined struct.
type IValidatorDataValidatorDataShort struct {
	Operator common.Address
	Key      [32]byte
	Stake    *big.Int
}

// IValidatorDataVaultDataSet is an auto generated low-level Go binding around an user-defined struct.
type IValidatorDataVaultDataSet struct {
	Vault           common.Address
	Stake           *big.Int
	PowerExpiresAt  *big.Int
	UnbondingTime   *big.Int
	UnbondingAmount *big.Int
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BaseMiddleware_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EpochCapture_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KeyManager256_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"Operators_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OzAccessControl_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SharedVaults_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addCollateralOracle\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collateralOracle\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCaptureTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpochDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpochStart\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRole\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPower\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorSet\",\"inputs\":[],\"outputs\":[{\"name\":\"validatorSet\",\"type\":\"tuple[]\",\"internalType\":\"structIValidatorData.ValidatorDataShort[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorSetWithUnbonding\",\"inputs\":[],\"outputs\":[{\"name\":\"validatorSet\",\"type\":\"tuple[]\",\"internalType\":\"structIValidatorData.ValidatorData[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vaultData\",\"type\":\"tuple[]\",\"internalType\":\"structIValidatorData.VaultDataSet[]\",\"components\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"powerExpiresAt\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"unbondingTime\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"unbondingAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVaultOracle\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slashingWindow\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"vaultRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorNetOptin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epochDuration\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"keyWasActiveAt\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"key_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorByKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorKey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseOperatorVault\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseSharedVault\",\"inputs\":[{\"name\":\"sharedVault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorVault\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerSharedVault\",\"inputs\":[{\"name\":\"sharedVault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeCollateralOracle\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinstake\",\"inputs\":[{\"name\":\"minStake_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeToPower\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"power\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"unpauseOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseOperatorVault\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseSharedVault\",\"inputs\":[{\"name\":\"sharedVault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterOperatorVault\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterSharedVault\",\"inputs\":[{\"name\":\"sharedVault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorKey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CollateralOracleAdded\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralOracleRemoved\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InstantSlash\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"subnetwork\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinStakeChanged\",\"inputs\":[{\"name\":\"oldMinstake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newMinstake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SelectorRoleSet\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VetoSlash\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"subnetwork\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadyEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CollateralAlreadyUseForVaults\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CollateralAlreadyUsedOtherOracle\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CollateralWithOracleAlreadyUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyAddressCollateralOrOracle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Enabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ImmutablePeriodNotPassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InactiveVaultSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxDisabledKeysReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoSlasher\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonVetoSlasher\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOperatorSpecificVault\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOperatorVault\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSupportedCollateral\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotVault\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotOptedIn\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorStakeIsSmall\",\"inputs\":[{\"name\":\"minStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TooOldTimestampSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnknownSlasherType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VaultAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VaultEpochTooShort\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VaultNotInitialized\",\"inputs\":[]}]",
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

// BaseMiddlewareVERSION is a free data retrieval call binding the contract method 0x32968557.
//
// Solidity: function BaseMiddleware_VERSION() view returns(uint64)
func (_Contracts *ContractsCaller) BaseMiddlewareVERSION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "BaseMiddleware_VERSION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BaseMiddlewareVERSION is a free data retrieval call binding the contract method 0x32968557.
//
// Solidity: function BaseMiddleware_VERSION() view returns(uint64)
func (_Contracts *ContractsSession) BaseMiddlewareVERSION() (uint64, error) {
	return _Contracts.Contract.BaseMiddlewareVERSION(&_Contracts.CallOpts)
}

// BaseMiddlewareVERSION is a free data retrieval call binding the contract method 0x32968557.
//
// Solidity: function BaseMiddleware_VERSION() view returns(uint64)
func (_Contracts *ContractsCallerSession) BaseMiddlewareVERSION() (uint64, error) {
	return _Contracts.Contract.BaseMiddlewareVERSION(&_Contracts.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Contracts.Contract.DEFAULTADMINROLE(&_Contracts.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Contracts.Contract.DEFAULTADMINROLE(&_Contracts.CallOpts)
}

// EpochCaptureVERSION is a free data retrieval call binding the contract method 0x64f084d5.
//
// Solidity: function EpochCapture_VERSION() view returns(uint64)
func (_Contracts *ContractsCaller) EpochCaptureVERSION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "EpochCapture_VERSION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// EpochCaptureVERSION is a free data retrieval call binding the contract method 0x64f084d5.
//
// Solidity: function EpochCapture_VERSION() view returns(uint64)
func (_Contracts *ContractsSession) EpochCaptureVERSION() (uint64, error) {
	return _Contracts.Contract.EpochCaptureVERSION(&_Contracts.CallOpts)
}

// EpochCaptureVERSION is a free data retrieval call binding the contract method 0x64f084d5.
//
// Solidity: function EpochCapture_VERSION() view returns(uint64)
func (_Contracts *ContractsCallerSession) EpochCaptureVERSION() (uint64, error) {
	return _Contracts.Contract.EpochCaptureVERSION(&_Contracts.CallOpts)
}

// KeyManager256VERSION is a free data retrieval call binding the contract method 0xe6989de7.
//
// Solidity: function KeyManager256_VERSION() view returns(uint64)
func (_Contracts *ContractsCaller) KeyManager256VERSION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "KeyManager256_VERSION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// KeyManager256VERSION is a free data retrieval call binding the contract method 0xe6989de7.
//
// Solidity: function KeyManager256_VERSION() view returns(uint64)
func (_Contracts *ContractsSession) KeyManager256VERSION() (uint64, error) {
	return _Contracts.Contract.KeyManager256VERSION(&_Contracts.CallOpts)
}

// KeyManager256VERSION is a free data retrieval call binding the contract method 0xe6989de7.
//
// Solidity: function KeyManager256_VERSION() view returns(uint64)
func (_Contracts *ContractsCallerSession) KeyManager256VERSION() (uint64, error) {
	return _Contracts.Contract.KeyManager256VERSION(&_Contracts.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) OPERATORROLE() ([32]byte, error) {
	return _Contracts.Contract.OPERATORROLE(&_Contracts.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Contracts.Contract.OPERATORROLE(&_Contracts.CallOpts)
}

// OperatorsVERSION is a free data retrieval call binding the contract method 0x409637fa.
//
// Solidity: function Operators_VERSION() view returns(uint64)
func (_Contracts *ContractsCaller) OperatorsVERSION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "Operators_VERSION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OperatorsVERSION is a free data retrieval call binding the contract method 0x409637fa.
//
// Solidity: function Operators_VERSION() view returns(uint64)
func (_Contracts *ContractsSession) OperatorsVERSION() (uint64, error) {
	return _Contracts.Contract.OperatorsVERSION(&_Contracts.CallOpts)
}

// OperatorsVERSION is a free data retrieval call binding the contract method 0x409637fa.
//
// Solidity: function Operators_VERSION() view returns(uint64)
func (_Contracts *ContractsCallerSession) OperatorsVERSION() (uint64, error) {
	return _Contracts.Contract.OperatorsVERSION(&_Contracts.CallOpts)
}

// OzAccessControlVERSION is a free data retrieval call binding the contract method 0xc52a6697.
//
// Solidity: function OzAccessControl_VERSION() view returns(uint64)
func (_Contracts *ContractsCaller) OzAccessControlVERSION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "OzAccessControl_VERSION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OzAccessControlVERSION is a free data retrieval call binding the contract method 0xc52a6697.
//
// Solidity: function OzAccessControl_VERSION() view returns(uint64)
func (_Contracts *ContractsSession) OzAccessControlVERSION() (uint64, error) {
	return _Contracts.Contract.OzAccessControlVERSION(&_Contracts.CallOpts)
}

// OzAccessControlVERSION is a free data retrieval call binding the contract method 0xc52a6697.
//
// Solidity: function OzAccessControl_VERSION() view returns(uint64)
func (_Contracts *ContractsCallerSession) OzAccessControlVERSION() (uint64, error) {
	return _Contracts.Contract.OzAccessControlVERSION(&_Contracts.CallOpts)
}

// SharedVaultsVERSION is a free data retrieval call binding the contract method 0x0a0530ea.
//
// Solidity: function SharedVaults_VERSION() view returns(uint64)
func (_Contracts *ContractsCaller) SharedVaultsVERSION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "SharedVaults_VERSION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SharedVaultsVERSION is a free data retrieval call binding the contract method 0x0a0530ea.
//
// Solidity: function SharedVaults_VERSION() view returns(uint64)
func (_Contracts *ContractsSession) SharedVaultsVERSION() (uint64, error) {
	return _Contracts.Contract.SharedVaultsVERSION(&_Contracts.CallOpts)
}

// SharedVaultsVERSION is a free data retrieval call binding the contract method 0x0a0530ea.
//
// Solidity: function SharedVaults_VERSION() view returns(uint64)
func (_Contracts *ContractsCallerSession) SharedVaultsVERSION() (uint64, error) {
	return _Contracts.Contract.SharedVaultsVERSION(&_Contracts.CallOpts)
}

// CollateralOracle is a free data retrieval call binding the contract method 0xe9144e73.
//
// Solidity: function collateralOracle(address collateral) view returns(address)
func (_Contracts *ContractsCaller) CollateralOracle(opts *bind.CallOpts, collateral common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "collateralOracle", collateral)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralOracle is a free data retrieval call binding the contract method 0xe9144e73.
//
// Solidity: function collateralOracle(address collateral) view returns(address)
func (_Contracts *ContractsSession) CollateralOracle(collateral common.Address) (common.Address, error) {
	return _Contracts.Contract.CollateralOracle(&_Contracts.CallOpts, collateral)
}

// CollateralOracle is a free data retrieval call binding the contract method 0xe9144e73.
//
// Solidity: function collateralOracle(address collateral) view returns(address)
func (_Contracts *ContractsCallerSession) CollateralOracle(collateral common.Address) (common.Address, error) {
	return _Contracts.Contract.CollateralOracle(&_Contracts.CallOpts, collateral)
}

// GetCaptureTimestamp is a free data retrieval call binding the contract method 0xdb3adf12.
//
// Solidity: function getCaptureTimestamp() view returns(uint48 timestamp)
func (_Contracts *ContractsCaller) GetCaptureTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCaptureTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCaptureTimestamp is a free data retrieval call binding the contract method 0xdb3adf12.
//
// Solidity: function getCaptureTimestamp() view returns(uint48 timestamp)
func (_Contracts *ContractsSession) GetCaptureTimestamp() (*big.Int, error) {
	return _Contracts.Contract.GetCaptureTimestamp(&_Contracts.CallOpts)
}

// GetCaptureTimestamp is a free data retrieval call binding the contract method 0xdb3adf12.
//
// Solidity: function getCaptureTimestamp() view returns(uint48 timestamp)
func (_Contracts *ContractsCallerSession) GetCaptureTimestamp() (*big.Int, error) {
	return _Contracts.Contract.GetCaptureTimestamp(&_Contracts.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint48)
func (_Contracts *ContractsCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint48)
func (_Contracts *ContractsSession) GetCurrentEpoch() (*big.Int, error) {
	return _Contracts.Contract.GetCurrentEpoch(&_Contracts.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint48)
func (_Contracts *ContractsCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _Contracts.Contract.GetCurrentEpoch(&_Contracts.CallOpts)
}

// GetEpochDuration is a free data retrieval call binding the contract method 0x5d3ea8f1.
//
// Solidity: function getEpochDuration() view returns(uint48)
func (_Contracts *ContractsCaller) GetEpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getEpochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochDuration is a free data retrieval call binding the contract method 0x5d3ea8f1.
//
// Solidity: function getEpochDuration() view returns(uint48)
func (_Contracts *ContractsSession) GetEpochDuration() (*big.Int, error) {
	return _Contracts.Contract.GetEpochDuration(&_Contracts.CallOpts)
}

// GetEpochDuration is a free data retrieval call binding the contract method 0x5d3ea8f1.
//
// Solidity: function getEpochDuration() view returns(uint48)
func (_Contracts *ContractsCallerSession) GetEpochDuration() (*big.Int, error) {
	return _Contracts.Contract.GetEpochDuration(&_Contracts.CallOpts)
}

// GetEpochStart is a free data retrieval call binding the contract method 0x246e158f.
//
// Solidity: function getEpochStart(uint48 epoch) view returns(uint48)
func (_Contracts *ContractsCaller) GetEpochStart(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getEpochStart", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochStart is a free data retrieval call binding the contract method 0x246e158f.
//
// Solidity: function getEpochStart(uint48 epoch) view returns(uint48)
func (_Contracts *ContractsSession) GetEpochStart(epoch *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetEpochStart(&_Contracts.CallOpts, epoch)
}

// GetEpochStart is a free data retrieval call binding the contract method 0x246e158f.
//
// Solidity: function getEpochStart(uint48 epoch) view returns(uint48)
func (_Contracts *ContractsCallerSession) GetEpochStart(epoch *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetEpochStart(&_Contracts.CallOpts, epoch)
}

// GetRole is a free data retrieval call binding the contract method 0xa846156d.
//
// Solidity: function getRole(bytes4 selector) view returns(bytes32)
func (_Contracts *ContractsCaller) GetRole(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getRole", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRole is a free data retrieval call binding the contract method 0xa846156d.
//
// Solidity: function getRole(bytes4 selector) view returns(bytes32)
func (_Contracts *ContractsSession) GetRole(selector [4]byte) ([32]byte, error) {
	return _Contracts.Contract.GetRole(&_Contracts.CallOpts, selector)
}

// GetRole is a free data retrieval call binding the contract method 0xa846156d.
//
// Solidity: function getRole(bytes4 selector) view returns(bytes32)
func (_Contracts *ContractsCallerSession) GetRole(selector [4]byte) ([32]byte, error) {
	return _Contracts.Contract.GetRole(&_Contracts.CallOpts, selector)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contracts *ContractsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contracts *ContractsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Contracts.Contract.GetRoleAdmin(&_Contracts.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contracts *ContractsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Contracts.Contract.GetRoleAdmin(&_Contracts.CallOpts, role)
}

// GetTotalPower is a free data retrieval call binding the contract method 0x53976a26.
//
// Solidity: function getTotalPower() view returns(uint256)
func (_Contracts *ContractsCaller) GetTotalPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getTotalPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPower is a free data retrieval call binding the contract method 0x53976a26.
//
// Solidity: function getTotalPower() view returns(uint256)
func (_Contracts *ContractsSession) GetTotalPower() (*big.Int, error) {
	return _Contracts.Contract.GetTotalPower(&_Contracts.CallOpts)
}

// GetTotalPower is a free data retrieval call binding the contract method 0x53976a26.
//
// Solidity: function getTotalPower() view returns(uint256)
func (_Contracts *ContractsCallerSession) GetTotalPower() (*big.Int, error) {
	return _Contracts.Contract.GetTotalPower(&_Contracts.CallOpts)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0xcf331250.
//
// Solidity: function getValidatorSet() view returns((address,bytes32,uint256)[] validatorSet)
func (_Contracts *ContractsCaller) GetValidatorSet(opts *bind.CallOpts) ([]IValidatorDataValidatorDataShort, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getValidatorSet")

	if err != nil {
		return *new([]IValidatorDataValidatorDataShort), err
	}

	out0 := *abi.ConvertType(out[0], new([]IValidatorDataValidatorDataShort)).(*[]IValidatorDataValidatorDataShort)

	return out0, err

}

// GetValidatorSet is a free data retrieval call binding the contract method 0xcf331250.
//
// Solidity: function getValidatorSet() view returns((address,bytes32,uint256)[] validatorSet)
func (_Contracts *ContractsSession) GetValidatorSet() ([]IValidatorDataValidatorDataShort, error) {
	return _Contracts.Contract.GetValidatorSet(&_Contracts.CallOpts)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0xcf331250.
//
// Solidity: function getValidatorSet() view returns((address,bytes32,uint256)[] validatorSet)
func (_Contracts *ContractsCallerSession) GetValidatorSet() ([]IValidatorDataValidatorDataShort, error) {
	return _Contracts.Contract.GetValidatorSet(&_Contracts.CallOpts)
}

// GetValidatorSetWithUnbonding is a free data retrieval call binding the contract method 0x4d1163a4.
//
// Solidity: function getValidatorSetWithUnbonding() view returns((address,bytes32,(address,uint256,uint48,uint48,uint256)[])[] validatorSet)
func (_Contracts *ContractsCaller) GetValidatorSetWithUnbonding(opts *bind.CallOpts) ([]IValidatorDataValidatorData, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getValidatorSetWithUnbonding")

	if err != nil {
		return *new([]IValidatorDataValidatorData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IValidatorDataValidatorData)).(*[]IValidatorDataValidatorData)

	return out0, err

}

// GetValidatorSetWithUnbonding is a free data retrieval call binding the contract method 0x4d1163a4.
//
// Solidity: function getValidatorSetWithUnbonding() view returns((address,bytes32,(address,uint256,uint48,uint48,uint256)[])[] validatorSet)
func (_Contracts *ContractsSession) GetValidatorSetWithUnbonding() ([]IValidatorDataValidatorData, error) {
	return _Contracts.Contract.GetValidatorSetWithUnbonding(&_Contracts.CallOpts)
}

// GetValidatorSetWithUnbonding is a free data retrieval call binding the contract method 0x4d1163a4.
//
// Solidity: function getValidatorSetWithUnbonding() view returns((address,bytes32,(address,uint256,uint48,uint48,uint256)[])[] validatorSet)
func (_Contracts *ContractsCallerSession) GetValidatorSetWithUnbonding() ([]IValidatorDataValidatorData, error) {
	return _Contracts.Contract.GetValidatorSetWithUnbonding(&_Contracts.CallOpts)
}

// GetVaultOracle is a free data retrieval call binding the contract method 0x542abf8b.
//
// Solidity: function getVaultOracle(address vault) view returns(address)
func (_Contracts *ContractsCaller) GetVaultOracle(opts *bind.CallOpts, vault common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getVaultOracle", vault)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVaultOracle is a free data retrieval call binding the contract method 0x542abf8b.
//
// Solidity: function getVaultOracle(address vault) view returns(address)
func (_Contracts *ContractsSession) GetVaultOracle(vault common.Address) (common.Address, error) {
	return _Contracts.Contract.GetVaultOracle(&_Contracts.CallOpts, vault)
}

// GetVaultOracle is a free data retrieval call binding the contract method 0x542abf8b.
//
// Solidity: function getVaultOracle(address vault) view returns(address)
func (_Contracts *ContractsCallerSession) GetVaultOracle(vault common.Address) (common.Address, error) {
	return _Contracts.Contract.GetVaultOracle(&_Contracts.CallOpts, vault)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contracts *ContractsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contracts *ContractsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Contracts.Contract.HasRole(&_Contracts.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contracts *ContractsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Contracts.Contract.HasRole(&_Contracts.CallOpts, role, account)
}

// KeyWasActiveAt is a free data retrieval call binding the contract method 0x1e0f2e1f.
//
// Solidity: function keyWasActiveAt(uint48 timestamp, bytes key_) view returns(bool)
func (_Contracts *ContractsCaller) KeyWasActiveAt(opts *bind.CallOpts, timestamp *big.Int, key_ []byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "keyWasActiveAt", timestamp, key_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeyWasActiveAt is a free data retrieval call binding the contract method 0x1e0f2e1f.
//
// Solidity: function keyWasActiveAt(uint48 timestamp, bytes key_) view returns(bool)
func (_Contracts *ContractsSession) KeyWasActiveAt(timestamp *big.Int, key_ []byte) (bool, error) {
	return _Contracts.Contract.KeyWasActiveAt(&_Contracts.CallOpts, timestamp, key_)
}

// KeyWasActiveAt is a free data retrieval call binding the contract method 0x1e0f2e1f.
//
// Solidity: function keyWasActiveAt(uint48 timestamp, bytes key_) view returns(bool)
func (_Contracts *ContractsCallerSession) KeyWasActiveAt(timestamp *big.Int, key_ []byte) (bool, error) {
	return _Contracts.Contract.KeyWasActiveAt(&_Contracts.CallOpts, timestamp, key_)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Contracts *ContractsCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Contracts *ContractsSession) MinStake() (*big.Int, error) {
	return _Contracts.Contract.MinStake(&_Contracts.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Contracts *ContractsCallerSession) MinStake() (*big.Int, error) {
	return _Contracts.Contract.MinStake(&_Contracts.CallOpts)
}

// OperatorByKey is a free data retrieval call binding the contract method 0x3e1ad83f.
//
// Solidity: function operatorByKey(bytes key) view returns(address)
func (_Contracts *ContractsCaller) OperatorByKey(opts *bind.CallOpts, key []byte) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "operatorByKey", key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorByKey is a free data retrieval call binding the contract method 0x3e1ad83f.
//
// Solidity: function operatorByKey(bytes key) view returns(address)
func (_Contracts *ContractsSession) OperatorByKey(key []byte) (common.Address, error) {
	return _Contracts.Contract.OperatorByKey(&_Contracts.CallOpts, key)
}

// OperatorByKey is a free data retrieval call binding the contract method 0x3e1ad83f.
//
// Solidity: function operatorByKey(bytes key) view returns(address)
func (_Contracts *ContractsCallerSession) OperatorByKey(key []byte) (common.Address, error) {
	return _Contracts.Contract.OperatorByKey(&_Contracts.CallOpts, key)
}

// OperatorKey is a free data retrieval call binding the contract method 0xb18125be.
//
// Solidity: function operatorKey(address operator) view returns(bytes)
func (_Contracts *ContractsCaller) OperatorKey(opts *bind.CallOpts, operator common.Address) ([]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "operatorKey", operator)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// OperatorKey is a free data retrieval call binding the contract method 0xb18125be.
//
// Solidity: function operatorKey(address operator) view returns(bytes)
func (_Contracts *ContractsSession) OperatorKey(operator common.Address) ([]byte, error) {
	return _Contracts.Contract.OperatorKey(&_Contracts.CallOpts, operator)
}

// OperatorKey is a free data retrieval call binding the contract method 0xb18125be.
//
// Solidity: function operatorKey(address operator) view returns(bytes)
func (_Contracts *ContractsCallerSession) OperatorKey(operator common.Address) ([]byte, error) {
	return _Contracts.Contract.OperatorKey(&_Contracts.CallOpts, operator)
}

// StakeToPower is a free data retrieval call binding the contract method 0x84af6324.
//
// Solidity: function stakeToPower(address vault, uint256 stake) pure returns(uint256 power)
func (_Contracts *ContractsCaller) StakeToPower(opts *bind.CallOpts, vault common.Address, stake *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "stakeToPower", vault, stake)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeToPower is a free data retrieval call binding the contract method 0x84af6324.
//
// Solidity: function stakeToPower(address vault, uint256 stake) pure returns(uint256 power)
func (_Contracts *ContractsSession) StakeToPower(vault common.Address, stake *big.Int) (*big.Int, error) {
	return _Contracts.Contract.StakeToPower(&_Contracts.CallOpts, vault, stake)
}

// StakeToPower is a free data retrieval call binding the contract method 0x84af6324.
//
// Solidity: function stakeToPower(address vault, uint256 stake) pure returns(uint256 power)
func (_Contracts *ContractsCallerSession) StakeToPower(vault common.Address, stake *big.Int) (*big.Int, error) {
	return _Contracts.Contract.StakeToPower(&_Contracts.CallOpts, vault, stake)
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

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contracts *ContractsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GrantRole(&_Contracts.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GrantRole(&_Contracts.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xa2244473.
//
// Solidity: function initialize(address network, uint48 slashingWindow, address vaultRegistry, address operatorRegistry, address operatorNetOptin, address reader, address owner, uint48 epochDuration) returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts, network common.Address, slashingWindow *big.Int, vaultRegistry common.Address, operatorRegistry common.Address, operatorNetOptin common.Address, reader common.Address, owner common.Address, epochDuration *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initialize", network, slashingWindow, vaultRegistry, operatorRegistry, operatorNetOptin, reader, owner, epochDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xa2244473.
//
// Solidity: function initialize(address network, uint48 slashingWindow, address vaultRegistry, address operatorRegistry, address operatorNetOptin, address reader, address owner, uint48 epochDuration) returns()
func (_Contracts *ContractsSession) Initialize(network common.Address, slashingWindow *big.Int, vaultRegistry common.Address, operatorRegistry common.Address, operatorNetOptin common.Address, reader common.Address, owner common.Address, epochDuration *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, network, slashingWindow, vaultRegistry, operatorRegistry, operatorNetOptin, reader, owner, epochDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xa2244473.
//
// Solidity: function initialize(address network, uint48 slashingWindow, address vaultRegistry, address operatorRegistry, address operatorNetOptin, address reader, address owner, uint48 epochDuration) returns()
func (_Contracts *ContractsTransactorSession) Initialize(network common.Address, slashingWindow *big.Int, vaultRegistry common.Address, operatorRegistry common.Address, operatorNetOptin common.Address, reader common.Address, owner common.Address, epochDuration *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, network, slashingWindow, vaultRegistry, operatorRegistry, operatorNetOptin, reader, owner, epochDuration)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address operator) returns()
func (_Contracts *ContractsTransactor) PauseOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "pauseOperator", operator)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address operator) returns()
func (_Contracts *ContractsSession) PauseOperator(operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.PauseOperator(&_Contracts.TransactOpts, operator)
}

// PauseOperator is a paid mutator transaction binding the contract method 0x72f9adab.
//
// Solidity: function pauseOperator(address operator) returns()
func (_Contracts *ContractsTransactorSession) PauseOperator(operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.PauseOperator(&_Contracts.TransactOpts, operator)
}

// PauseOperatorVault is a paid mutator transaction binding the contract method 0xc55041cf.
//
// Solidity: function pauseOperatorVault(address operator, address vault) returns()
func (_Contracts *ContractsTransactor) PauseOperatorVault(opts *bind.TransactOpts, operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "pauseOperatorVault", operator, vault)
}

// PauseOperatorVault is a paid mutator transaction binding the contract method 0xc55041cf.
//
// Solidity: function pauseOperatorVault(address operator, address vault) returns()
func (_Contracts *ContractsSession) PauseOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.PauseOperatorVault(&_Contracts.TransactOpts, operator, vault)
}

// PauseOperatorVault is a paid mutator transaction binding the contract method 0xc55041cf.
//
// Solidity: function pauseOperatorVault(address operator, address vault) returns()
func (_Contracts *ContractsTransactorSession) PauseOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.PauseOperatorVault(&_Contracts.TransactOpts, operator, vault)
}

// PauseSharedVault is a paid mutator transaction binding the contract method 0xb1630faa.
//
// Solidity: function pauseSharedVault(address sharedVault) returns()
func (_Contracts *ContractsTransactor) PauseSharedVault(opts *bind.TransactOpts, sharedVault common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "pauseSharedVault", sharedVault)
}

// PauseSharedVault is a paid mutator transaction binding the contract method 0xb1630faa.
//
// Solidity: function pauseSharedVault(address sharedVault) returns()
func (_Contracts *ContractsSession) PauseSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.PauseSharedVault(&_Contracts.TransactOpts, sharedVault)
}

// PauseSharedVault is a paid mutator transaction binding the contract method 0xb1630faa.
//
// Solidity: function pauseSharedVault(address sharedVault) returns()
func (_Contracts *ContractsTransactorSession) PauseSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.PauseSharedVault(&_Contracts.TransactOpts, sharedVault)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc1656d40.
//
// Solidity: function registerOperator(address operator, bytes key, address vault) returns()
func (_Contracts *ContractsTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, key []byte, vault common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerOperator", operator, key, vault)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc1656d40.
//
// Solidity: function registerOperator(address operator, bytes key, address vault) returns()
func (_Contracts *ContractsSession) RegisterOperator(operator common.Address, key []byte, vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterOperator(&_Contracts.TransactOpts, operator, key, vault)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc1656d40.
//
// Solidity: function registerOperator(address operator, bytes key, address vault) returns()
func (_Contracts *ContractsTransactorSession) RegisterOperator(operator common.Address, key []byte, vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterOperator(&_Contracts.TransactOpts, operator, key, vault)
}

// RegisterOperatorVault is a paid mutator transaction binding the contract method 0xb1a69fa2.
//
// Solidity: function registerOperatorVault(address operator, address vault) returns()
func (_Contracts *ContractsTransactor) RegisterOperatorVault(opts *bind.TransactOpts, operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerOperatorVault", operator, vault)
}

// RegisterOperatorVault is a paid mutator transaction binding the contract method 0xb1a69fa2.
//
// Solidity: function registerOperatorVault(address operator, address vault) returns()
func (_Contracts *ContractsSession) RegisterOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterOperatorVault(&_Contracts.TransactOpts, operator, vault)
}

// RegisterOperatorVault is a paid mutator transaction binding the contract method 0xb1a69fa2.
//
// Solidity: function registerOperatorVault(address operator, address vault) returns()
func (_Contracts *ContractsTransactorSession) RegisterOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterOperatorVault(&_Contracts.TransactOpts, operator, vault)
}

// RegisterSharedVault is a paid mutator transaction binding the contract method 0xca2d2a18.
//
// Solidity: function registerSharedVault(address sharedVault) returns()
func (_Contracts *ContractsTransactor) RegisterSharedVault(opts *bind.TransactOpts, sharedVault common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerSharedVault", sharedVault)
}

// RegisterSharedVault is a paid mutator transaction binding the contract method 0xca2d2a18.
//
// Solidity: function registerSharedVault(address sharedVault) returns()
func (_Contracts *ContractsSession) RegisterSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterSharedVault(&_Contracts.TransactOpts, sharedVault)
}

// RegisterSharedVault is a paid mutator transaction binding the contract method 0xca2d2a18.
//
// Solidity: function registerSharedVault(address sharedVault) returns()
func (_Contracts *ContractsTransactorSession) RegisterSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterSharedVault(&_Contracts.TransactOpts, sharedVault)
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

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Contracts *ContractsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Contracts *ContractsSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RenounceRole(&_Contracts.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Contracts *ContractsTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RenounceRole(&_Contracts.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contracts *ContractsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeRole(&_Contracts.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeRole(&_Contracts.TransactOpts, role, account)
}

// SetMinstake is a paid mutator transaction binding the contract method 0x6393aa17.
//
// Solidity: function setMinstake(uint256 minStake_) returns()
func (_Contracts *ContractsTransactor) SetMinstake(opts *bind.TransactOpts, minStake_ *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMinstake", minStake_)
}

// SetMinstake is a paid mutator transaction binding the contract method 0x6393aa17.
//
// Solidity: function setMinstake(uint256 minStake_) returns()
func (_Contracts *ContractsSession) SetMinstake(minStake_ *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMinstake(&_Contracts.TransactOpts, minStake_)
}

// SetMinstake is a paid mutator transaction binding the contract method 0x6393aa17.
//
// Solidity: function setMinstake(uint256 minStake_) returns()
func (_Contracts *ContractsTransactorSession) SetMinstake(minStake_ *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMinstake(&_Contracts.TransactOpts, minStake_)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address operator) returns()
func (_Contracts *ContractsTransactor) UnpauseOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unpauseOperator", operator)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address operator) returns()
func (_Contracts *ContractsSession) UnpauseOperator(operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnpauseOperator(&_Contracts.TransactOpts, operator)
}

// UnpauseOperator is a paid mutator transaction binding the contract method 0x2e5aaf33.
//
// Solidity: function unpauseOperator(address operator) returns()
func (_Contracts *ContractsTransactorSession) UnpauseOperator(operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnpauseOperator(&_Contracts.TransactOpts, operator)
}

// UnpauseOperatorVault is a paid mutator transaction binding the contract method 0x5aa59c4f.
//
// Solidity: function unpauseOperatorVault(address operator, address vault) returns()
func (_Contracts *ContractsTransactor) UnpauseOperatorVault(opts *bind.TransactOpts, operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unpauseOperatorVault", operator, vault)
}

// UnpauseOperatorVault is a paid mutator transaction binding the contract method 0x5aa59c4f.
//
// Solidity: function unpauseOperatorVault(address operator, address vault) returns()
func (_Contracts *ContractsSession) UnpauseOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnpauseOperatorVault(&_Contracts.TransactOpts, operator, vault)
}

// UnpauseOperatorVault is a paid mutator transaction binding the contract method 0x5aa59c4f.
//
// Solidity: function unpauseOperatorVault(address operator, address vault) returns()
func (_Contracts *ContractsTransactorSession) UnpauseOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnpauseOperatorVault(&_Contracts.TransactOpts, operator, vault)
}

// UnpauseSharedVault is a paid mutator transaction binding the contract method 0x08e809f0.
//
// Solidity: function unpauseSharedVault(address sharedVault) returns()
func (_Contracts *ContractsTransactor) UnpauseSharedVault(opts *bind.TransactOpts, sharedVault common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unpauseSharedVault", sharedVault)
}

// UnpauseSharedVault is a paid mutator transaction binding the contract method 0x08e809f0.
//
// Solidity: function unpauseSharedVault(address sharedVault) returns()
func (_Contracts *ContractsSession) UnpauseSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnpauseSharedVault(&_Contracts.TransactOpts, sharedVault)
}

// UnpauseSharedVault is a paid mutator transaction binding the contract method 0x08e809f0.
//
// Solidity: function unpauseSharedVault(address sharedVault) returns()
func (_Contracts *ContractsTransactorSession) UnpauseSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnpauseSharedVault(&_Contracts.TransactOpts, sharedVault)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address operator) returns()
func (_Contracts *ContractsTransactor) UnregisterOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unregisterOperator", operator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address operator) returns()
func (_Contracts *ContractsSession) UnregisterOperator(operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnregisterOperator(&_Contracts.TransactOpts, operator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address operator) returns()
func (_Contracts *ContractsTransactorSession) UnregisterOperator(operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnregisterOperator(&_Contracts.TransactOpts, operator)
}

// UnregisterOperatorVault is a paid mutator transaction binding the contract method 0xcb87ef6e.
//
// Solidity: function unregisterOperatorVault(address operator, address vault) returns()
func (_Contracts *ContractsTransactor) UnregisterOperatorVault(opts *bind.TransactOpts, operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unregisterOperatorVault", operator, vault)
}

// UnregisterOperatorVault is a paid mutator transaction binding the contract method 0xcb87ef6e.
//
// Solidity: function unregisterOperatorVault(address operator, address vault) returns()
func (_Contracts *ContractsSession) UnregisterOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnregisterOperatorVault(&_Contracts.TransactOpts, operator, vault)
}

// UnregisterOperatorVault is a paid mutator transaction binding the contract method 0xcb87ef6e.
//
// Solidity: function unregisterOperatorVault(address operator, address vault) returns()
func (_Contracts *ContractsTransactorSession) UnregisterOperatorVault(operator common.Address, vault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnregisterOperatorVault(&_Contracts.TransactOpts, operator, vault)
}

// UnregisterSharedVault is a paid mutator transaction binding the contract method 0x47449640.
//
// Solidity: function unregisterSharedVault(address sharedVault) returns()
func (_Contracts *ContractsTransactor) UnregisterSharedVault(opts *bind.TransactOpts, sharedVault common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unregisterSharedVault", sharedVault)
}

// UnregisterSharedVault is a paid mutator transaction binding the contract method 0x47449640.
//
// Solidity: function unregisterSharedVault(address sharedVault) returns()
func (_Contracts *ContractsSession) UnregisterSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnregisterSharedVault(&_Contracts.TransactOpts, sharedVault)
}

// UnregisterSharedVault is a paid mutator transaction binding the contract method 0x47449640.
//
// Solidity: function unregisterSharedVault(address sharedVault) returns()
func (_Contracts *ContractsTransactorSession) UnregisterSharedVault(sharedVault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnregisterSharedVault(&_Contracts.TransactOpts, sharedVault)
}

// UpdateOperatorKey is a paid mutator transaction binding the contract method 0x181d5cd6.
//
// Solidity: function updateOperatorKey(address operator, bytes key) returns()
func (_Contracts *ContractsTransactor) UpdateOperatorKey(opts *bind.TransactOpts, operator common.Address, key []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateOperatorKey", operator, key)
}

// UpdateOperatorKey is a paid mutator transaction binding the contract method 0x181d5cd6.
//
// Solidity: function updateOperatorKey(address operator, bytes key) returns()
func (_Contracts *ContractsSession) UpdateOperatorKey(operator common.Address, key []byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateOperatorKey(&_Contracts.TransactOpts, operator, key)
}

// UpdateOperatorKey is a paid mutator transaction binding the contract method 0x181d5cd6.
//
// Solidity: function updateOperatorKey(address operator, bytes key) returns()
func (_Contracts *ContractsTransactorSession) UpdateOperatorKey(operator common.Address, key []byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateOperatorKey(&_Contracts.TransactOpts, operator, key)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Contracts *ContractsTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Contracts.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Contracts *ContractsSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Fallback(&_Contracts.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Contracts *ContractsTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Fallback(&_Contracts.TransactOpts, calldata)
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

// ContractsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contracts contract.
type ContractsInitializedIterator struct {
	Event *ContractsInitialized // Event containing the contract specifics and raw log

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
func (it *ContractsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsInitialized)
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
		it.Event = new(ContractsInitialized)
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
func (it *ContractsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsInitialized represents a Initialized event raised by the Contracts contract.
type ContractsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contracts *ContractsFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractsInitializedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractsInitializedIterator{contract: _Contracts.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contracts *ContractsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractsInitialized) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsInitialized)
				if err := _Contracts.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contracts *ContractsFilterer) ParseInitialized(log types.Log) (*ContractsInitialized, error) {
	event := new(ContractsInitialized)
	if err := _Contracts.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsInstantSlashIterator is returned from FilterInstantSlash and is used to iterate over the raw logs and unpacked data for InstantSlash events raised by the Contracts contract.
type ContractsInstantSlashIterator struct {
	Event *ContractsInstantSlash // Event containing the contract specifics and raw log

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
func (it *ContractsInstantSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsInstantSlash)
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
		it.Event = new(ContractsInstantSlash)
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
func (it *ContractsInstantSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsInstantSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsInstantSlash represents a InstantSlash event raised by the Contracts contract.
type ContractsInstantSlash struct {
	Vault      common.Address
	Subnetwork [32]byte
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInstantSlash is a free log retrieval operation binding the contract event 0xa455bb45e23ed02807f6ef41727a47f3fcc85c9df0baa3570fd388f95b09b4da.
//
// Solidity: event InstantSlash(address vault, bytes32 subnetwork, uint256 amount)
func (_Contracts *ContractsFilterer) FilterInstantSlash(opts *bind.FilterOpts) (*ContractsInstantSlashIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "InstantSlash")
	if err != nil {
		return nil, err
	}
	return &ContractsInstantSlashIterator{contract: _Contracts.contract, event: "InstantSlash", logs: logs, sub: sub}, nil
}

// WatchInstantSlash is a free log subscription operation binding the contract event 0xa455bb45e23ed02807f6ef41727a47f3fcc85c9df0baa3570fd388f95b09b4da.
//
// Solidity: event InstantSlash(address vault, bytes32 subnetwork, uint256 amount)
func (_Contracts *ContractsFilterer) WatchInstantSlash(opts *bind.WatchOpts, sink chan<- *ContractsInstantSlash) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "InstantSlash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsInstantSlash)
				if err := _Contracts.contract.UnpackLog(event, "InstantSlash", log); err != nil {
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

// ParseInstantSlash is a log parse operation binding the contract event 0xa455bb45e23ed02807f6ef41727a47f3fcc85c9df0baa3570fd388f95b09b4da.
//
// Solidity: event InstantSlash(address vault, bytes32 subnetwork, uint256 amount)
func (_Contracts *ContractsFilterer) ParseInstantSlash(log types.Log) (*ContractsInstantSlash, error) {
	event := new(ContractsInstantSlash)
	if err := _Contracts.contract.UnpackLog(event, "InstantSlash", log); err != nil {
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

// ContractsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Contracts contract.
type ContractsRoleAdminChangedIterator struct {
	Event *ContractsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ContractsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRoleAdminChanged)
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
		it.Event = new(ContractsRoleAdminChanged)
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
func (it *ContractsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRoleAdminChanged represents a RoleAdminChanged event raised by the Contracts contract.
type ContractsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Contracts *ContractsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractsRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRoleAdminChangedIterator{contract: _Contracts.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Contracts *ContractsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRoleAdminChanged)
				if err := _Contracts.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseRoleAdminChanged(log types.Log) (*ContractsRoleAdminChanged, error) {
	event := new(ContractsRoleAdminChanged)
	if err := _Contracts.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Contracts contract.
type ContractsRoleGrantedIterator struct {
	Event *ContractsRoleGranted // Event containing the contract specifics and raw log

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
func (it *ContractsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRoleGranted)
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
		it.Event = new(ContractsRoleGranted)
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
func (it *ContractsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRoleGranted represents a RoleGranted event raised by the Contracts contract.
type ContractsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractsRoleGrantedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRoleGrantedIterator{contract: _Contracts.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRoleGranted)
				if err := _Contracts.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseRoleGranted(log types.Log) (*ContractsRoleGranted, error) {
	event := new(ContractsRoleGranted)
	if err := _Contracts.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Contracts contract.
type ContractsRoleRevokedIterator struct {
	Event *ContractsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ContractsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRoleRevoked)
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
		it.Event = new(ContractsRoleRevoked)
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
func (it *ContractsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRoleRevoked represents a RoleRevoked event raised by the Contracts contract.
type ContractsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractsRoleRevokedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRoleRevokedIterator{contract: _Contracts.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRoleRevoked)
				if err := _Contracts.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseRoleRevoked(log types.Log) (*ContractsRoleRevoked, error) {
	event := new(ContractsRoleRevoked)
	if err := _Contracts.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSelectorRoleSetIterator is returned from FilterSelectorRoleSet and is used to iterate over the raw logs and unpacked data for SelectorRoleSet events raised by the Contracts contract.
type ContractsSelectorRoleSetIterator struct {
	Event *ContractsSelectorRoleSet // Event containing the contract specifics and raw log

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
func (it *ContractsSelectorRoleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSelectorRoleSet)
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
		it.Event = new(ContractsSelectorRoleSet)
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
func (it *ContractsSelectorRoleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSelectorRoleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSelectorRoleSet represents a SelectorRoleSet event raised by the Contracts contract.
type ContractsSelectorRoleSet struct {
	Selector [4]byte
	Role     [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSelectorRoleSet is a free log retrieval operation binding the contract event 0xb579d5e7e95ac8795a9c9ecce0ee2e2d189dce9827bac2e35ebbd3a68be7d423.
//
// Solidity: event SelectorRoleSet(bytes4 indexed selector, bytes32 indexed role)
func (_Contracts *ContractsFilterer) FilterSelectorRoleSet(opts *bind.FilterOpts, selector [][4]byte, role [][32]byte) (*ContractsSelectorRoleSetIterator, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SelectorRoleSet", selectorRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSelectorRoleSetIterator{contract: _Contracts.contract, event: "SelectorRoleSet", logs: logs, sub: sub}, nil
}

// WatchSelectorRoleSet is a free log subscription operation binding the contract event 0xb579d5e7e95ac8795a9c9ecce0ee2e2d189dce9827bac2e35ebbd3a68be7d423.
//
// Solidity: event SelectorRoleSet(bytes4 indexed selector, bytes32 indexed role)
func (_Contracts *ContractsFilterer) WatchSelectorRoleSet(opts *bind.WatchOpts, sink chan<- *ContractsSelectorRoleSet, selector [][4]byte, role [][32]byte) (event.Subscription, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SelectorRoleSet", selectorRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSelectorRoleSet)
				if err := _Contracts.contract.UnpackLog(event, "SelectorRoleSet", log); err != nil {
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

// ParseSelectorRoleSet is a log parse operation binding the contract event 0xb579d5e7e95ac8795a9c9ecce0ee2e2d189dce9827bac2e35ebbd3a68be7d423.
//
// Solidity: event SelectorRoleSet(bytes4 indexed selector, bytes32 indexed role)
func (_Contracts *ContractsFilterer) ParseSelectorRoleSet(log types.Log) (*ContractsSelectorRoleSet, error) {
	event := new(ContractsSelectorRoleSet)
	if err := _Contracts.contract.UnpackLog(event, "SelectorRoleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsVetoSlashIterator is returned from FilterVetoSlash and is used to iterate over the raw logs and unpacked data for VetoSlash events raised by the Contracts contract.
type ContractsVetoSlashIterator struct {
	Event *ContractsVetoSlash // Event containing the contract specifics and raw log

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
func (it *ContractsVetoSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsVetoSlash)
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
		it.Event = new(ContractsVetoSlash)
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
func (it *ContractsVetoSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsVetoSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsVetoSlash represents a VetoSlash event raised by the Contracts contract.
type ContractsVetoSlash struct {
	Vault      common.Address
	Subnetwork [32]byte
	Index      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVetoSlash is a free log retrieval operation binding the contract event 0x4df99d47392012b66d459ea8fe495a8ce499b8faee622119c4cf353023b582fe.
//
// Solidity: event VetoSlash(address vault, bytes32 subnetwork, uint256 index)
func (_Contracts *ContractsFilterer) FilterVetoSlash(opts *bind.FilterOpts) (*ContractsVetoSlashIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "VetoSlash")
	if err != nil {
		return nil, err
	}
	return &ContractsVetoSlashIterator{contract: _Contracts.contract, event: "VetoSlash", logs: logs, sub: sub}, nil
}

// WatchVetoSlash is a free log subscription operation binding the contract event 0x4df99d47392012b66d459ea8fe495a8ce499b8faee622119c4cf353023b582fe.
//
// Solidity: event VetoSlash(address vault, bytes32 subnetwork, uint256 index)
func (_Contracts *ContractsFilterer) WatchVetoSlash(opts *bind.WatchOpts, sink chan<- *ContractsVetoSlash) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "VetoSlash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsVetoSlash)
				if err := _Contracts.contract.UnpackLog(event, "VetoSlash", log); err != nil {
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

// ParseVetoSlash is a log parse operation binding the contract event 0x4df99d47392012b66d459ea8fe495a8ce499b8faee622119c4cf353023b582fe.
//
// Solidity: event VetoSlash(address vault, bytes32 subnetwork, uint256 index)
func (_Contracts *ContractsFilterer) ParseVetoSlash(log types.Log) (*ContractsVetoSlash, error) {
	event := new(ContractsVetoSlash)
	if err := _Contracts.contract.UnpackLog(event, "VetoSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
