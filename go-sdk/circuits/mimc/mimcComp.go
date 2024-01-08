package mimc

import (
	"github.com/consensys/gnark/frontend"
)

type MimcCombined struct {
	Threshold frontend.Variable `gnark:",public"`
	Digest    frontend.Variable `gnark:",public"`
	Sender    frontend.Variable `gnark:",public"`
	Nonce     frontend.Variable `gnark:",public"`

	Input      frontend.Variable
	NoncePriv  frontend.Variable
	SenderPriv frontend.Variable
}

func (circuit *MimcCombined) Define(api frontend.API) error {
	api.AssertIsEqual(circuit.Sender, circuit.SenderPriv)
	api.AssertIsEqual(circuit.Nonce, circuit.NoncePriv)
	api.AssertIsLessOrEqual(circuit.Threshold, circuit.Input)
	mimcCircuit := &GnarkMimc{
		PreImage: circuit.Input,
		Hash:     circuit.Digest,
	}

	if err := mimcCircuit.Define(api); err != nil {
		return err
	}

	return nil
}
