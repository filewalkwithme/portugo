package main

import (
	"fmt"
	"math"
	"strconv"
)

type Resultado struct {
	valor string
	tipo string
}

func executaTree(tree *Node, simbolos map[string][]string) Resultado {
	//fmt.Println("exec", tree.token.id)
	if tree.token.tipo == "TIPOVAR" {
		tipoVariavel := tree.token.id
		for i := 0; i < len(tree.filhos); i++ {
			if tipoVariavel == "inteiro" {
				simbolos[tree.filhos[i].token.id] = []string{"0", tipoVariavel}
			}

			if tipoVariavel == "real" {
				simbolos[tree.filhos[i].token.id] = []string{"0.0", tipoVariavel}
			}

			if tipoVariavel == "lógico" {
				simbolos[tree.filhos[i].token.id] = []string{"falso", tipoVariavel}
			}
		}
	}

	if tree.token.tipo == "INTEIRO" {
		return Resultado{valor: tree.token.id, tipo: "inteiro"}
	}

	if tree.token.tipo == "REAL" {
		return Resultado{valor: tree.token.id}
	}

	if tree.token.tipo == "LOGICO" {
		return Resultado{valor: tree.token.id}
	}

	if tree.token.tipo == "v" {
		return Resultado{valor: simbolos[tree.token.id][0], tipo: simbolos[tree.token.id][1]}
	}

	if tree.token.tipo == "ESCREVA" {
		c := ""
		for i := 0; i < len(tree.filhos); i++ {
			c = c + executaTree(tree.filhos[i], simbolos).valor
		}
		fmt.Println(c)
	}

	if tree.token.tipo == "LEIA" {
		c := ""
		for i := 0; i < len(tree.filhos); i++ {
			//c = c + executaTree(tree.filhos[i], simbolos)
			fmt.Scanln(&c)
			simbolos[tree.filhos[i].token.id][0] = c
		}
	}

	if tree.token.tipo == "L5" {

		a := executaTree(tree.filhos[1], simbolos).valor
		if tree.filhos[0].token.id == "não" {
			if a == "verdadeiro" {
				a = "falso"
			} else {
				a = "verdadeiro"
			}
		}
		return Resultado{valor: a}
	}

	if tree.token.tipo == "M9" {
		resultadoInteiro := executaTree(tree.filhos[1], simbolos).tipo == "inteiro" 
		if resultadoInteiro {
			a, _ := strconv.ParseInt(executaTree(tree.filhos[1], simbolos).valor, 10, 0)
			if tree.filhos[0].token.id == "-" {
				a = a * -1
			}
			return Resultado{valor: strconv.FormatInt(a, 10), tipo: "inteiro"}
		}

		a, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos).valor, 64)
		if tree.filhos[0].token.id == "-" {
			a = a * -1
		}
		return Resultado{valor: strconv.FormatFloat(a, 'f', 6, 64), tipo: "real"}
	}

	if tree.token.tipo == "+-" {
		resultadoInteiro := executaTree(tree.filhos[0], simbolos).tipo == "inteiro" &&  executaTree(tree.filhos[1], simbolos).tipo == "inteiro" 
	
		if resultadoInteiro {
			a, _ := strconv.ParseInt(executaTree(tree.filhos[0], simbolos).valor, 10, 0)
			b, _ := strconv.ParseInt(executaTree(tree.filhos[1], simbolos).valor, 10, 0)

			var c int64
			c = 0
			if tree.token.id == "+" {
				c = b + a
			} else {
				c = b - a
			}

			return Resultado{valor: strconv.FormatInt(c, 10), tipo: "inteiro"}
		}
		a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos).valor, 64)
		b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos).valor, 64)

		c := 0.0
		if tree.token.id == "+" {
			c = b + a
		} else {
			c = b - a
		}

		return Resultado{valor: strconv.FormatFloat(c, 'f', 6, 64), tipo: "real"}
	}

	if tree.token.tipo == "**//" {
		a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos).valor, 64)
		b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos).valor, 64)

		c := 0.0
		if tree.token.id == "**" {
			c = math.Pow(a, b)
		}

		if tree.token.id == "//" {
			c = math.Pow(a, 1/b)
		}

		return Resultado{valor: strconv.FormatFloat(c, 'f', 6, 64), tipo: "real"}
	}

	if tree.token.tipo == "*/" {
		a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos).valor, 64)
		b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos).valor, 64)

		c := 0.0
		if tree.token.id == "*" {
			c = b * a
		} else {
			c = a / b
		}

		return Resultado{valor: strconv.FormatFloat(c, 'f', 6, 64), tipo: "real"}
	}

	if tree.token.tipo == "MOD" {
		resultadoInteiro := executaTree(tree.filhos[0], simbolos).tipo == "inteiro" &&  executaTree(tree.filhos[1], simbolos).tipo == "inteiro" 
		
		if resultadoInteiro {
			a, _ := strconv.ParseInt(executaTree(tree.filhos[0], simbolos).valor, 10, 0)
			b, _ := strconv.ParseInt(executaTree(tree.filhos[1], simbolos).valor, 10, 0)

			var c int64
			c = 0
			c = a % b

			return Resultado{valor: strconv.FormatInt(c, 10), tipo:"inteiro"}
		}
		return Resultado{valor: "Função MOD: requer 2 operandos do tipo inteiro", tipo:"erro"}
	}

	if tree.token.tipo == "DIV" {
		resultadoInteiro := executaTree(tree.filhos[0], simbolos).tipo == "inteiro" &&  executaTree(tree.filhos[1], simbolos).tipo == "inteiro" 
		
		if resultadoInteiro {
			a, _ := strconv.ParseInt(executaTree(tree.filhos[0], simbolos).valor, 10, 0)
			b, _ := strconv.ParseInt(executaTree(tree.filhos[1], simbolos).valor, 10, 0)

			var c int64
			c = 0
			c = a / b
	
			return Resultado{valor: strconv.FormatInt(c, 10), tipo:"inteiro"}
		}
		return Resultado{valor: "Função DIV: requer 2 operandos do tipo inteiro", tipo:"erro"}
	}


	if tree.token.tipo == "OP.LOGICO.XOU" {
		if tree.token.id == "xou" {
			a := executaTree(tree.filhos[0], simbolos).valor
			b := executaTree(tree.filhos[1], simbolos).valor

			c := ""
			if (a == "verdadeiro") != (b == "verdadeiro") {
				c = "verdadeiro"
			} else {
				c = "falso"
			}

			return Resultado{valor: c}
		}
	}

	if tree.token.tipo == "OP.LOGICO.E" {
		a := executaTree(tree.filhos[0], simbolos).valor
		b := executaTree(tree.filhos[1], simbolos).valor

		c := ""
		if a == "verdadeiro" && b == "verdadeiro" {
			c = "verdadeiro"
		} else {
			c = "falso"
		}

		return Resultado{valor: c}
	}

	if tree.token.tipo == "OP.LOGICO.OU" {
		a := executaTree(tree.filhos[0], simbolos).valor
		b := executaTree(tree.filhos[1], simbolos).valor

		c := ""
		if a == "verdadeiro" || b == "verdadeiro" {
			c = "verdadeiro"
		} else {
			c = "falso"
		}

		return Resultado{valor: c}
	}

	if tree.token.tipo == "OP.RELACIONAL" {
		c := ""
		if tree.token.id == ">" {
			a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos).valor, 64)
			b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos).valor, 64)

			if a > b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.token.id == ">=" {
			a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos).valor, 64)
			b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos).valor, 64)

			if a >= b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.token.id == "<" {
			a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos).valor, 64)
			b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos).valor, 64)

			if a < b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.token.id == "<=" {
			a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos).valor, 64)
			b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos).valor, 64)

			if a <= b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.token.id == "=" {
			a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos).valor, 64)
			b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos).valor, 64)

			if a == b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.token.id == "<>" {
			a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos).valor, 64)
			b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos).valor, 64)

			if a != b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		return Resultado{valor: c}
	}

	if tree.token.tipo == "<-" {
		a := tree.filhos[0].token.id
		b := executaTree(tree.filhos[1], simbolos).valor

		simbolos[a][0] = b
		simbolos[a][1] = executaTree(tree.filhos[1], simbolos).tipo

		return Resultado{valor: ""}
	}

	if len(tree.filhos) > 0 {
		for i := 0; i < len(tree.filhos); i++ {
			tree.filhos[i].pai = tree
			executaTree(tree.filhos[i], simbolos)
		}
	}

	return Resultado{valor: "EXPRESSAO NAO ENCONTRADA: {" + tree.token.tipo + "}"}
}
