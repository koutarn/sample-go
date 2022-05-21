package main

import "fmt"

func producer(first chan<- int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for v := range first {
		second <- v * 2
	}
}

func multi4(second <-chan int, third chan<- int) {
	defer close(third)
	for v := range second {
		third <- v * 4
	}
}

func main() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first)
	go multi2(first, second)
	go multi4(second, third)

	for result := range third {
		fmt.Println(result)
	}
}
