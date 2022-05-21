package main

import "fmt"

type Vertex struct {
	X, Y int64
}

func (v *Vertex) Plus() int64 {
	return v.X + v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Plus())
}
