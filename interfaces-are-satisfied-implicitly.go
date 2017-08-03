/*
Um tipo implementa uma interface através da implementação dos métodos. Não há declaração explícita de intenções, não há palavra-chave "implements".

Interfaces implícitas dissociam-se da definição de uma interface de sua implementação, que pode então aparecer em qualquer pacote sem pré-arranjamento.
*/

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
