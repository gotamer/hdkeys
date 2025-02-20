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
	flag_number int
	flag_count  int = 1
	flag_debug  bool
	command     string
	argsLength  int
)

func init() {
	argsLength = len(os.Args)
}

func hasNext(no int) bool {
	if argsLength > no {
		return true
	}
	return false
}

func getNext(no int) string {
	return os.Args[no+1]
}

func getNextArgAsInt(no, varsayılan int) int {
	if hasNext(no) {
		if i, err := strconv.Atoi(getNext(no)); err == nil {
			return i
		} else {
			Warn.Println(err)
		}
	} else {
		Warn.Println("Value expected, was not found")
	}
	return varsayılan
}

func flag() {

	if argsLength > 1 {

		// check for commands
		switch os.Args[1] {

		case "wof":
			command = "wof"
			break

		case "keyset":
			command = "keyset"
			break

		case "wif":
			if hasNext(1) {
				fromWifInput(getNext(1))
			} else {
				fromWifInput("")
			}
			os.Exit(0)
		}
	}

	// check for options
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

		case "-d", "--debug":
			debug(true)
			continue

		case "--verbose":
			debug(true)
			verbose(true)
			continue

		case "-no":
			Debug.Println("In case -n")
			flag_number = getNextArgAsInt(no, 0)
			Debug.Println("113 flag_number: ", flag_number)
			continue

		case "-mnemonic":
			if hasNext(no) {
				os.Setenv(ENV_MNEMONIC, getNext(no))
				Info.Printf("Setting env veriable %s\n", ENV_MNEMONIC)
			}
			continue

		case "-pass", "-password", "-passphrase":
			if hasNext(no) {
				os.Setenv(ENV_PASSWORD, getNext(no))
				Info.Printf("Setting env veriable %s\n", ENV_PASSWORD)
			} else {
				Error.Println("No passphrase provided on command line")
			}
			continue

		case "-count":
			Debug.Println("In case -count")
			flag_count = getNextArgAsInt(no, 1)
			continue
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

- hdkeys supports BIP39 passphrase protection.
- hdkeys creates Bitcoin and Nostr accounts from the same mnemonic seeds
- hdkeys can create WIF (Wallet Import Format), and decode private keys from WIF

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

COMMANDS
	wof
		Wall Of Fame, prints a whole set of keys
		COMMAND OPTIONS
			-count [int] (default = 1)
				Set number of keys to generate.

	wif [string] or [Environment variable]
		Decode the private key from wif(Wallet Import Format), then generate the address.

	keyset
		Gets a Bitcoin and Nostr key set with the same WIF (Wallet Import Format) as JSON.
		COMMAND OPTIONS
			-no [int] (default = 0)
				Nostr Account number to generate

OPTIONS
	-mnemonic [string]
		Mnemonic words

	-pass [string]
		Protect bip39 mnemonic with a passphrase via flag,
		or use environment variable,
		or you will be asked to enter at a prompt.

	-h or --help [bool]
		Report usage information and exit.

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
