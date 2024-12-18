package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/log"
	"github.com/qiaopengjun5162/go-ethereum-wallet/addresses"
)

func main() {
	addrInfo, err := addresses.CreateAddressFromPrivateKey()
	if err != nil {
		log.Error("create address failed", "err", err)
		return
	}
	log.Info("create address success", "address", addrInfo.Address)
	fmt.Printf("PrivateKey: %s\nPublicKey: %s\nAddress: %s\n", addrInfo.PrivateKey, addrInfo.PublicKey, addrInfo.Address)
}
