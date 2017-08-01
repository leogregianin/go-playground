/*
O tipo [n]T é uma matriz de n valor do tipo T.

A expressão "var a [10]int" declara a variável a com uma matriz de dez inteiros.

O tamanho da matriz é parte de seu tipo, logo, matrizes não podem ser redimensionadas. 

Isso parece uma limitação, mas não se preocupe; Go dispõe de uma maneira conveniente para trabalhar com matrizes.

*/

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}