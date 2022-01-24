package main

import "fmt"

/*
	São variáveis imutáveis
	Podem ser tipadas ou não
	As não tipadas só terão um tipo atribuido a elas
	quando forem usadas
	Pode declarar várias de uma vez.
*/
const b = 20

const (
	e = 12
	f = "olá"
)

var c int
var d float64

func main() {
	const a = 10
	c = b
	fmt.Println(c)
	d = a
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
