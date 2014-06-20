package lex

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"regexp"
	"strings"
	util "github.com/maiconio/portugo/util"
	core "github.com/maiconio/portugo/core"
)


func CarregaTokens(arquivo string) []core.Token {
	arq, _ := ioutil.ReadFile(arquivo)
	m := make(map[int]string)
	var i int = 0

	scanner := bufio.NewScanner(bytes.NewReader(arq))
	for scanner.Scan() {
		m[i] = scanner.Text()
		i++
	}

	listaTokens := []core.Token{}

	acaba := false
	for j := 0; j < len(m); j++ {
		acaba = false
		if len(m[j]) > 0 {
			for !acaba {
				id, tipoToken, frase := pegaToken(strings.TrimRight(m[j], " "))
				m[j] = frase

				token := core.Token{tipoToken, id}
				if token.Tipo == "+-" && token.Id == "-" {
					listaTokens = util.PushToken(listaTokens, core.Token{"+-", "+"})
				}

				listaTokens = util.PushToken(listaTokens, token)

				acaba = (len(m[j]) == 0) || (tipoToken == "ERRO")
			}
		}
	}

	return listaTokens
}

func pegaToken(linha string) (string, string, string) {
	re := regexp.MustCompile("^(\\s*)")

	//retira espaços
	linha = re.ReplaceAllString(linha, "")

	token := regexp.MustCompile("^((inteiro)|(real)|(caractere)|(lógico))(\\s*):").FindString(linha)
	if len(token) > 0 {
		token = regexp.MustCompile("(\\s*):").ReplaceAllString(token, "")
		return token, "TIPOVAR", linha[len(token):]
	}

	token = regexp.MustCompile("^:").FindString(linha)
	if len(token) > 0 {
		return token, ":", linha[len(token):]
	}

	token = regexp.MustCompile("^[.]").FindString(linha)
	if len(token) > 0 {
		return token, "PONTO", linha[len(token):]
	}

	token = regexp.MustCompile("^[(]").FindString(linha)
	if len(token) > 0 {
		return token, "(", linha[len(token):]
	}

	token = regexp.MustCompile("^[)]").FindString(linha)
	if len(token) > 0 {
		return token, ")", linha[len(token):]
	}

	token = regexp.MustCompile("^,").FindString(linha)
	if len(token) > 0 {
		return token, ",", linha[len(token):]
	}

	token = regexp.MustCompile("^;").FindString(linha)
	if len(token) > 0 {
		return token, ";", linha[len(token):]
	}

	token = regexp.MustCompile("^(div)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "div"
		return token, "DIV", linha[len(token):]
	}

	token = regexp.MustCompile("^(mod)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "mod"
		return token, "MOD", linha[len(token):]
	}

	token = regexp.MustCompile("^(sen)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "sen"
		return token, "FUNCMAT", linha[len(token):]
	}

	token = regexp.MustCompile("^(cos)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "cos"
		return token, "FUNCMAT", linha[len(token):]
	}

	token = regexp.MustCompile("^(tg)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "tg"
		return token, "FUNCMAT", linha[len(token):]
	}

	token = regexp.MustCompile("^(arctg)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "arctg"
		return token, "FUNCMAT", linha[len(token):]
	}

	token = regexp.MustCompile("^(arccos)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "arccos"
		return token, "FUNCMAT", linha[len(token):]
	}

	token = regexp.MustCompile("^(arcsen)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "arcsen"
		return token, "FUNCMAT", linha[len(token):]
	}

	token = regexp.MustCompile("^(abs)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "abs"
		return token, "FUNCMAT", linha[len(token):]
	}

	token = regexp.MustCompile("^(int)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "int"
		return token, "FUNCMAT", linha[len(token):]
	}

	token = regexp.MustCompile("^(frac)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "frac"
		return token, "FUNCMAT", linha[len(token):]
	}

	token = regexp.MustCompile("^(ard)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "ard"
		return token, "FUNCMAT", linha[len(token):]
	}

	token = regexp.MustCompile("^(RND)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "ard"
		return token, "FUNCMAT", linha[len(token):]
	}

	token = regexp.MustCompile("^(e)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "e"
		return token, "OP.LOGICO.E", linha[len(token):]
	}

	token = regexp.MustCompile("^(ou)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "ou"
		return token, "OP.LOGICO.OU", linha[len(token):]
	}

	token = regexp.MustCompile("^(xou)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "xou"
		return token, "OP.LOGICO.XOU", linha[len(token):]
	}

	token = regexp.MustCompile("^(não)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "não"
		return token, "OP.LOGICO.UN", linha[len(token):]
	}

	token = regexp.MustCompile("^(sinal)([^À-úA-z]+|$)").FindString(linha)
	if len(token) > 0 {
		token = "sinal"
		return token, "FUNCMAT", linha[len(token):]
	}

	token = regexp.MustCompile("^(<-)").FindString(linha)
	if len(token) > 0 {
		return token, "<-", linha[len(token):]
	}

	token = regexp.MustCompile("^(<[ ]*=)").FindString(linha)
	if len(token) > 0 {
		return token, "OP.RELACIONAL", linha[len(token):]
	}

	token = regexp.MustCompile("^(>[ ]*=)").FindString(linha)
	if len(token) > 0 {
		return token, "OP.RELACIONAL", linha[len(token):]
	}

	token = regexp.MustCompile("^(<[ ]*>)").FindString(linha)
	if len(token) > 0 {
		return token, "OP.RELACIONAL", linha[len(token):]
	}

	token = regexp.MustCompile("^(<)").FindString(linha)
	if len(token) > 0 {
		return token, "OP.RELACIONAL", linha[len(token):]
	}

	token = regexp.MustCompile("^(>)").FindString(linha)
	if len(token) > 0 {
		return token, "OP.RELACIONAL", linha[len(token):]
	}

	token = regexp.MustCompile("^(=)").FindString(linha)
	if len(token) > 0 {
		return token, "OP.RELACIONAL", linha[len(token):]
	}

	token = regexp.MustCompile("^(leia)").FindString(linha)
	if len(token) > 0 {
		return token, "LEIA", linha[len(token):]
	}

	token = regexp.MustCompile("^(início)").FindString(linha)
	if len(token) > 0 {
		return token, "INICIO", linha[len(token):]
	}

	token = regexp.MustCompile("^(fim)").FindString(linha)
	if len(token) > 0 {
		return token, "FIM", linha[len(token):]
	}

	token = regexp.MustCompile("^(escreva)").FindString(linha)
	if len(token) > 0 {
		return token, "ESCREVA", linha[len(token):]
	}

	token = regexp.MustCompile("^[+]").FindString(linha)
	if len(token) > 0 {
		return token, "+-", linha[len(token):]
	}

	token = regexp.MustCompile("^[-]").FindString(linha)
	if len(token) > 0 {
		return token, "+-", linha[len(token):]
	}

	token = regexp.MustCompile("^[*]{2}").FindString(linha)
	if len(token) > 0 {
		return token, "**//", linha[len(token):]
	}

	token = regexp.MustCompile("^[/]{2}").FindString(linha)
	if len(token) > 0 {
		return token, "**//", linha[len(token):]
	}

	token = regexp.MustCompile("^[*]").FindString(linha)
	if len(token) > 0 {
		return token, "*/", linha[len(token):]
	}

	token = regexp.MustCompile("^[/]").FindString(linha)
	if len(token) > 0 {
		return token, "*/", linha[len(token):]
	}

	token = regexp.MustCompile("^[0-9]+[.][0-9]+[À-úA-z]+").FindString(linha)
	if len(token) > 0 {
		return token, "REAL INVALIDO", linha[len(token):]
	}

	token = regexp.MustCompile("^[0-9]+[.][0-9]+").FindString(linha)
	if len(token) > 0 {
		return token, "REAL", linha[len(token):]
	}

	token = regexp.MustCompile("^[0-9]+").FindString(linha)
	if len(token) > 0 {
		return token, "INTEIRO", linha[len(token):]
	}

	//token = regexp.MustCompile("^(\".*\")[À-úA-z]+").FindString(linha)
	//if (len(token) > 0) {
	//	return token, "STRING INVALIDO", linha[len(token):]
	//}

	token = regexp.MustCompile("^\"[^\"]*\"").FindString(linha)
	if len(token) > 0 {
		return token[1:len(token)-1], "STRING", linha[len(token):]
	}

	token = regexp.MustCompile("^{.*}").FindString(linha)
	if len(token) > 0 {
		return token, "COMENTARIO", linha[len(token):]
	}

	token = regexp.MustCompile("^(verdadeiro|falso)+").FindString(linha)
	if len(token) > 0 {
		return token, "LOGICO", linha[len(token):]
	}

	token = regexp.MustCompile("^[A-z]+[A-zÀ-ú0-9]*").FindString(linha)
	if len(token) > 0 {
		return token, "v", linha[len(token):]
	}

	return linha, "ERRO", linha
}
