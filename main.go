package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	flag.Parse()
	arquivo := flag.Arg(0)

	if len(arquivo) > 0 {
		bufferArquivo, err := ioutil.ReadFile(arquivo)

		if err == nil {
			fmt.Printf("%v\n", string(bufferArquivo))
		} else {
			fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
		}
	}
}
