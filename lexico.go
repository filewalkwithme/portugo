package main

import "regexp"

type token struct {
	tipo  string
	valor string
}

var palavrasReservadas = []string{"verdadeiro", "falso", "inteiro", "caractere", "real", "lógico"}

func verificaDigito(simbolo string) bool {
	re := regexp.MustCompile("[[:digit:]]")
	return re.MatchString(simbolo)
}

func verificaLetra(simbolo string) bool {
	//caracteres pt-BR
	re := regexp.MustCompile("[[:alpha:]áàâãÀÁÂÃéÉíÍóÓúÚçÇ]")
	return re.MatchString(simbolo)
}

func verificaLetraMaiuscula(simbolo string) bool {
	//caracteres pt-BR
	re := regexp.MustCompile("[[:upper:]ÀÁÂÃÉÍÓÚÇ]")
	return re.MatchString(simbolo)
}

func verificaPalavraReservada(alvo string) bool {
	for _, palavra := range palavrasReservadas {
		if palavra == alvo {
			return true
		}
	}
	return false
}

func extraiConstanteInteira(texto string) (token, string) {
	tipoToken := ""
	valorToken := ""
	vTextoRestante := ""

	if len(texto) > 0 {
		continua := verificaDigito(string(texto[0]))

		for i := 0; continua; i++ {
			tipoToken = "CONSTANTE_INTEIRA"

			if verificaDigito(string(texto[i])) {
				valorToken = valorToken + string(texto[i])
				vTextoRestante = texto[i+1:]
			}

			//ignora real
			if string(texto[i]) == "." {
				valorToken = ""
				tipoToken = ""
				break
			}

			continua = i < len(texto)-1 && verificaDigito(string(texto[i]))
		}
	}

	if tipoToken != "CONSTANTE_INTEIRA" {
		vTextoRestante = texto
	}
	return token{tipo: tipoToken, valor: valorToken}, vTextoRestante
}

func extraiConstanteReal(texto string) (token, string) {
	tipoToken := ""
	valorToken := ""
	vTextoRestante := ""

	if len(texto) > 0 {
		//0 parte inteira
		//1 parte decimal
		parte := 0
		continua := verificaDigito(string(texto[0]))

		for i := 0; continua; i++ {
			if verificaDigito(string(texto[i])) {
				valorToken = valorToken + string(texto[i])
				vTextoRestante = texto[i+1:]
			}

			if parte == 1 && string(texto[i]) == "." {
				valorToken = ""
				tipoToken = ""
				break
			}

			if parte == 0 && string(texto[i]) == "." {
				tipoToken = "CONSTANTE_REAL"
				parte = 1
				valorToken = valorToken + "."
				vTextoRestante = texto[i+1:]

				//se o proximo simbolo nao for um digito, então este token não é do tipo real
				if i == len(texto) || verificaDigito(string(texto[i+1])) == false {
					valorToken = ""
					tipoToken = ""
					break
				}
			}

			continua = i < len(texto)-1 && (verificaDigito(string(texto[i])) || string(texto[i]) == ".")
		}
	}

	if tipoToken == "" {
		vTextoRestante = texto
	}
	return token{tipo: tipoToken, valor: valorToken}, vTextoRestante
}

func extraiConstanteLogica(texto string) (token, string) {
	tipoToken := ""
	valorToken := ""
	vTextoRestante := ""

	if len(texto) == 10 {
		if texto[0:10] == "verdadeiro" {
			tipoToken = "CONSTANTE_LOGICA"
			valorToken = "verdadeiro"
		}
	}

	if len(texto) > 10 {
		if texto[0:10] == "verdadeiro" && verificaDigito(string(texto[10])) == false && verificaLetra(string(texto[10])) == false {
			tipoToken = "CONSTANTE_LOGICA"
			valorToken = "verdadeiro"
			vTextoRestante = texto[10:]
		}
	}

	if len(texto) == 5 {
		if texto[0:5] == "falso" {
			tipoToken = "CONSTANTE_LOGICA"
			valorToken = "falso"
		}
	}

	if len(texto) > 5 {
		if texto[0:5] == "falso" && verificaDigito(string(texto[5])) == false && verificaLetra(string(texto[5])) == false {
			tipoToken = "CONSTANTE_LOGICA"
			valorToken = "falso"
			vTextoRestante = texto[5:]
		}
	}

	if tipoToken == "" {
		vTextoRestante = texto
	}

	return token{tipo: tipoToken, valor: valorToken}, vTextoRestante
}

func extraiConstanteCaractere(texto string) (token, string) {
	tipoToken := ""
	valorToken := ""
	vCaractereAnterior := ""
	vTextoRestante := ""
	var aspas = "\""
	var barra = "\\"
	aspasEncerramento := false

	if len(texto) > 0 {
		continua := string(texto[0]) == aspas

		for i := 0; continua; i++ {
			tipoToken = "CONSTANTE_CARACTERE"

			valorToken = valorToken + string(texto[i])
			vTextoRestante = texto[i+1:]

			aspasEncerramento = string(texto[i]) == aspas && vCaractereAnterior != barra && i > 0
			if aspasEncerramento {
				tipoToken = "CONSTANTE_CARACTERE"
				break
			}

			continua = i < len(texto)-1 && !aspasEncerramento
			vCaractereAnterior = string(texto[i])
		}
	}

	if !aspasEncerramento {
		tipoToken = ""
		valorToken = ""
	}

	if tipoToken == "" {
		vTextoRestante = texto
	}
	return token{tipo: tipoToken, valor: valorToken}, vTextoRestante
}

func extraiVariavel(texto string) (token, string) {
	tipoToken := ""
	valorToken := ""
	vTextoRestante := ""

	if len(texto) > 0 {
		continua := verificaLetraMaiuscula(string(texto[0]))

		for i := 0; continua; i++ {
			tipoToken = "VARIAVEL"

			if verificaLetraMaiuscula(string(texto[i])) || verificaDigito(string(texto[i])) {
				valorToken = valorToken + string(texto[i])
				vTextoRestante = texto[i+1:]
			}

			continua = i < len(texto)-1 && (verificaDigito(string(texto[i])) || verificaLetraMaiuscula(string(texto[i])))
		}
	}

	if tipoToken == "VARIAVEL" && verificaPalavraReservada(valorToken) {
		valorToken = ""
		tipoToken = ""
	}

	if tipoToken == "" {
		vTextoRestante = texto
	}
	return token{tipo: tipoToken, valor: valorToken}, vTextoRestante
}

func extraiTipoVariavel(texto string) (token, string) {
	tipoToken := ""
	valorToken := ""
	vTextoRestante := ""

	if len(texto) > 0 {
		if len(texto) == 4 {
			if texto[0:4] == "real" {
				tipoToken = "TIPO_VARIAVEL"
				valorToken = "real"
				vTextoRestante = texto[4:]
			}
		}

		if len(texto) > 4 && verificaDigito(string(texto[4])) == false && verificaLetra(string(texto[4])) == false {
			if texto[0:4] == "real" {
				tipoToken = "TIPO_VARIAVEL"
				valorToken = "real"
				vTextoRestante = texto[4:]
			}
		}

		if len(texto) == 7 {
			if texto[0:7] == "inteiro" {
				tipoToken = "TIPO_VARIAVEL"
				valorToken = "inteiro"
				vTextoRestante = texto[7:]
			}
		}

		if len(texto) > 7 {
			if texto[0:7] == "inteiro" && verificaDigito(string(texto[7])) == false && verificaLetra(string(texto[7])) == false {
				tipoToken = "TIPO_VARIAVEL"
				valorToken = "inteiro"
				vTextoRestante = texto[7:]
			}
		}

		if len(texto) == 7 {
			if texto[0:7] == "lógico" {
				tipoToken = "TIPO_VARIAVEL"
				valorToken = "lógico"
				vTextoRestante = texto[7:]
			}
		}

		if len(texto) > 7 && verificaDigito(string(texto[7])) == false && verificaLetra(string(texto[7])) == false {
			if texto[0:7] == "lógico" {
				tipoToken = "TIPO_VARIAVEL"
				valorToken = "lógico"
				vTextoRestante = texto[7:]
			}
		}

		if len(texto) == 9 {
			if texto[0:9] == "caractere" {
				tipoToken = "TIPO_VARIAVEL"
				valorToken = "caractere"
				vTextoRestante = texto[9:]
			}
		}

		if len(texto) > 9 {
			if texto[0:9] == "caractere" && verificaDigito(string(texto[9])) == false && verificaLetra(string(texto[9])) == false {
				tipoToken = "TIPO_VARIAVEL"
				valorToken = "caractere"
				vTextoRestante = texto[9:]
			}
		}

	}

	if tipoToken == "" {
		vTextoRestante = texto
	}
	return token{tipo: tipoToken, valor: valorToken}, vTextoRestante
}

func extraiCaractereUnico(texto string, caractere string, tipoAlvo string) (token, string) {
	tipoToken := ""
	valorToken := ""
	vTextoRestante := ""

	if len(texto) > 0 {
		if texto[0:1] == caractere {
			tipoToken = tipoAlvo
			valorToken = caractere
			vTextoRestante = texto[1:]
		}
	}

	if tipoToken == "" {
		vTextoRestante = texto
	}
	return token{tipo: tipoToken, valor: valorToken}, vTextoRestante
}

func extraiDoisPontos(texto string) (token, string) {
	return extraiCaractereUnico(texto, ":", "DOIS_PONTOS")
}

func extraiPontoEVirgula(texto string) (token, string) {
	return extraiCaractereUnico(texto, ";", "PONTO_E_VIRGULA")
}

func extraiVirgula(texto string) (token, string) {
	return extraiCaractereUnico(texto, ",", "VIRGULA")
}
