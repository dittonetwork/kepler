package main

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/crypto"
)

const argsCount = 2

// main function to decode secp256k1 base64-encoded public key and print the Ethereum address.
func main() {
	if len(os.Args) < argsCount {
		log.Print("Usage: ethaddr <base64-encoded-pubkey>")
		os.Exit(1)
	}

	pubkeyBase64 := os.Args[1]
	// Decode the base64 public key
	pubkeyBytes, err := base64.StdEncoding.DecodeString(pubkeyBase64)
	if err != nil {
		log.Printf("Error decoding public key: %v", err)
		os.Exit(1)
	}

	pubKey, err := btcec.ParsePubKey(pubkeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate Ethereum address from public key
	address := crypto.PubkeyToAddress(*pubKey.ToECDSA())
	log.Printf("Ethereum address: %s", address.Hex())
}
