package main

import "fmt"

func main() {

	c := make([]int, 5)
	d := make([]int, 0, 5)

	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Printf("len %d cap %d v %v\n", len(c), cap(c), c)
	}

	for i := 0; i < 5; i++ {
		d = append(d, i)
		fmt.Printf("len %d cap %d v %v\n", len(d), cap(d), d)
	}
}
