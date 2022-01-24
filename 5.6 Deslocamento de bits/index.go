package main

import "fmt"

/*
Deslocamento de bits é quando deslocamos dígitos binários
para a direita ou esquerda.

<< ou >>
*/
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) /* 1024 // 1 << (1 * 10)  ou seja o número 1
	// vai deslocar seu número em binário 10 casas para a esquerda */
	MB
	GB
	TB
)

func main() {
	fmt.Println("Binary\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t\t\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t\t\t", TB)
	fmt.Printf("%d\n", TB)

}
