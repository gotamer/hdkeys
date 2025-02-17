// flag - cli flags
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gotamer/hdkeys/build"
)

const (
	ENV_PASSWORD = "HDKEYS_PASSWORD"
	ENV_MNEMONIC = "HDKEYS_MNEMONIC"
	ENV_WIF      = "HDKEYS_WIF"
)
const FILE_NAME_MNEMONIC = "mnemonic-words"
const compress = true // generate a compressed public key

var (
	flag_number int = 1
	flag_debug      = false
)

func flag() {

	for no, arg := range os.Args {
		switch arg {

		case "-h", "--help":
			print(help)
			os.Exit(0)

		case "-v":
			fmt.Printf("%s\n", build.Version)
			os.Exit(0)

		case "--version":
			fmt.Printf("%s\n", build.Info())
			os.Exit(0)

		case "-mnemonic":
			var mnemonic = os.Args[no+1]
			os.Setenv(ENV_MNEMONIC, mnemonic)
			Info.Printf("Setting env veriable %s\n", ENV_MNEMONIC)
			continue

		case "-pass":
			var pass = os.Args[no+1]
			os.Setenv(ENV_PASSWORD, pass)
			Info.Printf("Setting env veriable %s\n", ENV_PASSWORD)
			continue

		case "-debug":
			debug(true)
			continue

		case "-verbose":
			verbose(true)
			continue

		case "-n":
			if num, err := strconv.Atoi(os.Args[no+1]); err == nil {
				flag_number = num
				Debug.Println("flag_number: ", flag_number)
			} else {
				Error.Println("Flag -n needs a number. 1 is the default.")
				Error.Fatal(err)
			}
			continue

		case "wif":
			fromWifInput()
			os.Exit(0)

		case "-wif", "--wif":
			var a = os.Args[no+1]
			os.Setenv(ENV_WIF, a)
			Info.Printf("Setting env veriable %s\n", ENV_WIF)
			fromWifInput()
			os.Exit(0)
		}
	}
}

const help = `
hdkeys - Mnemonic seeds and Hierarchical Deterministic (HD) addresses.

SYNOPSIS
	hdkeys [OPTIONS]
	hdkeys [OPTIONS] [COMMAND] [COMMAND OPTIONS]

ENVIRONMENT VARIABLE OPTIONS
	HDKEYS_PASSWORD
	HDKEYS_MNEMONIC
	HDKEYS_WIF

DESCRIPTION
hdkeys allows for the creation of mnemonic seeds, and Hierarchical Deterministic (HD) addresses.

    BIP32 - Hierarchical Deterministic Wallets
    BIP39 - Mnemonic code for generating deterministic keys
    BIP43 - Purpose Field for Deterministic Wallets
    BIP44 - Multi-Account Hierarchy for Deterministic Wallets
    BIP49 - Derivation scheme for P2WPKH-nested-in-P2SH based accounts
    BIP84 - Derivation scheme for P2WPKH based accounts
    BIP86 - Derivation scheme for Pay-to-Taproot (P2TR) based accounts
    BIP173 - Base32 address format for native v0-16 witness outputs
    SLIP44 - Registered coin types for BIP-0044
    NIP06 - Basic key derivation from mnemonic seed phrase
	...

COMMANDS:

	wif [prompt] or [Environment variable]
		Decode the private key from wif(Wallet Import Format), then generate the address.

OPTIONS:

	-h or --help [bool]
		Report usage information and exit.

	-wif [string]
		Decode the private key from wif(Wallet Import Format), then generate the address.

	-pass [string]
		Protect bip39 mnemonic with a passphrase via flag,
		or use environment variable,
		or you will be asked to enter at a prompt.

	-n [int] (default = 1)
		Set number of keys to generate.

	-v [bool]
		Print version tag and exit.

	--version [bool]
		Print detailed version and exit.

	-d or --debug [bool]
		Print debug information about the process.

	--verbose [bool]
		Print extra verbose debug information about the process.

REPORTING BUGS
	npub12jjczvd2mzstyhr468fyas7vzmsm5d2x3tv5l9tev6q0jakk9djqx4uk7x

REPO
	https://github.com/gotamer/hdkeys

COPYRIGHT
	© 2025 WTFPL – Do What the Fuck You Want to Public License. (http://www.wtfpl.net)
`
