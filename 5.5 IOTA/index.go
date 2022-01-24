package main

import "fmt"

/*
Num identicador de constantes, o identificador iota representa
números sequenciais
pode usar o _ pra pular a sequencia do próxima
pode ocultar o iota das constantes seguintes q ele repete a formulação deixada
*/
const (
	a = iota
	_ = iota
	b
	c = iota * 10
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
