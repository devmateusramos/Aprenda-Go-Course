package main

import "fmt"

var a int
var b float64
var c string
var d bool

// declara uma variável e não utilizar ela vai ter por padrão o valor zero

func main() {
	fmt.Printf("a: %v %T \n", a, a)
	fmt.Printf("b: %v %T \n", b, b)
	fmt.Printf("c: %v %T \n", c, c)
	fmt.Printf("d: %v %T \n", d, d)
}
