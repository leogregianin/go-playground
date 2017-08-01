// A declaração if parece com o que o de C ou Java fazem, exceto que as ( ) se foram (não são nem mesmo opcionais) e os { } são obrigatórios.

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}