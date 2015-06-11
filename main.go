package main

import "regexp"

func main() {

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
