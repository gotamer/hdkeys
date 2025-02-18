package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/gotamer/hdkeys"
	"golang.org/x/term"
)

var (
	km *hdkeys.KeyManager
)

func main() {
	var err error
	flag()

	var pass = readPassword()
	var mnemonic = readMnemonic()

	km, err = hdkeys.NewKeyManager(mnemonic, pass)
	if err != nil {
		Error.Fatal(err)
	}
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
	for i := 0; i < flag_number; i++ {
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
	for i := 0; i < flag_number; i++ {
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
	for i := 0; i < flag_number; i++ {
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
	for i := 0; i < flag_number; i++ {
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
	for i := 0; i < flag_number; i++ {
		key, err := km.GetKey(hdkeys.PurposeBIP44, hdkeys.CoinTypeNostr, uint32(i), 0, 0)
		if err != nil {
			Error.Fatal(err)
		}
		wif, _, _, _, _, err := key.Calculate(compress)
		if err != nil {
			Error.Fatal(err)
		}

		if nsec, err := hdkeys.NostrGetPrivateKey(key.Bip32Key.Key); err == nil {
			fmt.Printf("%-18s WIF:   %s\n", key.Path, wif)
			fmt.Printf("%-18s Nostr: %s\n", key.Path, nsec)
			if npub, err := hdkeys.NostrGetPublicKey(key.Bip32Key.Key); err != nil {
				Error.Println(err)
			} else {
				fmt.Printf("%-18s Nostr: %s\n\n", key.Path, npub)
			}
		} else {
			Info.Println(err)
		}
	}
	fmt.Println()

}

func readPassword() string {

	if password, ok := os.LookupEnv(ENV_PASSWORD); ok {
		Info.Println("using password from env")
		return password
	}

	var pass1, pass2 string
PASS:
	fmt.Print("\nEnter Password: ")
	if p, err := term.ReadPassword(int(syscall.Stdin)); err != nil {
		fmt.Println(err)
		goto PASS
	} else {
		pass1 = strings.TrimSpace(string(p))
	}

	//TODO: Validate password length, ASCII etc

	fmt.Print("\nVerify password: ")
	if p, err := term.ReadPassword(int(syscall.Stdin)); err != nil {
		fmt.Println(err)
		goto PASS
	} else {
		pass2 = strings.TrimSpace(string(p))
	}

	if pass1 != pass2 {
		fmt.Println("\nPasswords do NOT match! Try again.")
		goto PASS
	}
	fmt.Println("\n")
	return pass1
}

func readWordsFile() (words string) {
	var mw []byte
	var err error
	var dir string
	if dir, err = os.Getwd(); err == nil {
		var path = filepath.Join(dir, "mnemonic-words")
		//var path = string(p)
		if _, err = os.Stat(path); err == nil {
			if mw, err = os.ReadFile(path); err == nil {
				words = string(mw)
				Info.Println("mnemonic-words from pwdr")
			}
		}
	}
	Info.Println(err)
	if dir, err = os.UserConfigDir(); err == nil {
		var path = filepath.Join(dir, "mnemonic-words")
		//var path = string(p)
		if _, err = os.Stat(path); err == nil {
			if mw, err = os.ReadFile(path); err == nil {
				words = string(mw)
				Info.Println("mnemonic-words from UserConfigDir")
			}
		}
	}
	Info.Println(err)
	return string(words)
}

func readMnemonic() (words string) {

	words, _ = os.LookupEnv(ENV_MNEMONIC)

	if len(words) < 10 {
		words = readWordsFile()
	}

	words = hdkeys.CleanWords(words)
	Debug.Println(words)

	if len(words) < 10 {
		words = ""
		return
	}

	if !hdkeys.ValidateWords(words) {
		Info.Println(words)
		Error.Fatal("Could not validate Mnemonic words!")
	}
	return
}

func fromWifInput() {
	var input string

	Info.Println("\n\nDecodeing private key from wif.")
	input, _ = os.LookupEnv(ENV_WIF)
	if len(input) < 5 {
		fmt.Print("Enter wif key: ")
		if p, err := term.ReadPassword(int(syscall.Stdin)); err != nil {
			Error.Fatal(err)
		} else {
			input = strings.TrimSpace(string(p))
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
	fmt.Println()
}
