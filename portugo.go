package main

import (
	"fmt"
	core "github.com/maiconio/portugo/core"
	lex "github.com/maiconio/portugo/lex"
	sintatico "github.com/maiconio/portugo/sintatico"
	util "github.com/maiconio/portugo/util"
	ast "github.com/maiconio/portugo/ast"
	exec "github.com/maiconio/portugo/exec"
	"flag"
)

func main() {
	var arquivo = flag.String("arq", "", "Arquivo a ser processado [obrigatório]")	
	flag.Parse()

	if (*arquivo != "") {
		listaTokens := lex.CarregaTokens(*arquivo)
		parseTree := core.Node{nil, nil, "P", 0, 0, core.Token{"", ""}}
		sintatico.MontaParsingTree(&parseTree, listaTokens)
		ast.ConfiguraAST(&parseTree)
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
