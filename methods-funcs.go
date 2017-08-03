/*
Métodos são funções
Lembre-se: um método é apenas uma função com um argumento receptor.

Aqui está Abs escrito como uma função regular, sem qualquer alteração na funcionalidade.
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
