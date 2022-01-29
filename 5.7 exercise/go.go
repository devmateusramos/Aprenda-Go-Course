package main

import "fmt"

const (
	_ = 2021 + iota
	a
	b
	c
	d
)

func main() {
	x := 10
	fmt.Printf("%d, %#x, %b\n", x, x, x)
	x = x >> 1
	fmt.Printf("%d, %#x, %b\n", x, x, x)
	y := `isto
 				é uma coisa
		muita doida`
	fmt.Printf("%v", y)
	fmt.Println(a, b, c, d)
}
