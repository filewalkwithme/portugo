package main

import "regexp"

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

func verificaPalavraReservada(alvo string) bool {
	for _, palavra := range palavrasReservadas {
		if palavra == alvo {
			return true
		}
	}
	return false
}

func extraiConstanteInteira(texto string) (bool, string, string) {
	bInteiro := false
	vInteiro := ""
	vTextoRestante := ""

	if len(texto) > 0 {
		continua := verificaDigito(string(texto[0]))

		for i := 0; continua; i++ {
			bInteiro = true

			if verificaDigito(string(texto[i])) {
				vInteiro = vInteiro + string(texto[i])
				vTextoRestante = texto[i+1:]
			}

			//ignora real
			if string(texto[i]) == "." {
				vInteiro = ""
				bInteiro = false
				break
			}

			continua = i < len(texto)-1 && verificaDigito(string(texto[i]))
		}
	}

	if !bInteiro {
		vTextoRestante = texto
	}
	return bInteiro, vInteiro, vTextoRestante
}

func extraiConstanteReal(texto string) (bool, string, string) {
	bReal := false
	vReal := ""
	vTextoRestante := ""

	if len(texto) > 0 {
		//0 parte inteira
		//1 parte decimal
		parte := 0
		continua := verificaDigito(string(texto[0]))

		for i := 0; continua; i++ {
			if verificaDigito(string(texto[i])) {
				vReal = vReal + string(texto[i])
				vTextoRestante = texto[i+1:]
			}

			if parte == 1 && string(texto[i]) == "." {
				vReal = ""
				bReal = false
				break
			}

			if parte == 0 && string(texto[i]) == "." {
				bReal = true
				parte = 1
				vReal = vReal + "."
				vTextoRestante = texto[i+1:]

				//se o proximo simbolo nao for um digito, então este token não é do tipo real
				if i == len(texto) || verificaDigito(string(texto[i+1])) == false {
					vReal = ""
					bReal = false
					break
				}
			}

			continua = i < len(texto)-1 && (verificaDigito(string(texto[i])) || string(texto[i]) == ".")
		}
	}

	if !bReal {
		vTextoRestante = texto
	}
	return bReal, vReal, vTextoRestante
}

func extraiConstanteLogica(texto string) (bool, string, string) {
	bLogico := false
	vLogico := ""
	vTextoRestante := ""

	if len(texto) == 10 {
		if texto[0:10] == "verdadeiro" {
			bLogico = true
			vLogico = "verdadeiro"
		}
	}

	if len(texto) > 10 {
		if texto[0:10] == "verdadeiro" && verificaDigito(string(texto[10])) == false && verificaLetra(string(texto[10])) == false {
			bLogico = true
			vLogico = "verdadeiro"
			vTextoRestante = texto[10:]
		}
	}

	if len(texto) == 5 {
		if texto[0:5] == "falso" {
			bLogico = true
			vLogico = "falso"
		}
	}

	if len(texto) > 5 {
		if texto[0:5] == "falso" && verificaDigito(string(texto[5])) == false && verificaLetra(string(texto[5])) == false {
			bLogico = true
			vLogico = "falso"
			vTextoRestante = texto[5:]
		}
	}

	if !bLogico {
		vTextoRestante = texto
	}

	return bLogico, vLogico, vTextoRestante
}

func extraiConstanteCaractere(texto string) (bool, string, string) {
	bCaractere := false
	vCaractere := ""
	vCaractereAnterior := ""
	vTextoRestante := ""
	var aspas = "\""
	var barra = "\\"
	aspasEncerramento := false

	if len(texto) > 0 {
		continua := string(texto[0]) == aspas

		for i := 0; continua; i++ {
			bCaractere = true

			vCaractere = vCaractere + string(texto[i])
			vTextoRestante = texto[i+1:]

			aspasEncerramento = string(texto[i]) == aspas && vCaractereAnterior != barra && i > 0
			if aspasEncerramento {
				bCaractere = true
				break
			}

			continua = i < len(texto)-1 && !aspasEncerramento
			vCaractereAnterior = string(texto[i])
		}
	}

	if !aspasEncerramento {
		bCaractere = false
		vCaractere = ""
	}

	if !bCaractere {
		vTextoRestante = texto
	}
	return bCaractere, vCaractere, vTextoRestante
}

func extraiVariavel(texto string) (bool, string, string) {
	bVariavel := false
	vVariavel := ""
	vTextoRestante := ""

	if len(texto) > 0 {
		continua := verificaLetra(string(texto[0]))

		for i := 0; continua; i++ {
			bVariavel = true

			if verificaLetra(string(texto[i])) || verificaDigito(string(texto[i])) {
				vVariavel = vVariavel + string(texto[i])
				vTextoRestante = texto[i+1:]
			}

			continua = i < len(texto)-1 && (verificaDigito(string(texto[i])) || verificaLetra(string(texto[i])))
		}
	}

	if bVariavel && verificaPalavraReservada(vVariavel) {
		vVariavel = ""
		bVariavel = false
	}

	if !bVariavel {
		vTextoRestante = texto
	}
	return bVariavel, vVariavel, vTextoRestante
}

func extraiTipoVariavel(texto string) (bool, string, string) {
	bTipoVariavel := false
	vTipoVariavel := ""
	vTextoRestante := ""

	if len(texto) > 0 {
		if len(texto) == 4 {
			if texto[0:4] == "real" {
				bTipoVariavel = true
				vTipoVariavel = "real"
				vTextoRestante = texto[4:]
			}
		}

		if len(texto) > 4 && verificaDigito(string(texto[4])) == false && verificaLetra(string(texto[4])) == false {
			if texto[0:4] == "real" {
				bTipoVariavel = true
				vTipoVariavel = "real"
				vTextoRestante = texto[4:]
			}
		}

		if len(texto) == 7 {
			if texto[0:7] == "inteiro" {
				bTipoVariavel = true
				vTipoVariavel = "inteiro"
				vTextoRestante = texto[7:]
			}
		}

		if len(texto) > 7 {
			if texto[0:7] == "inteiro" && verificaDigito(string(texto[7])) == false && verificaLetra(string(texto[7])) == false {
				bTipoVariavel = true
				vTipoVariavel = "inteiro"
				vTextoRestante = texto[7:]
			}
		}

		if len(texto) == 7 {
			if texto[0:7] == "lógico" {
				bTipoVariavel = true
				vTipoVariavel = "lógico"
				vTextoRestante = texto[7:]
			}
		}

		if len(texto) > 7 && verificaDigito(string(texto[7])) == false && verificaLetra(string(texto[7])) == false {
			if texto[0:7] == "lógico" {
				bTipoVariavel = true
				vTipoVariavel = "lógico"
				vTextoRestante = texto[7:]
			}
		}

		if len(texto) == 9 {
			if texto[0:9] == "caractere" {
				bTipoVariavel = true
				vTipoVariavel = "caractere"
				vTextoRestante = texto[9:]
			}
		}

		if len(texto) > 9 {
			if texto[0:9] == "caractere" && verificaDigito(string(texto[9])) == false && verificaLetra(string(texto[9])) == false {
				bTipoVariavel = true
				vTipoVariavel = "caractere"
				vTextoRestante = texto[9:]
			}
		}

	}

	if !bTipoVariavel {
		vTextoRestante = texto
	}
	return bTipoVariavel, vTipoVariavel, vTextoRestante
}

func extraiDoisPontos(texto string) (bool, string, string) {
	b := false
	v := ""
	vTextoRestante := ""

	if len(texto) > 0 {
		if texto[0:1] == ":" {
			b = true
			v = ":"
			vTextoRestante = texto[1:]
		}
	}

	if !b {
		vTextoRestante = texto
	}
	return b, v, vTextoRestante
}

func extraiPontoEVirgula(texto string) (bool, string, string) {
	b := false
	v := ""
	vTextoRestante := ""

	if len(texto) > 0 {
		if texto[0:1] == ";" {
			b = true
			v = ";"
			vTextoRestante = texto[1:]
		}
	}

	if !b {
		vTextoRestante = texto
	}
	return b, v, vTextoRestante
}
