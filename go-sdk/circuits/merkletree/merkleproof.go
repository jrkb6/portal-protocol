package merkletree

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

// MerkleProofCircuit used for testing only
type MerkleProofCircuit struct {
	M    MerkleProof
	Leaf frontend.Variable
}

func (mp *MerkleProofCircuit) Define(api frontend.API) error {

	h, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}
	mp.M.VerifyProof(api, &h, mp.Leaf)

	return nil
}
