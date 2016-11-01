package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink from its water bowl",
		"The dark bird of prey lands on a small tree after hunting for fish",
	}

	histogram := make(map[string]int)
	stopCh := make(chan struct{})

	words := words(stopCh, data) // returns handle to channel
	for word := range words {
		if histogram["the"] == 3 {
			close(stopCh)
		}
		histogram[word]++
	}

	for k, v := range histogram {
		fmt.Printf("%s\t(%d)\n", k, v)
	}
}

// generator function that produces data
func words(stopCh chan struct{}, data []string) <-chan string {
	out := make(chan string)

	// splits line and emit words
	go func() {
		defer close(out) // closes channel upon fn return
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				select {
				case out <- word:
				case <-stopCh:
					return
				}
			}
		}
	}()

	return out
}
