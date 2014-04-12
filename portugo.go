package main

import (
	"fmt"
)

type Token struct {
	tipo, id string
}

type Node struct {
	pai     *Node
	filhos  []*Node
	valor   string
	indice  int
	deletar int
	token   Token
}

func main() {
	listaTokens := carregaTokens("texte2")
	parseTree := Node{nil, nil, "P", 0, 0, Token{"", ""}}
	montaParsingTree(&parseTree, listaTokens)
	configuraAST(&parseTree)
	mostraTree(&parseTree)

	fmt.Println("\n\n-----------inicio do programa-------")
	simbolos := make(map[string][]string)
	executaTree(&parseTree, simbolos)
	//fmt.Println(simbolos)
	fmt.Println("-----------fim do programa----------\n\n")
}
