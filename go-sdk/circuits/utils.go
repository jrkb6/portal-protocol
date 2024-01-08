package circuits

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/kzg"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/constraint"
	"go-sdk/client"
	"log"
)

var (
	curve = ecc.BN254
)

type CircuitMetadata struct {
	Name            string   `json:"name"`
	Statement       string   `json:"statement"`
	ContractAddress string   `json:"contractAddress"`
	Fields          []string `json:"fields"`
}
type CircuitOnIpfs struct {
	Metadata CircuitMetadata `json:"metadata"`
	Ccs      []byte          `json:"ccs"`
	Srs      []byte          `json:"srs"`
	Pk       []byte          `json:"pk"`
	Vk       []byte          `json:"vk"`
	Proof    []byte          `json:"proof"`
	Wt       []byte          `json:"wt"`
}

type ProofSystem struct {
	Ccs   constraint.ConstraintSystem
	Srs   kzg.SRS
	Pk    plonk.ProvingKey
	Vk    plonk.VerifyingKey
	Proof plonk.Proof
	Wt    witness.Witness
}

func SerializeCCS(circuit constraint.ConstraintSystem) (*bytes.Buffer, error) {
	// Serialize the circuit to bytes
	var buf bytes.Buffer
	_, err := circuit.WriteTo(&buf)
	if err != nil {
		return nil, err
	}

	// Write the serialized data to a file
	return &buf, err
}

func SerializeProvingKey(pk plonk.ProvingKey) (*bytes.Buffer, error) {
	// Serialize the proving key to bytes without point compression
	var buf bytes.Buffer
	_, err := pk.WriteRawTo(&buf)
	if err != nil {
		log.Fatal("Error writing ser buff", err)

		return nil, err
	}

	return &buf, err
}

func Serialize(ccs constraint.ConstraintSystem, srs kzg.SRS, wt witness.Witness, vk plonk.VerifyingKey,
	pk plonk.ProvingKey, proof plonk.Proof, name string, address string) CircuitOnIpfs {
	ret := CircuitOnIpfs{}
	buf, err := SerializeCCS(ccs)
	if err != nil {
		fmt.Println("Failed to serialize the constraints:", err)
	}
	ret.Ccs = buf.Bytes()

	data, err := wt.MarshalBinary()
	if err != nil {
		log.Fatal("wt marshal failed", err)

	}
	ret.Wt = data

	// Serialize the verifier key
	var bufVK bytes.Buffer
	_, _ = vk.WriteRawTo(&bufVK)
	ret.Vk = bufVK.Bytes()

	// serialize the srs
	var bufSRS bytes.Buffer
	_, _ = srs.WriteTo(&bufSRS)
	ret.Srs = bufSRS.Bytes()

	// Serialize the proving key without point compression
	pkData, err := SerializeProvingKey(pk)
	if err != nil {
		fmt.Println("Failed to serialize the proving key:", err)
	}
	ret.Pk = pkData.Bytes()

	//serialize proof as well
	var proofBuf bytes.Buffer
	_, err = proof.WriteTo(&proofBuf)
	ret.Proof = proofBuf.Bytes()
	ret.Metadata.Name = name
	ret.Metadata.ContractAddress = address
	return ret
}

func PublishCircuitOnIpfs(ipfsClient *client.IpfsClient, circuit *CircuitOnIpfs) (string, error) {

	// Serialize the struct to JSON.
	data, err := json.Marshal(circuit)
	if err != nil {
		log.Fatal(err)
	}

	circuitAddr, err := ipfsClient.AddAndPublish(data)
	if err != nil {
		log.Fatal("Failed to publish circuit to ipfs: ", err)
		return "", nil
	}
	return circuitAddr["Name"], nil

}

func BuildCircuitFrom(ipfsURI string, ipfsClient *client.IpfsClient) *ProofSystem {
	// Retrieve the circuit from IPFS
	circuitData, err := ipfsClient.RetrieveRaw(ipfsURI)
	if err != nil {
		log.Fatal("Failed to retrieve circuit from ipfs: ", err)
		return nil
	}

	// Deserialize the JSON back into a struct.
	var circuit CircuitOnIpfs
	if err := json.Unmarshal(circuitData, &circuit); err != nil {
		log.Fatal(err)
	}
	ccs := plonk.NewCS(curve)
	_, err = ccs.ReadFrom(bytes.NewReader(circuit.Ccs))

	srs := kzg.NewSRS(curve)
	_, err = srs.ReadFrom(bytes.NewReader(circuit.Srs))
	// Reconstruct witness
	newWitness, _ := witness.New(curve.ScalarField())

	// Extract the public part only
	err = newWitness.UnmarshalBinary(circuit.Wt)

	reconstructedPK := plonk.NewProvingKey(curve)
	_, err = reconstructedPK.ReadFrom(bytes.NewReader(circuit.Pk))

	reconstructedProof := plonk.NewProof(curve)
	_, err = reconstructedProof.ReadFrom(bytes.NewReader(circuit.Proof))

	reconstructedVK := plonk.NewVerifyingKey(curve)
	_, err = reconstructedVK.ReadFrom(bytes.NewReader(circuit.Vk))

	return &ProofSystem{Ccs: ccs, Srs: srs, Pk: reconstructedPK, Vk: reconstructedVK, Proof: reconstructedProof, Wt: newWitness}

}
