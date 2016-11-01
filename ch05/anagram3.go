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

var (
	emptyFileNameError = errors.New("File name cannot be empty")
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
func load(fname string) ([]string, error) {
	if fname == "" {
		return nil, errors.New("Dictionary file name is empty")
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

// write writes the map of anagrams to a file file specified by fname.
// The use of os.OpenFile (instead of os.Create) is to demonstrate the use
// of panic() which is generated when the OpenFile returns an error.
func write(fname string, anagrams map[string][]string) {
	if anagrams == nil {
		panic("Unable to write, anagrams missing.")
	}

	file, err := os.OpenFile(fname, os.O_WRONLY+os.O_CREATE+os.O_EXCL, 0644)
	if err != nil {
		msg := fmt.Sprintf("Unable to open outpput file for creation: %v", err)
		panic(msg)
	}
	defer file.Close()
	for k, v := range anagrams {
		output := fmt.Sprintf("%s -> %v\n", k, v)
		file.WriteString(output)
	}
}

// mapWords maps each word to its associated signature
func mapWords(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	for _, word := range words {
		wordSig := sortRunes(word)
		anagrams[wordSig] = append(anagrams[wordSig], word)
	}

	return anagrams
}

// makeAnagrams creates the anagram and writes them to a file.
// It also capture any panic situtaion and recovers gracefully.
func makeAnagrams(words []string, fname string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Failed to make anagram:", r)
		}
	}()

	anagrams := mapWords(words)
	write(fname, anagrams)
}

func main() {
	words, err := load("dict.txt")
	if err != nil {
		fmt.Println("Unable to load file:", err)
		os.Exit(1)
	}

	// generate anagrams and write to file
	makeAnagrams(words, "out.txt")
}
