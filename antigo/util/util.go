package util

import (
	core "github.com/maiconio/portugo/core"
)

func MostraTree(tree *core.Node) string {
	for tree.Pai != nil {
		tree = tree.Pai
	}

	//achou o pai
	tmp := mostraRec(tree, "", 0)
	return tmp
}

func mostraRec(tree *core.Node, tab string, index int) string {
	tmp := tab + tree.Valor + "[" + tree.Token.Id + "]\n"
	//fmt.Printf(tab + tree.valor + "[" + tree.token.id + "]["+strconv.Itoa(index)+"]\n")
	tab = tab + "  "
	for i := 0; i < len(tree.Filhos); i++ {
		tmp = tmp + mostraRec(tree.Filhos[i], tab, i)
	}
	return tmp
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
