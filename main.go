package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

var exibeTokens = flag.Bool("tokens", false, "Exibe lista dos tokens extraídos do arquivo")
var arquivo string

func init() {
	flag.Parse()
	arquivo = flag.Arg(0)
}

func main() {

	if len(arquivo) > 0 {
		bufferArquivo, err := ioutil.ReadFile(arquivo)

		if err != nil {
			fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
		} else {

			conteudoArquivo := string(bufferArquivo)

			texto := conteudoArquivo
			token := token{tipo: "inicio"}

			if *exibeTokens {
				for token.tipo != "" {
					token, texto = extraiToken(texto)
					fmt.Printf("%v\t\t\t\t%v\n", token.valor, token.tipo)
				}
			}
		}
	} else {
		fmt.Printf("Nenhum arquivo foi especificado. Ex.: \n portugo arquivo.txt \n")
	}
}
