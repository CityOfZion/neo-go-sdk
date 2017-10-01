package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/CityOfZion/neo-go-sdk/neo"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "--wif" {
		outputUsage()
	}

	wif := os.Args[2]

	if len(wif) != 52 {
		outputError("WIF entered is not 52 characters long.")
	}

	privateKey, err := neo.NewPrivateKeyFromWIF(wif)
	if err != nil {
		outputError("Unable to convert WIF to private key.")
	}

	publicAddress, err := privateKey.PublicAddress()
	if err != nil {
		outputError("Error when deriving public address from private key.")
	}

	publicKeyBytes, err := privateKey.PublicKey()
	if err != nil {
		outputError("Error when deriving public key from private key.")
	}
	publicKey := hex.EncodeToString(publicKeyBytes)

	fmt.Println("Details:")
	fmt.Printf(" - NEO address (compressed): \t%s\n", publicAddress)
	fmt.Printf(" - Private key: \t\t%s\n", privateKey.Output())
	fmt.Printf(" - Private key (base64): \t%s\n", privateKey.OutputBase64())
	fmt.Printf(" - Public key (compressed): \t%s\n", publicKey)
	fmt.Printf(" - WIF (compressed): \t\t%s\n", wif)
}

func outputError(message string) {
	fmt.Printf("[ERROR] %s\n", message)
	fmt.Println("Try: neo-go-sdk --help")
	os.Exit(1)
}

func outputUsage() {
	fmt.Println("Tool to help debug a NEO public and private key pair.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("\tneo-go-sdk --wif <WIF>")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("\t--wif\tWIF (wallet import format) for NEO private key.")
	fmt.Println("\t--help\tPrint usage.")
	os.Exit(1)
}
