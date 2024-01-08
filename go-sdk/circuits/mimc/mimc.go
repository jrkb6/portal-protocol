package mimc

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

type GnarkMimc struct {
	PreImage frontend.Variable
	Hash     frontend.Variable `gnark:",public"`
}

func (circuit *GnarkMimc) Define(api frontend.API) error {
	// hash function
	mimc, _ := mimc.NewMiMC(api)
	mimc.Write(circuit.PreImage)
	sum := mimc.Sum()

	api.Println("Sum: ", sum)
	api.AssertIsEqual(sum, circuit.Hash)

	return nil
}

type ZkOpenWrapper struct {
	InMap [][32]frontend.Variable
	Hash  frontend.Variable `gnark:“,public”`
}

func (circuit *ZkOpenWrapper) Define(api frontend.API) error {
	mimc, _ := mimc.NewMiMC(api)

	// loop over bytes
	for i := 0; i < len(circuit.InMap); i++ {
		// rearrange input to match mimc input requirements
		ddd := make([]frontend.Variable, 256)
		for j := 0; j < 32; j++ {
			// get bits of ecb input, little endian!
			myBits := api.ToBinary(circuit.InMap[i][j], 8)
			for k := 7; k >= 0; k-- {
				ddd[(31-j)*8+(k)] = myBits[k]
			}
		}
		// input data into mimc
		varSum := api.FromBinary(ddd...)
		mimc.Write(varSum)
	}
	// mimc hash constraints check
	result := mimc.Sum()
	api.Println("Sum: ", result)
	api.AssertIsEqual(circuit.Hash, result)
	return nil
}
