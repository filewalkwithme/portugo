package main

import (
	"fmt"
	"math"
	"strconv"
)

func executaTree(tree *Node, simbolos map[string][]string) string {
	//fmt.Println("exec", tree.token.id)
	if tree.token.tipo == "TIPOVAR" {
		tipoVariavel := tree.token.id
		for i := 0; i < len(tree.filhos); i++ {
			if tipoVariavel == "inteiro" {
				simbolos[tree.filhos[i].token.id] = []string{"0", tipoVariavel}
			}

			if tipoVariavel == "lógico" {
				simbolos[tree.filhos[i].token.id] = []string{"falso", tipoVariavel}
			}
		}
	}

	if tree.token.tipo == "INTEIRO" {
		return tree.token.id
	}

	if tree.token.tipo == "LOGICO" {
		return tree.token.id
	}

	if tree.token.tipo == "v" {
		return simbolos[tree.token.id][0]
	}

	if tree.token.tipo == "ESCREVA" {
		c := ""
		for i := 0; i < len(tree.filhos); i++ {
			c = c + executaTree(tree.filhos[i], simbolos)
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

	if tree.token.tipo == "M9" {

		a, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos), 64)
		if tree.filhos[0].token.id == "-" {
			a = a * -1
		}
		return strconv.FormatFloat(a, 'f', 6, 64)
	}

	if tree.token.tipo == "L3" {

		a := executaTree(tree.filhos[1], simbolos)
		if tree.filhos[0].token.id == "não" {
			if a == "verdadeiro" {
				a = "falso"
			} else {
				a = "verdadeiro"
			}
		}
		return a
	}

	if tree.token.tipo == "+-" {
		a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos), 64)
		b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos), 64)
		//fmt.Println("a:", a, ",b:", b)
		c := 0.0
		if tree.token.id == "+" {
			c = b + a
		} else {
			c = b - a
		}

		return strconv.FormatFloat(c, 'f', 6, 64)
	}

	if tree.token.tipo == "MOD" {
		a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos), 64)
		b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos), 64)
		fmt.Println("a:", a, ",b:", b)
		c := 0.0
		c = math.Mod(a, b)

		return strconv.FormatFloat(c, 'f', 6, 64)
	}

	if tree.token.tipo == "DIV" {
		a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos), 64)
		b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos), 64)
		fmt.Println("a:", a, ",b:", b)
		c := 0.0
		c = math.Trunc(a / b)

		return strconv.FormatFloat(c, 'f', 6, 64)
	}

	if tree.token.tipo == "**//" {
		a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos), 64)
		b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos), 64)
		fmt.Println("a:", a, ",b:", b)
		c := 0.0
		if tree.token.id == "**" {
			c = math.Pow(a, b)
		}

		if tree.token.id == "//" {
			c = math.Pow(a, 1/b)
		}

		return strconv.FormatFloat(c, 'f', 6, 64)
	}

	if tree.token.tipo == "*/" {
		a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos), 64)
		b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos), 64)

		c := 0.0
		if tree.token.id == "*" {
			c = b * a
		} else {
			c = a / b
		}

		return strconv.FormatFloat(c, 'f', 6, 64)
	}

	if tree.token.tipo == "OP.LOGICO.XOU" {
		if tree.token.id == "xou" {
			a := executaTree(tree.filhos[0], simbolos)
			b := executaTree(tree.filhos[1], simbolos)

			c := ""
			if (a == "verdadeiro") != (b == "verdadeiro") {
				c = "verdadeiro"
			} else {
				c = "falso"
			}

			return c
		}
	}

	if tree.token.tipo == "OP.LOGICO" {
		if tree.token.id == "ou" {
			a := executaTree(tree.filhos[0], simbolos)
			b := executaTree(tree.filhos[1], simbolos)

			c := ""
			if a == "verdadeiro" || b == "verdadeiro" {
				c = "verdadeiro"
			} else {
				c = "falso"
			}

			return c
		}

		if tree.token.id == "e" {
			a := executaTree(tree.filhos[0], simbolos)
			b := executaTree(tree.filhos[1], simbolos)

			c := ""
			if a == "verdadeiro" && b == "verdadeiro" {
				c = "verdadeiro"
			} else {
				c = "falso"
			}

			return c
		}
	}

	if tree.token.tipo == "OP.RELACIONAL" {
		c := ""
		if tree.token.id == ">" {
			a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos), 64)
			b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos), 64)

			if a > b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.token.id == ">=" {
			a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos), 64)
			b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos), 64)

			if a >= b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.token.id == "<" {
			a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos), 64)
			b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos), 64)

			if a < b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.token.id == "<=" {
			a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos), 64)
			b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos), 64)

			if a <= b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.token.id == "=" {
			a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos), 64)
			b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos), 64)

			if a == b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.token.id == "<>" {
			a, _ := strconv.ParseFloat(executaTree(tree.filhos[0], simbolos), 64)
			b, _ := strconv.ParseFloat(executaTree(tree.filhos[1], simbolos), 64)

			if a != b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		return c
	}

	if tree.token.tipo == "<-" {
		a := tree.filhos[0].token.id
		b := executaTree(tree.filhos[1], simbolos)

		simbolos[a][0] = b

		return ""
	}

	if len(tree.filhos) > 0 {
		for i := 0; i < len(tree.filhos); i++ {
			tree.filhos[i].pai = tree
			executaTree(tree.filhos[i], simbolos)
		}
	}

	return "EXPRESSAO NAO ENCONTRADA"
}
