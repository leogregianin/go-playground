/*
O pacote io especifica a interface io.Reader, que representa o fim da leitura de um fluxo de dados.

A biblioteca padrão contém várias implementações destas interfaces, incluindo arquivos, conexões de rede, compressores, cifras e outros.

A interface io.Reader tem um método Read:

func (T) Read(b []byte) (n int, err error)
Read popula um slice de bytes passados com dados e retorna o número de bytes populados e um valor de erro. Este retorna um erro io.EOF quando o fluxo termina.

O exemplo de código cria um strings.Reader. e consome sua saída 8 bytes cada vez.
*/

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 10)
	for {
		n, err := r.Read(b)
		fmt.Printf("%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
