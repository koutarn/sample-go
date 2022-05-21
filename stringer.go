package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %v.\n I'm %v old\n", p.Name, p.Age)
}

func main() {
	mike := Person{"Mike", 20}
	fmt.Println(mike)
}
