package main

var digitos = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func verificaDigito(simbolo string) bool {
	encontrou := false
	for _, digito := range digitos {
		if digito == simbolo {
			encontrou = true
			break
		}
	}
	return encontrou
}

func extraiInteiro(texto string) (bool, string, string) {
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

func extraiReal(texto string) (bool, string, string) {
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
