package main

import "fmt"

func main() {

	// s := make([]byte, 5)
	// s[0] = 10
	// s[1] = 13
	// printSlice("s", s)

	// t := make([]byte, len(s), (cap(s)+1)*2)
	// printSlice("t", t)

	// for i := range s {
	// 	t[i] = s[i]
	// }
	// // t = s
	// t[2] = 12
	// t[3] = 14
	// printSlice("s", s)
	// printSlice("t", t)

	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	printSlice("p", p)

	d := []byte{2, 3, 5}
	printSlice("pre d", d)
	d = append(d, 7, 11, 13)
	printSlice("after d", d)

}

func printSlice(name string, t []byte) {
	fmt.Println(name, "=", t, "len(t)=", len(t), "cap(t)=", cap(t))
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
