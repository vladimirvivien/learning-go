package main

import "fmt"

var ch = make(chan int) // create channel

func Pings() {
	for i := 0; i < 3; i++ {
		fmt.Println("Ping")
	}
	ch <- 1 // send channel value
}

func main() {
	go Pings() // concurrent rexecution
	fmt.Println("Ping")
	<-ch // wait for channel value
	fmt.Println("Pong")
}
