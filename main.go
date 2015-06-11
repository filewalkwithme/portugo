package main

import "regexp"

func main() {

}

func verificaCaractere(token string) bool {
	matched, _ := regexp.MatchString("^\".*\"", token)
	return matched
}

func verificaInteiro(token string) bool {
	matched, _ := regexp.MatchString("^[0-9]+$", token)
	return matched
}

func verificaReal(token string) bool {
	matched, _ := regexp.MatchString("^[0-9]+,[0-9]+$", token)
	return matched
}
