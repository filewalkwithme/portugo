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

type NoArvore struct {
	tipo   string
	valor  string
	filhos []NoArvore
}

func main() {

	if len(arquivo) > 0 {
		bufferArquivo, err := ioutil.ReadFile(arquivo)

		if err != nil {
			fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
		} else {

			conteudoArquivo := string(bufferArquivo)

			texto := conteudoArquivo
			tokenTmp := token{tipo: "inicio"}
			tokens := []token{}

			for tokenTmp.tipo != "" {
				tokenTmp, texto = extraiToken(texto)
				tokens = append(tokens, tokenTmp)
				if *exibeTokens {
					fmt.Printf("%v\t\t\t\t%v\n", tokenTmp.valor, tokenTmp.tipo)
				}
			}

			Parse(tokens)
		}
	} else {
		fmt.Printf("Nenhum arquivo foi especificado. Ex.: \n portugo arquivo.txt \n")
	}
}

func Parse(tokens []token) {
	for i, token := range tokens {
		fmt.Printf("[%v]%v\n", i, token)
	}
}
