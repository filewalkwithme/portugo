package core

type Node struct {
	Pai     *Node
	Filhos  []*Node
	Valor   string
	Indice  int
	Deletar int
	Token   Token
}

type Token struct {
	Tipo, Id string
}
