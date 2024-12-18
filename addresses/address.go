package addresses

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EthAddress struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`
}

// CreateAddressFromPrivateKey generates a new Ethereum address from a randomly generated private key.
//
// This function creates a new ECDSA private key, derives the corresponding public key,
// and calculates the Ethereum address. It returns an EthAddress struct containing
// the hexadecimal representations of the private key, public key, and the Ethereum address.
//
// Returns:
//   - *EthAddress: A pointer to an EthAddress struct containing the generated address information.
//   - error: An error if key generation fails, nil otherwise.
func CreateAddressFromPrivateKey() (*EthAddress, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	address := &EthAddress{
		PrivateKey: hex.EncodeToString(crypto.FromECDSA(privateKey)),
		PublicKey:  hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey)),
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey).String(),
	}
	return address, nil
}

func PublicKeyToAddress(publicKey string) (string, error) {
	publicKeyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		return "", err
	}
	addressCommon := common.BytesToAddress(crypto.Keccak256(publicKeyBytes[1:])[12:])
	return addressCommon.String(), nil
}
