/*
Slices são como referências para matrizes

Uma slice não armazena todos os dados, apenas descreve uma seção de uma matriz subjacente.

Alterando os elementos de uma slice modifica os elementos correspondentes da sua matriz subjacente.

Outras slices que partilham a mesma matriz subjacente vão ver essas mudanças.

*/

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
