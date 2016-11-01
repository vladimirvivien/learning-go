package main

import "fmt"

type channelWriter chan byte

func (c channelWriter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	count := 0
	wait := make(chan struct{})
	go func() {
		for _, b := range p {
			c <- b
			count++
		}
		close(c)
		close(wait)
	}()
	<-wait
	return count, nil
}

func main() {
	data := []byte("Stream me!")
	cw := channelWriter(make(chan byte, len(data)))
	cw.Write(data)
	for c := range cw {
		fmt.Println(c)
	}
}
