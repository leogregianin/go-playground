/*
Você pode declarar um método em tipos que não são struct também.

Neste exemplo, vemos um tipo numérico MyFloat com um método Abs.

Só pode declarar um método com um receptor cujo tipo é definido no mesmo pacote como o método. 

Você não pode declarar um método com um receptor cujo tipo é definido em outro pacote (o qual inclui os tipos padrões tais como int).
*/

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}