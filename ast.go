package main

import ()

func RemoveNodeFilho(tree *Node, i int) {
	copy(tree.filhos[i:], tree.filhos[i+1:])
	tree.filhos[len(tree.filhos)-1] = nil // or the zero value of T
	tree.filhos = tree.filhos[:len(tree.filhos)-1]
}

func AdicionaNodeFilho(tree *Node, i int, filho *Node) {
	tree.filhos = append(tree.filhos, &Node{})
	copy(tree.filhos[i+1:], tree.filhos[i:])
	tree.filhos[i] = filho
}

func RetiraNodesVazios(tree *Node, res int) int {
	r := res

	for i := 0; i < len(tree.filhos); i++ {
		if tree.filhos[i].valor == "_" {
			RemoveNodeFilho(tree, i)
			r++
		} else {
			if len(tree.filhos[i].filhos) == 0 {
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
					"D3":     true,
				}

				if deletar[tree.filhos[i].valor] {
					RemoveNodeFilho(tree, i)
					r++
				}
			}
		}
	}

	for i := 0; i < len(tree.filhos); i++ {
		r = r + RetiraNodesVazios(tree.filhos[i], 0)
	}

	return r
}

func PromoveNodeSimples(tree *Node, res int) int {
	r := res

	naoPromover := map[string]bool{
		"+-": true,
	}

	if len(tree.filhos) == 1 {
		if naoPromover[tree.valor] == false {
			pai := &Node{}
			pai = tree.pai
			*tree = *tree.filhos[0]
			tree.pai = pai
			r++
		}
	}

	for i := 0; i < len(tree.filhos); i++ {
		r = r + PromoveNodeSimples(tree.filhos[i], 0)
	}

	return r
}

func ComandoOperador(tree *Node, res int) int {
	r := res

	operador := map[string]bool{
		"+-":            true,
		"*/":            true,
		"MOD":           true,
		"DIV":           true,
		"**//":          true,
		"OP.RELACIONAL": true,
		"OP.LOGICO":     true,
		"OP.LOGICO.XOU": true,
	}

	for i := 1; i < len(tree.filhos); i++ {
		if len(tree.filhos[i].filhos) == 2 && tree.filhos[i-1].valor != "<-" && tree.filhos[i].valor != "M9" {

			if operador[tree.filhos[i].filhos[0].valor] {
				AdicionaNodeFilho(tree, i+1, &Node{nil, nil, tree.filhos[i].filhos[0].valor, -1, 0, tree.filhos[i].filhos[0].token})
				filho1 := &Node{}
				*filho1 = *tree.filhos[i-1]
				AdicionaNodeFilho(tree.filhos[i+1], 0, filho1)

				filho2 := &Node{}
				*filho2 = *tree.filhos[i].filhos[1]
				AdicionaNodeFilho(tree.filhos[i+1], 1, filho2)

				RemoveNodeFilho(tree, i-1)
				RemoveNodeFilho(tree, i-1)
				r++
			}
		}
	}

	for i := 0; i < len(tree.filhos); i++ {
		r = r + ComandoOperador(tree.filhos[i], 0)
	}

	return r
}

func ComandoAtribuicao(tree *Node, res int) int {
	r := res

	if len(tree.filhos) == 3 {
		if tree.filhos[1].valor == "<-" {
			AdicionaNodeFilho(tree, 3, &Node{nil, nil, tree.filhos[1].valor, -1, 0, tree.filhos[1].token})

			filho1 := &Node{}
			*filho1 = *tree.filhos[0]
			AdicionaNodeFilho(tree.filhos[3], 0, filho1)

			filho2 := &Node{}
			*filho2 = *tree.filhos[2]
			AdicionaNodeFilho(tree.filhos[3], 1, filho2)

			RemoveNodeFilho(tree, 0)
			RemoveNodeFilho(tree, 0)
			RemoveNodeFilho(tree, 0)
			r++
		}
	}

	for i := 0; i < len(tree.filhos); i++ {
		r = r + ComandoAtribuicao(tree.filhos[i], 0)
	}

	return r
}

func ComandoEscreva(tree *Node, res int) int {
	r := res

	n := len(tree.filhos)
	if n >= 2 {
		if tree.filhos[0].valor == "ESCREVA" {
			AdicionaNodeFilho(tree, n, &Node{nil, nil, tree.filhos[0].valor, -1, 0, tree.filhos[0].token})
			for i := 1; i < n; i++ {
				filho1 := &Node{}
				*filho1 = *tree.filhos[i]
				AdicionaNodeFilho(tree.filhos[n], len(tree.filhos[n].filhos), filho1)
				r++
			}

			for i := 0; i < n; i++ {
				RemoveNodeFilho(tree, 0)
			}
		}
	}

	n = len(tree.filhos)
	if n >= 2 {
		for i := 0; i < n; i++ {
			if tree.filhos[i].valor == "ESCV2" {
				for j := 0; j < len(tree.filhos[i].filhos); j++ {
					filho1 := &Node{}
					*filho1 = *tree.filhos[i].filhos[j]
					AdicionaNodeFilho(tree, len(tree.filhos), filho1)
					r++
				}

				RemoveNodeFilho(tree, i)
			}
		}
	}

	for i := 0; i < len(tree.filhos); i++ {
		r = r + ComandoEscreva(tree.filhos[i], 0)
	}

	return r
}

func ComandoLeia(tree *Node, res int) int {
	r := res

	n := len(tree.filhos)
	if n >= 2 {
		if tree.filhos[0].valor == "LEIA" {
			AdicionaNodeFilho(tree, n, &Node{nil, nil, tree.filhos[0].valor, -1, 0, tree.filhos[0].token})
			for i := 1; i < n; i++ {
				filho1 := &Node{}
				*filho1 = *tree.filhos[i]
				AdicionaNodeFilho(tree.filhos[n], len(tree.filhos[n].filhos), filho1)
				r++
			}

			for i := 0; i < n; i++ {
				RemoveNodeFilho(tree, 0)
			}
		}
	}

	n = len(tree.filhos)
	if n >= 2 {
		for i := 0; i < n; i++ {
			if tree.filhos[i].valor == "LEIA2" {
				for j := 0; j < len(tree.filhos[i].filhos); j++ {
					filho1 := &Node{}
					*filho1 = *tree.filhos[i].filhos[j]
					AdicionaNodeFilho(tree, len(tree.filhos), filho1)
					r++
				}

				RemoveNodeFilho(tree, i)
			}
		}
	}

	for i := 0; i < len(tree.filhos); i++ {
		r = r + ComandoLeia(tree.filhos[i], 0)
	}

	return r
}

func ComandoDeclaraVariavel(tree *Node, res int) int {
	r := res

	n := len(tree.filhos)
	if n >= 2 {
		if tree.filhos[0].valor == "TIPOVAR" {
			AdicionaNodeFilho(tree, n, &Node{nil, nil, tree.filhos[0].valor, -1, 0, tree.filhos[0].token})
			for i := 1; i < n; i++ {
				filho1 := &Node{}
				*filho1 = *tree.filhos[i]
				AdicionaNodeFilho(tree.filhos[n], len(tree.filhos[n].filhos), filho1)
			}

			for i := 0; i < n; i++ {
				RemoveNodeFilho(tree, 0)
			}
		}
	}

	for i := 0; i < len(tree.filhos); i++ {
		if tree.filhos[i].valor == "D2" {
			for j := 0; j < len(tree.filhos[i].filhos); j++ {
				filho1 := &Node{}
				*filho1 = *tree.filhos[i].filhos[j]
				AdicionaNodeFilho(tree, len(tree.filhos), filho1)
			}
			RemoveNodeFilho(tree, i)
		}
	}

	for i := 0; i < len(tree.filhos); i++ {
		r = r + ComandoDeclaraVariavel(tree.filhos[i], 0)
	}

	return r
}

func PromoveAcoes(tree *Node, res int) int {
	r := res

	for i := 0; i < len(tree.filhos); i++ {
		r = r + PromoveAcoes(tree.filhos[i], 0)
	}

	for i := 0; i < len(tree.filhos); i++ {
		if tree.filhos[i].valor == "AC1" || tree.filhos[i].valor == "V1" {

			for j := 0; j < len(tree.filhos[i].filhos); j++ {
				filho1 := &Node{}
				*filho1 = *tree.filhos[i].filhos[j]
				AdicionaNodeFilho(tree, len(tree.filhos), filho1)
				r++
			}
			RemoveNodeFilho(tree, i)
		}
	}

	return r
}

func configuraAST(tree *Node) {
	for tree.pai != nil {
		tree = tree.pai
	}

	for RetiraNodesVazios(tree, 0) > 0 {
	}

	for PromoveNodeSimples(tree, 0) > 0 {
	}

	for ComandoOperador(tree, 0) > 0 {
	}

	for PromoveNodeSimples(tree, 0) > 0 {
	}

	for ComandoAtribuicao(tree, 0) > 0 {
	}

	for ComandoEscreva(tree, 0) > 0 {
	}

	for ComandoLeia(tree, 0) > 0 {
	}

	for ComandoDeclaraVariavel(tree, 0) > 0 {
	}

	for PromoveAcoes(tree, 0) > 0 {
	}
}
