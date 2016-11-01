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

// mapWords loads the content of the specified file's name into and uses the specified
// function to generate a tuple containng the word-class signature and the original word.
func mapWords(fname string, f func(word string) (string, string)) [][]string {
	file, err := os.Open(fname)
	if err != nil {
		panic("Unable to load file " + fname)
	}
	defer file.Close()

	var tuples [][]string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		key, word := f(scanner.Text())
		tuples = append(tuples, []string{key, word})
	}
	if scanner.Err() != nil {
		panic("Error parsing file " + fname)
	}
	return tuples
}

// bySigKv returns a word-class signature for a given word.
// It is used as a parameter to the mapWords method defined earlier.
func bySigKv(str string) (string, string) {
	return sortRunes(str), str
}

// reduceWords creates a map of the array of tuples produced the mapping step.
// Each word is collected and mapped to its associated signature.
func reduceWords(tuples [][]string) map[string][]string {
	anagrams := make(map[string][]string)
	for _, tuples := range tuples {
		anagrams[tuples[0]] = append(anagrams[tuples[0]], tuples[1])
	}
	return anagrams
}

func main() {
	tuples := mapWords("dict.txt", bySigKv)
	anagrams := reduceWords(tuples)
	for k, v := range anagrams {
		fmt.Println(k, "->", v)
	}

	fmt.Println("Anagrams", len(anagrams))
}
