package merkletree

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash"
)

// MerkleProof stores the path, the root hash and a helper for the Merkle proof.
type MerkleProof struct {

	//root of the Merkle tree
	RootHash frontend.Variable `gnark:",public"`

	// path of the Merkle proof
	Path []frontend.Variable
}

// LeafSum returns the hash created from data inserted to form a leaf.
// Without domain separation.
func LeafSum(h hash.FieldHasher, data frontend.Variable) frontend.Variable {

	h.Reset()
	h.Write(data)
	res := h.Sum()

	return res
}

// nodeSum returns the hash created from data inserted to form a leaf.
// Without domain separation.
func nodeSum(h hash.FieldHasher, a, b frontend.Variable) frontend.Variable {

	h.Reset()
	h.Write(a, b)
	res := h.Sum()

	return res
}

// VerifyProof takes a Merkle root, a proofSet, and a proofIndex and returns
// true if the first element of the proof set is a leaf of data in the Merkle
// root. False is returned if the proof set or Merkle root is nil, and if
// 'numLeaves' equals 0.
func (mp *MerkleProof) VerifyProof(api frontend.API, h hash.FieldHasher, leaf frontend.Variable) {

	depth := len(mp.Path) - 1
	sum := LeafSum(h, mp.Path[0])

	// The binary decomposition is the bitwise negation of the order of hashes ->
	// If the path in the plain go code is 					0 1 1 0 1 0
	// The binary decomposition of the leaf index will be 	1 0 0 1 0 1 (little endian)
	binLeaf := api.ToBinary(leaf, depth)

	for i := 1; i < len(mp.Path); i++ { // the size of the loop is fixed -> one circuit per size
		d1 := api.Select(binLeaf[i-1], mp.Path[i], sum)
		d2 := api.Select(binLeaf[i-1], sum, mp.Path[i])
		sum = nodeSum(h, d1, d2)
	}

	// Compare our calculated Merkle root to the desired Merkle root.
	api.AssertIsEqual(sum, mp.RootHash)
}
