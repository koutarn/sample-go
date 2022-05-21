package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["Mike"] = 20
	m["Nancy"] = 24
	m["Messi"] = 30
	fmt.Printf("%T %v", m, m)

}
