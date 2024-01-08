// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IdentityRegistry

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

// IdentityRegistryMetaData contains all meta data concerning the IdentityRegistry contract.
var IdentityRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"NewManager\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"circuits\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"deploymentType\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"deploymentAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"deregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_circuitId\",\"type\":\"string\"}],\"name\":\"getCircuit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"identities\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identityManagerContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"running\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_circuitId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_deploymentType\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_deploymentAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_ipfsHash\",\"type\":\"string\"}],\"name\":\"setCircuit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600160035f6101000a81548160ff021916908315150217905550348015610029575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611be4806100765f395ff3fe608060405234801561000f575f80fd5b50600436106100cd575f3560e01c80639ea2cd381161008a578063d85bd52611610064578063d85bd526146101fb578063e4baeda314610219578063f653b81e1461024b578063f6a3d24e1461027b576100cd565b80639ea2cd381461017d578063d0ebdbe7146101af578063d4d2e7f2146101cb576100cd565b806314afd79e146100d157806324b8fbf6146101015780635208ef6a1461011d57806383197ef01461013957806384ac33ec146101435780638da5cb5b1461015f575b5f80fd5b6100eb60048036038101906100e691906110b9565b6102ab565b6040516100f891906110f3565b60405180910390f35b61011b60048036038101906101169190611248565b610310565b005b61013760048036038101906101329190611340565b61056c565b005b6101416106d8565b005b61015d600480360381019061015891906110b9565b610749565b005b610167610891565b60405161017491906110f3565b60405180910390f35b610197600480360381019061019291906113f8565b6108b4565b6040516101a6939291906114b9565b60405180910390f35b6101c960048036038101906101c491906110b9565b610a1d565b005b6101e560048036038101906101e091906110b9565b610bcf565b6040516101f291906110f3565b60405180910390f35b610203610c34565b6040516102109190611516565b60405180910390f35b610233600480360381019061022e91906113f8565b610c46565b604051610242939291906114b9565b60405180910390f35b610265600480360381019061026091906110b9565b610dee565b60405161027291906110f3565b60405180910390f35b610295600480360381019061029091906110b9565b610e1e565b6040516102a29190611516565b60405180910390f35b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60035f9054906101000a900460ff1661035e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035590611579565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610428576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041f9061162d565b60405180910390fd5b5f6104338333610eb2565b905061045f815f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684610ee4565b61049e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049590611695565b60405180910390fd5b8260015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f770e6248a70b6ac757edf422766216da592c37e3112db900fe0da8984191831b8460405161055f91906110f3565b60405180910390a2505050565b60035f9054906101000a900460ff166105ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b190611579565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610610575f80fd5b60405180606001604052808481526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281525060028560405161065191906116ed565b90815260200160405180910390205f820151815f0190816106729190611906565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190816106ce9190611906565b5090505050505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461072e575f80fd5b5f60035f6101000a81548160ff021916908315150217905550565b60035f9054906101000a900460ff16610797576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078e90611579565b60405180910390fd5b5f3390503373ffffffffffffffffffffffffffffffffffffffff1660015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461082e575f80fd5b60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6002818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0180546108eb90611730565b80601f016020809104026020016040519081016040528092919081815260200182805461091790611730565b80156109625780601f1061093957610100808354040283529160200191610962565b820191905f5260205f20905b81548152906001019060200180831161094557829003601f168201915b505050505090806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600201805461099c90611730565b80601f01602080910402602001604051908101604052809291908181526020018280546109c890611730565b8015610a135780601f106109ea57610100808354040283529160200191610a13565b820191905f5260205f20905b8154815290600101906020018083116109f657829003601f168201915b5050505050905083565b60035f9054906101000a900460ff16610a6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6290611579565b60405180910390fd5b5f3390503373ffffffffffffffffffffffffffffffffffffffff1660015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b02575f80fd5b8160015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f770e6248a70b6ac757edf422766216da592c37e3112db900fe0da8984191831b83604051610bc391906110f3565b60405180910390a25050565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60035f9054906101000a900460ff1681565b60605f6060600284604051610c5b91906116ed565b90815260200160405180910390205f01600285604051610c7b91906116ed565b90815260200160405180910390206001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600286604051610cbc91906116ed565b9081526020016040518091039020600201828054610cd990611730565b80601f0160208091040260200160405190810160405280929190818152602001828054610d0590611730565b8015610d505780601f10610d2757610100808354040283529160200191610d50565b820191905f5260205f20905b815481529060010190602001808311610d3357829003601f168201915b50505050509250808054610d6390611730565b80601f0160208091040260200160405190810160405280929190818152602001828054610d8f90611730565b8015610dda5780601f10610db157610100808354040283529160200191610dda565b820191905f5260205f20905b815481529060010190602001808311610dbd57829003601f168201915b505050505090509250925092509193909250565b6001602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b5f8282604051602001610ec6929190611a1a565b60405160208183030381529060405280519060200120905092915050565b5f806040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a33320000000081525090505f8186604051602001610f32929190611ab2565b6040516020818303038152906040528051906020012090505f805f610f5687610fe9565b9250925092505f6001858386866040515f8152602001604052604051610f7f9493929190611b03565b6020604051602081039080840390855afa158015610f9f573d5f803e3d5ffd5b5050506020604051035190508073ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff161496505050505050509392505050565b5f805f6041845114611030576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102790611b90565b60405180910390fd5b602084015192506040840151915060608401515f1a90509193909250565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6110888261105f565b9050919050565b6110988161107e565b81146110a2575f80fd5b50565b5f813590506110b38161108f565b92915050565b5f602082840312156110ce576110cd611057565b5b5f6110db848285016110a5565b91505092915050565b6110ed8161107e565b82525050565b5f6020820190506111065f8301846110e4565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61115a82611114565b810181811067ffffffffffffffff8211171561117957611178611124565b5b80604052505050565b5f61118b61104e565b90506111978282611151565b919050565b5f67ffffffffffffffff8211156111b6576111b5611124565b5b6111bf82611114565b9050602081019050919050565b828183375f83830152505050565b5f6111ec6111e78461119c565b611182565b90508281526020810184848401111561120857611207611110565b5b6112138482856111cc565b509392505050565b5f82601f83011261122f5761122e61110c565b5b813561123f8482602086016111da565b91505092915050565b5f806040838503121561125e5761125d611057565b5b5f61126b858286016110a5565b925050602083013567ffffffffffffffff81111561128c5761128b61105b565b5b6112988582860161121b565b9150509250929050565b5f67ffffffffffffffff8211156112bc576112bb611124565b5b6112c582611114565b9050602081019050919050565b5f6112e46112df846112a2565b611182565b905082815260208101848484011115611300576112ff611110565b5b61130b8482856111cc565b509392505050565b5f82601f8301126113275761132661110c565b5b81356113378482602086016112d2565b91505092915050565b5f805f806080858703121561135857611357611057565b5b5f85013567ffffffffffffffff8111156113755761137461105b565b5b61138187828801611313565b945050602085013567ffffffffffffffff8111156113a2576113a161105b565b5b6113ae87828801611313565b93505060406113bf878288016110a5565b925050606085013567ffffffffffffffff8111156113e0576113df61105b565b5b6113ec87828801611313565b91505092959194509250565b5f6020828403121561140d5761140c611057565b5b5f82013567ffffffffffffffff81111561142a5761142961105b565b5b61143684828501611313565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561147657808201518184015260208101905061145b565b5f8484015250505050565b5f61148b8261143f565b6114958185611449565b93506114a5818560208601611459565b6114ae81611114565b840191505092915050565b5f6060820190508181035f8301526114d18186611481565b90506114e060208301856110e4565b81810360408301526114f28184611481565b9050949350505050565b5f8115159050919050565b611510816114fc565b82525050565b5f6020820190506115295f830184611507565b92915050565b7f436f6e7472616374206973206e6f742072756e6e696e670000000000000000005f82015250565b5f611563601783611449565b915061156e8261152f565b602082019050919050565b5f6020820190508181035f83015261159081611557565b9050919050565b7f416c726561647920726567697374657265642e20557365207365744d616e61675f8201527f65722063616c6c20746f20757064617465206d616e616765722061646472657360208201527f732e000000000000000000000000000000000000000000000000000000000000604082015250565b5f611617604283611449565b915061162282611597565b606082019050919050565b5f6020820190508181035f8301526116448161160b565b9050919050565b7f496e76616c6964207369676e61747572650000000000000000000000000000005f82015250565b5f61167f601183611449565b915061168a8261164b565b602082019050919050565b5f6020820190508181035f8301526116ac81611673565b9050919050565b5f81905092915050565b5f6116c78261143f565b6116d181856116b3565b93506116e1818560208601611459565b80840191505092915050565b5f6116f882846116bd565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061174757607f821691505b60208210810361175a57611759611703565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026117bc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611781565b6117c68683611781565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61180a611805611800846117de565b6117e7565b6117de565b9050919050565b5f819050919050565b611823836117f0565b61183761182f82611811565b84845461178d565b825550505050565b5f90565b61184b61183f565b61185681848461181a565b505050565b5b818110156118795761186e5f82611843565b60018101905061185c565b5050565b601f8211156118be5761188f81611760565b61189884611772565b810160208510156118a7578190505b6118bb6118b385611772565b83018261185b565b50505b505050565b5f82821c905092915050565b5f6118de5f19846008026118c3565b1980831691505092915050565b5f6118f683836118cf565b9150826002028217905092915050565b61190f8261143f565b67ffffffffffffffff81111561192857611927611124565b5b6119328254611730565b61193d82828561187d565b5f60209050601f83116001811461196e575f841561195c578287015190505b61196685826118eb565b8655506119cd565b601f19841661197c86611760565b5f5b828110156119a35784890151825560018201915060208501945060208101905061197e565b868310156119c057848901516119bc601f8916826118cf565b8355505b6001600288020188555050505b505050505050565b5f8160601b9050919050565b5f6119eb826119d5565b9050919050565b5f6119fc826119e1565b9050919050565b611a14611a0f8261107e565b6119f2565b82525050565b5f611a258285611a03565b601482019150611a358284611a03565b6014820191508190509392505050565b5f81519050919050565b5f81905092915050565b5f611a6382611a45565b611a6d8185611a4f565b9350611a7d818560208601611459565b80840191505092915050565b5f819050919050565b5f819050919050565b611aac611aa782611a89565b611a92565b82525050565b5f611abd8285611a59565b9150611ac98284611a9b565b6020820191508190509392505050565b611ae281611a89565b82525050565b5f60ff82169050919050565b611afd81611ae8565b82525050565b5f608082019050611b165f830187611ad9565b611b236020830186611af4565b611b306040830185611ad9565b611b3d6060830184611ad9565b95945050505050565b7f696e76616c6964207369676e6174757265206c656e67746800000000000000005f82015250565b5f611b7a601883611449565b9150611b8582611b46565b602082019050919050565b5f6020820190508181035f830152611ba781611b6e565b905091905056fea26469706673582212207a6a39ecd11e693de75daa3b1d3082de8d6669ef04e06f14e456cd8403b2194164736f6c63430008140033",
}

// IdentityRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityRegistryMetaData.ABI instead.
var IdentityRegistryABI = IdentityRegistryMetaData.ABI

// IdentityRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IdentityRegistryMetaData.Bin instead.
var IdentityRegistryBin = IdentityRegistryMetaData.Bin

// DeployIdentityRegistry deploys a new Ethereum contract, binding an instance of IdentityRegistry to it.
func DeployIdentityRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IdentityRegistry, error) {
	parsed, err := IdentityRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IdentityRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// IdentityRegistry is an auto generated Go binding around an Ethereum contract.
type IdentityRegistry struct {
	IdentityRegistryCaller     // Read-only binding to the contract
	IdentityRegistryTransactor // Write-only binding to the contract
	IdentityRegistryFilterer   // Log filterer for contract events
}

// IdentityRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityRegistrySession struct {
	Contract     *IdentityRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityRegistryCallerSession struct {
	Contract *IdentityRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IdentityRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityRegistryTransactorSession struct {
	Contract     *IdentityRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IdentityRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRegistryRaw struct {
	Contract *IdentityRegistry // Generic contract binding to access the raw methods on
}

// IdentityRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityRegistryCallerRaw struct {
	Contract *IdentityRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactorRaw struct {
	Contract *IdentityRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityRegistry creates a new instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistry(address common.Address, backend bind.ContractBackend) (*IdentityRegistry, error) {
	contract, err := bindIdentityRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// NewIdentityRegistryCaller creates a new read-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryCaller(address common.Address, caller bind.ContractCaller) (*IdentityRegistryCaller, error) {
	contract, err := bindIdentityRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryCaller{contract: contract}, nil
}

// NewIdentityRegistryTransactor creates a new write-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityRegistryTransactor, error) {
	contract, err := bindIdentityRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryTransactor{contract: contract}, nil
}

// NewIdentityRegistryFilterer creates a new log filterer instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityRegistryFilterer, error) {
	contract, err := bindIdentityRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFilterer{contract: contract}, nil
}

// bindIdentityRegistry binds a generic wrapper to an already deployed contract.
func bindIdentityRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.IdentityRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transact(opts, method, params...)
}

// Circuits is a free data retrieval call binding the contract method 0x9ea2cd38.
//
// Solidity: function circuits(string ) view returns(string deploymentType, address deploymentAddress, string ipfsHash)
func (_IdentityRegistry *IdentityRegistryCaller) Circuits(opts *bind.CallOpts, arg0 string) (struct {
	DeploymentType    string
	DeploymentAddress common.Address
	IpfsHash          string
}, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "circuits", arg0)

	outstruct := new(struct {
		DeploymentType    string
		DeploymentAddress common.Address
		IpfsHash          string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DeploymentType = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DeploymentAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.IpfsHash = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Circuits is a free data retrieval call binding the contract method 0x9ea2cd38.
//
// Solidity: function circuits(string ) view returns(string deploymentType, address deploymentAddress, string ipfsHash)
func (_IdentityRegistry *IdentityRegistrySession) Circuits(arg0 string) (struct {
	DeploymentType    string
	DeploymentAddress common.Address
	IpfsHash          string
}, error) {
	return _IdentityRegistry.Contract.Circuits(&_IdentityRegistry.CallOpts, arg0)
}

// Circuits is a free data retrieval call binding the contract method 0x9ea2cd38.
//
// Solidity: function circuits(string ) view returns(string deploymentType, address deploymentAddress, string ipfsHash)
func (_IdentityRegistry *IdentityRegistryCallerSession) Circuits(arg0 string) (struct {
	DeploymentType    string
	DeploymentAddress common.Address
	IpfsHash          string
}, error) {
	return _IdentityRegistry.Contract.Circuits(&_IdentityRegistry.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address node) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) Exists(opts *bind.CallOpts, node common.Address) (bool, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "exists", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address node) view returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) Exists(node common.Address) (bool, error) {
	return _IdentityRegistry.Contract.Exists(&_IdentityRegistry.CallOpts, node)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address node) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) Exists(node common.Address) (bool, error) {
	return _IdentityRegistry.Contract.Exists(&_IdentityRegistry.CallOpts, node)
}

// GetCircuit is a free data retrieval call binding the contract method 0xe4baeda3.
//
// Solidity: function getCircuit(string _circuitId) view returns(string, address, string)
func (_IdentityRegistry *IdentityRegistryCaller) GetCircuit(opts *bind.CallOpts, _circuitId string) (string, common.Address, string, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "getCircuit", _circuitId)

	if err != nil {
		return *new(string), *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// GetCircuit is a free data retrieval call binding the contract method 0xe4baeda3.
//
// Solidity: function getCircuit(string _circuitId) view returns(string, address, string)
func (_IdentityRegistry *IdentityRegistrySession) GetCircuit(_circuitId string) (string, common.Address, string, error) {
	return _IdentityRegistry.Contract.GetCircuit(&_IdentityRegistry.CallOpts, _circuitId)
}

// GetCircuit is a free data retrieval call binding the contract method 0xe4baeda3.
//
// Solidity: function getCircuit(string _circuitId) view returns(string, address, string)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetCircuit(_circuitId string) (string, common.Address, string, error) {
	return _IdentityRegistry.Contract.GetCircuit(&_IdentityRegistry.CallOpts, _circuitId)
}

// Identities is a free data retrieval call binding the contract method 0xf653b81e.
//
// Solidity: function identities(address ) view returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) Identities(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "identities", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Identities is a free data retrieval call binding the contract method 0xf653b81e.
//
// Solidity: function identities(address ) view returns(address)
func (_IdentityRegistry *IdentityRegistrySession) Identities(arg0 common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.Identities(&_IdentityRegistry.CallOpts, arg0)
}

// Identities is a free data retrieval call binding the contract method 0xf653b81e.
//
// Solidity: function identities(address ) view returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) Identities(arg0 common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.Identities(&_IdentityRegistry.CallOpts, arg0)
}

// Manager is a free data retrieval call binding the contract method 0xd4d2e7f2.
//
// Solidity: function manager(address node) view returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) Manager(opts *bind.CallOpts, node common.Address) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "manager", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0xd4d2e7f2.
//
// Solidity: function manager(address node) view returns(address)
func (_IdentityRegistry *IdentityRegistrySession) Manager(node common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.Manager(&_IdentityRegistry.CallOpts, node)
}

// Manager is a free data retrieval call binding the contract method 0xd4d2e7f2.
//
// Solidity: function manager(address node) view returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) Manager(node common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.Manager(&_IdentityRegistry.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityRegistry *IdentityRegistrySession) Owner() (common.Address, error) {
	return _IdentityRegistry.Contract.Owner(&_IdentityRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) Owner() (common.Address, error) {
	return _IdentityRegistry.Contract.Owner(&_IdentityRegistry.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address node) view returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) OwnerOf(opts *bind.CallOpts, node common.Address) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "ownerOf", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address node) view returns(address)
func (_IdentityRegistry *IdentityRegistrySession) OwnerOf(node common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.OwnerOf(&_IdentityRegistry.CallOpts, node)
}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address node) view returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) OwnerOf(node common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.OwnerOf(&_IdentityRegistry.CallOpts, node)
}

// Running is a free data retrieval call binding the contract method 0xd85bd526.
//
// Solidity: function running() view returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) Running(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "running")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Running is a free data retrieval call binding the contract method 0xd85bd526.
//
// Solidity: function running() view returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) Running() (bool, error) {
	return _IdentityRegistry.Contract.Running(&_IdentityRegistry.CallOpts)
}

// Running is a free data retrieval call binding the contract method 0xd85bd526.
//
// Solidity: function running() view returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) Running() (bool, error) {
	return _IdentityRegistry.Contract.Running(&_IdentityRegistry.CallOpts)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address node) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) Deregister(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "deregister", node)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address node) returns()
func (_IdentityRegistry *IdentityRegistrySession) Deregister(node common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Deregister(&_IdentityRegistry.TransactOpts, node)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address node) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) Deregister(node common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Deregister(&_IdentityRegistry.TransactOpts, node)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_IdentityRegistry *IdentityRegistryTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_IdentityRegistry *IdentityRegistrySession) Destroy() (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Destroy(&_IdentityRegistry.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) Destroy() (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Destroy(&_IdentityRegistry.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address identityManagerContract, bytes signature) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) Register(opts *bind.TransactOpts, identityManagerContract common.Address, signature []byte) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "register", identityManagerContract, signature)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address identityManagerContract, bytes signature) returns()
func (_IdentityRegistry *IdentityRegistrySession) Register(identityManagerContract common.Address, signature []byte) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Register(&_IdentityRegistry.TransactOpts, identityManagerContract, signature)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address identityManagerContract, bytes signature) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) Register(identityManagerContract common.Address, signature []byte) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Register(&_IdentityRegistry.TransactOpts, identityManagerContract, signature)
}

// SetCircuit is a paid mutator transaction binding the contract method 0x5208ef6a.
//
// Solidity: function setCircuit(string _circuitId, string _deploymentType, address _deploymentAddress, string _ipfsHash) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) SetCircuit(opts *bind.TransactOpts, _circuitId string, _deploymentType string, _deploymentAddress common.Address, _ipfsHash string) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "setCircuit", _circuitId, _deploymentType, _deploymentAddress, _ipfsHash)
}

// SetCircuit is a paid mutator transaction binding the contract method 0x5208ef6a.
//
// Solidity: function setCircuit(string _circuitId, string _deploymentType, address _deploymentAddress, string _ipfsHash) returns()
func (_IdentityRegistry *IdentityRegistrySession) SetCircuit(_circuitId string, _deploymentType string, _deploymentAddress common.Address, _ipfsHash string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetCircuit(&_IdentityRegistry.TransactOpts, _circuitId, _deploymentType, _deploymentAddress, _ipfsHash)
}

// SetCircuit is a paid mutator transaction binding the contract method 0x5208ef6a.
//
// Solidity: function setCircuit(string _circuitId, string _deploymentType, address _deploymentAddress, string _ipfsHash) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) SetCircuit(_circuitId string, _deploymentType string, _deploymentAddress common.Address, _ipfsHash string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetCircuit(&_IdentityRegistry.TransactOpts, _circuitId, _deploymentType, _deploymentAddress, _ipfsHash)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) SetManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "setManager", manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager) returns()
func (_IdentityRegistry *IdentityRegistrySession) SetManager(manager common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetManager(&_IdentityRegistry.TransactOpts, manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) SetManager(manager common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetManager(&_IdentityRegistry.TransactOpts, manager)
}

// IdentityRegistryNewManagerIterator is returned from FilterNewManager and is used to iterate over the raw logs and unpacked data for NewManager events raised by the IdentityRegistry contract.
type IdentityRegistryNewManagerIterator struct {
	Event *IdentityRegistryNewManager // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryNewManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryNewManager)
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
		it.Event = new(IdentityRegistryNewManager)
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
func (it *IdentityRegistryNewManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryNewManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryNewManager represents a NewManager event raised by the IdentityRegistry contract.
type IdentityRegistryNewManager struct {
	Node    common.Address
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewManager is a free log retrieval operation binding the contract event 0x770e6248a70b6ac757edf422766216da592c37e3112db900fe0da8984191831b.
//
// Solidity: event NewManager(address indexed node, address manager)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterNewManager(opts *bind.FilterOpts, node []common.Address) (*IdentityRegistryNewManagerIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "NewManager", nodeRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryNewManagerIterator{contract: _IdentityRegistry.contract, event: "NewManager", logs: logs, sub: sub}, nil
}

// WatchNewManager is a free log subscription operation binding the contract event 0x770e6248a70b6ac757edf422766216da592c37e3112db900fe0da8984191831b.
//
// Solidity: event NewManager(address indexed node, address manager)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchNewManager(opts *bind.WatchOpts, sink chan<- *IdentityRegistryNewManager, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "NewManager", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryNewManager)
				if err := _IdentityRegistry.contract.UnpackLog(event, "NewManager", log); err != nil {
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
func (_IdentityRegistry *IdentityRegistryFilterer) ParseNewManager(log types.Log) (*IdentityRegistryNewManager, error) {
	event := new(IdentityRegistryNewManager)
	if err := _IdentityRegistry.contract.UnpackLog(event, "NewManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
