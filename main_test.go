package main

import "testing"

func TestVerificaCaractere(t *testing.T) {
	vInteiro := "123"
	vReal := "123,47"
	vCaractere := "\"Absf cde\""
	vLogicoVerdadeiro := "verdadeiro"
	vLogicoFalso := "falso"

	if verificaCaractere(vInteiro) {
		t.Errorf("Token do tipo inteiro reconhecido como caractere: %v", vInteiro)
	}

	if verificaCaractere(vReal) {
		t.Errorf("Token do tipo real reconhecido como caractere: %v", vReal)
	}

	if verificaCaractere(vCaractere) == false {
		t.Errorf("Token do tipo caractere não reconhecido: %v", vCaractere)
	}

	if verificaCaractere(vLogicoVerdadeiro) {
		t.Errorf("Token do tipo inteiro não reconhecido como caractere: %v", vLogicoVerdadeiro)
	}

	if verificaCaractere(vLogicoFalso) {
		t.Errorf("Token do tipo inteiro não reconhecido como caractere: %v", vLogicoFalso)
	}
}

func TestVerificaInteiro(t *testing.T) {
	vInteiro := "123"
	vReal := "123,47"
	vCaractere := "\"Abcde\""
	vLogicoVerdadeiro := "verdadeiro"
	vLogicoFalso := "falso"

	if verificaInteiro(vInteiro) == false {
		t.Errorf("Token do tipo inteiro não reconhecido: %v", vInteiro)
	}

	if verificaInteiro(vReal) {
		t.Errorf("Token do tipo real reconhecido como inteiro: %v", vReal)
	}

	if verificaInteiro(vCaractere) {
		t.Errorf("Token do tipo caractere reconhecido como inteiro: %v", vCaractere)
	}

	if verificaInteiro(vLogicoVerdadeiro) {
		t.Errorf("Token do tipo inteiro não reconhecido como inteiro: %v", vLogicoVerdadeiro)
	}

	if verificaInteiro(vLogicoFalso) {
		t.Errorf("Token do tipo inteiro não reconhecido como inteiro: %v", vLogicoFalso)
	}
}

func TestVerificaReal(t *testing.T) {
	vInteiro := "123"
	vReal := "123,47"
	vCaractere := "\"Abcde\""
	vLogicoVerdadeiro := "verdadeiro"
	vLogicoFalso := "falso"

	if verificaReal(vInteiro) {
		t.Errorf("Token do tipo inteiro reconhecido como real: %v", vInteiro)
	}

	if verificaReal(vReal) == false {
		t.Errorf("Token do tipo real não reconhecido: %v", vReal)
	}

	if verificaReal(vCaractere) {
		t.Errorf("Token do tipo caractere reconhecido como real: %v", vCaractere)
	}

	if verificaReal(vLogicoVerdadeiro) {
		t.Errorf("Token do tipo inteiro não reconhecido como real: %v", vLogicoVerdadeiro)
	}

	if verificaReal(vLogicoFalso) {
		t.Errorf("Token do tipo inteiro não reconhecido como real: %v", vLogicoFalso)
	}
}
