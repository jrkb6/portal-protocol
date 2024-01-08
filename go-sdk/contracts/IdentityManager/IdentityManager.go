// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IdentityManager

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

// IdentityManagerMetaData contains all meta data concerning the IdentityManager contract.
var IdentityManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIdentityInterface\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"AttestationRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NewAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPrivate\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"version\",\"type\":\"int256\"}],\"name\":\"NewClaim\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"attestations\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"attestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"claimId\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isPrivate\",\"type\":\"bool\"}],\"name\":\"deleteClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"deleteClaimURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ipfsClaims\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"privateClaims\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"statement\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsCircuitMetadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"eventHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"version\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"publicClaims\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"statement\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"version\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIdentityInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"revocations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"attestedTo\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"attestationId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"attestedTo\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"attestationId\",\"type\":\"string\"}],\"name\":\"revokeAttestation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"attestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"claimId\",\"type\":\"string\"}],\"name\":\"setAttestation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setClaimURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"statement\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"eventHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"setPrivateClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"statement\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"setPublicClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162002a8b38038062002a8b833981810160405281019062000036919062000134565b3360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000164565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620000eb82620000c0565b9050919050565b5f620000fe82620000df565b9050919050565b6200011081620000f2565b81146200011b575f80fd5b50565b5f815190506200012e8162000105565b92915050565b5f602082840312156200014c576200014b620000bc565b5b5f6200015b848285016200011e565b91505092915050565b61291980620001725f395ff3fe608060405234801561000f575f80fd5b50600436106100e8575f3560e01c80637b1039991161008a57806389285c211161006457806389285c21146102625780638da5cb5b1461027e578063c03bc8b51461029c578063ee10c7e2146102b8576100e8565b80637b103999146101f55780638107769b1461021357806384ce394c14610246576100e8565b806345ddaf5d116100c657806345ddaf5d146101585780634665b258146101885780636066b5c3146101a4578063755af686146101c0576100e8565b8063249d026b146100ec57806327f541f81461012057806331b6399a1461013c575b5f80fd5b610106600480360381019061010191906117c0565b6102eb565b60405161011795949392919061192a565b60405180910390f35b61013a60048036038101906101359190611a10565b610460565b005b61015660048036038101906101519190611b53565b61063e565b005b610172600480360381019061016d91906117c0565b61088a565b60405161017f9190611c2a565b60405180910390f35b6101a2600480360381019061019d9190611c7f565b61093d565b005b6101be60048036038101906101b99190611cdc565b610a60565b005b6101da60048036038101906101d591906117c0565b610cb1565b6040516101ec96959493929190611db8565b60405180910390f35b6101fd610f19565b60405161020a9190611e8e565b60405180910390f35b61022d600480360381019061022891906117c0565b610f3c565b60405161023d9493929190611ea7565b60405180910390f35b610260600480360381019061025b9190611ef8565b6110ab565b005b61027c60048036038101906102779190611f43565b611132565b005b6102866113d7565b6040516102939190612069565b60405180910390f35b6102b660048036038101906102b19190612082565b6113fc565b005b6102d260048036038101906102cd91906117c0565b6114cb565b6040516102e29493929190612100565b60405180910390f35b6006818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0180546103229061217e565b80601f016020809104026020016040519081016040528092919081815260200182805461034e9061217e565b80156103995780601f1061037057610100808354040283529160200191610399565b820191905f5260205f20905b81548152906001019060200180831161037c57829003601f168201915b505050505090806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060040180546103df9061217e565b80601f016020809104026020016040519081016040528092919081815260200182805461040b9061217e565b80156104565780601f1061042d57610100808354040283529160200191610456565b820191905f5260205f20905b81548152906001019060200180831161043957829003601f168201915b5050505050905085565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104b8575f80fd5b5f6001600389896040516104cd9291906121dc565b9081526020016040518091039020600301546104e99190612221565b9050604051806080016040528087878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200183815260200182815250600389896040516105a49291906121dc565b90815260200160405180910390205f820151815f0190816105c591906123f6565b5060208201518160010190816105db91906123f6565b5060408201518160020155606082015181600301559050507f26f9b425c95c58e8bec57a1841c65b09f8b1b4408004758fcebf088adc4b3d62888888885f8660405161062c96959493929190612500565b60405180910390a15050505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610696575f80fd5b5f600689896040516106a99291906121dc565b908152602001604051809103902060030154146106c4575f80fd5b6040518060a0016040528085858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018773ffffffffffffffffffffffffffffffffffffffff16815260200186815260200142815260200183838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815250600689896040516107999291906121dc565b90815260200160405180910390205f820151815f0190816107ba91906125ad565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015560608201518160030155608082015181600401908161082a91906123f6565b509050508573ffffffffffffffffffffffffffffffffffffffff167f9476865afb170ca0b75c734e92ad4be2fa603b741a4c4542361fe02f8ff9ae4a8989886040516108789392919061267c565b60405180910390a25050505050505050565b6004818051602081018201805184825260208301602085012081835280955050505050505f9150905080546108be9061217e565b80601f01602080910402602001604051908101604052809291908181526020018280546108ea9061217e565b80156109355780601f1061090c57610100808354040283529160200191610935565b820191905f5260205f20905b81548152906001019060200180831161091857829003601f168201915b505050505081565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610995575f80fd5b8015610a0c57600283836040516109ad9291906121dc565b90815260200160405180910390205f8082015f6109ca919061161b565b600182015f6109d9919061161b565b600282015f6109e8919061161b565b600382015f6109f7919061161b565b600482015f9055600582015f90555050610a5b565b60038383604051610a1e9291906121dc565b90815260200160405180910390205f8082015f610a3b919061161b565b600182015f610a4a919061161b565b600282015f9055600382015f905550505b505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ab8575f80fd5b60405180608001604052808473ffffffffffffffffffffffffffffffffffffffff16815260200183838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020014281525060058888604051610b879291906121dc565b90815260200160405180910390205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081610bee91906123f6565b506040820151816002019081610c0491906123f6565b506060820151816003015590505060068787604051610c249291906121dc565b90815260200160405180910390206001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f1e926d3070e51ea520634ecc30c671c6d1363a87da2c8af5f2aa75b6bcfae1a288888888604051610ca094939291906126ac565b60405180910390a250505050505050565b6002818051602081018201805184825260208301602085012081835280955050505050505f91509050805f018054610ce89061217e565b80601f0160208091040260200160405190810160405280929190818152602001828054610d149061217e565b8015610d5f5780601f10610d3657610100808354040283529160200191610d5f565b820191905f5260205f20905b815481529060010190602001808311610d4257829003601f168201915b505050505090806001018054610d749061217e565b80601f0160208091040260200160405190810160405280929190818152602001828054610da09061217e565b8015610deb5780601f10610dc257610100808354040283529160200191610deb565b820191905f5260205f20905b815481529060010190602001808311610dce57829003601f168201915b505050505090806002018054610e009061217e565b80601f0160208091040260200160405190810160405280929190818152602001828054610e2c9061217e565b8015610e775780601f10610e4e57610100808354040283529160200191610e77565b820191905f5260205f20905b815481529060010190602001808311610e5a57829003601f168201915b505050505090806003018054610e8c9061217e565b80601f0160208091040260200160405190810160405280929190818152602001828054610eb89061217e565b8015610f035780601f10610eda57610100808354040283529160200191610f03565b820191905f5260205f20905b815481529060010190602001808311610ee657829003601f168201915b5050505050908060040154908060050154905086565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6005818051602081018201805184825260208301602085012081835280955050505050505f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054610f989061217e565b80601f0160208091040260200160405190810160405280929190818152602001828054610fc49061217e565b801561100f5780601f10610fe65761010080835404028352916020019161100f565b820191905f5260205f20905b815481529060010190602001808311610ff257829003601f168201915b5050505050908060020180546110249061217e565b80601f01602080910402602001604051908101604052809291908181526020018280546110509061217e565b801561109b5780601f106110725761010080835404028352916020019161109b565b820191905f5260205f20905b81548152906001019060200180831161107e57829003601f168201915b5050505050908060030154905084565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611103575f80fd5b600482826040516111159291906121dc565b90815260200160405180910390205f61112e919061161b565b5050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461118a575f80fd5b5f60028b60405161119b9190612715565b90815260200160405180910390206005015490506040518060c001604052808b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200189898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018381526020016001836112f19190612221565b81525060028c6040516113049190612715565b90815260200160405180910390205f820151815f01908161132591906123f6565b50602082015181600101908161133b91906123f6565b50604082015181600201908161135191906123f6565b50606082015181600301908161136791906123f6565b506080820151816004015560a082015181600501559050507f26f9b425c95c58e8bec57a1841c65b09f8b1b4408004758fcebf088adc4b3d628b8b8b600180866113b19190612221565b6040516113c295949392919061272b565b60405180910390a15050505050505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611454575f80fd5b8181600486866040516114689291906121dc565b90815260200160405180910390209182611483929190612788565b507f26f9b425c95c58e8bec57a1841c65b09f8b1b4408004758fcebf088adc4b3d62848484845f806040516114bd9695949392919061288e565b60405180910390a150505050565b6003818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0180546115029061217e565b80601f016020809104026020016040519081016040528092919081815260200182805461152e9061217e565b80156115795780601f1061155057610100808354040283529160200191611579565b820191905f5260205f20905b81548152906001019060200180831161155c57829003601f168201915b50505050509080600101805461158e9061217e565b80601f01602080910402602001604051908101604052809291908181526020018280546115ba9061217e565b80156116055780601f106115dc57610100808354040283529160200191611605565b820191905f5260205f20905b8154815290600101906020018083116115e857829003601f168201915b5050505050908060020154908060030154905084565b5080546116279061217e565b5f825580601f106116385750611655565b601f0160209004905f5260205f20908101906116549190611658565b5b50565b5b8082111561166f575f815f905550600101611659565b5090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6116d28261168c565b810181811067ffffffffffffffff821117156116f1576116f061169c565b5b80604052505050565b5f611703611673565b905061170f82826116c9565b919050565b5f67ffffffffffffffff82111561172e5761172d61169c565b5b6117378261168c565b9050602081019050919050565b828183375f83830152505050565b5f61176461175f84611714565b6116fa565b9050828152602081018484840111156117805761177f611688565b5b61178b848285611744565b509392505050565b5f82601f8301126117a7576117a6611684565b5b81356117b7848260208601611752565b91505092915050565b5f602082840312156117d5576117d461167c565b5b5f82013567ffffffffffffffff8111156117f2576117f1611680565b5b6117fe84828501611793565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561183e578082015181840152602081019050611823565b5f8484015250505050565b5f61185382611807565b61185d8185611811565b935061186d818560208601611821565b6118768161168c565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6118aa82611881565b9050919050565b6118ba816118a0565b82525050565b5f819050919050565b6118d2816118c0565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f6118fc826118d8565b61190681856118e2565b9350611916818560208601611821565b61191f8161168c565b840191505092915050565b5f60a0820190508181035f8301526119428188611849565b905061195160208301876118b1565b61195e60408301866118c9565b61196b60608301856118c9565b818103608083015261197d81846118f2565b90509695505050505050565b5f80fd5b5f80fd5b5f8083601f8401126119a6576119a5611684565b5b8235905067ffffffffffffffff8111156119c3576119c2611989565b5b6020830191508360018202830111156119df576119de61198d565b5b9250929050565b6119ef816118c0565b81146119f9575f80fd5b50565b5f81359050611a0a816119e6565b92915050565b5f805f805f805f6080888a031215611a2b57611a2a61167c565b5b5f88013567ffffffffffffffff811115611a4857611a47611680565b5b611a548a828b01611991565b9750975050602088013567ffffffffffffffff811115611a7757611a76611680565b5b611a838a828b01611991565b9550955050604088013567ffffffffffffffff811115611aa657611aa5611680565b5b611ab28a828b01611991565b93509350506060611ac58a828b016119fc565b91505092959891949750929550565b611add816118a0565b8114611ae7575f80fd5b50565b5f81359050611af881611ad4565b92915050565b5f8083601f840112611b1357611b12611684565b5b8235905067ffffffffffffffff811115611b3057611b2f611989565b5b602083019150836001820283011115611b4c57611b4b61198d565b5b9250929050565b5f805f805f805f8060a0898b031215611b6f57611b6e61167c565b5b5f89013567ffffffffffffffff811115611b8c57611b8b611680565b5b611b988b828c01611991565b98509850506020611bab8b828c01611aea565b9650506040611bbc8b828c016119fc565b955050606089013567ffffffffffffffff811115611bdd57611bdc611680565b5b611be98b828c01611afe565b9450945050608089013567ffffffffffffffff811115611c0c57611c0b611680565b5b611c188b828c01611991565b92509250509295985092959890939650565b5f6020820190508181035f830152611c4281846118f2565b905092915050565b5f8115159050919050565b611c5e81611c4a565b8114611c68575f80fd5b50565b5f81359050611c7981611c55565b92915050565b5f805f60408486031215611c9657611c9561167c565b5b5f84013567ffffffffffffffff811115611cb357611cb2611680565b5b611cbf86828701611991565b93509350506020611cd286828701611c6b565b9150509250925092565b5f805f805f805f6080888a031215611cf757611cf661167c565b5b5f88013567ffffffffffffffff811115611d1457611d13611680565b5b611d208a828b01611991565b9750975050602088013567ffffffffffffffff811115611d4357611d42611680565b5b611d4f8a828b01611991565b95509550506040611d628a828b01611aea565b935050606088013567ffffffffffffffff811115611d8357611d82611680565b5b611d8f8a828b01611991565b925092505092959891949750929550565b5f819050919050565b611db281611da0565b82525050565b5f60c0820190508181035f830152611dd081896118f2565b90508181036020830152611de481886118f2565b90508181036040830152611df881876118f2565b90508181036060830152611e0c81866118f2565b9050611e1b60808301856118c9565b611e2860a0830184611da9565b979650505050505050565b5f819050919050565b5f611e56611e51611e4c84611881565b611e33565b611881565b9050919050565b5f611e6782611e3c565b9050919050565b5f611e7882611e5d565b9050919050565b611e8881611e6e565b82525050565b5f602082019050611ea15f830184611e7f565b92915050565b5f608082019050611eba5f8301876118b1565b8181036020830152611ecc81866118f2565b90508181036040830152611ee081856118f2565b9050611eef60608301846118c9565b95945050505050565b5f8060208385031215611f0e57611f0d61167c565b5b5f83013567ffffffffffffffff811115611f2b57611f2a611680565b5b611f3785828601611991565b92509250509250929050565b5f805f805f805f805f8060c08b8d031215611f6157611f6061167c565b5b5f8b013567ffffffffffffffff811115611f7e57611f7d611680565b5b611f8a8d828e01611793565b9a505060208b013567ffffffffffffffff811115611fab57611faa611680565b5b611fb78d828e01611991565b995099505060408b013567ffffffffffffffff811115611fda57611fd9611680565b5b611fe68d828e01611991565b975097505060608b013567ffffffffffffffff81111561200957612008611680565b5b6120158d828e01611991565b955095505060808b013567ffffffffffffffff81111561203857612037611680565b5b6120448d828e01611991565b935093505060a06120578d828e016119fc565b9150509295989b9194979a5092959850565b5f60208201905061207c5f8301846118b1565b92915050565b5f805f806040858703121561209a5761209961167c565b5b5f85013567ffffffffffffffff8111156120b7576120b6611680565b5b6120c387828801611991565b9450945050602085013567ffffffffffffffff8111156120e6576120e5611680565b5b6120f287828801611991565b925092505092959194509250565b5f6080820190508181035f83015261211881876118f2565b9050818103602083015261212c81866118f2565b905061213b60408301856118c9565b6121486060830184611da9565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061219557607f821691505b6020821081036121a8576121a7612151565b5b50919050565b5f81905092915050565b5f6121c383856121ae565b93506121d0838584611744565b82840190509392505050565b5f6121e88284866121b8565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61222b82611da0565b915061223683611da0565b92508282019050828112155f8312168382125f84121516171561225c5761225b6121f4565b5b92915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026122be7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612283565b6122c88683612283565b95508019841693508086168417925050509392505050565b5f6122fa6122f56122f0846118c0565b611e33565b6118c0565b9050919050565b5f819050919050565b612313836122e0565b61232761231f82612301565b84845461228f565b825550505050565b5f90565b61233b61232f565b61234681848461230a565b505050565b5b818110156123695761235e5f82612333565b60018101905061234c565b5050565b601f8211156123ae5761237f81612262565b61238884612274565b81016020851015612397578190505b6123ab6123a385612274565b83018261234b565b50505b505050565b5f82821c905092915050565b5f6123ce5f19846008026123b3565b1980831691505092915050565b5f6123e683836123bf565b9150826002028217905092915050565b6123ff826118d8565b67ffffffffffffffff8111156124185761241761169c565b5b612422825461217e565b61242d82828561236d565b5f60209050601f83116001811461245e575f841561244c578287015190505b61245685826123db565b8655506124bd565b601f19841661246c86612262565b5f5b828110156124935784890151825560018201915060208501945060208101905061246e565b868310156124b057848901516124ac601f8916826123bf565b8355505b6001600288020188555050505b505050505050565b5f6124d083856118e2565b93506124dd838584611744565b6124e68361168c565b840190509392505050565b6124fa81611c4a565b82525050565b5f6080820190508181035f83015261251981888a6124c5565b9050818103602083015261252e8186886124c5565b905061253d60408301856124f1565b61254a6060830184611da9565b979650505050505050565b5f819050815f5260205f209050919050565b601f8211156125a85761257981612555565b61258284612274565b81016020851015612591578190505b6125a561259d85612274565b83018261234b565b50505b505050565b6125b682611807565b67ffffffffffffffff8111156125cf576125ce61169c565b5b6125d9825461217e565b6125e4828285612567565b5f60209050601f831160018114612615575f8415612603578287015190505b61260d85826123db565b865550612674565b601f19841661262386612555565b5f5b8281101561264a57848901518255600182019150602085019450602081019050612625565b868310156126675784890151612663601f8916826123bf565b8355505b6001600288020188555050505b505050505050565b5f6040820190508181035f8301526126958185876124c5565b90506126a460208301846118c9565b949350505050565b5f6040820190508181035f8301526126c58186886124c5565b905081810360208301526126da8184866124c5565b905095945050505050565b5f6126ef826118d8565b6126f981856121ae565b9350612709818560208601611821565b80840191505092915050565b5f61272082846126e5565b915081905092915050565b5f6080820190508181035f83015261274381886118f2565b905081810360208301526127588186886124c5565b905061276760408301856124f1565b6127746060830184611da9565b9695505050505050565b5f82905092915050565b612792838361277e565b67ffffffffffffffff8111156127ab576127aa61169c565b5b6127b5825461217e565b6127c082828561236d565b5f601f8311600181146127ed575f84156127db578287013590505b6127e585826123db565b86555061284c565b601f1984166127fb86612262565b5f5b82811015612822578489013582556001820191506020850194506020810190506127fd565b8683101561283f578489013561283b601f8916826123bf565b8355505b6001600288020188555050505b50505050505050565b5f819050919050565b5f61287861287361286e84612855565b611e33565b611da0565b9050919050565b6128888161285e565b82525050565b5f6080820190508181035f8301526128a781888a6124c5565b905081810360208301526128bc8186886124c5565b90506128cb60408301856124f1565b6128d8606083018461287f565b97965050505050505056fea26469706673582212201385d680dbdd0fc7dbfaca2f9f4c61fd56bf31203958a0f146f4d2a73c982f5464736f6c63430008140033",
}

// IdentityManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityManagerMetaData.ABI instead.
var IdentityManagerABI = IdentityManagerMetaData.ABI

// IdentityManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IdentityManagerMetaData.Bin instead.
var IdentityManagerBin = IdentityManagerMetaData.Bin

// DeployIdentityManager deploys a new Ethereum contract, binding an instance of IdentityManager to it.
func DeployIdentityManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registry common.Address) (common.Address, *types.Transaction, *IdentityManager, error) {
	parsed, err := IdentityManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IdentityManagerBin), backend, _registry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityManager{IdentityManagerCaller: IdentityManagerCaller{contract: contract}, IdentityManagerTransactor: IdentityManagerTransactor{contract: contract}, IdentityManagerFilterer: IdentityManagerFilterer{contract: contract}}, nil
}

// IdentityManager is an auto generated Go binding around an Ethereum contract.
type IdentityManager struct {
	IdentityManagerCaller     // Read-only binding to the contract
	IdentityManagerTransactor // Write-only binding to the contract
	IdentityManagerFilterer   // Log filterer for contract events
}

// IdentityManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityManagerSession struct {
	Contract     *IdentityManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityManagerCallerSession struct {
	Contract *IdentityManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IdentityManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityManagerTransactorSession struct {
	Contract     *IdentityManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IdentityManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityManagerRaw struct {
	Contract *IdentityManager // Generic contract binding to access the raw methods on
}

// IdentityManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityManagerCallerRaw struct {
	Contract *IdentityManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityManagerTransactorRaw struct {
	Contract *IdentityManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityManager creates a new instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManager(address common.Address, backend bind.ContractBackend) (*IdentityManager, error) {
	contract, err := bindIdentityManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityManager{IdentityManagerCaller: IdentityManagerCaller{contract: contract}, IdentityManagerTransactor: IdentityManagerTransactor{contract: contract}, IdentityManagerFilterer: IdentityManagerFilterer{contract: contract}}, nil
}

// NewIdentityManagerCaller creates a new read-only instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManagerCaller(address common.Address, caller bind.ContractCaller) (*IdentityManagerCaller, error) {
	contract, err := bindIdentityManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerCaller{contract: contract}, nil
}

// NewIdentityManagerTransactor creates a new write-only instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityManagerTransactor, error) {
	contract, err := bindIdentityManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerTransactor{contract: contract}, nil
}

// NewIdentityManagerFilterer creates a new log filterer instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityManagerFilterer, error) {
	contract, err := bindIdentityManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerFilterer{contract: contract}, nil
}

// bindIdentityManager binds a generic wrapper to an already deployed contract.
func bindIdentityManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityManager *IdentityManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityManager.Contract.IdentityManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityManager *IdentityManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityManager.Contract.IdentityManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityManager *IdentityManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityManager.Contract.IdentityManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityManager *IdentityManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityManager *IdentityManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityManager *IdentityManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityManager.Contract.contract.Transact(opts, method, params...)
}

// Attestations is a free data retrieval call binding the contract method 0x249d026b.
//
// Solidity: function attestations(string ) view returns(bytes signature, address attestor, uint256 expires, uint256 timestamp, string claimId)
func (_IdentityManager *IdentityManagerCaller) Attestations(opts *bind.CallOpts, arg0 string) (struct {
	Signature []byte
	Attestor  common.Address
	Expires   *big.Int
	Timestamp *big.Int
	ClaimId   string
}, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "attestations", arg0)

	outstruct := new(struct {
		Signature []byte
		Attestor  common.Address
		Expires   *big.Int
		Timestamp *big.Int
		ClaimId   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Signature = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Attestor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Expires = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ClaimId = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// Attestations is a free data retrieval call binding the contract method 0x249d026b.
//
// Solidity: function attestations(string ) view returns(bytes signature, address attestor, uint256 expires, uint256 timestamp, string claimId)
func (_IdentityManager *IdentityManagerSession) Attestations(arg0 string) (struct {
	Signature []byte
	Attestor  common.Address
	Expires   *big.Int
	Timestamp *big.Int
	ClaimId   string
}, error) {
	return _IdentityManager.Contract.Attestations(&_IdentityManager.CallOpts, arg0)
}

// Attestations is a free data retrieval call binding the contract method 0x249d026b.
//
// Solidity: function attestations(string ) view returns(bytes signature, address attestor, uint256 expires, uint256 timestamp, string claimId)
func (_IdentityManager *IdentityManagerCallerSession) Attestations(arg0 string) (struct {
	Signature []byte
	Attestor  common.Address
	Expires   *big.Int
	Timestamp *big.Int
	ClaimId   string
}, error) {
	return _IdentityManager.Contract.Attestations(&_IdentityManager.CallOpts, arg0)
}

// IpfsClaims is a free data retrieval call binding the contract method 0x45ddaf5d.
//
// Solidity: function ipfsClaims(string ) view returns(string)
func (_IdentityManager *IdentityManagerCaller) IpfsClaims(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "ipfsClaims", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// IpfsClaims is a free data retrieval call binding the contract method 0x45ddaf5d.
//
// Solidity: function ipfsClaims(string ) view returns(string)
func (_IdentityManager *IdentityManagerSession) IpfsClaims(arg0 string) (string, error) {
	return _IdentityManager.Contract.IpfsClaims(&_IdentityManager.CallOpts, arg0)
}

// IpfsClaims is a free data retrieval call binding the contract method 0x45ddaf5d.
//
// Solidity: function ipfsClaims(string ) view returns(string)
func (_IdentityManager *IdentityManagerCallerSession) IpfsClaims(arg0 string) (string, error) {
	return _IdentityManager.Contract.IpfsClaims(&_IdentityManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityManager *IdentityManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityManager *IdentityManagerSession) Owner() (common.Address, error) {
	return _IdentityManager.Contract.Owner(&_IdentityManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityManager *IdentityManagerCallerSession) Owner() (common.Address, error) {
	return _IdentityManager.Contract.Owner(&_IdentityManager.CallOpts)
}

// PrivateClaims is a free data retrieval call binding the contract method 0x755af686.
//
// Solidity: function privateClaims(string ) view returns(string value, string statement, string ipfsCircuitMetadata, string eventHash, uint256 timestamp, int256 version)
func (_IdentityManager *IdentityManagerCaller) PrivateClaims(opts *bind.CallOpts, arg0 string) (struct {
	Value               string
	Statement           string
	IpfsCircuitMetadata string
	EventHash           string
	Timestamp           *big.Int
	Version             *big.Int
}, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "privateClaims", arg0)

	outstruct := new(struct {
		Value               string
		Statement           string
		IpfsCircuitMetadata string
		EventHash           string
		Timestamp           *big.Int
		Version             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Statement = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.IpfsCircuitMetadata = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.EventHash = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Version = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PrivateClaims is a free data retrieval call binding the contract method 0x755af686.
//
// Solidity: function privateClaims(string ) view returns(string value, string statement, string ipfsCircuitMetadata, string eventHash, uint256 timestamp, int256 version)
func (_IdentityManager *IdentityManagerSession) PrivateClaims(arg0 string) (struct {
	Value               string
	Statement           string
	IpfsCircuitMetadata string
	EventHash           string
	Timestamp           *big.Int
	Version             *big.Int
}, error) {
	return _IdentityManager.Contract.PrivateClaims(&_IdentityManager.CallOpts, arg0)
}

// PrivateClaims is a free data retrieval call binding the contract method 0x755af686.
//
// Solidity: function privateClaims(string ) view returns(string value, string statement, string ipfsCircuitMetadata, string eventHash, uint256 timestamp, int256 version)
func (_IdentityManager *IdentityManagerCallerSession) PrivateClaims(arg0 string) (struct {
	Value               string
	Statement           string
	IpfsCircuitMetadata string
	EventHash           string
	Timestamp           *big.Int
	Version             *big.Int
}, error) {
	return _IdentityManager.Contract.PrivateClaims(&_IdentityManager.CallOpts, arg0)
}

// PublicClaims is a free data retrieval call binding the contract method 0xee10c7e2.
//
// Solidity: function publicClaims(string ) view returns(string value, string statement, uint256 timestamp, int256 version)
func (_IdentityManager *IdentityManagerCaller) PublicClaims(opts *bind.CallOpts, arg0 string) (struct {
	Value     string
	Statement string
	Timestamp *big.Int
	Version   *big.Int
}, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "publicClaims", arg0)

	outstruct := new(struct {
		Value     string
		Statement string
		Timestamp *big.Int
		Version   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Statement = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Version = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PublicClaims is a free data retrieval call binding the contract method 0xee10c7e2.
//
// Solidity: function publicClaims(string ) view returns(string value, string statement, uint256 timestamp, int256 version)
func (_IdentityManager *IdentityManagerSession) PublicClaims(arg0 string) (struct {
	Value     string
	Statement string
	Timestamp *big.Int
	Version   *big.Int
}, error) {
	return _IdentityManager.Contract.PublicClaims(&_IdentityManager.CallOpts, arg0)
}

// PublicClaims is a free data retrieval call binding the contract method 0xee10c7e2.
//
// Solidity: function publicClaims(string ) view returns(string value, string statement, uint256 timestamp, int256 version)
func (_IdentityManager *IdentityManagerCallerSession) PublicClaims(arg0 string) (struct {
	Value     string
	Statement string
	Timestamp *big.Int
	Version   *big.Int
}, error) {
	return _IdentityManager.Contract.PublicClaims(&_IdentityManager.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_IdentityManager *IdentityManagerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_IdentityManager *IdentityManagerSession) Registry() (common.Address, error) {
	return _IdentityManager.Contract.Registry(&_IdentityManager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_IdentityManager *IdentityManagerCallerSession) Registry() (common.Address, error) {
	return _IdentityManager.Contract.Registry(&_IdentityManager.CallOpts)
}

// Revocations is a free data retrieval call binding the contract method 0x8107769b.
//
// Solidity: function revocations(string ) view returns(address attestedTo, string attestationId, string status, uint256 timestamp)
func (_IdentityManager *IdentityManagerCaller) Revocations(opts *bind.CallOpts, arg0 string) (struct {
	AttestedTo    common.Address
	AttestationId string
	Status        string
	Timestamp     *big.Int
}, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "revocations", arg0)

	outstruct := new(struct {
		AttestedTo    common.Address
		AttestationId string
		Status        string
		Timestamp     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttestedTo = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AttestationId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Revocations is a free data retrieval call binding the contract method 0x8107769b.
//
// Solidity: function revocations(string ) view returns(address attestedTo, string attestationId, string status, uint256 timestamp)
func (_IdentityManager *IdentityManagerSession) Revocations(arg0 string) (struct {
	AttestedTo    common.Address
	AttestationId string
	Status        string
	Timestamp     *big.Int
}, error) {
	return _IdentityManager.Contract.Revocations(&_IdentityManager.CallOpts, arg0)
}

// Revocations is a free data retrieval call binding the contract method 0x8107769b.
//
// Solidity: function revocations(string ) view returns(address attestedTo, string attestationId, string status, uint256 timestamp)
func (_IdentityManager *IdentityManagerCallerSession) Revocations(arg0 string) (struct {
	AttestedTo    common.Address
	AttestationId string
	Status        string
	Timestamp     *big.Int
}, error) {
	return _IdentityManager.Contract.Revocations(&_IdentityManager.CallOpts, arg0)
}

// DeleteClaim is a paid mutator transaction binding the contract method 0x4665b258.
//
// Solidity: function deleteClaim(string key, bool isPrivate) returns()
func (_IdentityManager *IdentityManagerTransactor) DeleteClaim(opts *bind.TransactOpts, key string, isPrivate bool) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "deleteClaim", key, isPrivate)
}

// DeleteClaim is a paid mutator transaction binding the contract method 0x4665b258.
//
// Solidity: function deleteClaim(string key, bool isPrivate) returns()
func (_IdentityManager *IdentityManagerSession) DeleteClaim(key string, isPrivate bool) (*types.Transaction, error) {
	return _IdentityManager.Contract.DeleteClaim(&_IdentityManager.TransactOpts, key, isPrivate)
}

// DeleteClaim is a paid mutator transaction binding the contract method 0x4665b258.
//
// Solidity: function deleteClaim(string key, bool isPrivate) returns()
func (_IdentityManager *IdentityManagerTransactorSession) DeleteClaim(key string, isPrivate bool) (*types.Transaction, error) {
	return _IdentityManager.Contract.DeleteClaim(&_IdentityManager.TransactOpts, key, isPrivate)
}

// DeleteClaimURI is a paid mutator transaction binding the contract method 0x84ce394c.
//
// Solidity: function deleteClaimURI(string key) returns()
func (_IdentityManager *IdentityManagerTransactor) DeleteClaimURI(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "deleteClaimURI", key)
}

// DeleteClaimURI is a paid mutator transaction binding the contract method 0x84ce394c.
//
// Solidity: function deleteClaimURI(string key) returns()
func (_IdentityManager *IdentityManagerSession) DeleteClaimURI(key string) (*types.Transaction, error) {
	return _IdentityManager.Contract.DeleteClaimURI(&_IdentityManager.TransactOpts, key)
}

// DeleteClaimURI is a paid mutator transaction binding the contract method 0x84ce394c.
//
// Solidity: function deleteClaimURI(string key) returns()
func (_IdentityManager *IdentityManagerTransactorSession) DeleteClaimURI(key string) (*types.Transaction, error) {
	return _IdentityManager.Contract.DeleteClaimURI(&_IdentityManager.TransactOpts, key)
}

// RevokeAttestation is a paid mutator transaction binding the contract method 0x6066b5c3.
//
// Solidity: function revokeAttestation(string key, string reason, address attestedTo, string attestationId) returns()
func (_IdentityManager *IdentityManagerTransactor) RevokeAttestation(opts *bind.TransactOpts, key string, reason string, attestedTo common.Address, attestationId string) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "revokeAttestation", key, reason, attestedTo, attestationId)
}

// RevokeAttestation is a paid mutator transaction binding the contract method 0x6066b5c3.
//
// Solidity: function revokeAttestation(string key, string reason, address attestedTo, string attestationId) returns()
func (_IdentityManager *IdentityManagerSession) RevokeAttestation(key string, reason string, attestedTo common.Address, attestationId string) (*types.Transaction, error) {
	return _IdentityManager.Contract.RevokeAttestation(&_IdentityManager.TransactOpts, key, reason, attestedTo, attestationId)
}

// RevokeAttestation is a paid mutator transaction binding the contract method 0x6066b5c3.
//
// Solidity: function revokeAttestation(string key, string reason, address attestedTo, string attestationId) returns()
func (_IdentityManager *IdentityManagerTransactorSession) RevokeAttestation(key string, reason string, attestedTo common.Address, attestationId string) (*types.Transaction, error) {
	return _IdentityManager.Contract.RevokeAttestation(&_IdentityManager.TransactOpts, key, reason, attestedTo, attestationId)
}

// SetAttestation is a paid mutator transaction binding the contract method 0x31b6399a.
//
// Solidity: function setAttestation(string key, address attestor, uint256 expires, bytes signature, string claimId) returns()
func (_IdentityManager *IdentityManagerTransactor) SetAttestation(opts *bind.TransactOpts, key string, attestor common.Address, expires *big.Int, signature []byte, claimId string) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "setAttestation", key, attestor, expires, signature, claimId)
}

// SetAttestation is a paid mutator transaction binding the contract method 0x31b6399a.
//
// Solidity: function setAttestation(string key, address attestor, uint256 expires, bytes signature, string claimId) returns()
func (_IdentityManager *IdentityManagerSession) SetAttestation(key string, attestor common.Address, expires *big.Int, signature []byte, claimId string) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetAttestation(&_IdentityManager.TransactOpts, key, attestor, expires, signature, claimId)
}

// SetAttestation is a paid mutator transaction binding the contract method 0x31b6399a.
//
// Solidity: function setAttestation(string key, address attestor, uint256 expires, bytes signature, string claimId) returns()
func (_IdentityManager *IdentityManagerTransactorSession) SetAttestation(key string, attestor common.Address, expires *big.Int, signature []byte, claimId string) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetAttestation(&_IdentityManager.TransactOpts, key, attestor, expires, signature, claimId)
}

// SetClaimURI is a paid mutator transaction binding the contract method 0xc03bc8b5.
//
// Solidity: function setClaimURI(string key, string value) returns()
func (_IdentityManager *IdentityManagerTransactor) SetClaimURI(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "setClaimURI", key, value)
}

// SetClaimURI is a paid mutator transaction binding the contract method 0xc03bc8b5.
//
// Solidity: function setClaimURI(string key, string value) returns()
func (_IdentityManager *IdentityManagerSession) SetClaimURI(key string, value string) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetClaimURI(&_IdentityManager.TransactOpts, key, value)
}

// SetClaimURI is a paid mutator transaction binding the contract method 0xc03bc8b5.
//
// Solidity: function setClaimURI(string key, string value) returns()
func (_IdentityManager *IdentityManagerTransactorSession) SetClaimURI(key string, value string) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetClaimURI(&_IdentityManager.TransactOpts, key, value)
}

// SetPrivateClaim is a paid mutator transaction binding the contract method 0x89285c21.
//
// Solidity: function setPrivateClaim(string key, string value, string statement, string ipfsURI, string eventHash, uint256 timestamp) returns()
func (_IdentityManager *IdentityManagerTransactor) SetPrivateClaim(opts *bind.TransactOpts, key string, value string, statement string, ipfsURI string, eventHash string, timestamp *big.Int) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "setPrivateClaim", key, value, statement, ipfsURI, eventHash, timestamp)
}

// SetPrivateClaim is a paid mutator transaction binding the contract method 0x89285c21.
//
// Solidity: function setPrivateClaim(string key, string value, string statement, string ipfsURI, string eventHash, uint256 timestamp) returns()
func (_IdentityManager *IdentityManagerSession) SetPrivateClaim(key string, value string, statement string, ipfsURI string, eventHash string, timestamp *big.Int) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetPrivateClaim(&_IdentityManager.TransactOpts, key, value, statement, ipfsURI, eventHash, timestamp)
}

// SetPrivateClaim is a paid mutator transaction binding the contract method 0x89285c21.
//
// Solidity: function setPrivateClaim(string key, string value, string statement, string ipfsURI, string eventHash, uint256 timestamp) returns()
func (_IdentityManager *IdentityManagerTransactorSession) SetPrivateClaim(key string, value string, statement string, ipfsURI string, eventHash string, timestamp *big.Int) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetPrivateClaim(&_IdentityManager.TransactOpts, key, value, statement, ipfsURI, eventHash, timestamp)
}

// SetPublicClaim is a paid mutator transaction binding the contract method 0x27f541f8.
//
// Solidity: function setPublicClaim(string key, string value, string statement, uint256 timestamp) returns()
func (_IdentityManager *IdentityManagerTransactor) SetPublicClaim(opts *bind.TransactOpts, key string, value string, statement string, timestamp *big.Int) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "setPublicClaim", key, value, statement, timestamp)
}

// SetPublicClaim is a paid mutator transaction binding the contract method 0x27f541f8.
//
// Solidity: function setPublicClaim(string key, string value, string statement, uint256 timestamp) returns()
func (_IdentityManager *IdentityManagerSession) SetPublicClaim(key string, value string, statement string, timestamp *big.Int) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetPublicClaim(&_IdentityManager.TransactOpts, key, value, statement, timestamp)
}

// SetPublicClaim is a paid mutator transaction binding the contract method 0x27f541f8.
//
// Solidity: function setPublicClaim(string key, string value, string statement, uint256 timestamp) returns()
func (_IdentityManager *IdentityManagerTransactorSession) SetPublicClaim(key string, value string, statement string, timestamp *big.Int) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetPublicClaim(&_IdentityManager.TransactOpts, key, value, statement, timestamp)
}

// IdentityManagerAttestationRevokedIterator is returned from FilterAttestationRevoked and is used to iterate over the raw logs and unpacked data for AttestationRevoked events raised by the IdentityManager contract.
type IdentityManagerAttestationRevokedIterator struct {
	Event *IdentityManagerAttestationRevoked // Event containing the contract specifics and raw log

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
func (it *IdentityManagerAttestationRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityManagerAttestationRevoked)
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
		it.Event = new(IdentityManagerAttestationRevoked)
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
func (it *IdentityManagerAttestationRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityManagerAttestationRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityManagerAttestationRevoked represents a AttestationRevoked event raised by the IdentityManager contract.
type IdentityManagerAttestationRevoked struct {
	Identity common.Address
	Id       string
	Reason   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAttestationRevoked is a free log retrieval operation binding the contract event 0x1e926d3070e51ea520634ecc30c671c6d1363a87da2c8af5f2aa75b6bcfae1a2.
//
// Solidity: event AttestationRevoked(address indexed identity, string id, string reason)
func (_IdentityManager *IdentityManagerFilterer) FilterAttestationRevoked(opts *bind.FilterOpts, identity []common.Address) (*IdentityManagerAttestationRevokedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityManager.contract.FilterLogs(opts, "AttestationRevoked", identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerAttestationRevokedIterator{contract: _IdentityManager.contract, event: "AttestationRevoked", logs: logs, sub: sub}, nil
}

// WatchAttestationRevoked is a free log subscription operation binding the contract event 0x1e926d3070e51ea520634ecc30c671c6d1363a87da2c8af5f2aa75b6bcfae1a2.
//
// Solidity: event AttestationRevoked(address indexed identity, string id, string reason)
func (_IdentityManager *IdentityManagerFilterer) WatchAttestationRevoked(opts *bind.WatchOpts, sink chan<- *IdentityManagerAttestationRevoked, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityManager.contract.WatchLogs(opts, "AttestationRevoked", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityManagerAttestationRevoked)
				if err := _IdentityManager.contract.UnpackLog(event, "AttestationRevoked", log); err != nil {
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

// ParseAttestationRevoked is a log parse operation binding the contract event 0x1e926d3070e51ea520634ecc30c671c6d1363a87da2c8af5f2aa75b6bcfae1a2.
//
// Solidity: event AttestationRevoked(address indexed identity, string id, string reason)
func (_IdentityManager *IdentityManagerFilterer) ParseAttestationRevoked(log types.Log) (*IdentityManagerAttestationRevoked, error) {
	event := new(IdentityManagerAttestationRevoked)
	if err := _IdentityManager.contract.UnpackLog(event, "AttestationRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityManagerNewAttestationIterator is returned from FilterNewAttestation and is used to iterate over the raw logs and unpacked data for NewAttestation events raised by the IdentityManager contract.
type IdentityManagerNewAttestationIterator struct {
	Event *IdentityManagerNewAttestation // Event containing the contract specifics and raw log

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
func (it *IdentityManagerNewAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityManagerNewAttestation)
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
		it.Event = new(IdentityManagerNewAttestation)
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
func (it *IdentityManagerNewAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityManagerNewAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityManagerNewAttestation represents a NewAttestation event raised by the IdentityManager contract.
type IdentityManagerNewAttestation struct {
	Identity  common.Address
	Key       string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewAttestation is a free log retrieval operation binding the contract event 0x9476865afb170ca0b75c734e92ad4be2fa603b741a4c4542361fe02f8ff9ae4a.
//
// Solidity: event NewAttestation(address indexed identity, string key, uint256 timestamp)
func (_IdentityManager *IdentityManagerFilterer) FilterNewAttestation(opts *bind.FilterOpts, identity []common.Address) (*IdentityManagerNewAttestationIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityManager.contract.FilterLogs(opts, "NewAttestation", identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerNewAttestationIterator{contract: _IdentityManager.contract, event: "NewAttestation", logs: logs, sub: sub}, nil
}

// WatchNewAttestation is a free log subscription operation binding the contract event 0x9476865afb170ca0b75c734e92ad4be2fa603b741a4c4542361fe02f8ff9ae4a.
//
// Solidity: event NewAttestation(address indexed identity, string key, uint256 timestamp)
func (_IdentityManager *IdentityManagerFilterer) WatchNewAttestation(opts *bind.WatchOpts, sink chan<- *IdentityManagerNewAttestation, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityManager.contract.WatchLogs(opts, "NewAttestation", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityManagerNewAttestation)
				if err := _IdentityManager.contract.UnpackLog(event, "NewAttestation", log); err != nil {
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

// ParseNewAttestation is a log parse operation binding the contract event 0x9476865afb170ca0b75c734e92ad4be2fa603b741a4c4542361fe02f8ff9ae4a.
//
// Solidity: event NewAttestation(address indexed identity, string key, uint256 timestamp)
func (_IdentityManager *IdentityManagerFilterer) ParseNewAttestation(log types.Log) (*IdentityManagerNewAttestation, error) {
	event := new(IdentityManagerNewAttestation)
	if err := _IdentityManager.contract.UnpackLog(event, "NewAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityManagerNewClaimIterator is returned from FilterNewClaim and is used to iterate over the raw logs and unpacked data for NewClaim events raised by the IdentityManager contract.
type IdentityManagerNewClaimIterator struct {
	Event *IdentityManagerNewClaim // Event containing the contract specifics and raw log

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
func (it *IdentityManagerNewClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityManagerNewClaim)
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
		it.Event = new(IdentityManagerNewClaim)
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
func (it *IdentityManagerNewClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityManagerNewClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityManagerNewClaim represents a NewClaim event raised by the IdentityManager contract.
type IdentityManagerNewClaim struct {
	Key       string
	Value     string
	IsPrivate bool
	Version   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewClaim is a free log retrieval operation binding the contract event 0x26f9b425c95c58e8bec57a1841c65b09f8b1b4408004758fcebf088adc4b3d62.
//
// Solidity: event NewClaim(string key, string value, bool isPrivate, int256 version)
func (_IdentityManager *IdentityManagerFilterer) FilterNewClaim(opts *bind.FilterOpts) (*IdentityManagerNewClaimIterator, error) {

	logs, sub, err := _IdentityManager.contract.FilterLogs(opts, "NewClaim")
	if err != nil {
		return nil, err
	}
	return &IdentityManagerNewClaimIterator{contract: _IdentityManager.contract, event: "NewClaim", logs: logs, sub: sub}, nil
}

// WatchNewClaim is a free log subscription operation binding the contract event 0x26f9b425c95c58e8bec57a1841c65b09f8b1b4408004758fcebf088adc4b3d62.
//
// Solidity: event NewClaim(string key, string value, bool isPrivate, int256 version)
func (_IdentityManager *IdentityManagerFilterer) WatchNewClaim(opts *bind.WatchOpts, sink chan<- *IdentityManagerNewClaim) (event.Subscription, error) {

	logs, sub, err := _IdentityManager.contract.WatchLogs(opts, "NewClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityManagerNewClaim)
				if err := _IdentityManager.contract.UnpackLog(event, "NewClaim", log); err != nil {
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

// ParseNewClaim is a log parse operation binding the contract event 0x26f9b425c95c58e8bec57a1841c65b09f8b1b4408004758fcebf088adc4b3d62.
//
// Solidity: event NewClaim(string key, string value, bool isPrivate, int256 version)
func (_IdentityManager *IdentityManagerFilterer) ParseNewClaim(log types.Log) (*IdentityManagerNewClaim, error) {
	event := new(IdentityManagerNewClaim)
	if err := _IdentityManager.contract.UnpackLog(event, "NewClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
