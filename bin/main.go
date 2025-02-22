package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/gotamer/hdkeys"
	"golang.org/x/term"
)

var (
	km *hdkeys.KeyManager
)

func main() {
	flag()

		var pass = readPassword()
		var mnemonic = readMnemonic()

		var err error
		km, err = hdkeys.NewKeyManager(mnemonic, pass)
		if err != nil {
			Error.Fatal(err)
		}

	switch command {

	case "wof":
		wof(flag_count)
		return

	case "keyset":
		Info.Println("flag_number: ", flag_number)
		if json, err := GetAllJson(uint32(flag_number)); err != nil {
			Error.Println(err)
		} else {
			fmt.Println(json)
		}
		return
	}
}

func readPassword() string {

	if password, ok := os.LookupEnv(ENV_PASSWORD); ok {
		Info.Println("using password from env: ", password)
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
	Warn.Println(err)
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
	Warn.Println(err)
	return string(words)
}

func readMnemonic() (words string) {

	words, _ = os.LookupEnv(ENV_MNEMONIC)

	if len(words) < 10 {
		Debug.Println("Mnemonic to short: ", words)
		words = readWordsFile()
	}

	words = hdkeys.CleanWords(words)
	Debug.Println("Mnemonic clean: ", words)

	if len(words) < 10 {
		Debug.Println("Mnemonic to short: ", words)
		words = ""
		return
	}

	if !hdkeys.ValidateWords(words) {
		Debug.Println("Mnemonic Verify: ", words)
		Error.Fatal("Could not validate Mnemonic words!")
	}
	return
}
