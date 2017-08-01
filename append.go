/*

É comum acrescentar novos elementos em uma slice, então Go dispões de uma função padrão para isso o append. A documentação do pacote padrão descreve melhor o append.

func append(s []T, vs ...T) []T
O primeiro parâmetro s do append é uma slice do tipo T, e o resto são valores de T para acrescentar na slice.

O valor resultante do append é uma slice contendo todos os elementos da slice original mais os valores providos.

Se a matriz anterior de s for muito pequena para caber todos os valores uma matriz gigante será alocada. A slice retornada apontará para a nova matriz alocada.

*/

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
