package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innnerLen := i + 1
		twoD[i] = make([]int, innnerLen)
		for j := 0; j < innnerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	arr := [2]int{1, 2}
	s2 := arr[:]
	fmt.Println(s2, len(s2), cap(s2)) // [1 2] 2 2

	s2[0] = 99
	fmt.Println(arr) // [99 2] (s2 の underlying array は arr なのでそちらも更新されている)

	s2 = append(s2, 3)
	s2[0] = 999
	fmt.Println(s2, len(s2), cap(s2)) // [999 2 3] 3 4
	fmt.Println(arr)                  // append での拡張時に underlying array 用に別のメモリ領域が確保されているので、999 に上書きされていない

}
