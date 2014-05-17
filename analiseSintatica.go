package main

import (
	"fmt"
)

func montaParsingTree(tree *Node, listaTokens []Token) {
	producao := make(map[string][]string)
	tab := make(map[string]map[string]string)

	producao["LOG01"] = []string{"L", "L1", "L2"}
	producao["LOG02"] = []string{"L1", "L3", "L4"}
	producao["LOG03"] = []string{"L2", "OP.LOGICO.XOU", "L1", "L2"}
	producao["LOG04"] = []string{"L2", "_"}
	producao["LOG05"] = []string{"L3", "L5", "L6"}
	producao["LOG06"] = []string{"L4", "OP.LOGICO.E", "L3", "L4"}
	producao["LOG07"] = []string{"L4", "_"}

	producao["LOG08"] = []string{"L5", "OP.LOGICO.UN", "L5"}
	producao["LOG09"] = []string{"L5", "R1", "R2"}
	producao["LOG10"] = []string{"L5", "LOGICO"}
	producao["LOG11"] = []string{"L5", "v"}
	producao["LOG12"] = []string{"L5", "(", "L", ")"}
	producao["LOG13"] = []string{"L6", "OP.LOGICO.OU", "L5", "L6"}
	producao["LOG14"] = []string{"L6", "_"}

	producao["REL01"] = []string{"R1", "M1", "M2"}
	producao["REL02"] = []string{"R2", "OP.RELACIONAL", "R1"}
	producao["REL03"] = []string{"R2", "_"}
	producao["REL04"] = []string{"R1", "(", "L", ")"}

	producao["MAT01"] = []string{"M1", "M3", "M4"}
	producao["MAT02"] = []string{"M2", "+-", "M1", "M2"}
	producao["MAT03"] = []string{"M2", "_"}

	producao["MAT04"] = []string{"M3", "M5", "M6"}
	producao["MAT05"] = []string{"M4", "DIV", "M3", "M4"}
	producao["MAT06"] = []string{"M4", "_"}

	producao["MAT07"] = []string{"M5", "M7", "M8"}
	producao["MAT08"] = []string{"M6", "MOD", "M5", "M6"}
	producao["MAT09"] = []string{"M6", "_"}

	producao["MAT10"] = []string{"M7", "M9", "M10"}
	producao["MAT11"] = []string{"M8", "*/", "M7", "M8"}
	producao["MAT12"] = []string{"M8", "_"}

	producao["MAT13"] = []string{"M9", "INTEIRO"}
	producao["MAT14"] = []string{"M9", "REAL"}
	producao["MAT15"] = []string{"M9", "v"}
	producao["MAT16"] = []string{"M9", "(", "L", ")"}
	producao["MAT17"] = []string{"M9", "+-", "M1"}

	producao["MAT18"] = []string{"M10", "**//", "M9", "M10"}
	producao["MAT19"] = []string{"M10", "_"}

	producao["MAT20"] = []string{"FUNCMAT", "FUNCMAT", "(", "M1", "M2", ")"}
	producao["MAT21"] = []string{"M9", "+-", "M1"}

	tab["L"] = make(map[string]string)
	tab["L"]["("] = "LOG01"
	tab["L"]["v"] = "LOG01"
	tab["L"]["INTEIRO"] = "LOG01"
	tab["L"]["+-"] = "LOG01"
	tab["L"]["FUNCMAT"] = "LOG01"
	tab["L"]["LOGICO"] = "LOG01"
	tab["L"]["OP.LOGICO.E"] = "LOG01"
	tab["L"]["OP.LOGICO.UN"] = "LOG01"

	tab["L1"] = make(map[string]string)
	tab["L1"]["("] = "LOG02"
	tab["L1"]["v"] = "LOG02"
	tab["L1"]["INTEIRO"] = "LOG02"
	tab["L1"]["+-"] = "LOG02"
	tab["L1"]["LOGICO"] = "LOG02"
	tab["L1"]["FUNCMAT"] = "LOG02"
	tab["L1"]["OP.LOGICO.E"] = "LOG02"
	tab["L1"]["OP.LOGICO.UN"] = "LOG02"

	tab["L2"] = make(map[string]string)
	tab["L2"]["OP.LOGICO.XOU"] = "LOG03"
	tab["L2"][")"] = "LOG04"
	tab["L2"][","] = "LOG04"
	tab["L2"][";"] = "LOG04"

	tab["L3"] = make(map[string]string)
	tab["L3"]["OP.LOGICO.UN"] = "LOG05"
	tab["L3"]["v"] = "LOG05"
	tab["L3"]["INTEIRO"] = "LOG05"
	tab["L3"]["+-"] = "LOG05"
	tab["L3"]["FUNCMAT"] = "LOG05"
	tab["L3"]["("] = "LOG05"
	tab["L3"]["LOGICO"] = "LOG05"

	tab["L4"] = make(map[string]string)
	tab["L4"]["OP.LOGICO.E"] = "LOG06"
	tab["L4"][")"] = "LOG07"
	tab["L4"][","] = "LOG07"
	tab["L4"][";"] = "LOG07"
	tab["L4"]["OP.LOGICO.XOU"] = "LOG07"

	tab["L5"] = make(map[string]string)
	tab["L5"]["OP.LOGICO.UN"] = "LOG08"
	tab["L5"]["INTEIRO"] = "LOG09"
	tab["L5"]["+-"] = "LOG09"
	tab["L5"]["FUNCMAT"] = "LOG09"
	tab["L5"]["LOGICO"] = "LOG10"
	tab["L5"]["v"] = "LOG11"
	tab["L5"]["("] = "LOG12"

	tab["L6"] = make(map[string]string)
	tab["L6"]["OP.LOGICO.OU"] = "LOG13"
	tab["L6"][")"] = "LOG14"
	tab["L6"][","] = "LOG14"
	tab["L6"][";"] = "LOG14"
	tab["L6"]["OP.LOGICO.E"] = "LOG14"
	tab["L6"]["OP.LOGICO.XOU"] = "LOG14"

	tab["R1"] = make(map[string]string)
	tab["R1"]["v"] = "REL01"
	tab["R1"]["INTEIRO"] = "REL01"
	tab["R1"]["+-"] = "REL01"
	tab["R1"]["FUNCMAT"] = "REL01"
	tab["R1"]["("] = "REL01"

	tab["R2"] = make(map[string]string)
	tab["R2"]["OP.RELACIONAL"] = "REL02"
	tab["R2"]["OP.LOGICO.E"] = "REL03"
	tab["R2"][")"] = "REL03"
	tab["R2"][","] = "REL03"
	tab["R2"][";"] = "REL03"

	tab["M1"] = make(map[string]string)
	tab["M1"]["v"] = "MAT01"
	tab["M1"]["("] = "MAT01"
	tab["M1"]["INTEIRO"] = "MAT01"
	tab["M1"]["FUNCMAT"] = "MAT01"
	tab["M1"]["+-"] = "MAT01"

	tab["M2"] = make(map[string]string)
	tab["M2"]["OP.LOGICO.XOU"] = "MAT03"
	tab["M2"]["OP.LOGICO.E"] = "MAT03"
	tab["M2"]["OP.LOGICO.OU"] = "MAT03"
	tab["M2"]["OP.RELACIONAL"] = "MAT03"
	tab["M2"]["+-"] = "MAT02"
	tab["M2"][")"] = "MAT03"
	tab["M2"][","] = "MAT03"
	tab["M2"][";"] = "MAT03"

	tab["M3"] = make(map[string]string)
	tab["M3"]["v"] = "MAT04"
	tab["M3"]["("] = "MAT04"
	tab["M3"]["INTEIRO"] = "MAT04"
	tab["M3"]["FUNCMAT"] = "MAT04"
	tab["M3"]["+-"] = "MAT04"

	tab["M4"] = make(map[string]string)
	tab["M4"]["DIV"] = "MAT05"
	tab["M4"]["OP.LOGICO.XOU"] = "MAT06"
	tab["M4"]["OP.LOGICO.E"] = "MAT06"
	tab["M4"]["OP.LOGICO.OU"] = "MAT06"
	tab["M4"]["OP.RELACIONAL"] = "MAT06"
	tab["M4"]["+-"] = "MAT06"
	tab["M4"][")"] = "MAT06"
	tab["M4"][","] = "MAT06"
	tab["M4"][";"] = "MAT06"

	tab["M5"] = make(map[string]string)
	tab["M5"]["v"] = "MAT07"
	tab["M5"]["("] = "MAT07"
	tab["M5"]["INTEIRO"] = "MAT07"
	tab["M5"]["FUNCMAT"] = "MAT07"
	tab["M5"]["+-"] = "MAT07"

	tab["M6"] = make(map[string]string)
	tab["M6"]["MOD"] = "MAT08"
	tab["M6"]["OP.LOGICO.XOU"] = "MAT09"
	tab["M6"]["OP.LOGICO.E"] = "MAT09"
	tab["M6"]["OP.LOGICO.OU"] = "MAT09"
	tab["M6"]["OP.RELACIONAL"] = "MAT09"
	tab["M6"]["DIV"] = "MAT09"
	tab["M6"]["+-"] = "MAT09"
	tab["M6"][")"] = "MAT09"
	tab["M6"][","] = "MAT09"
	tab["M6"][";"] = "MAT09"

	tab["M7"] = make(map[string]string)
	tab["M7"]["v"] = "MAT10"
	tab["M7"]["("] = "MAT10"
	tab["M7"]["INTEIRO"] = "MAT10"
	tab["M7"]["FUNCMAT"] = "MAT10"
	tab["M7"]["+-"] = "MAT10"

	tab["M8"] = make(map[string]string)
	tab["M8"]["*/"] = "MAT11"
	tab["M8"]["OP.LOGICO.XOU"] = "MAT12"
	tab["M8"]["OP.LOGICO.E"] = "MAT12"
	tab["M8"]["OP.LOGICO.OU"] = "MAT12"
	tab["M8"]["OP.RELACIONAL"] = "MAT12"
	tab["M8"]["DIV"] = "MAT12"
	tab["M8"]["MOD"] = "MAT12"
	tab["M8"]["+-"] = "MAT12"
	tab["M8"][")"] = "MAT12"
	tab["M8"][","] = "MAT12"
	tab["M8"][";"] = "MAT12"

	tab["M9"] = make(map[string]string)
	tab["M9"]["INTEIRO"] = "MAT13"
	tab["M9"]["FUNCMAT"] = "MAT20"
	tab["M9"]["v"] = "MAT15"
	tab["M9"]["("] = "MAT16"
	tab["M9"]["+-"] = "MAT21"

	tab["M10"] = make(map[string]string)
	tab["M10"]["**//"] = "MAT18"
	tab["M10"]["OP.RELACIONAL"] = "MAT19"
	tab["M10"]["OP.LOGICO.E"] = "MAT19"
	tab["M10"]["OP.LOGICO.OU"] = "MAT19"
	tab["M10"]["OP.LOGICO.XOU"] = "MAT19"
	tab["M10"]["DIV"] = "MAT19"
	tab["M10"]["MOD"] = "MAT19"
	tab["M10"]["*/"] = "MAT19"
	tab["M10"]["+-"] = "MAT19"
	tab["M10"][")"] = "MAT19"
	tab["M10"][","] = "MAT19"
	tab["M10"][";"] = "MAT19"

	producao["PROG01"] = []string{"P", "INICIO", "V", "AC", "FIM", "PONTO"}
	tab["P"] = make(map[string]string)
	tab["P"]["v"] = "PROG01"
	tab["P"]["TIPOVAR"] = "PROG01"
	tab["P"]["LEIA"] = "PROG01"
	tab["P"]["ESCREVA"] = "PROG01"
	tab["P"]["INICIO"] = "PROG01"

	producao["ACOE01"] = []string{"AC", "AC1"}
	producao["ACOE02"] = []string{"AC1", "A", ";", "AC1"}
	producao["ACOE03"] = []string{"AC1", "_"}

	tab["AC"] = make(map[string]string)
	tab["AC"]["v"] = "ACOE01"
	tab["AC"]["LEIA"] = "ACOE01"
	tab["AC"]["ESCREVA"] = "ACOE01"
	tab["AC"][";"] = "ACOE01"
	tab["AC"]["FIM"] = "ACOE01"

	tab["AC1"] = make(map[string]string)
	tab["AC1"]["v"] = "ACOE02"
	tab["AC1"]["LEIA"] = "ACOE02"
	tab["AC1"]["ESCREVA"] = "ACOE02"
	tab["AC1"][";"] = "ACOE03"
	tab["AC1"]["FIM"] = "ACOE03"

	producao["ATRV01"] = []string{"A", "v", "<-", "L"}

	producao["LEIA01"] = []string{"A", "LEIA", "(", "LEIA1", "LEIA2", ")"}
	producao["LEIA02"] = []string{"LEIA1", "v"}
	producao["LEIA03"] = []string{"LEIA2", ",", "LEIA1", "LEIA2"}
	producao["LEIA04"] = []string{"LEIA2", "_"}

	producao["ESCV01"] = []string{"A", "ESCREVA", "(", "ESCV1", "ESCV2", ")"}
	producao["ESCV02"] = []string{"ESCV1", "L"}
	producao["ESCV03"] = []string{"ESCV2", ",", "ESCV1", "ESCV2"}
	producao["ESCV04"] = []string{"ESCV2", "_"}

	tab["A"] = make(map[string]string)
	tab["A"]["v"] = "ATRV01"
	tab["A"]["LEIA"] = "LEIA01"
	tab["A"]["ESCREVA"] = "ESCV01"

	tab["LEIA1"] = make(map[string]string)
	tab["LEIA1"]["v"] = "LEIA02"

	tab["LEIA2"] = make(map[string]string)
	tab["LEIA2"][","] = "LEIA03"
	tab["LEIA2"][")"] = "LEIA04"

	tab["ESCV1"] = make(map[string]string)
	tab["ESCV1"]["INTEIRO"] = "ESCV02"
	tab["ESCV1"]["LOGICO"] = "ESCV02"
	tab["ESCV1"]["v"] = "ESCV02"
	tab["ESCV1"]["("] = "ESCV02"
	tab["ESCV1"]["OP.LOGICO.UN"] = "ESCV02"
	tab["ESCV1"]["FUNCMAT"] = "ESCV02"

	tab["ESCV2"] = make(map[string]string)
	tab["ESCV2"][","] = "ESCV03"
	tab["ESCV2"][")"] = "ESCV04"

	producao["VARS01"] = []string{"V", "V1"}
	producao["VARS02"] = []string{"V1", "D", ";", "V1"}
	producao["VARS03"] = []string{"V1", "_"}

	tab["V"] = make(map[string]string)
	tab["V"]["TIPOVAR"] = "VARS01"
	tab["V"]["v"] = "VARS01"
	tab["V"]["LEIA"] = "VARS01"
	tab["V"]["ESCREVA"] = "VARS01"

	tab["V1"] = make(map[string]string)
	tab["V1"]["TIPOVAR"] = "VARS02"
	tab["V1"][";"] = "VARS03"
	tab["V1"]["FIM"] = "VARS03"
	tab["V1"]["v"] = "VARS03"
	tab["V1"]["LEIA"] = "VARS03"
	tab["V1"]["ESCREVA"] = "VARS03"

	producao["DECV01"] = []string{"D", "D1", "D2"}
	producao["DECV02"] = []string{"D1", "TIPOVAR", ":"}
	producao["DECV03"] = []string{"D2", "v", "D3"}
	producao["DECV04"] = []string{"D3", ",", "D2"}
	producao["DECV05"] = []string{"D3", "_"}

	tab["D"] = make(map[string]string)
	tab["D"]["TIPOVAR"] = "DECV01"

	tab["D1"] = make(map[string]string)
	tab["D1"]["TIPOVAR"] = "DECV02"

	tab["D2"] = make(map[string]string)
	tab["D2"]["v"] = "DECV03"

	tab["D3"] = make(map[string]string)
	tab["D3"][","] = "DECV04"
	tab["D3"][";"] = "DECV05"

	//posicao inicial na lista de tokens
	posToken := 0

	//inicializa a pilha para reconhecimento de sentenca
	pilha := []string{}
	//pilha = push(pilha, ";")
	pilha = push(pilha, "P")

	acaba := false
	i := 0
	z := 0
	for acaba == false {
		a := topo(pilha)
		b := listaTokens[posToken]

		//fmt.Println(i, pilha)
		//fmt.Println(i, listaTokens)
		//fmt.Println("")
		i++

		if topo(pilha) == "_" {
			pilha, _ = pop(pilha)
		} else {
			if topo(pilha) == listaTokens[0].tipo {
				//fmt.Println("----",listaTokens[0].tipo, listaTokens[0].id)
				pilha, _ = pop(pilha)
				listaTokens = listaTokens[1:]
			} else {
				p := tab[a][b.tipo]

				if len(p) > 0 {
					pilha, _ = pop(pilha)

					tmp := producao[p]
					tmpToken := []Token{}

					for w := len(tmp) - 1; w >= 1; w-- {
						pilha = push(pilha, tmp[w])
					}

					for w := 1; w < len(tmp); w++ {
						executa := len(listaTokens) > 0
						token := Token{tmp[w], tmp[w]}
						for y := 0; executa && y < len(listaTokens); y++ {
							if listaTokens[y].tipo == tmp[w] {
								token = Token{listaTokens[y].tipo, listaTokens[y].id}
								executa = false
							}
						}
						tmpToken = pushToken(tmpToken, token)
						//fmt.Println("....",tmp[w],tmpToken)

					}

					//fmt.Println("adiciona:", tmp[1:])
					adicionarItem(tree, tmpToken, z)
					z = z + len(tmp[1:])
					//mostraTree(tree)
				} else {
					acaba = true
					fmt.Println(a, b.tipo, "ERRO!")
					tree = &Node{nil, nil, "E", 0, 0, Token{"", ""}}
				}
			}
		}
		if len(listaTokens) == 0 {
			acaba = true
		}
		//fmt.Println(listaTokens)
	}

	//mostraTree(tree)
	if len(pilha) == 0 {
		//	mostraTree(tree)
	} else {
		fmt.Println(pilha)
		fmt.Println("ERRO SINTATICO")
	}
}

func adicionarItem(tree *Node, valores []Token, z int) {
	base := maiorPositivo(tree, -1)
	//fmt.Println(valores);

	//fmt.Println(base);
	adicionar(tree, valores, base, z)
}

//elemento mais a esquerda, sem expansao
func adicionar(tree *Node, valores []Token, base int, z int) {
	ast := tree
	if ast.indice == base {
		//marca indice como -1
		ast.indice = -1

		//adiciona
		if len(valores) > 0 {
			j := z + len(valores)
			for i := 0; i < len(valores); i++ {

				filho := Node{tree, nil, valores[i].tipo, -1, 0, Token{valores[i].tipo, valores[i].id}}
				if valores[i].tipo != "TIPOVAR" &&
					valores[i].tipo != ":" &&
					valores[i].tipo != "v" &&
					valores[i].tipo != "," &&
					valores[i].tipo != "_" &&
					valores[i].tipo != "+-" &&
					valores[i].tipo != "MOD" &&
					valores[i].tipo != "DIV" &&
					valores[i].tipo != "*/" &&
					valores[i].tipo != "**//" &&
					valores[i].tipo != "OP.LOGICO.E" &&
					valores[i].tipo != "OP.LOGICO.OU" &&
					valores[i].tipo != "OP.LOGICO.UN" &&
					valores[i].tipo != "OP.LOGICO.XOU" &&
					valores[i].tipo != "OP.RELACIONAL" &&
					valores[i].tipo != "INTEIRO" &&
					valores[i].tipo != "LOGICO" &&
					valores[i].tipo != "REAL" &&
					valores[i].tipo != "FUNCMAT" &&
					valores[i].tipo != "LEIA" &&
					valores[i].tipo != "ESCREVA" &&
					valores[i].tipo != "INICIO" &&
					valores[i].tipo != "FIM" &&
					valores[i].tipo != "PONTO" &&
					valores[i].tipo != "<-" &&
					valores[i].tipo != "(" &&
					valores[i].tipo != ")" &&
					valores[i].tipo != ";" {
					filho = Node{tree, nil, valores[i].tipo, j, 0, Token{valores[i].tipo, valores[i].id}}
					j = j - 1
				}
				tree.filhos = append(tree.filhos[:len(tree.filhos)], &filho)
			}
		}
	} else {
		if len(ast.filhos) > 0 {
			for i := 0; i < len(ast.filhos); i++ {
				adicionar(ast.filhos[i], valores, base, z)
			}
		}
	}

}

func maiorPositivo(tree *Node, maior int) int {
	ast := tree
	if ast.indice >= 0 {
		if ast.indice > maior {
			maior = ast.indice
		}
	}

	if len(ast.filhos) > 0 {
		for i := 0; i < len(ast.filhos); i++ {
			tmp2 := maiorPositivo(ast.filhos[i], maior)
			if tmp2 > maior {
				maior = tmp2
			}
		}
	}

	return maior
}
