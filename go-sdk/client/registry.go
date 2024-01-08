package client

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"go-sdk/contracts/IdentityRegistry"
	"log"
)

type RegistryClient struct {
	Registry *IdentityRegistry.IdentityRegistry
}

type CircuitMeta struct {
	DeploymentType string
	Address        common.Address
	IpfsURI        string
}

func NewRegistryClient(contractAddr string, reader TxReader) *RegistryClient {

	contract := common.HexToAddress(contractAddr)
	instance, err := IdentityRegistry.NewIdentityRegistry(contract, reader.Client)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &RegistryClient{
		Registry: instance,
	}
}

func RegisterIdentity(client *RegistryClient, user *Signer, registryOwner *Signer, userIdentityAddr common.Address) (*types.Receipt, error) {
	auth := user.BindTxOpts()
	userPub := user.PublicKey
	// sign message by contract registryOwner for integrity and authenticity check for registering a new identity
	_, signature, err := CalculateSignature(userPub, registryOwner, userIdentityAddr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("signature sending to contract: 0x%s\n", common.Bytes2Hex(signature))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	tx, err := client.Registry.Register(auth, userIdentityAddr, signature)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	receipt, err := user.WaitForReceipt(tx.Hash())
	if err != nil {
		log.Fatal(err)
		return nil, err

	}
	return receipt, nil
}
func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}
func CalculateSignature(userAddr []byte, registryOwner *Signer, userIdentityAddr common.Address) ([]byte, []byte, error) {
	signDataInBytes := encodePacked(userIdentityAddr.Bytes(), userAddr)
	hash := crypto.Keccak256(signDataInBytes)
	log.Printf("MessageHash 0x%s\n", common.Bytes2Hex(hash))

	var buf bytes.Buffer
	buf.Write([]byte("\x19Ethereum Signed Message:\n32")) // For byte slice
	buf.Write(hash)

	finalHash := crypto.Keccak256Hash(buf.Bytes())
	fBytes := finalHash.Bytes()

	signature, err := registryOwner.Sign(fBytes)
	signature[64] = signature[64] + 27

	////get v, r, s from signature
	//r := signature[:32]
	//s := signature[32:64]
	//v := signature[64:65]
	//verify signature locally
	// create pubkey from registryOwner private key
	//pubKey := registryOwner.PrivateKey.Public()
	//pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("error casting public key to ECDSA")
	//}
	//verified := crypto.VerifySignature(crypto.CompressPubkey(pubKeyECDSA), fBytes, signature[:64])
	//fmt.Println("verified", verified)
	return fBytes, signature, err
}

func GetIdentity(client *RegistryClient, address common.Address) (common.Address, error) {
	read, err := client.Registry.OwnerOf(nil, address)
	if err != nil {
		log.Fatal(err)
	}
	return read, nil
}

func GetManager(client *RegistryClient, address common.Address) (common.Address, error) {

	read, err := client.Registry.Manager(nil, address)
	if err != nil {
		log.Fatal(err)
	}
	return read, nil
}

func RegisterCircuit(client *RegistryClient, signer *Signer, circuitId string, c *CircuitMeta) (*types.Receipt, error) {
	auth := signer.BindTxOpts()
	tx, err := client.Registry.SetCircuit(auth, circuitId, c.DeploymentType, c.Address, c.IpfsURI)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	receipt, err := signer.WaitForReceipt(tx.Hash())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return receipt, nil
}

func GetCircuit(client *RegistryClient, circuitId string) (*CircuitMeta, error) {
	read, err := client.Registry.Circuits(nil, circuitId)
	if err != nil {
		log.Fatal(err)
	}
	return &CircuitMeta{DeploymentType: read.DeploymentType, Address: read.DeploymentAddress, IpfsURI: read.IpfsHash}, nil
}
