package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

type Config struct {
	Rawurl  string
	ChainId *big.Int
}

type Signer struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Client     *ethclient.Client
	GasLimit   uint64
	value      int64
}

type TxReader struct {
	Client *ethclient.Client
}

func NewSigner(key ecdsa.PrivateKey, publicKey []byte, config Config) *Signer {
	client, err := ethclient.Dial(config.Rawurl)
	if err != nil {
		log.Fatal(err)
	}

	return &Signer{
		PrivateKey: &key,
		PublicKey:  publicKey,
		Client:     client,
		GasLimit:   10000000,
		value:      0,
	}
}

func NewReader(config Config) *TxReader {
	client, err := ethclient.Dial(config.Rawurl)
	if err != nil {
		log.Fatal(err)
	}
	return &TxReader{
		Client: client,
	}
}

func (s *Signer) BindTxOpts() *bind.TransactOpts {

	publicKey := s.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := s.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SuggestGasPrice %d\n", gasPrice)
	/* increase gas price by 50% */
	//gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(15))
	//gasPrice = new(big.Int).Div(gasPrice, big.NewInt(10))

	chainId, _ := s.Client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(s.PrivateKey, chainId)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(s.value)
	auth.GasLimit = s.GasLimit
	auth.GasPrice = gasPrice
	return auth
}

func (s *Signer) Sign(digestHash []byte) ([]byte, error) {
	return crypto.Sign(digestHash, s.PrivateKey)
}
func (s *Signer) WaitForReceipt(txHash common.Hash) (*types.Receipt, error) {
	receiptChan := make(chan *types.Receipt)

	go func() {
		pendingTime := 0
		for {
			receipt, err := s.Client.TransactionReceipt(context.Background(), txHash)
			if err == nil {
				receiptChan <- receipt
				log.Printf("Waiting for transaction receipt for %d seconds\n", pendingTime)
				return
			}

			pendingTime++
			time.Sleep(1 * time.Second)
		}
	}()
	return <-receiptChan, nil
}
