package main

import (
	"bytes"
	crand "crypto/rand"
	"crypto/sha256"
	"flag"
	"fmt"
	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/hash"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/frontend/cs/scs"
	"github.com/consensys/gnark/test"
	"github.com/joho/godotenv"
	merklecombined "go-sdk/circuits/merklecombined"
	"go-sdk/simulation"
	"os"

	"github.com/consensys/gnark/frontend"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	circuitutils "go-sdk/circuits"
	mymerkletree "go-sdk/circuits/merkletree"
	"go-sdk/circuits/mimc"
	"go-sdk/client"
	"go-sdk/deployment"
	"log"
	"math/big"
	"strconv"
)

var (
	fBindings    = flag.Bool("bindings", false, "use go run main.go -bindings true to compile the solidity proof verification contract and create golang contract bindings")
	fDeploy      = flag.Bool("deploy", false, "use go run main.go -deploy true to deploy the solidity verification contract on-chain")
	fRegister    = flag.Bool("register", false, "use go run main.go -register true to registry identity to the registry contract")
	fClaim       = flag.Bool("claim", false, "use go run main.go -claim true to set public claim for identity")
	fIpfs        = flag.Bool("ipfs", false, "use go run main.go -ipfs true to set public claim for identity on ipfs")
	fIpfsOnChain = flag.Bool("ipfs-on-chain", false, "use go run main.go -ipfs-on-chain true to set public claim uri on manager for identity contract ")
	fSimulate    = flag.Bool("sim", false, "use go run main.go -sim true to simulate public claim on-chain set ")
	fCircuits    = flag.Bool("circuits", false, "use go run main.go -circuits true to run zkp verification circuits")
	fM           = flag.Bool("m", false, "use go run main.go -m true to run zkp verification circuits")
)

const (
	contractsPath = "../contracts"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	var (
		curve               = ecc.BN254
		conn                = os.Getenv("RPC_URL")
		ipfsConn            = os.Getenv("IPFS_URL") //"http://localhost:5001
		chainStr            = os.Getenv("CHAIN_ID")
		registryAddr        = os.Getenv("REGISTRY_CONTRACT")
		userManagerAddr     = os.Getenv("USER_IDENTITY_CONTRACT")
		ownerManagerAddr    = os.Getenv("ATTESTOR_IDENTITY_CONTRACT")
		ownerWalletAddr     = os.Getenv("OWNER_WALLET")
		userWalletAddr      = os.Getenv("USER_WALLET")
		attestorManagerAddr = os.Getenv("ATTESTOR_IDENTITY_CONTRACT")
	)
	// read private key from env
	chainId, _ := new(big.Int).SetString(chainStr, 10)
	ownerPrivStr := os.Getenv("OWNER_PRIVATE_KEY")
	userPrivStr := os.Getenv("USER_PRIVATE_KEY")
	managerAddr := common.HexToAddress(userManagerAddr)
	ownerManager := common.HexToAddress(ownerManagerAddr)
	ownerAccount := common.HexToAddress(ownerWalletAddr)
	userAccount := common.HexToAddress(userWalletAddr)
	ownerPriv, _ := crypto.HexToECDSA(ownerPrivStr)
	userPriv, _ := crypto.HexToECDSA(userPrivStr)
	fmt.Printf("Loaded env variables rpc url: %s,  ipfsUrl: %s, ownerPriv: %s,"+
		" userPriv: %s \n", conn, ipfsConn, ownerPrivStr, userPrivStr)

	reader := client.NewReader(client.Config{Rawurl: conn, ChainId: chainId})
	ownerSigner := client.NewSigner(*ownerPriv, ownerAccount.Bytes(), client.Config{
		Rawurl:  conn,
		ChainId: chainId,
	})

	userSigner := client.NewSigner(*userPriv, userAccount.Bytes(), client.Config{
		Rawurl:  conn,
		ChainId: chainId,
	})

	flag.Parse()

	if *fBindings {
		deployment.CreateBindings(contractsPath, "IdentityManager")
		deployment.CreateBindings(contractsPath, "IdentityRegistry")
		deployment.CreateBindings(contractsPath, "IdentityInterface")
		return
	}

	if *fDeploy {

		fmt.Println("Deploying contracts from", ownerAccount.Hex())

		// ownerAccount deploys the registry
		registryContract, _ := deployment.DeployRegistry(ownerSigner)

		// userAccount deploys it's own manager
		userIdentity, _ := deployment.DeployManager(userSigner, registryContract)
		attestorIdentity, _ := deployment.DeployManager(ownerSigner, registryContract)

		fmt.Println("Registry contract deployed at:", registryContract)
		fmt.Println("UserIdentity contract deployed at:", userIdentity)
		fmt.Println("AttestorIdentity contract deployed at:", attestorIdentity)

		fmt.Println("Make sure to update .env file with new contract addresses!")

	}
	if *fRegister {
		// register identity tx is sent by user
		regClient := client.NewRegistryClient(registryAddr, *reader)
		receipt, err := client.RegisterIdentity(regClient, userSigner, ownerSigner, managerAddr)
		if err != nil {
			log.Fatal("Could not register identity", err)
		}
		receipt, err = userSigner.WaitForReceipt(receipt.TxHash)
		if err != nil {
			return
		}

		fmt.Println("Identity registered for user:", userAccount.Hex())
		receipt, err = client.RegisterIdentity(regClient, ownerSigner, ownerSigner, ownerManager)
		receipt, err = userSigner.WaitForReceipt(receipt.TxHash)
		if err != nil {
			return
		}
		fmt.Println("Identity registered for attestor:", ownerAccount.Hex())

	}

	if *fClaim {
		regClient := client.NewRegistryClient(registryAddr, *reader)
		// check identity is registered
		manager, err := client.GetManager(regClient, userAccount)
		if err != nil {
			fmt.Println("Error getting manager:", err)
			return
		}
		log.Printf("Manager address for user %s: %s\n", userAccount.Hex(), manager)

		// now user can set public claim by it's own manager
		manClient := client.NewManagerClient(manager.Hex(), *reader, "")
		//jsonStr := generateDummyClaim()

		_, err = client.SetPublicClaim(manClient, userSigner, "birthDate", "01/01/2000", "age > 22")

		if err != nil {
			log.Fatal("Could not set claim", err)
			return
		}
		// get public claim here
		value, err := client.GetPublicClaim(manClient, "birthDate")
		if err != nil {
			return
		}

		fmt.Println("Claim:", value)
	}

	if *fIpfs {
		cli := client.NewIpfsClient(ipfsConn)

		respDict := createClaimOnIpfs(cli)

		//retrieve data from ipfs
		resp, err := cli.Retrieve(respDict["Name"], false)
		if err != nil {
			log.Fatal("Could not retrieve data", err)
			return
		}
		fmt.Println("Retrieved data:", resp)

		// this part is disabled since we use ipfs to store claim data (not ipns)
		//// update claim on ipfs
		//updateResp, err := cli.Update(respDict["Name"], generateDummyClaim())
		//if err != nil {
		//	log.Fatal("Could not update data", err)
		//	return
		//}
		//fmt.Println("Updated data:", updateResp)
		//
		//// retrieve updated content with hash
		//resp, err = cli.Retrieve(updateResp["Name"], false)
		//if err != nil {
		//	log.Fatal("Could not retrieve data", err)
		//	return
		//}
		//fmt.Println("Retrieved updated data:", resp)

	}

	if *fIpfsOnChain {

		regClient := client.NewRegistryClient(registryAddr, *reader)
		// check identity is registered
		manager, err := client.GetManager(regClient, userAccount)
		if err != nil {
			fmt.Println("Error getting manager:", err)
			return
		}
		log.Printf("Manager address for user %s: %s\n", userAccount.Hex(), manager)
		// now user can set public claim by it's own manager
		manClient := client.NewManagerClient(manager.Hex(), *reader, ipfsConn)
		claimDict := simulation.GenerateDummyClaim()

		_, err = client.PublishClaim(manClient, userSigner, "age", claimDict)

		if err != nil {
			log.Fatal("Could not set claim", err)
			return
		}

		//get public claim here
		value, err := client.GetClaimURI(manClient, "age")
		if err != nil {
			log.Fatal("Could not get claim", err)
			return
		}
		fmt.Println("Claim uri read from contract:", value)
		// query ipfs to return claim data
		resp, err := manClient.IpfsClient.Retrieve(value, false)
		if err != nil {
			log.Fatal("Could not retrieve data on ipfs", err)
			return
		}

		fmt.Println("Claim value retrieved from ipfs:", resp)

	}

	if *fSimulate {
		simType := flag.Arg(0)
		fmt.Printf("Reading .env file: Registry address: %s, Manager address: %s\n", registryAddr, userManagerAddr)
		switch simType {
		default:
			fmt.Println("No type specified. Please specify type: ipfs, onchain, ipfs-on-chain, deployment. e.g.: go run main.go -sim onchain")
			return

		//case "ipfs-on-chain":
		//	simulation.SimulateIpfsOnChainPublicClaim(registryAddr, userAccount, userSigner, reader, ipfsConn, false)
		//case "onchain":
		//	simulation.SimulateOnChainPublicClaims(registryAddr, userAccount, userSigner, reader, false)
		case "ipfs":
			simulation.SimulateIPFSConnection(ipfsConn)
		case "deployment":
			simulation.SimulateContractDeployment(userSigner, ownerSigner, reader)
		case "mimcDeployment":
			simulation.SimulateMimcCircuitDeployment(ownerSigner, ipfsConn, registryAddr, reader)
		case "merkleDeployment":
			simulation.SimulateMerkleCircuitDeployment(ownerSigner, ipfsConn, registryAddr, reader)
		//case "mimcClaim":
		//	simulation.SimulatePrivateClaim(userSigner, reader, ipfsConn, registryAddr, userManagerAddr, "MimcCombined")
		//case "merkleClaim":
		//	simulation.SimulatePrivateMerkleClaim(userSigner, reader, ipfsConn, registryAddr, userManagerAddr, "MerkleCombined")
		case "claim":
			claimType := flag.Arg(1)
			if claimType == "privateMerkle" {
				simulation.SimulatePrivateMerkleClaim(userSigner, reader, ipfsConn, registryAddr, userManagerAddr, "MerkleCombined")
			} else if claimType == "privateMerkleLive" {
				simulation.SimulateMerkleLiveVerification(userSigner, reader, ipfsConn, userManagerAddr)
			} else if claimType == "privateMimc" {
				simulation.SimulatePrivateClaim(userSigner, reader, ipfsConn, registryAddr, userManagerAddr, "MimcCombined")
			} else if claimType == "publicOnChain" {
				simulation.SimulateOnChainPublicClaims(registryAddr, userAccount, userSigner, reader)
			} else if claimType == "publicIpfs" {
				simulation.SimulateIpfsOnChainPublicClaim(registryAddr, userAccount, userSigner, reader, ipfsConn)
			}
		case "attestation":
			claimType := flag.Arg(1)
			if claimType == "private" {
				fmt.Println("Running attestation simulation... make sure you have deployed registry and manager contracts on chain")
				simulation.SimulateAttestation(userSigner, ownerSigner, reader, ipfsConn, registryAddr, userManagerAddr)
			} else if claimType == "publicOnChain" {
				fmt.Println("Running attestation on-chain simulation... make sure you have deployed registry and manager contracts on chain")
				simulation.SimulateAttestationPublicOnChain(userSigner, ownerSigner, reader, ipfsConn, registryAddr, userManagerAddr)
			} else if claimType == "publicIpfs" {
				fmt.Println("Running attestation on-ipfs simulation... make sure you have deployed registry and manager contracts on chain")
				simulation.SimulateAttestationPublicIpfs(userSigner, ownerSigner, reader, ipfsConn, registryAddr, userManagerAddr)
			}

		case "revocation":
			fmt.Println("Running attestation simulation... make sure you have deployed registry and manager contracts on chain")
			simulation.SimulateRevocation(userSigner, ownerSigner, reader, ipfsConn, userManagerAddr, attestorManagerAddr)
		}

	}

	if *fCircuits {
		var (
			pre       = 100
			threshold = 22
			digest    = "20370067689261511688289967978544823130432235585709842144916192060767982363628"
		)

		ipfsCli := client.NewIpfsClient(ipfsConn)
		var circuit mimc.MimcCombined

		// generate CompiledConstraintSystem
		ccs, err := frontend.Compile(curve.ScalarField(), scs.NewBuilder, &circuit)
		if err != nil {
			log.Fatal("frontend.Compile")
		}

		/*
			Prover part
		*/
		assignment := mimc.MimcCombined{
			Input:      pre,
			Threshold:  threshold,
			Digest:     digest,
			Nonce:      0,
			Sender:     0,
			NoncePriv:  0,
			SenderPriv: 0,
		}

		wt, err := frontend.NewWitness(&assignment, curve.ScalarField())
		if err != nil {
			log.Fatal("wt creation failed", err)

		}
		publicWit, _ := wt.Public()

		// plonk zkSNARK: Setup

		srs, err := test.NewKZGSRS(ccs)
		if err != nil {
			log.Fatal("test.NewKZGSRS(ccs)")
		}
		pk, vk, err := plonk.Setup(ccs, srs)
		if err != nil {
			log.Fatal("plonk.Setup", err)
		}

		// plonk: Prove
		proof, err := plonk.Prove(ccs, pk, wt)
		if err != nil {
			log.Fatal("prove computation failed...", err)
			return
		}

		err = plonk.Verify(proof, vk, publicWit)
		if err != nil {
			log.Fatal("plonk verify failed on local...")
			return
		}

		/*
			Serialize part for sending to remote verifier
		*/
		mimcCompSerialize := circuitutils.Serialize(ccs, srs, wt, vk, pk, proof, "mimc", "0x0")
		mimcCompMapping, _ := circuitutils.PublishCircuitOnIpfs(ipfsCli, &mimcCompSerialize)
		deserializedMimcComp := circuitutils.BuildCircuitFrom(mimcCompMapping, ipfsCli)
		mimcCompWitness, err := deserializedMimcComp.Wt.Public()
		if err != nil {
			return
		}
		err = plonk.Verify(deserializedMimcComp.Proof, deserializedMimcComp.Vk, mimcCompWitness)
		if err != nil {
			log.Fatal("Verify failed from remote verifier", err)
			return
		}

		// merkle proof

		// write the proof of mimc comparison into merkle tree and verify

		numLeaves := 32
		depth := 5
		const segmentSize = 32
		var merkleCircuit mymerkletree.MerkleProofCircuit
		merkleCircuit.M.Path = make([]frontend.Variable, depth+1)

		cc, err := frontend.Compile(curve.ScalarField(), scs.NewBuilder, &merkleCircuit)
		if err != nil {
			log.Fatal(err)
		}

		mod := curve.ScalarField()
		modNbBytes := len(mod.Bytes())

		// proof size is 256 bytes
		// we allow 32 bytes for each leaf

		var buf bytes.Buffer
		var tmp bytes.Buffer
		_, err = proof.WriteRawTo(&tmp)
		if err != nil {
			log.Fatal("proof write failed", err)
			return
		}

		sha := sha256.New()
		sha.Write(tmp.Bytes())
		sum := sha.Sum(nil)

		// normalize the sum with mod BN254
		sumMod := new(big.Int).SetBytes(sum)
		sumMod.Mod(sumMod, curve.BaseField())

		sumModBytes := sumMod.Bytes()
		buf.Write(make([]byte, modNbBytes-len(sumModBytes)))
		buf.Write(sumModBytes)

		// some random elements for tree
		for i := 1; i < numLeaves; i++ {
			leaf, _ := crand.Int(crand.Reader, mod)

			b := leaf.Bytes()

			buf.Write(make([]byte, modNbBytes-len(b)))
			buf.Write(b)
		}

		hGo := hash.MIMC_BN254.New()
		merkleRoot, proofPath, leaves, err := merkletree.BuildReaderProof(&buf, hGo, segmentSize, uint64(0))

		if err != nil {
			log.Fatal(err)
		}

		// verify the proof in plain go (data part)
		verified := merkletree.VerifyProof(hGo, merkleRoot, proofPath, uint64(0), leaves)
		if !verified {
			log.Print("The merkle proof in plain go should pass", err)
		}

		// wt
		var merkleWt mymerkletree.MerkleProofCircuit
		merkleWt.Leaf = 0
		merkleWt.M.RootHash = merkleRoot
		merkleWt.M.Path = make([]frontend.Variable, depth+1)

		for i := 0; i < depth+1; i++ {
			merkleWt.M.Path[i] = proofPath[i]
		}

		w, err := frontend.NewWitness(&merkleWt, curve.ScalarField())
		if err != nil {
			log.Fatal("Couldn't set witness ", err)
		}

		if err != nil {
			return
		}

		// plonk zkSNARK: Setup
		pk, vk, err = plonk.Setup(cc, srs)
		if err != nil {
			log.Fatal("plonk.Setup")
		}
		// plonk: Prove & Verify
		proof, err = plonk.Prove(cc, pk, w)
		if err != nil {
			log.Fatal("prove computation failed...")
		}

		merkleSerialize := circuitutils.Serialize(cc, srs, w, vk, pk, proof, "Merkle", "0x0")
		merkleMapping, _ := circuitutils.PublishCircuitOnIpfs(ipfsCli, &merkleSerialize)
		deserializedMerkle := circuitutils.BuildCircuitFrom(merkleMapping, ipfsCli)
		merklePublicWit, err := deserializedMerkle.Wt.Public()
		if err != nil {
			return
		}
		err = plonk.Verify(deserializedMerkle.Proof, deserializedMerkle.Vk, merklePublicWit)
		if err != nil {
			log.Fatal("Verify failed from remote verifier", err)
			return
		}
	}
	if *fM {
		var (
			pre       = 100
			threshold = 22
			digest    = "20370067689261511688289967978544823130432235585709842144916192060767982363628"
		)
		numLeaves := 32
		segmentSize := 32
		mod := curve.ScalarField()
		modNbBytes := len(mod.Bytes())
		depth := 5

		var merkleCircuit merklecombined.MerkleCombined
		merkleCircuit.Path = make([]frontend.Variable, depth+1)

		cc, err := frontend.Compile(curve.ScalarField(), scs.NewBuilder, &merkleCircuit)
		if err != nil {
			log.Fatal(err)
		}

		var buf bytes.Buffer
		mimcGo := hash.MIMC_BN254.New()
		// convert pre to bytes with padding to 32 bytes
		inpBytes := make([]byte, 32)
		preBytes := []byte(strconv.Itoa(pre))
		copy(inpBytes[32-len(preBytes):], preBytes)

		mimcGo.Write(inpBytes)
		sum := mimcGo.Sum(nil)
		// normalize the input with mod BN254
		inputMod := new(big.Int).SetBytes(sum)
		inputMod.Mod(inputMod, curve.BaseField())

		sumModBytes := inputMod.Bytes()
		buf.Write(make([]byte, modNbBytes-len(sumModBytes)))
		buf.Write(sumModBytes)

		//// some random elements for tree
		for i := 1; i < numLeaves; i++ {
			leaf, _ := crand.Int(crand.Reader, mod)

			b := leaf.Bytes()

			buf.Write(make([]byte, modNbBytes-len(b)))
			buf.Write(b)
		}

		hGo := hash.MIMC_BN254.New()

		//merkleRoot, proofPath, index, leaves := tree.Prove()
		merkleRoot, proofPath, leaves, err := merkletree.BuildReaderProof(&buf, hGo, segmentSize, uint64(0))

		if err != nil {
			log.Fatal(err)
		}

		// verify the proof in plain go (data part)
		verified := merkletree.VerifyProof(hGo, merkleRoot, proofPath, 0, leaves)
		if !verified {
			log.Print("The merkle proof in plain go should pass", err)
		}

		// wt
		var merkleWt merklecombined.MerkleCombined
		merkleWt.Input = pre
		merkleWt.Threshold = threshold
		merkleWt.Digest = digest
		merkleWt.RootHash = merkleRoot
		merkleWt.Leaf = 0
		merkleWt.NoncePriv = 0
		merkleWt.SenderPriv = 0
		merkleWt.Nonce = 0
		merkleWt.Sender = 0
		merkleWt.Path = make([]frontend.Variable, depth+1)
		for i := 0; i < depth+1; i++ {
			merkleWt.Path[i] = proofPath[i]
		}

		w, err := frontend.NewWitness(&merkleWt, curve.ScalarField())
		if err != nil {
			log.Fatal("Couldn't set witness ", err)
		}

		srs, err := test.NewKZGSRS(cc)
		if err != nil {
			log.Fatal("test.NewKZGSRS(ccs)")
		}
		// groth16 zkSNARK: Setup
		pk, vk, err := plonk.Setup(cc, srs)
		if err != nil {
			log.Fatal("plonk.Setup", err)
		}
		// groth16: Prove & Verify
		proof, err := plonk.Prove(cc, pk, w)
		if err != nil {
			log.Fatal("prove computation failed...", err)
		}

		publicWit, _ := w.Public()
		err = plonk.Verify(proof, vk, publicWit)
		if err != nil {
			log.Fatal("Verify failed from local verifier", err)
			return
		}

	}

}

func assertNoError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createClaimOnIpfs(cli *client.IpfsClient) map[string]string {
	claim := simulation.GenerateDummyClaim()

	resp, err := cli.AddAndPublish(claim)

	if err != nil {
		log.Fatal("Could not store data", err)
		return nil
	}

	fmt.Printf("Stored at: %s , link: http://localhost:8080%s\n", resp["Name"], resp["Value"])
	return resp
}
