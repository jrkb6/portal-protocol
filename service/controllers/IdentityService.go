package controllers

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gofiber/fiber/v2"
	"log"
	"service/client"
	"service/config"
	"time"
)

type OwnerKeys struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
}

var signer *OwnerKeys

func InitOwnerKeys() {

	privateKey, err := crypto.HexToECDSA(config.CONFIG.OwnerPrivateKey)
	if err != nil {
		log.Fatal("error casting private key to ECDSA", err)
	}
	publicKey := crypto.PubkeyToAddress(privateKey.PublicKey)

	signer = &OwnerKeys{
		PrivateKey: privateKey,
		PublicKey:  publicKey.Bytes(),
	}
	log.Printf("Using owner wallet: %s\n", publicKey.String())
}

func Sign(c *fiber.Ctx) error {

	sender := c.Params("sender", "0x000000000000000000000000000000000000000000")
	identity := c.Query("identity")
	fbytes, signature, err := CalculateSignature([]byte(sender), signer, common.HexToAddress(identity))
	if err != nil {
		return err
	}
	resp := map[string]interface{}{
		"signature": signature,
		"hash":      fbytes,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

const (
	ErrGettingIdentity     = "error getting identity"
	ErrAttestationNotFound = "Attestation not found"
	ErrAttestationExpired  = "Attestation is expired"
	ErrInvalidCredential   = "Invalid credential"
	ErrAttestationRevoked  = "Attestation is revoked"
)

func VerifyAttestation(c *fiber.Ctx) error {

	sender := c.Params("sender", "0x000000000000000000000000000000000000000000")
	attestationName := c.Query("attestationName")
	claimType := c.Query("type")

	reader := client.NewReader(client.Config{Rawurl: config.CONFIG.RpcURL, ChainId: config.CONFIG.ChainId})
	regClient := client.NewRegistryClient(config.CONFIG.RegistryAddress, *reader)

	startTime := time.Now()
	identity, err := getIdentityOrRespondError(regClient, sender, c)
	if err != nil {
		return err
	}

	manClient := client.NewManagerClient(identity.String(), *reader, config.CONFIG.IpfsURL)
	attestation, err := getAttestationOrRespondError(manClient, attestationName, c)
	if err != nil {
		return err
	}
	if attestation.ClaimId == "" {
		log.Println("attestation not found")
		return respondWithError(c, fiber.StatusBadRequest, ErrAttestationNotFound)
	}

	if !isAttestationValid(attestation) {
		return respondWithError(c, fiber.StatusBadRequest, ErrAttestationExpired)
	}
	if claimSig, err := verifyClaimType(manClient, attestation, claimType); err != nil {
		return respondWithError(c, fiber.StatusBadRequest, err.Error())
	} else if !claimSig {
		return respondWithError(c, fiber.StatusBadRequest, ErrInvalidCredential)
	}
	if verified, err := checkRevocation(regClient, reader, attestation, attestationName); err != nil {
		return respondWithError(c, fiber.StatusBadRequest, err.Error())
	} else if !verified {
		return respondWithError(c, fiber.StatusBadRequest, ErrAttestationRevoked)
	}
	resp := map[string]interface{}{
		"attestation": attestationName,
		"verified":    true,
	}
	log.Printf("Verified attestation %s in %s\n", attestationName, time.Since(startTime))
	return c.Status(fiber.StatusOK).JSON(resp)

}
func VerifyClaim(c *fiber.Ctx) error {

	sender := c.Params("sender", "0x0000000")
	claimName := c.Query("claimName")
	reader := client.NewReader(client.Config{Rawurl: config.CONFIG.RpcURL, ChainId: config.CONFIG.ChainId})
	regClient := client.NewRegistryClient(config.CONFIG.RegistryAddress, *reader)

	identity, err := getIdentityOrRespondError(regClient, sender, c)
	if err != nil {
		return err
	}

	manClient := client.NewManagerClient(identity.String(), *reader, config.CONFIG.IpfsURL)
	claim, err := getClaimOrRespondError(manClient, claimName, c)
	if err != nil {
		return err
	}

	verified, _ := verifyPrivateClaim(*claim)
	resp := map[string]interface{}{
		"claim":    claim,
		"verified": verified,
	}
	return c.Status(fiber.StatusOK).JSON(resp)

}

func getIdentityOrRespondError(regClient *client.RegistryClient, sender string, c *fiber.Ctx) (common.Address, error) {
	identity, err := client.GetIdentity(regClient, common.HexToAddress(sender))
	if err != nil {
		log.Println(ErrGettingIdentity, err)
		resp := map[string]interface{}{
			"verified": false,
			"reason":   ErrGettingIdentity,
		}
		return common.HexToAddress("0x0"), c.Status(fiber.StatusBadRequest).JSON(resp)
	}
	return identity, nil
}

func getAttestationOrRespondError(manClient *client.ManagerClient, attestationName string, c *fiber.Ctx) (*client.Attestation, error) {
	attestation, err := client.GetAttestation(manClient, attestationName)
	if err != nil {
		log.Println(ErrAttestationNotFound, err)
		return nil, respondWithError(c, fiber.StatusBadRequest, ErrAttestationNotFound)
	}
	return attestation, nil
}

func getClaimOrRespondError(manClient *client.ManagerClient, claimName string, c *fiber.Ctx) (*client.PrivateClaim, error) {
	claim, err := client.GetPrivateClaim(manClient, claimName)
	if err != nil {
		log.Println(ErrAttestationNotFound, err)
		return nil, respondWithError(c, fiber.StatusBadRequest, ErrAttestationNotFound)
	}
	return claim, nil
}

func checkRevocation(regClient *client.RegistryClient,
	reader *client.TxReader,
	attestation *client.Attestation,
	attestationName string) (bool, error) {
	attestorIdentity, _ := client.GetIdentity(regClient, common.HexToAddress(attestation.Attestor))
	attestorManClient := client.NewManagerClient(attestorIdentity.String(), *reader, "")
	revoked, err := client.GetRevocation(attestorManClient, attestationName)
	if err != nil {
		log.Println("error getting revocation", err)
		return false, err
	}
	return revoked.AttestationId == "", nil
}
func respondWithError(c *fiber.Ctx, statusCode int, reason string) error {
	resp := map[string]interface{}{
		"attestation": "",
		"verified":    false,
		"reason":      reason,
	}
	return c.Status(statusCode).JSON(resp)
}
func verifyClaimType(manClient *client.ManagerClient,
	attestation *client.Attestation,
	claimType string) (bool, error) {

	var claimSig bool
	var err error
	switch claimType {
	case "public", "ipfs", "private":
		claimSig, err = verifyClaimBasedOnType(manClient, attestation, claimType)
		if err != nil {
			log.Println("error getting claim", err)
		}
	default:
		err = fmt.Errorf("invalid claim type")
		log.Println(err)
	}
	return claimSig, err
}

func verifyClaimBasedOnType(manClient *client.ManagerClient,
	attestation *client.Attestation,
	claimType string) (bool, error) {

	claimSig := false
	switch claimType {
	case "public":
		claim, err := client.GetPublicClaim(manClient, attestation.ClaimId)
		if err != nil {
			log.Println(ErrInvalidCredential, err)
			return false, err
		}
		// sign the claim with attestor {t, claim={commitment/public_value}, statement}
		claimSig, _ = verifyClaim(claim.ClaimMeta, *attestation)
	case "ipfs":
		claim, err := client.GetIpfsClaim(manClient, attestation.ClaimId)
		if err != nil {
			log.Println(ErrInvalidCredential, err)
			return false, err
		}
		claimSig, _ = verifyClaim(claim.ClaimMeta, *attestation)
	case "private":
		claim, err := client.GetPrivateClaim(manClient, attestation.ClaimId)
		if err != nil {
			log.Println(ErrInvalidCredential, err)
			return false, err
		}
		claimSig, _ = verifyClaim(claim.ClaimMeta, *attestation)

	}
	return claimSig, nil
}

func isAttestationValid(attestation *client.Attestation) bool {
	return attestation.ExpiresAt.Int64() > time.Now().Unix()
}

func CalculateSignature(userAddr []byte,
	registryOwner *OwnerKeys,
	userIdentityAddr common.Address) ([]byte, []byte, error) {
	signDataInBytes := encodePacked(userIdentityAddr.Bytes(), userAddr)
	hash := crypto.Keccak256(signDataInBytes)

	var buf bytes.Buffer
	buf.Write([]byte("\x19Ethereum Signed Message:\n32"))
	buf.Write(hash)

	finalHash := crypto.Keccak256Hash(buf.Bytes())
	fBytes := finalHash.Bytes()

	signature, err := registryOwner.Sign(fBytes)
	signature[64] = signature[64] + 27

	return fBytes, signature, err
}

func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}
func (s *OwnerKeys) Sign(digestHash []byte) ([]byte, error) {
	return crypto.Sign(digestHash, s.PrivateKey)
}
func verifyClaim(meta client.ClaimMeta, attestation client.Attestation) (bool, error) {
	// assert claim signature is valid and matches attestation
	// sign the claim with attestor {t, claim={commitment/public_value}, statement}
	jsonMeta, err := json.Marshal(meta)
	if err != nil {
		log.Println("error marshalling claim meta", err)
		return false, err
	}

	sha := sha256.New()
	sha.Write(jsonMeta)
	sum := sha.Sum(nil)

	if err != nil {
		return false, err
	}
	publicKeyECDSA, err := crypto.SigToPub(sum, attestation.Signature)

	// Serialize the public key to bytes
	recoveredAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Convert the address to a string and compare it to the attestorAddress
	if recoveredAddr.String() != attestation.Attestor {
		log.Println("attestation signature is invalid")
		return false, nil
	}
	return true, nil
}
func verifyPrivateClaim(claim client.PrivateClaim) (bool, error) {
	ipfsClient := client.NewIpfsClient(config.CONFIG.IpfsURL)
	return verifyLive(claim, ipfsClient), nil
}
