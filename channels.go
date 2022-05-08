package main

import "fmt"

func main() {
	channnel := make(chan string)

	go func() { channnel <- "done" }()

	msg := <-channnel
	fmt.Println(msg)
}
