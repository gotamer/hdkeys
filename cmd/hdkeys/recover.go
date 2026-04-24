package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"go.hansaray.pw/hdkeys/lib"
)

func recover(from, to uint32) {

	var mnemonic = readMnemonic()

	targetKey := os.Getenv("HDKEYS_TARGET_KEY")
	if targetKey == "" {
		fmt.Println("Error: HDKEYS_TARGET_KEY not set")
		return
	}

	dirRoot := os.Getenv("HDKEYS_ROOT_FOLDER")
	if dirRoot == "" {
		fmt.Println("Error: HDKEYS_ROOT_FOLDER not set")
		return
	}

	// 2. Setup Concurrency
	jobs := make(chan string, 100)
	results := make(chan string)
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()

	// Start Workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for password := range jobs {
				if path, found := checkPassword(mnemonic, password, targetKey, from, to); found {
					results <- fmt.Sprintf("\nMATCH FOUND!\nPass: %s\nPath: %s", password, path)
				}
			}
		}()
	}

	// 3. Stream File into Job Channel
	go func() {
		file, _ := os.Open(filepath.Join(dirRoot, "word_list.txt"))
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jobs <- scanner.Text()
		}
		close(jobs)
		wg.Wait()
		close(results)
	}()

	// 4. Wait for Match or End
	fmt.Printf("Searching with %d workers...\n", numWorkers)
	for res := range results {
		fmt.Println(res)
		os.Exit(0) // Stop immediately on first match
	}
	fmt.Println("Search complete. No matches.")
}

func checkPassword(mnemonic, password, target string, from, to uint32) (string, bool) {
	km, err := hdkeys.NewKeyManager(mnemonic, password)
	if err != nil {
		return "", false
	}

	var purpose uint32
	var getAddress func(k *hdkeys.Key) (string, error)

	switch {
	case strings.HasPrefix(target, "bc1p"):
		purpose = hdkeys.PurposeBIP86
		getAddress = func(k *hdkeys.Key) (string, error) {
			_, _, _, _, tap, err := k.Calculate(true)
			return tap, err
		}
	case strings.HasPrefix(target, "bc1q"):
		purpose = hdkeys.PurposeBIP84
		getAddress = func(k *hdkeys.Key) (string, error) {
			_, _, bech, _, _, err := k.Calculate(true)
			return bech, err
		}
	case strings.HasPrefix(target, "3"):
		purpose = hdkeys.PurposeBIP49
		getAddress = func(k *hdkeys.Key) (string, error) {
			_, _, _, nested, _, err := k.Calculate(true)
			return nested, err
		}
	case strings.HasPrefix(target, "1"):
		purpose = hdkeys.PurposeBIP44
		getAddress = func(k *hdkeys.Key) (string, error) {
			_, addr, _, _, _, err := k.Calculate(true)
			return addr, err
		}
	default:
		return "", false
	}

	for i := from; i <= to; i++ {
		// GetKey handles m/purpose'/0'/0'/0/i
		k, err := km.GetKey(purpose, 0+hdkeys.Apostrophe, 0, 0, i)
		if err != nil {
			continue
		}

		addr, err := getAddress(k)
		if err == nil && addr == target {
			return k.Path, true
		}
	}

	return "", false
}
