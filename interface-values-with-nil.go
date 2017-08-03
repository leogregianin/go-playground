/*
Se o valor concreto no interior da própria interface é nulo, o método será chamado com um receptor nulo.

Em algumas linguagens isso iria desencadear uma exceção de ponteiro nulo, mas em Go é comum para escrever métodos que normalmente lidam sendo chamado com um receptor nil (como acontece com o método M neste exemplo).

Nota-se que um valor de interface que contém um valor nulo concreto é por si não-nulo.
*/

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
