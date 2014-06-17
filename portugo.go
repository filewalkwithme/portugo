package main

import (
	"fmt"
	lex "github.com/maiconio/portugo/lex"
	sintatico "github.com/maiconio/portugo/sintatico"
	core "github.com/maiconio/portugo/core"
	util "github.com/maiconio/portugo/util"
)

func main() {
	listaTokens := lex.CarregaTokens("texte2")
	parseTree := core.Node{nil, nil, "P", 0, 0, core.Token{"", ""}}
	sintatico.MontaParsingTree(&parseTree, listaTokens)
	configuraAST(&parseTree)
	util.MostraTree(&parseTree)

	fmt.Println("\n\n-----------inicio do programa-------")
	simbolos := make(map[string][]string)
	executaTree(&parseTree, simbolos)
	//fmt.Println(simbolos)
	fmt.Println("-----------fim do programa----------\n\n")
}
