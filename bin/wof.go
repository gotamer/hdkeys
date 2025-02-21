package main

import (
	"fmt"
	"strings"

	"github.com/gotamer/hdkeys"
)

func wof(count int) {

	masterKey, err := km.GetMasterKey()
	if err != nil {
		Error.Fatal(err)
	}

	passphrase := km.Passphrase
	if passphrase == "" {
		passphrase = "<none>"
	}

	fmt.Printf("\n%-18s %s\n", "BIP39 Mnemonic:", km.Mnemonic)
	fmt.Printf("%-18s %s\n", "BIP39 Passphrase:", passphrase)
	fmt.Printf("%-18s %x\n", "BIP39 Seed:", km.GetSeed())
	fmt.Printf("%-18s %s\n", "BIP32 Root Key:", masterKey.B58Serialize())

	fmt.Printf("\n%-18s %-34s %-52s\n", "Path(BIP44)", "Legacy(P2PKH, compresed)", "WIF(Wallet Import Format)")
	fmt.Println(strings.Repeat("-", 106))
	for i := 0; i < count; i++ {
		key, err := km.GetKey(hdkeys.PurposeBIP44, hdkeys.CoinTypeBTC, 0, 0, uint32(i))
		if err != nil {
			Error.Fatal(err)
		}
		wif, address, _, _, _, err := key.Calculate(compress)
		if err != nil {
			Error.Fatal(err)
		}

		fmt.Printf("%-18s %-34s %s\n", key.Path, address, wif)
	}

	fmt.Printf("\n%-18s %-34s %s\n", "Path(BIP49)", "SegWit(P2WPKH-nested-in-P2SH)", "WIF(Wallet Import Format)")
	fmt.Println(strings.Repeat("-", 106))
	for i := 0; i < count; i++ {
		key, err := km.GetKey(hdkeys.PurposeBIP49, hdkeys.CoinTypeBTC, 0, 0, uint32(i))
		if err != nil {
			Error.Fatal(err)
		}
		wif, _, _, segwitNested, _, err := key.Calculate(compress)
		if err != nil {
			Error.Fatal(err)
		}
		fmt.Printf("%-18s %s %s\n", key.Path, segwitNested, wif)
	}

	fmt.Printf("\n%-18s %-42s %s\n", "Path(BIP84)", "SegWit(P2WPKH, bech32)", "WIF(Wallet Import Format)")
	fmt.Println(strings.Repeat("-", 114))
	for i := 0; i < count; i++ {
		key, err := km.GetKey(hdkeys.PurposeBIP84, hdkeys.CoinTypeBTC, 0, 0, uint32(i))
		if err != nil {
			Error.Fatal(err)
		}
		wif, _, segwitBech32, _, _, err := key.Calculate(compress)
		if err != nil {
			Error.Fatal(err)
		}

		fmt.Printf("%-18s %s %s\n", key.Path, segwitBech32, wif)
	}

	fmt.Printf("\n%-18s %-62s %s\n", "Path(BIP86)", "Taproot(P2TR, bech32m)", "WIF(Wallet Import Format)")
	fmt.Println(strings.Repeat("-", 134))
	for i := 0; i < count; i++ {
		key, err := km.GetKey(hdkeys.PurposeBIP86, hdkeys.CoinTypeBTC, 0, 0, uint32(i))
		if err != nil {
			Error.Fatal(err)
		}
		wif, _, _, _, taproot, err := key.Calculate(compress)
		if err != nil {
			Error.Fatal(err)
		}

		fmt.Printf("%-18s %s %s\n", key.Path, taproot, wif)
	}

	fmt.Printf("\n%-18s %-42s\n", "Path(BIP44)", "Nostr")
	fmt.Println(strings.Repeat("-", 80))
	for i := 0; i < count; i++ {

		nostr, err := km.GetNostr(uint32(i))
		if err != nil {
			Error.Fatal(err)
		}
		fmt.Printf("%-18s WIF:   %s\n", nostr.Path, nostr.Wif)
		fmt.Printf("%-18s Nostr: %s\n", nostr.Path, nostr.NSec)
		fmt.Printf("%-18s Nostr: %s\n", nostr.Path, nostr.NPub)
		fmt.Printf("%-18s PHEX: %s\n\n", nostr.Path, nostr.PHex)
	}
	fmt.Println()

}
