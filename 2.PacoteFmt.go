package main

import "fmt"

func main() {
	x := "oi bom dia\ncomo vai?\tespero que tudo bem."
	fmt.Println(x)
	y := `oi bom dia\ncomo vai?\tespero que tudo bem.`
	fmt.Println(y)
	// Template raw no segundo caso não vai ser interpretado e mostrará do jeito que for
	// feito, inclusive se der enter e etc.
	/*	println adiciona uma linha no final e tem seu retorno a informação e um erro mais a
		linha no final.
		Sprint ele pega o que foi passado e simplesmente retorna isso numa string, ele só adiciona
		espaço se não com string
		Fprint retorna algo relaciona a file, bytecode essas coisas chegar documentação
		deve ser bem usado pra desenvolvimento
	*/
}
