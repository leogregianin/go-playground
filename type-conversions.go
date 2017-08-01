/*
A expressão T(v) converte o valor v para o tipo T.

Algumas conversões numéricas:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Ou, de uma forma mais simples:

i := 42
f := float64(i)
u := uint(f)
Diferente de C, em Go atribuição entre os itens de tipo diferente requer uma conversão explícita. Tente remover as conversões float64 ou int no exemplo e veja o que acontece.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
