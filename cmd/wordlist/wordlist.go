package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const pathRoot = "HDKEYS_ROOT_FOLDER"
const pathTarget = "HDKEYS_TARGET_KEY"

func main() {
	dir_root := os.Getenv(pathRoot)
	target_key := os.Getenv(pathTarget)
	sep := os.Getenv("HDKEYS_PASSWORD_WORDS_SEPERATOR")
	l := os.Getenv("HDKEYS_PASSWORD_MAX_LENGTH_WORDS")
	if dir_root == "" || target_key == "" || sep == "" || l == "" {
		fmt.Println("Error: HDKEYS_ env vars not set")
		return
	}

	maxLen, err := strconv.Atoi(l)
	if err != nil {
		fmt.Printf("error HDKEYS_PASSWORD_MAX_LENGTH_WORDS: %s\n", err)
	}

	fileWords := filepath.Join(dir_root, "words.txt")
	fileList := filepath.Join(dir_root, "word_list.txt")
	// History file unique to the public key
	fileHistory := filepath.Join(dir_root, "history_"+target_key+".txt")

	// 1. Load History into a map for fast lookup
	seen := make(map[string]bool)
	if _, err := os.Stat(fileHistory); err == nil {
		hfile, _ := os.Open(fileHistory)
		scanner := bufio.NewScanner(hfile)
		for scanner.Scan() {
			seen[scanner.Text()] = true
		}
		hfile.Close()
	}

	baseWords, _ := readWords(fileWords)

	// 2. Open word_list.txt (overwrite) and history file (append)
	outFile, _ := os.Create(fileList)
	histFile, _ := os.OpenFile(fileHistory, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer outFile.Close()
	defer histFile.Close()

	writer := bufio.NewWriter(outFile)
	hWriter := bufio.NewWriter(histFile)
	defer writer.Flush()
	defer hWriter.Flush()

	// maxLen := 4
	// sep := "@"
	usedBase := make(map[int]bool)

	// Modified stream function to handle "seen" logic
	writeUnique := func(s string) {
		if !seen[s] {
			writer.WriteString(s + "\n")
			hWriter.WriteString(s + "\n")
			seen[s] = true // prevent duplicates within the same run
		}
	}

	for r := 1; r <= maxLen && r <= len(baseWords); r++ {
		generateWithHistory(baseWords, r, []string{}, usedBase, sep, writeUnique)
	}

	fmt.Println("Done! Only new combinations were added to word_list.txt")
}

func generateWithHistory(baseWords []string, r int, current []string, usedBase map[int]bool, sep string, callback func(string)) {
	if len(current) == r {
		// No separator
		callback(strings.Join(current, ""))

		// One separator at each junction
		if len(current) > 1 {
			for i := 1; i < len(current); i++ {
				left := strings.Join(current[:i], "")
				right := strings.Join(current[i:], "")
				callback(left + sep + right)
			}
		}
		return
	}

	for i, word := range baseWords {
		if !usedBase[i] {
			usedBase[i] = true
			lower := strings.ToLower(word)

			// Lowercase branch
			generateWithHistory(baseWords, r, append(current, lower), usedBase, sep, callback)

			// Title case branch
			title := strings.Title(lower)
			if title != lower {
				generateWithHistory(baseWords, r, append(current, title), usedBase, sep, callback)
			}
			usedBase[i] = false
		}
	}
}

func readWords(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		w := strings.TrimSpace(scanner.Text())
		if w != "" {
			words = append(words, w)
		}
	}
	return words, scanner.Err()
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// const pathRoot = "HDKEYS_ROOT_FOLDER"

// func main() {
// 	// 1. Get file path from environment variable
// 	dir_root := os.Getenv(pathRoot)
// 	if dir_root == "" {
// 		fmt.Println("Error: HDKEYS_ROOT_FOLDER variable not set")
// 		return
// 	}

// 	// INPUT FILE: List of words to be combined in to possible passwords
// 	fileWords := filepath.Join(dir_root, "words.txt")
// 	// OUTPUT FILE: List of possible passwords
// 	fileList := filepath.Join(dir_root, "word_list.txt")

// 	baseWords, err := readWords(fileWords)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		return
// 	}

// 	// 2. Setup streaming output
// 	outFile, err := os.Create(fileList)
// 	if err != nil {
// 		fmt.Printf("Error creating file: %v\n", err)
// 		return
// 	}
// 	defer outFile.Close()
// 	writer := bufio.NewWriter(outFile)
// 	defer writer.Flush()

// 	maxLen := 4
// 	sep := "@"
// 	usedBase := make(map[int]bool)

// 	fmt.Printf("Processing %d words. Max length: %d. Saving to %s\n", len(baseWords), maxLen, fileList)

// 	for r := 1; r <= maxLen && r <= len(baseWords); r++ {
// 		streamPermute(baseWords, r, []string{}, usedBase, writer, sep)
// 	}

// 	fmt.Println("Done!")
// }

// func streamPermute(baseWords []string, r int, current []string, usedBase map[int]bool, writer *bufio.Writer, sep string) {
// 	if len(current) == r {
// 		// No separator
// 		writer.WriteString(strings.Join(current, "") + "\n")

// 		// One separator at each junction
// 		if len(current) > 1 {
// 			for i := 1; i < len(current); i++ {
// 				left := strings.Join(current[:i], "")
// 				right := strings.Join(current[i:], "")
// 				writer.WriteString(left + sep + right + "\n")
// 			}
// 		}
// 		return
// 	}

// 	for i, word := range baseWords {
// 		if !usedBase[i] {
// 			usedBase[i] = true

// 			lower := strings.ToLower(word)
// 			// Branch 1: Lowercase
// 			streamPermute(baseWords, r, append(current, lower), usedBase, writer, sep)

// 			// Branch 2: Title Case (only if word actually changes)
// 			title := strings.Title(lower)
// 			if title != lower {
// 				streamPermute(baseWords, r, append(current, title), usedBase, writer, sep)
// 			}

// 			usedBase[i] = false
// 		}
// 	}
// }
