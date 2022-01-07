package main

import "fmt"

type creation int

var x creation

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v", x)
}
