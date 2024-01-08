// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IdentityInterface

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

// IdentityInterfaceMetaData contains all meta data concerning the IdentityInterface contract.
var IdentityInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"NewManager\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identityManagerContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IdentityInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityInterfaceMetaData.ABI instead.
var IdentityInterfaceABI = IdentityInterfaceMetaData.ABI

// IdentityInterface is an auto generated Go binding around an Ethereum contract.
type IdentityInterface struct {
	IdentityInterfaceCaller     // Read-only binding to the contract
	IdentityInterfaceTransactor // Write-only binding to the contract
	IdentityInterfaceFilterer   // Log filterer for contract events
}

// IdentityInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityInterfaceSession struct {
	Contract     *IdentityInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IdentityInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityInterfaceCallerSession struct {
	Contract *IdentityInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IdentityInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityInterfaceTransactorSession struct {
	Contract     *IdentityInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IdentityInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityInterfaceRaw struct {
	Contract *IdentityInterface // Generic contract binding to access the raw methods on
}

// IdentityInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityInterfaceCallerRaw struct {
	Contract *IdentityInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityInterfaceTransactorRaw struct {
	Contract *IdentityInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityInterface creates a new instance of IdentityInterface, bound to a specific deployed contract.
func NewIdentityInterface(address common.Address, backend bind.ContractBackend) (*IdentityInterface, error) {
	contract, err := bindIdentityInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityInterface{IdentityInterfaceCaller: IdentityInterfaceCaller{contract: contract}, IdentityInterfaceTransactor: IdentityInterfaceTransactor{contract: contract}, IdentityInterfaceFilterer: IdentityInterfaceFilterer{contract: contract}}, nil
}

// NewIdentityInterfaceCaller creates a new read-only instance of IdentityInterface, bound to a specific deployed contract.
func NewIdentityInterfaceCaller(address common.Address, caller bind.ContractCaller) (*IdentityInterfaceCaller, error) {
	contract, err := bindIdentityInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityInterfaceCaller{contract: contract}, nil
}

// NewIdentityInterfaceTransactor creates a new write-only instance of IdentityInterface, bound to a specific deployed contract.
func NewIdentityInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityInterfaceTransactor, error) {
	contract, err := bindIdentityInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityInterfaceTransactor{contract: contract}, nil
}

// NewIdentityInterfaceFilterer creates a new log filterer instance of IdentityInterface, bound to a specific deployed contract.
func NewIdentityInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityInterfaceFilterer, error) {
	contract, err := bindIdentityInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityInterfaceFilterer{contract: contract}, nil
}

// bindIdentityInterface binds a generic wrapper to an already deployed contract.
func bindIdentityInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityInterface *IdentityInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityInterface.Contract.IdentityInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityInterface *IdentityInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityInterface.Contract.IdentityInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityInterface *IdentityInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityInterface.Contract.IdentityInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityInterface *IdentityInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityInterface *IdentityInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityInterface *IdentityInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityInterface.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address node) view returns(bool)
func (_IdentityInterface *IdentityInterfaceCaller) Exists(opts *bind.CallOpts, node common.Address) (bool, error) {
	var out []interface{}
	err := _IdentityInterface.contract.Call(opts, &out, "exists", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address node) view returns(bool)
func (_IdentityInterface *IdentityInterfaceSession) Exists(node common.Address) (bool, error) {
	return _IdentityInterface.Contract.Exists(&_IdentityInterface.CallOpts, node)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address node) view returns(bool)
func (_IdentityInterface *IdentityInterfaceCallerSession) Exists(node common.Address) (bool, error) {
	return _IdentityInterface.Contract.Exists(&_IdentityInterface.CallOpts, node)
}

// Manager is a free data retrieval call binding the contract method 0xd4d2e7f2.
//
// Solidity: function manager(address node) view returns(address)
func (_IdentityInterface *IdentityInterfaceCaller) Manager(opts *bind.CallOpts, node common.Address) (common.Address, error) {
	var out []interface{}
	err := _IdentityInterface.contract.Call(opts, &out, "manager", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0xd4d2e7f2.
//
// Solidity: function manager(address node) view returns(address)
func (_IdentityInterface *IdentityInterfaceSession) Manager(node common.Address) (common.Address, error) {
	return _IdentityInterface.Contract.Manager(&_IdentityInterface.CallOpts, node)
}

// Manager is a free data retrieval call binding the contract method 0xd4d2e7f2.
//
// Solidity: function manager(address node) view returns(address)
func (_IdentityInterface *IdentityInterfaceCallerSession) Manager(node common.Address) (common.Address, error) {
	return _IdentityInterface.Contract.Manager(&_IdentityInterface.CallOpts, node)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address identityManagerContract, bytes signature) returns()
func (_IdentityInterface *IdentityInterfaceTransactor) Register(opts *bind.TransactOpts, identityManagerContract common.Address, signature []byte) (*types.Transaction, error) {
	return _IdentityInterface.contract.Transact(opts, "register", identityManagerContract, signature)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address identityManagerContract, bytes signature) returns()
func (_IdentityInterface *IdentityInterfaceSession) Register(identityManagerContract common.Address, signature []byte) (*types.Transaction, error) {
	return _IdentityInterface.Contract.Register(&_IdentityInterface.TransactOpts, identityManagerContract, signature)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address identityManagerContract, bytes signature) returns()
func (_IdentityInterface *IdentityInterfaceTransactorSession) Register(identityManagerContract common.Address, signature []byte) (*types.Transaction, error) {
	return _IdentityInterface.Contract.Register(&_IdentityInterface.TransactOpts, identityManagerContract, signature)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager) returns()
func (_IdentityInterface *IdentityInterfaceTransactor) SetManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _IdentityInterface.contract.Transact(opts, "setManager", manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager) returns()
func (_IdentityInterface *IdentityInterfaceSession) SetManager(manager common.Address) (*types.Transaction, error) {
	return _IdentityInterface.Contract.SetManager(&_IdentityInterface.TransactOpts, manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager) returns()
func (_IdentityInterface *IdentityInterfaceTransactorSession) SetManager(manager common.Address) (*types.Transaction, error) {
	return _IdentityInterface.Contract.SetManager(&_IdentityInterface.TransactOpts, manager)
}

// IdentityInterfaceNewManagerIterator is returned from FilterNewManager and is used to iterate over the raw logs and unpacked data for NewManager events raised by the IdentityInterface contract.
type IdentityInterfaceNewManagerIterator struct {
	Event *IdentityInterfaceNewManager // Event containing the contract specifics and raw log

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
func (it *IdentityInterfaceNewManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityInterfaceNewManager)
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
		it.Event = new(IdentityInterfaceNewManager)
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
func (it *IdentityInterfaceNewManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityInterfaceNewManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityInterfaceNewManager represents a NewManager event raised by the IdentityInterface contract.
type IdentityInterfaceNewManager struct {
	Node    common.Address
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewManager is a free log retrieval operation binding the contract event 0x770e6248a70b6ac757edf422766216da592c37e3112db900fe0da8984191831b.
//
// Solidity: event NewManager(address indexed node, address manager)
func (_IdentityInterface *IdentityInterfaceFilterer) FilterNewManager(opts *bind.FilterOpts, node []common.Address) (*IdentityInterfaceNewManagerIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IdentityInterface.contract.FilterLogs(opts, "NewManager", nodeRule)
	if err != nil {
		return nil, err
	}
	return &IdentityInterfaceNewManagerIterator{contract: _IdentityInterface.contract, event: "NewManager", logs: logs, sub: sub}, nil
}

// WatchNewManager is a free log subscription operation binding the contract event 0x770e6248a70b6ac757edf422766216da592c37e3112db900fe0da8984191831b.
//
// Solidity: event NewManager(address indexed node, address manager)
func (_IdentityInterface *IdentityInterfaceFilterer) WatchNewManager(opts *bind.WatchOpts, sink chan<- *IdentityInterfaceNewManager, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IdentityInterface.contract.WatchLogs(opts, "NewManager", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityInterfaceNewManager)
				if err := _IdentityInterface.contract.UnpackLog(event, "NewManager", log); err != nil {
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

// ParseNewManager is a log parse operation binding the contract event 0x770e6248a70b6ac757edf422766216da592c37e3112db900fe0da8984191831b.
//
// Solidity: event NewManager(address indexed node, address manager)
func (_IdentityInterface *IdentityInterfaceFilterer) ParseNewManager(log types.Log) (*IdentityInterfaceNewManager, error) {
	event := new(IdentityInterfaceNewManager)
	if err := _IdentityInterface.contract.UnpackLog(event, "NewManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
