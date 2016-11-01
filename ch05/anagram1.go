// anagram.go
//
// This is an implementation of the anagram problem (Column 2)
// from the popular Programming Pearls book (second ed) from
// author Jon Bentley.
// https://books.google.com/books?id=kse_7qbWbjsC
//
// The data - dict.txt
// The data file is a subset of the Linux dictionary file
// found in /usr/share/dict/american-english. It is scrubbed of
// all posessive nounds and made all lowe rcase.
package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
)

// sortRunes is a simple insertion sort that sorts the runes of the given string.
// By sorting the unicode chars of the string (i.e. "morning" -> "gimnnor")
// the result can be used as a signature to create words in the same class.
func sortRunes(str string) string {
	runes := bytes.Runes([]byte(str))
	var temp rune
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[j] < runes[i] {
				temp = runes[i]
				runes[i], runes[j] = runes[j], temp
			}

		}
	}
	return string(runes)
}

// load loads the content of the specified file's name into
// memory as an slice (array) of strings.  Notice the function
// also returns an error value that will be non-nil if an error
// occured during the load process.
// Scanner.Split() uses a function that knows
// to split the file records.  Here the code uses a provided function bufio.ScanLines
// to do the job of splitting the records from the file and returns a word
// for each line read and saves in slice of words.
func load(fname string) ([]string, error) {
	if fname == "" {
		return nil, errors.New("Dictionary file name cannot be emtpy.")
	}

	// attempt to load dict file.
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	// loads dictionary words in memory
	words, err := load("dict.txt")

	// test err for non-nil value.
	// If err != nil, then the error is handled.
	// In this context, the program ends.
	if err != nil {
		fmt.Println("Unable to load file:", err)
		os.Exit(1)
	}

	// the next code snippet goes through the words slice
	// and does the followings for each word in the slice:
	// -  Creates a word signature by soriting the letters in the word itself.
	// -  Maps the word to its signature
	// So, words with the same signature will be mapped together (i.e. abers -> sabre, bares, bears, saber, etc)
	anagrams := make(map[string][]string)
	for _, word := range words {
		wordSig := sortRunes(word)
		anagrams[wordSig] = append(anagrams[wordSig], word)
	}

	// prints the result
	for k, v := range anagrams {
		fmt.Println(k, "->", v)
	}
}
