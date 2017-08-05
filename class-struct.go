package main

import (
	"fmt"
	"encoding/json"
)

type Pessoa struct {
    Nome string
    Sobrenome string
    Empresa Empresa
}

type Empresa struct {
    Nome string
    Pais string
}

type Marshaler interface {
    MarshalJSON() ([]byte, error)
}

func (p Pessoa) String() string {
    return "Name: " + p.Nome + " " + p.Sobrenome + " - " +
           "Empresa: " + p.Empresa.Nome + " - " +
           "País: " + p.Empresa.Pais
}

func (p Pessoa) MarshalJSON() ([]byte, error) {
    return []byte(`{"Nome":"` + p.Nome + ` ` + p.Sobrenome +
                  `","Empresa":"` + p.Empresa.Nome + ` ` + `","País":"`    + p.Empresa.Pais + `"}`), nil
}

func main() {
    Brasileiro := Pessoa{
        Nome: "Joao",
        Sobrenome: "da Silva",
        Empresa: Empresa{
            Nome: "Petrobras",
            Pais: "Brasil"}}

    Americano := Pessoa{
        Nome: "John",
        Sobrenome: "McFly",
        Empresa: Empresa{
            Nome: "NY Times",
            Pais: "Estados Unidos"}}

    fmt.Println(Brasileiro)
    j, _ := json.Marshal(Americano)
    fmt.Println(string(j))
}
