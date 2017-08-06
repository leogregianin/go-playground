/*
Programas Go expressam estados de erro com valores error.

O tipo error é uma interface built-in simliar a fmt.Stringer:

type error interface {
    Error() string
}
(Tal como acontece com fmt.Stringer, o pacote fmt procura a interface de error quando mostra valores.)

Funções frequentemente retornam um valor error e a chamada do código deve lidar com erros testando se o erro é igual a nil.

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("Não conseguiu converter número: %v\n", err)
}
fmt.Println("Inteiro convertido:", i)
Um error nil indica sucesso; um error não-nil denota fracasso.

*/

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
