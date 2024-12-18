package addresses

import (
	"fmt"
	"testing"
)

func TestCreateAddressFromPrivateKey(t *testing.T) {
	address, err := CreateAddressFromPrivateKey()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("PrivateKey: %s\nPublicKey: %s\nAddress: %s\n", address.PrivateKey, address.PublicKey, address.Address)
}

//	{
//			"address":"0x6Fe908602d5606D6a83257D3e054688c24E39072",
//			"publicKey":"0x0292f95b732a3085a0bd08c6bbadb8a04a32da974d4094d2fd9ee3ba2db175f8b5",
//			"privateKey":"0xdc2d6117326e9953bc997045df045ea87ebb1c974b49580fafa92fe9a7336ef9"
//	}
func TestPublicKeyToAddress(t *testing.T) {
	address, err := PublicKeyToAddress("0292f95b732a3085a0bd08c6bbadb8a04a32da974d4094d2fd9ee3ba2db175f8b5")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Address: %s\n", address)
}
