/*
Você pode declarar métodos com ponteiro receptor .

Isso significa que o tipo de receptor tem a sintaxe literal *t por algum tipo t. (Além disso, t não pode ser ele próprio, um ponteiro como *int.)

Por exemplo, o método Scale aqui é definido no *Vertex.

Métodos com ponteiro receptores pode modificar o valor ao qual o receptor Os pontos (como Scale faz aqui). Desde métodos muitas vezes precisam modificar seu receptor, receptores de ponteiro são mais comuns do que receptores de valor.

Tente remover o * a partir da declaração da função Scale na linha 16 e observar como o comportamento do programa muda.

Com um receptor de valor, o método Scale opera sobre uma cópia do original `Valor Vertex`. (Este é o mesmo comportamento para qualquer outro argumento da função.) O método Scale deve ter um receptor ponteiro para mudar o `valor Vertex` declarados na função main.
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
