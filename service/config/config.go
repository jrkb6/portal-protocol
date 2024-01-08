package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"math/big"
	"os"
)

type Config struct {
	RegistryAddress string   `env:"REGISTRY_ADDRESS,required"`
	RpcURL          string   `env:"RPC_URL,required"`
	IpfsURL         string   `env:"IPFS_URL,required"`
	OwnerPrivateKey string   `env:"OWNER_PRIVATE_KEY,required"`
	ChainId         *big.Int `env:"CHAIN_ID,required"`
}

var CONFIG Config

func LoadConfig() {

	err := godotenv.Load(".env")
	CONFIG.IpfsURL = os.Getenv("IPFS_URL")
	CONFIG.RpcURL = os.Getenv("RPC_URL")
	CONFIG.RegistryAddress = os.Getenv("REGISTRY_ADDRESS")
	CONFIG.OwnerPrivateKey = os.Getenv("OWNER_PRIVATE_KEY")
	CONFIG.ChainId, _ = new(big.Int).SetString(os.Getenv("CHAIN_ID"), 10)

	if err != nil {
		fmt.Print("Error loading .env file")
	}

}
