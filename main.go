package main

import (
	"fmt"
	"os"

	"github.com/docker/libtrust"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <keyfile>\n", os.Args[0])
		os.Exit(0)
	}

	cmd := os.Args[1]

	if cmd == "-h" || cmd == "--help" {
		fmt.Printf("Usage: %s <keyfile>\n", os.Args[0])
		os.Exit(0)
	}

	keyfilePath := os.Args[1]
	var pubkey libtrust.PublicKey

	if !isExists(keyfilePath) {
		fmt.Println("key file does not exists")
		os.Exit(1)
	}

	// Attempts to read as Private Key
	prikey, err := libtrust.LoadKeyFile(keyfilePath)

	// Attempts to read as Public Key if not private key
	if prikey == nil {
		pubkey, err = libtrust.LoadPublicKeyFile(keyfilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	} else {
		pubkey = prikey.PublicKey()
	}

	fmt.Printf("IsPrivateKey: %t\n", prikey != nil)
	fmt.Printf("FilePath: %s\n", keyfilePath)
	fmt.Printf("KeyType: %s\n", pubkey.KeyType())
	fmt.Printf("KeyID: %s\n", pubkey.KeyID())
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
