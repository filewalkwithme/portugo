package exec

import (
	"fmt"
	core "github.com/maiconio/portugo/core"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Resultado struct {
	Valor string
	Tipo  string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ExecutaTree(tree *core.Node, simbolos map[string][]string) Resultado {
	//fmt.Println("exec", tree.Token.Id)
	if tree.Token.Tipo == "TIPOVAR" {
		tipoVariavel := tree.Token.Id
		for i := 0; i < len(tree.Filhos); i++ {
			if tipoVariavel == "inteiro" {
				simbolos[tree.Filhos[i].Token.Id] = []string{"0", tipoVariavel}
			}

			if tipoVariavel == "real" {
				simbolos[tree.Filhos[i].Token.Id] = []string{"0.0", tipoVariavel}
			}

			if tipoVariavel == "lógico" {
				simbolos[tree.Filhos[i].Token.Id] = []string{"falso", tipoVariavel}
			}

			if tipoVariavel == "caractere" {
				simbolos[tree.Filhos[i].Token.Id] = []string{"", tipoVariavel}
			}
		}
	}

	if tree.Token.Tipo == "INTEIRO" {
		return Resultado{Valor: tree.Token.Id, Tipo: "inteiro"}
	}

	if tree.Token.Tipo == "REAL" {
		return Resultado{Valor: tree.Token.Id, Tipo: "real"}
	}

	if tree.Token.Tipo == "LOGICO" {
		return Resultado{Valor: tree.Token.Id, Tipo: "lógico"}
	}

	if tree.Token.Tipo == "STRING" {
		return Resultado{Valor: tree.Token.Id, Tipo: "string"}
	}

	if tree.Token.Tipo == "v" {
		return Resultado{Valor: simbolos[tree.Token.Id][0], Tipo: simbolos[tree.Token.Id][1]}
	}

	if tree.Token.Tipo == "ESCREVA" {
		c := ""
		for i := 0; i < len(tree.Filhos); i++ {
			c = c + ExecutaTree(tree.Filhos[i], simbolos).Valor
		}
		fmt.Println(c)
	}

	if tree.Token.Tipo == "LEIA" {
		c := ""
		for i := 0; i < len(tree.Filhos); i++ {
			//c = c + ExecutaTree(tree.Filhos[i], simbolos)
			fmt.Scanln(&c)
			simbolos[tree.Filhos[i].Token.Id][0] = c
		}
	}

	if tree.Token.Tipo == "L5" {

		a := ExecutaTree(tree.Filhos[1], simbolos).Valor
		if tree.Filhos[0].Token.Id == "não" {
			if a == "verdadeiro" {
				a = "falso"
			} else {
				a = "verdadeiro"
			}
		}
		return Resultado{Valor: a}
	}

	//
	if tree.Token.Tipo == "M9" {
		resultadoA := ExecutaTree(tree.Filhos[1], simbolos)

		resultadoInteiro := resultadoA.Tipo == "inteiro"
		if resultadoInteiro {
			a, _ := strconv.ParseInt(resultadoA.Valor, 10, 0)
			if tree.Filhos[0].Token.Id == "-" {
				a = a * -1
			}
			return Resultado{Valor: strconv.FormatInt(a, 10), Tipo: "inteiro"}
		}

		a, _ := strconv.ParseFloat(resultadoA.Valor, 64)
		if tree.Filhos[0].Token.Id == "-" {
			a = a * -1
		}
		return Resultado{Valor: strconv.FormatFloat(a, 'f', 6, 64), Tipo: "real"}
	}

	if tree.Token.Tipo == "+-" {
		resultadoA := ExecutaTree(tree.Filhos[0], simbolos)
		resultadoB := ExecutaTree(tree.Filhos[1], simbolos)

		resultadoInteiro := resultadoA.Tipo == "inteiro" && resultadoB.Tipo == "inteiro"

		if resultadoInteiro {
			a, _ := strconv.ParseInt(resultadoA.Valor, 10, 0)
			b, _ := strconv.ParseInt(resultadoB.Valor, 10, 0)

			var c int64
			c = 0
			if tree.Token.Id == "+" {
				c = b + a
			} else {
				c = b - a
			}

			return Resultado{Valor: strconv.FormatInt(c, 10), Tipo: "inteiro"}
		}
		a, _ := strconv.ParseFloat(resultadoA.Valor, 64)
		b, _ := strconv.ParseFloat(resultadoB.Valor, 64)

		c := 0.0
		if tree.Token.Id == "+" {
			c = b + a
		} else {
			c = b - a
		}

		return Resultado{Valor: strconv.FormatFloat(c, 'f', 6, 64), Tipo: "real"}
	}

	if tree.Token.Tipo == "**//" {
		a, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[0], simbolos).Valor, 64)
		b, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[1], simbolos).Valor, 64)

		c := 0.0
		if tree.Token.Id == "**" {
			c = math.Pow(a, b)
		}

		if tree.Token.Id == "//" {
			c = math.Pow(a, 1/b)
		}

		return Resultado{Valor: strconv.FormatFloat(c, 'f', 6, 64), Tipo: "real"}
	}

	if tree.Token.Tipo == "*/" {
		a, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[0], simbolos).Valor, 64)
		b, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[1], simbolos).Valor, 64)

		c := 0.0
		if tree.Token.Id == "*" {
			c = b * a
		} else {
			c = a / b
		}

		return Resultado{Valor: strconv.FormatFloat(c, 'f', 6, 64), Tipo: "real"}
	}

	if tree.Token.Tipo == "MOD" {
		resultadoA := ExecutaTree(tree.Filhos[0], simbolos)
		resultadoB := ExecutaTree(tree.Filhos[1], simbolos)

		resultadoInteiro := resultadoA.Tipo == "inteiro" && resultadoB.Tipo == "inteiro"

		if resultadoInteiro {
			a, _ := strconv.ParseInt(resultadoA.Valor, 10, 0)
			b, _ := strconv.ParseInt(resultadoB.Valor, 10, 0)

			var c int64
			c = 0
			c = a % b

			return Resultado{Valor: strconv.FormatInt(c, 10), Tipo: "inteiro"}
		}
		return Resultado{Valor: "Função MOD: requer 2 operandos do tipo inteiro", Tipo: "erro"}
	}

	if tree.Token.Tipo == "DIV" {
		resultadoA := ExecutaTree(tree.Filhos[0], simbolos)
		resultadoB := ExecutaTree(tree.Filhos[1], simbolos)

		resultadoInteiro := resultadoA.Tipo == "inteiro" && resultadoB.Tipo == "inteiro"

		if resultadoInteiro {
			a, _ := strconv.ParseInt(resultadoA.Valor, 10, 0)
			b, _ := strconv.ParseInt(resultadoB.Valor, 10, 0)

			var c int64
			c = 0
			c = a / b

			return Resultado{Valor: strconv.FormatInt(c, 10), Tipo: "inteiro"}
		}
		return Resultado{Valor: "Função DIV: requer 2 operandos do tipo inteiro", Tipo: "erro"}
	}

	if tree.Token.Tipo == "OP.LOGICO.XOU" {
		if tree.Token.Id == "xou" {
			a := ExecutaTree(tree.Filhos[0], simbolos).Valor
			b := ExecutaTree(tree.Filhos[1], simbolos).Valor

			c := ""
			if (a == "verdadeiro") != (b == "verdadeiro") {
				c = "verdadeiro"
			} else {
				c = "falso"
			}

			return Resultado{Valor: c}
		}
	}

	if tree.Token.Tipo == "OP.LOGICO.E" {
		a := ExecutaTree(tree.Filhos[0], simbolos).Valor
		b := ExecutaTree(tree.Filhos[1], simbolos).Valor

		c := ""
		if a == "verdadeiro" && b == "verdadeiro" {
			c = "verdadeiro"
		} else {
			c = "falso"
		}

		return Resultado{Valor: c}
	}

	if tree.Token.Tipo == "OP.LOGICO.OU" {
		a := ExecutaTree(tree.Filhos[0], simbolos).Valor
		b := ExecutaTree(tree.Filhos[1], simbolos).Valor

		c := ""
		if a == "verdadeiro" || b == "verdadeiro" {
			c = "verdadeiro"
		} else {
			c = "falso"
		}

		return Resultado{Valor: c}
	}

	if tree.Token.Tipo == "OP.RELACIONAL" {
		c := ""
		if tree.Token.Id == ">" {
			a, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[0], simbolos).Valor, 64)
			b, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[1], simbolos).Valor, 64)

			if a > b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.Token.Id == ">=" {
			a, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[0], simbolos).Valor, 64)
			b, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[1], simbolos).Valor, 64)

			if a >= b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.Token.Id == "<" {
			a, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[0], simbolos).Valor, 64)
			b, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[1], simbolos).Valor, 64)

			if a < b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.Token.Id == "<=" {
			a, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[0], simbolos).Valor, 64)
			b, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[1], simbolos).Valor, 64)

			if a <= b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.Token.Id == "=" {
			a, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[0], simbolos).Valor, 64)
			b, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[1], simbolos).Valor, 64)

			if a == b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		if tree.Token.Id == "<>" {
			a, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[0], simbolos).Valor, 64)
			b, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[1], simbolos).Valor, 64)

			if a != b {
				c = "verdadeiro"
			} else {
				c = "falso"
			}
		}

		return Resultado{Valor: c}
	}

	if tree.Token.Tipo == "<-" {
		a := tree.Filhos[0].Token.Id
		b := ExecutaTree(tree.Filhos[1], simbolos).Valor

		simbolos[a][0] = b
		simbolos[a][1] = ExecutaTree(tree.Filhos[1], simbolos).Tipo

		return Resultado{Valor: ""}
	}

	if tree.Token.Tipo == "FUNCMAT" {

		retornaReal := map[string]bool{
			"sen":    true,
			"cos":    true,
			"tg":     true,
			"arctg":  true,
			"arccos": true,
			"arcsen": true,
			"abs":    true,
			"frac":   true,
		}

		if retornaReal[tree.Token.Id] {
			a, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[0], simbolos).Valor, 64)
			c := 0.0

			if tree.Token.Id == "sen" {
				c = math.Sin(a)
			}

			if tree.Token.Id == "cos" {
				c = math.Cos(a)
			}

			if tree.Token.Id == "tg" {
				c = math.Tan(a)
			}

			if tree.Token.Id == "arctg" {
				c = math.Atan(a)
			}

			if tree.Token.Id == "arccos" {
				c = math.Acos(a)
			}

			if tree.Token.Id == "arcsen" {
				c = math.Asin(a)
			}

			if tree.Token.Id == "abs" {
				c = math.Abs(a)
			}

			if tree.Token.Id == "frac" {
				c = a - math.Trunc(a)
			}
			return Resultado{Valor: strconv.FormatFloat(c, 'f', 6, 64), Tipo: "real"}
		} else {
			if tree.Token.Id == "int" {
				a, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[0], simbolos).Valor, 64)
				c := 0.0
				c = math.Trunc(a)
				return Resultado{Valor: strconv.FormatFloat(c, 'f', 0, 64), Tipo: "inteiro"}
			}

			if tree.Token.Id == "ard" {
				a, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[0], simbolos).Valor, 64)
				frac := a - math.Trunc(a)
				c := 0.0
				if frac < 0.5 {
					c = math.Floor(a)
				} else {
					c = math.Ceil(a)
				}

				return Resultado{Valor: strconv.FormatFloat(c, 'f', 0, 64), Tipo: "inteiro"}
			}

			if tree.Token.Id == "sinal" {
				a, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[0], simbolos).Valor, 64)

				c := 0.0
				if c != a {
					if a < c {
						c = -1.0
					} else {
						c = 1.0
					}
				}
				return Resultado{Valor: strconv.FormatFloat(c, 'f', 0, 64), Tipo: "inteiro"}
			}

			if tree.Token.Id == "rnd" {
				//a, _ := strconv.ParseFloat(ExecutaTree(tree.Filhos[0], simbolos).Valor, 64)
				a, _ := strconv.ParseInt(ExecutaTree(tree.Filhos[0], simbolos).Valor, 10, 0)
				c := rand.Int63n(a)
				return Resultado{Valor: strconv.FormatInt(c, 10), Tipo: "inteiro"}
			}
		}
		return Resultado{Valor: strconv.FormatFloat(0.0, 'f', 6, 64), Tipo: "real"}

	}

	if len(tree.Filhos) > 0 {
		for i := 0; i < len(tree.Filhos); i++ {
			tree.Filhos[i].Pai = tree
			ExecutaTree(tree.Filhos[i], simbolos)
		}
	}

	return Resultado{Valor: "EXPRESSAO NAO ENCONTRADA: {" + tree.Token.Tipo + "}", Tipo: "erro"}
}
