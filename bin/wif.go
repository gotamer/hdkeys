package main

import (
	"fmt"
	"os"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/gotamer/hdkeys"
)

func fromWifInput(input string) {

	Info.Println("\n\nDecodeing private key from wif.")
	if len(input) < 50 {
		input, _ = os.LookupEnv(ENV_WIF)
		if len(input) < 50 {
			Error.Fatal("Not a WIF (Wallet Import Format)")
		}
	}

	wif, err := btcutil.DecodeWIF(input)
	if err != nil {
		Error.Fatal(err)
	}

	wifCompressed, addressCompressed, segwitBech32, segwitNested, taproot, err := hdkeys.CalculateFromPrivateKey(wif.PrivKey, true)
	if err != nil {
		Error.Fatal(err)
	}

	wifUncompressed, addressUncompressed, _, _, _, err := hdkeys.CalculateFromPrivateKey(wif.PrivKey, false)
	if err != nil {
		Error.Fatal(err)
	}

	fmt.Println("\n Wallet Import Format:")
	fmt.Printf(" *   %-24s %s\n", "WIF(compressed):", wifCompressed)
	fmt.Printf(" *   %-24s %s\n", "WIF(uncompressed):", wifUncompressed)

	fmt.Println("\n Public Addresses:")
	fmt.Printf(" *   %-24s %s\n", "Legacy(compresed):", addressCompressed)
	fmt.Printf(" *   %-24s %s\n", "Legacy(uncompressed):", addressUncompressed)
	fmt.Printf(" *   %-24s %s\n", "SegWit(nested):", segwitNested)
	fmt.Printf(" *   %-24s %s\n", "SegWit(bech32):", segwitBech32)
	fmt.Printf(" *   %-24s %s\n", "Taproot(bech32m):", taproot)
}
