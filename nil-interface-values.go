/*
Um valor de interface nil detém nem o valor nem tipo concreto.

Chamar um método em uma interface nula resulta em um erro de tempo de execução porque não há um tipo dentro da tupla de interface para indicar qual o método concreto a chamar.

*/

package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M() // <-- error
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
