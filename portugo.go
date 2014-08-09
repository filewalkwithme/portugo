package main

import (
	"flag"
	"fmt"
	ast "github.com/maiconio/portugo/ast"
	core "github.com/maiconio/portugo/core"
	exec "github.com/maiconio/portugo/exec"
	lex "github.com/maiconio/portugo/lex"
	sintatico "github.com/maiconio/portugo/sintatico"
	util "github.com/maiconio/portugo/util"
)

func main() {
	var arquivo = flag.String("arq", "", "Arquivo a ser processado [obrigatório]")
	flag.Parse()

	if *arquivo != "" {
		listaTokens := lex.CarregaTokens(*arquivo)
		fmt.Println("TOKENS:\n------>")
		fmt.Println(listaTokens)

		parseTree := core.Node{nil, nil, "P", 0, 0, core.Token{"", ""}}
		sintatico.MontaParsingTree(&parseTree, listaTokens)
		fmt.Println("\n\nÁRVORE SINTÁTICA:\n------>")
		util.MostraTree(&parseTree)

		ast.ConfiguraAST(&parseTree)
		fmt.Println("\n\nAST - ÁRVORE SINTÁTICA ABSTRATA:\n------>")
		util.MostraTree(&parseTree)

		fmt.Println("\n\n-----------inicio do programa-------")
		simbolos := make(map[string][]string)
		exec.ExecutaTree(&parseTree, simbolos)
		//fmt.Println(simbolos)
		fmt.Println("-----------fim do programa----------\n\n")
	} else {
		flag.PrintDefaults()
	}
}
