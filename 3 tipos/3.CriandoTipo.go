package main

import "fmt"

type hotdog int

var b hotdog = 1

func main() {
	x := 10
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", b)
	b = 3
	fmt.Println(b)
	/*
		Mas aqui não posso passar b = x mas porque não serão do mesmo tipo, pois x será int
		em vez de hotdog
	*/

	/* Type Conversion em go só tem isso, não tem casting, um tipo pode ser convertido dessa forma
	 */
	x = int(b)
	fmt.Println(x)

}
