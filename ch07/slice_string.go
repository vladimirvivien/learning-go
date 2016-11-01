package main

import "fmt"

func main() {
	msg := "Bobsayshelloworld!"
	fmt.Println(msg[:3], msg[3:7], msg[7:12], msg[12:17], msg[len(msg)-1:])
	fmt.Println(sort(msg))
	printBytes(msg)
}

func sort(str string) string {
	bytes := []byte(str)
	var temp byte
	for i := range bytes {
		for j := i + 1; j < len(bytes); j++ {
			if bytes[j] < bytes[i] {
				temp = bytes[i]
				bytes[i], bytes[j] = bytes[j], temp
			}
		}
	}
	return string(bytes)
}

func printBytes(str string) {
	for i := range str {
		fmt.Println(str[i])
	}
}
