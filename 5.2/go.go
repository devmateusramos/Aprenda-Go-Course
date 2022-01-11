package main

import (
	"fmt"
	"runtime"
)

var i uint16

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	a := "e"
	b := "é"
	c := []byte(a)
	d := []byte(b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	i = 65535 // se Adicionar 65536 vai lançar erro, e se só somar 1 vai voltar pra 0
	// isso é o overflow
	fmt.Println(i)
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)
	i = 340
	fmt.Println(i)
}
