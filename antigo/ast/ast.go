//Pacote responsável por converter a Árvore Sintática em uma Árvore Sintática Abstrata (AST, abstract syntax tree).
//Nesta etapa são removidas todas as informações não essenciais e também são realizadas
//manipulações nos filhos da Árvore Sintática com o objetivo de viabilizar a execução do programa
package ast

import (
	core "github.com/maiconio/portugo/core"
)

func RemoveNodeFilho(tree *core.Node, i int) {
	copy(tree.Filhos[i:], tree.Filhos[i+1:])
	tree.Filhos[len(tree.Filhos)-1] = nil // or the zero value of T
	tree.Filhos = tree.Filhos[:len(tree.Filhos)-1]
}

func AdicionaNodeFilho(tree *core.Node, i int, filho *core.Node) {
	tree.Filhos = append(tree.Filhos, &core.Node{})
	copy(tree.Filhos[i+1:], tree.Filhos[i:])
	tree.Filhos[i] = filho
}

func RetiraNodesVazios(tree *core.Node, res int) int {
	r := res

	for i := 0; i < len(tree.Filhos); i++ {
		if tree.Filhos[i].Valor == "_" {
			RemoveNodeFilho(tree, i)
			r++
		} else {
			if len(tree.Filhos[i].Filhos) == 0 {
				deletar := map[string]bool{
					"INICIO": true,
					"FIM":    true,
					"PONTO":  true,
					",":      true,
					";":      true,
					":":      true,
					"(":      true,
					")":      true,
					"ESCV2":  true,
					"AC1":    true,
					"V1":     true,
					"R2":     true,
					"M10":    true,
					"M8":     true,
					"M6":     true,
					"M4":     true,
					"M2":     true,
					"L2":     true,
					"L4":     true,
					"L6":     true,
					"D3":     true,
				}

				if deletar[tree.Filhos[i].Valor] {
					RemoveNodeFilho(tree, i)
					r++
				}
			}
		}
	}

	for i := 0; i < len(tree.Filhos); i++ {
		r = r + RetiraNodesVazios(tree.Filhos[i], 0)
	}

	return r
}

func PromoveNodeSimples(tree *core.Node, res int) int {
	r := res

	naoPromover := map[string]bool{
		"+-":      true,
		"FUNCMAT": true,
	}

	if len(tree.Filhos) == 1 {
		if naoPromover[tree.Valor] == false {
			pai := &core.Node{}
			pai = tree.Pai
			*tree = *tree.Filhos[0]
			tree.Pai = pai
			r++
		}
	}

	for i := 0; i < len(tree.Filhos); i++ {
		r = r + PromoveNodeSimples(tree.Filhos[i], 0)
	}

	return r
}

func ComandoOperador(tree *core.Node, res int) int {
	r := res

	operador := map[string]bool{
		"+-":            true,
		"*/":            true,
		"MOD":           true,
		"DIV":           true,
		"**//":          true,
		"OP.RELACIONAL": true,
		"OP.LOGICO.E":   true,
		"OP.LOGICO.OU":  true,
		"OP.LOGICO.XOU": true,
	}

	for i := 1; i < len(tree.Filhos); i++ {
		if len(tree.Filhos[i].Filhos) == 2 && tree.Filhos[i-1].Valor != "<-" && tree.Filhos[i].Valor != "M9" {

			if operador[tree.Filhos[i].Filhos[0].Valor] {
				AdicionaNodeFilho(tree, i+1, &core.Node{nil, nil, tree.Filhos[i].Filhos[0].Valor, -1, 0, tree.Filhos[i].Filhos[0].Token})
				filho1 := &core.Node{}
				*filho1 = *tree.Filhos[i-1]
				AdicionaNodeFilho(tree.Filhos[i+1], 0, filho1)

				filho2 := &core.Node{}
				*filho2 = *tree.Filhos[i].Filhos[1]
				AdicionaNodeFilho(tree.Filhos[i+1], 1, filho2)

				RemoveNodeFilho(tree, i-1)
				RemoveNodeFilho(tree, i-1)
				r++
			}
		}
	}

	for i := 0; i < len(tree.Filhos); i++ {
		r = r + ComandoOperador(tree.Filhos[i], 0)
	}

	return r
}

func ComandoAtribuicao(tree *core.Node, res int) int {
	r := res

	if len(tree.Filhos) == 3 {
		if tree.Filhos[1].Valor == "<-" {
			AdicionaNodeFilho(tree, 3, &core.Node{nil, nil, tree.Filhos[1].Valor, -1, 0, tree.Filhos[1].Token})

			filho1 := &core.Node{}
			*filho1 = *tree.Filhos[0]
			AdicionaNodeFilho(tree.Filhos[3], 0, filho1)

			filho2 := &core.Node{}
			*filho2 = *tree.Filhos[2]
			AdicionaNodeFilho(tree.Filhos[3], 1, filho2)

			RemoveNodeFilho(tree, 0)
			RemoveNodeFilho(tree, 0)
			RemoveNodeFilho(tree, 0)
			r++
		}
	}

	for i := 0; i < len(tree.Filhos); i++ {
		r = r + ComandoAtribuicao(tree.Filhos[i], 0)
	}

	return r
}

func ComandoEscreva(tree *core.Node, res int) int {
	r := res

	n := len(tree.Filhos)
	if n >= 2 {
		if tree.Filhos[0].Valor == "ESCREVA" {
			AdicionaNodeFilho(tree, n, &core.Node{nil, nil, tree.Filhos[0].Valor, -1, 0, tree.Filhos[0].Token})
			for i := 1; i < n; i++ {
				filho1 := &core.Node{}
				*filho1 = *tree.Filhos[i]
				AdicionaNodeFilho(tree.Filhos[n], len(tree.Filhos[n].Filhos), filho1)
				r++
			}

			for i := 0; i < n; i++ {
				RemoveNodeFilho(tree, 0)
			}
		}
	}

	n = len(tree.Filhos)
	if n >= 2 {
		for i := 0; i < n; i++ {
			if tree.Filhos[i].Valor == "ESCV2" {
				for j := 0; j < len(tree.Filhos[i].Filhos); j++ {
					filho1 := &core.Node{}
					*filho1 = *tree.Filhos[i].Filhos[j]
					AdicionaNodeFilho(tree, len(tree.Filhos), filho1)
					r++
				}

				RemoveNodeFilho(tree, i)
			}
		}
	}

	for i := 0; i < len(tree.Filhos); i++ {
		r = r + ComandoEscreva(tree.Filhos[i], 0)
	}

	return r
}

func ComandoLeia(tree *core.Node, res int) int {
	r := res

	n := len(tree.Filhos)
	if n >= 2 {
		if tree.Filhos[0].Valor == "LEIA" {
			AdicionaNodeFilho(tree, n, &core.Node{nil, nil, tree.Filhos[0].Valor, -1, 0, tree.Filhos[0].Token})
			for i := 1; i < n; i++ {
				filho1 := &core.Node{}
				*filho1 = *tree.Filhos[i]
				AdicionaNodeFilho(tree.Filhos[n], len(tree.Filhos[n].Filhos), filho1)
				r++
			}

			for i := 0; i < n; i++ {
				RemoveNodeFilho(tree, 0)
			}
		}
	}

	n = len(tree.Filhos)
	if n >= 2 {
		for i := 0; i < n; i++ {
			if tree.Filhos[i].Valor == "LEIA2" {
				for j := 0; j < len(tree.Filhos[i].Filhos); j++ {
					filho1 := &core.Node{}
					*filho1 = *tree.Filhos[i].Filhos[j]
					AdicionaNodeFilho(tree, len(tree.Filhos), filho1)
					r++
				}

				RemoveNodeFilho(tree, i)
			}
		}
	}

	for i := 0; i < len(tree.Filhos); i++ {
		r = r + ComandoLeia(tree.Filhos[i], 0)
	}

	return r
}

func ComandoFuncaoMatematica(tree *core.Node, res int) int {
	r := res

	n := len(tree.Filhos)
	if n >= 2 {
		if tree.Filhos[0].Valor == "FUNCMAT" {
			AdicionaNodeFilho(tree, n, &core.Node{nil, nil, tree.Filhos[0].Valor, -1, 0, tree.Filhos[0].Token})
			for i := 1; i < n; i++ {
				filho1 := &core.Node{}
				*filho1 = *tree.Filhos[i]
				AdicionaNodeFilho(tree.Filhos[n], len(tree.Filhos[n].Filhos), filho1)
				r++
			}

			for i := 0; i < n; i++ {
				RemoveNodeFilho(tree, 0)
			}
		}
	}

	for i := 0; i < len(tree.Filhos); i++ {
		r = r + ComandoFuncaoMatematica(tree.Filhos[i], 0)
	}

	return r
}

func ComandoDeclaraVariavel(tree *core.Node, res int) int {
	r := res

	n := len(tree.Filhos)
	if n >= 2 {
		if tree.Filhos[0].Valor == "TIPOVAR" {
			AdicionaNodeFilho(tree, n, &core.Node{nil, nil, tree.Filhos[0].Valor, -1, 0, tree.Filhos[0].Token})
			for i := 1; i < n; i++ {
				filho1 := &core.Node{}
				*filho1 = *tree.Filhos[i]
				AdicionaNodeFilho(tree.Filhos[n], len(tree.Filhos[n].Filhos), filho1)
			}

			for i := 0; i < n; i++ {
				RemoveNodeFilho(tree, 0)
			}
		}
	}

	for i := 0; i < len(tree.Filhos); i++ {
		if tree.Filhos[i].Valor == "D2" {
			for j := 0; j < len(tree.Filhos[i].Filhos); j++ {
				filho1 := &core.Node{}
				*filho1 = *tree.Filhos[i].Filhos[j]
				AdicionaNodeFilho(tree, len(tree.Filhos), filho1)
			}
			RemoveNodeFilho(tree, i)
		}
	}

	for i := 0; i < len(tree.Filhos); i++ {
		r = r + ComandoDeclaraVariavel(tree.Filhos[i], 0)
	}

	return r
}

func PromoveAcoes(tree *core.Node, res int) int {
	r := res

	for i := 0; i < len(tree.Filhos); i++ {
		r = r + PromoveAcoes(tree.Filhos[i], 0)
	}

	for i := 0; i < len(tree.Filhos); i++ {
		if tree.Filhos[i].Valor == "AC1" || tree.Filhos[i].Valor == "V1" {

			for j := len(tree.Filhos[i].Filhos) - 1; j >= 0; j-- {
				filho1 := &core.Node{}
				*filho1 = *tree.Filhos[i].Filhos[j]
				AdicionaNodeFilho(tree, i+1, filho1)
				r++
			}
			RemoveNodeFilho(tree, i)
		}
	}

	return r
}

func ConfiguraAST(tree *core.Node) {
	for tree.Pai != nil {
		tree = tree.Pai
	}

	for RetiraNodesVazios(tree, 0) > 0 {
	}

	for PromoveNodeSimples(tree, 0) > 0 {
	}
	for ComandoOperador(tree, 0) > 0 {
	}

	ComandoFuncaoMatematica(tree, 0)

	for PromoveNodeSimples(tree, 0) > 0 {
	}

	for ComandoAtribuicao(tree, 0) > 0 {
	}

	for ComandoEscreva(tree, 0) > 0 {
	}

	for ComandoLeia(tree, 0) > 0 {
	}

	for ComandoFuncaoMatematica(tree, 0) > 0 {
	}

	for ComandoDeclaraVariavel(tree, 0) > 0 {
	}

	for PromoveAcoes(tree, 0) > 0 {
	}
	/**/
}
