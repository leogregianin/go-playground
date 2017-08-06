/*
Um type switch é uma construção que permite diversas type assertations em série.

Um type switch é como uma instrução switch regular, mas os casos em um type switch especificam os tipos (e não valores), e esses valores são comparados com o tipo dos valores determinados da interface informada.

switch v := i.(type) {
case T:
    // aqui v tem o tipo T
case S:
    // aqui v tem o tipo S
default:
    // sem correspondênte; aqui V tem o mesmo tipo I
}
A declaração em um type switch tem a mesma sintaxe como uma afirmação de tipo de i.(T), mas o tipo específico t é substituído com a palavra-chave type.

Nisso a instrução switch testa se o valor de interface i detém um valor de tipo T ou S. Em cada um dos casos T e S, a variável v será do tipo T ou S respectivamente e guarda o valor mantido por i. No caso padrão (onde não houver correspondência), a variável v é do mesmo tipo e valor da interface i.
*/

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
