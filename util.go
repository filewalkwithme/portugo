package main

import (
	"fmt"
)

func mostraTree(tree *Node) {
	for tree.pai != nil {
		tree = tree.pai
	}

	//achou o pai
	mostraRec(tree, "", 0)

}

func mostraRec(tree *Node, tab string, index int) {
	fmt.Printf(tab + tree.valor + "[" + tree.token.id + "]\n")
	//fmt.Printf(tab + tree.valor + "[" + tree.token.id + "]["+strconv.Itoa(index)+"]\n")
	tab = tab + "  "
	for i := 0; i < len(tree.filhos); i++ {
		mostraRec(tree.filhos[i], tab, i)
	}
}

func push(pilha []string, valor string) []string {
	pilha = append(pilha[:len(pilha)], valor)
	return pilha
}

func pushToken(pilha []Token, valor Token) []Token {
	pilha = append(pilha[:len(pilha)], valor)
	return pilha
}

func pop(pilha []string) ([]string, string) {
	valor := pilha[len(pilha)-1]
	pilha = pilha[:len(pilha)-1]
	return pilha, valor
}

func topo(pilha []string) string {
	return pilha[len(pilha)-1]
}
