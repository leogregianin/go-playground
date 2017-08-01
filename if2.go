/*
If com uma curta declaração
Como o for, a instrução if pode começar com uma breve declaração antes de executar a condição.

Variáveis ​​declaradas pela instrução são válidas somente no escopo até o final do if.

*/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}