/*
Variáveis declaradas sem um valor inicial explicitado darão seu valor zero.

O valor zero é:

0 para tipos numéricos,
false para tipos boleanos, e
"" (string vazia) para strings.
*/

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}