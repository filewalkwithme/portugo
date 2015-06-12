package main

import (
	"fmt"
	"regexp"
)

func pegaTokens(texto string) {
	texto, tipo, conteudo := pegaInteiro(texto)
	fmt.Printf("tipo = %v \n", tipo)
	fmt.Printf("conteudo = %v \n", conteudo)
	fmt.Printf("texto = %v \n", texto)
}

func pegaInteiro(texto string) (string, string, string) {
	conteudoToken := regexp.MustCompile("^[0-9]+").FindString(texto)
	if len(conteudoToken) > 0 {
		return texto[len(conteudoToken):], "INTEIRO", conteudoToken
	}
	return texto, "", ""
}

func verificaInteiro(token string) bool {
	matched, _ := regexp.MatchString("^[0-9]+$", token)
	return matched
}

func verificaReal(token string) bool {
	matched, _ := regexp.MatchString("^[0-9]+,[0-9]+$", token)
	return matched
}

func verificaCaractere(token string) bool {
	matched, _ := regexp.MatchString("^\".*\"", token)
	return matched
}

func verificaLogico(token string) bool {
	matched, _ := regexp.MatchString("^(verdadeiro|falso)$", token)
	return matched
}

func verificaVariavel(token string) bool {
	matched, _ := regexp.MatchString("^[A-z][A-z0-9]+$", token)

	return matched && !verificaLogico(token)
}
