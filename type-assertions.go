/*
A type assertation fornece acesso ao valor concreto subjacente de um valor de interface.

t := i.(T)
Esta declaração afirma que o valor de interface i detém o tipo concreto T e atribui o valor de T subjacente à variável t.

Se i não detém uma T, a declaração irá desencadear um erro pânico.

Para testar se um valor de interface possui um tipo específico, uma type assertation pode retornar dois valores: o valor subjacente e um valor booleano que informa se a afirmação é correta.

t, ok: = i (T).
Se i detém T, então t será o valor subjacente e ok será verdade.

Se não, ok será falso e t será o valor zero do tipo T, e sem ocorrer erro pânico.

Note a semelhança entre esta sintaxe e da leitura de um map.

*/

package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	f = i.(float64) // panic: interface conversion: interface {} is string, not float64
	fmt.Println(f)
}
