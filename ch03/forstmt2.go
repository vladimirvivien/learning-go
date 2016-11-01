package main

import (
	"fmt"
	"math/rand"
)

var list1 = []string{"break", "lake", "go", "right", "strong", "kite", "hello"}
var list2 = []string{"fix", "river", "stop", "left", "weak", "flight", "bye"}

func main() {
	rand.Seed(31)
	for w1, w2 := nextPair(); w1 != "go" && w2 != "stop"; w1, w2 = nextPair() {
		fmt.Printf("Word Pair -> [%s, %s]\n", w1, w2)
	}
}

func nextPair() (w1, w2 string) {
	pos := rand.Intn(len(list1))
	return list1[pos], list2[pos]
}
