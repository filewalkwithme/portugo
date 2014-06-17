package util

import (
	"fmt"
	core "github.com/maiconio/portugo/core"
)

func MostraTree(tree *core.Node) {
	for tree.Pai != nil {
		tree = tree.Pai
	}

	//achou o pai
	mostraRec(tree, "", 0)

}

func mostraRec(tree *core.Node, tab string, index int) {
	fmt.Printf(tab + tree.Valor + "[" + tree.Token.Id + "]\n")
	//fmt.Printf(tab + tree.valor + "[" + tree.token.id + "]["+strconv.Itoa(index)+"]\n")
	tab = tab + "  "
	for i := 0; i < len(tree.Filhos); i++ {
		mostraRec(tree.Filhos[i], tab, i)
	}
}

func Push(pilha []string, valor string) []string {
	pilha = append(pilha[:len(pilha)], valor)
	return pilha
}

func PushToken(pilha []core.Token, valor core.Token) []core.Token {
	pilha = append(pilha[:len(pilha)], valor)
	return pilha
}

func Pop(pilha []string) ([]string, string) {
	valor := pilha[len(pilha)-1]
	pilha = pilha[:len(pilha)-1]
	return pilha, valor
}

func Topo(pilha []string) string {
	return pilha[len(pilha)-1]
}
