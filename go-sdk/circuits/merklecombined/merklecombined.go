package merklecombined

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
	"go-sdk/circuits/merkletree"
)

type MerkleCombined struct {
	Threshold frontend.Variable `gnark:",public"`
	Digest    frontend.Variable `gnark:",public"`
	RootHash  frontend.Variable `gnark:",public"`
	Sender    frontend.Variable `gnark:",public"`
	Nonce     frontend.Variable `gnark:",public"`

	NoncePriv  frontend.Variable
	SenderPriv frontend.Variable

	Input frontend.Variable
	Path  []frontend.Variable
	Leaf  frontend.Variable
}

func (mp *MerkleCombined) Define(api frontend.API) error {
	api.AssertIsEqual(mp.Sender, mp.SenderPriv)
	api.AssertIsEqual(mp.Nonce, mp.NoncePriv)

	// compare input to threshold
	api.AssertIsLessOrEqual(mp.Threshold, mp.Input)

	h, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}
	h.Write(mp.Input)
	sum := h.Sum()
	api.Println("Sum: ", sum)
	api.AssertIsEqual(mp.Digest, sum)

	// build merkle tree from input
	var merkleCirc merkletree.MerkleProofCircuit
	merkleCirc.Leaf = mp.Leaf
	merkleCirc.M.RootHash = mp.RootHash
	merkleCirc.M.Path = make([]frontend.Variable, len(mp.Path))
	for i := 0; i < len(mp.Path); i++ {
		merkleCirc.M.Path[i] = mp.Path[i]
	}
	return merkleCirc.Define(api)

}
