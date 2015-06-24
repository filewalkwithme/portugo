package main

import (
	"fmt"
	"testing"
)

func TestVerificaDigito(t *testing.T) {
	testaLetras := []string{"a", "A", "b", "B", "ç", "Ç", "é", "É", "ó", "#", "$"}
	for _, letra := range testaLetras {
		if verificaDigito(letra) == true {
			t.Errorf("verificaDigito('%v') Experado: [false] --> Obtido: [true]\n", letra)
		}
	}

	testaNumeros := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, numero := range testaNumeros {
		if verificaDigito(numero) == false {
			t.Errorf("verificaDigito('%v') Experado: [true] --> Obtido: [false]\n", numero)
		}
	}
}

func TestVerificaLetra(t *testing.T) {
	testaLetras := []string{"a", "A", "b", "B", "ç", "Ç", "é", "É", "ó"}
	for _, letra := range testaLetras {
		if verificaLetra(letra) == false {
			t.Errorf("verificaLetra('%v') Experado: [true] --> Obtido: [false]\n", letra)
		}
	}

	testaNumeros := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, numero := range testaNumeros {
		if verificaLetra(numero) == true {
			t.Errorf("verificaLetra('%v') Experado: [false] --> Obtido: [true]\n", numero)
		}
	}

	testaOutrosSimbolos := []string{"Õ", "&", "#", "@", "%", "ü"}
	for _, simbolo := range testaOutrosSimbolos {
		if verificaLetra(simbolo) == true {
			t.Errorf("verificaLetra('%v') Experado: [false] --> Obtido: [true]\n", simbolo)
		}
	}
}

func TestVerificaLetraMaiuscula(t *testing.T) {
	testaLetrasMaiusculas := []string{"A", "B", "Ç", "É"}
	for _, letra := range testaLetrasMaiusculas {
		if verificaLetraMaiuscula(letra) == false {
			t.Errorf("verificaLetraMaiuscula('%v') Experado: [true] --> Obtido: [false]\n", letra)
		}
	}

	testaLetrasMinusculas := []string{"a", "b", "ç", "é"}
	for _, letra := range testaLetrasMinusculas {
		if verificaLetraMaiuscula(letra) == true {
			t.Errorf("verificaLetraMaiuscula('%v') Experado: [false] --> Obtido: [false]\n", letra)
		}
	}

	testaNumeros := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, numero := range testaNumeros {
		if verificaLetraMaiuscula(numero) == true {
			t.Errorf("verificaLetraMaiuscula('%v') Experado: [false] --> Obtido: [true]\n", numero)
		}
	}

	testaOutrosSimbolos := []string{"Õ", "&", "#", "@", "%", "ü"}
	for _, simbolo := range testaOutrosSimbolos {
		if verificaLetraMaiuscula(simbolo) == true {
			t.Errorf("verificaLetraMaiuscula('%v') Experado: [false] --> Obtido: [true]\n", simbolo)
		}
	}
}

func TestVerificaPalavraReservada(t *testing.T) {
	testaPalavrasReservadas := []string{"verdadeiro", "falso"}
	for _, palavra := range testaPalavrasReservadas {
		if verificaPalavraReservada(palavra) == false {
			t.Errorf("verificaPalavraReservada('%v') Experado: [true] --> Obtido: [false]\n", palavra)
		}
	}

	testaPalavrasComuns := []string{"abc", "123", "+", " ", "%", " verdadeiro", " falso123"}
	for _, palavra := range testaPalavrasComuns {
		if verificaPalavraReservada(palavra) == true {
			t.Errorf("verificaPalavraReservada('%v') Experado: [true] --> Obtido: [false]\n", palavra)
		}
	}
}

func TestExtraiConstanteInteira(t *testing.T) {
	token, r := extraiConstanteInteira("")
	if !(token.tipo == "" && token.valor == "" && r == "") {
		t.Errorf("extraiConstanteInteira('') Experado: tipo[], valor[], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteInteira("123")
	if !(token.tipo == "CONSTANTE_INTEIRA" && token.valor == "123" && r == "") {
		t.Errorf("extraiConstanteInteira('123') Experado: tipo[CONSTANTE_INTEIRA], valor[123], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteInteira("123abc")
	if !(token.tipo == "CONSTANTE_INTEIRA" && token.valor == "123" && r == "abc") {
		t.Errorf("extraiConstanteInteira('123abc') Experado: tipo[CONSTANTE_INTEIRA], valor[123], resto[abc] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteInteira("123.45")
	if !(token.tipo == "" && token.valor == "" && r == "123.45") {
		t.Errorf("extraiConstanteInteira('123.45') Experado: tipo[], valor[], resto[123.45] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteInteira("abc123")
	if !(token.tipo == "" && token.valor == "" && r == "abc123") {
		t.Errorf("extraiConstanteInteira('abc123') Experado: tipo[], valor[], resto[abc123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}
}

func TestExtraiConstanteReal(t *testing.T) {
	token, r := extraiConstanteReal("")
	if !(token.tipo == "" && token.valor == "" && r == "") {
		t.Errorf("extraiConstanteReal('') Experado tipo[], valor[], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]. \n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteReal("123")
	if !(token.tipo == "" && token.valor == "123" && r == "123") {
		t.Errorf("extraiConstanteReal('123') Experado tipo[], valor[123], resto[123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteReal("123abc")
	if !(token.tipo == "" && token.valor == "123" && r == "123abc") {
		t.Errorf("extraiConstanteReal('123abc') Experado: tipo[], valor[123], resto[123abc] --> Obtido: tipo[%v], valor[%v], resto[%v] \n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteReal("123.45")
	if !(token.tipo == "CONSTANTE_REAL" && token.valor == "123.45" && r == "") {
		t.Errorf("extraiConstanteReal('123.45') Experado: tipo[CONSTANTE_REAL], valor[123.45], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteReal("123.45abc")
	if !(token.tipo == "CONSTANTE_REAL" && token.valor == "123.45" && r == "abc") {
		t.Errorf("extraiConstanteReal('123.45abc') Experado: tipo[CONSTANTE_REAL], valor[123.45], resto[abc] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteReal("123.45.67")
	if !(token.tipo == "" && token.valor == "" && r == "123.45.67") {
		t.Errorf("extraiConstanteReal('123.45.67') Experado: tipo[], valor[], resto[123.45.67] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteReal("123.45.abc")
	if !(token.tipo == "" && token.valor == "" && r == "123.45.abc") {
		t.Errorf("extraiConstanteReal('123.45.abc') Experado: tipo[], valor[], resto[123.45.abc] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteReal("abc123")
	if !(token.tipo == "" && token.valor == "" && r == "abc123") {
		t.Errorf("extraiConstanteReal('abc123') Experado: tipo[], valor[], resto[abc123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}
}

func TestExtraiConstanteLogica(t *testing.T) {
	token, r := extraiConstanteLogica("")
	if !(token.tipo == "" && token.valor == "" && r == "") {
		t.Errorf("extraiConstanteLogica('') Experado: tipo[], valor[], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica(" ")
	if !(token.tipo == "" && token.valor == "" && r == " ") {
		t.Errorf("extraiConstanteLogica(' ') Experado: tipo[], valor[], resto[ ] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("falso")
	if !(token.tipo == "CONSTANTE_LOGICA" && token.valor == "falso" && r == "") {
		t.Errorf("extraiConstanteLogica('falso') Experado: tipo[CONSTANTE_LOGICA], valor[falso], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("falso123")
	if !(token.tipo == "" && token.valor == "" && r == "falso123") {
		t.Errorf("extraiConstanteLogica('falso123') Experado: tipo[], valor[], resto[falso123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("falsoabcde")
	if !(token.tipo == "" && token.valor == "" && r == "falsoabcde") {
		t.Errorf("extraiConstanteLogica('falsoabcde') Experado: tipo[], valor[], resto[falsoabcde] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("falsoAbcde")
	if !(token.tipo == "" && token.valor == "" && r == "falsoAbcde") {
		t.Errorf("extraiConstanteLogica('falsoAbcde') Experado: tipo[], valor[], resto[falsoAbcde] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("falso#")
	if !(token.tipo == "CONSTANTE_LOGICA" && token.valor == "falso" && r == "#") {
		t.Errorf("extraiConstanteLogica('falso#') Experado: tipo[CONSTANTE_LOGICA], valor[falso], resto[#] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("falso ")
	if !(token.tipo == "CONSTANTE_LOGICA" && token.valor == "falso" && r == " ") {
		t.Errorf("extraiConstanteLogica('falso ') Experado: tipo[CONSTANTE_LOGICA], valor[falso], resto[ ] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("123falso")
	if !(token.tipo == "" && token.valor == "" && r == "123falso") {
		t.Errorf("extraiConstanteLogica('123falso') Experado: tipo[], valor[], resto[123falso] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("verdadeiro")
	if !(token.tipo == "CONSTANTE_LOGICA" && token.valor == "verdadeiro" && r == "") {
		t.Errorf("extraiConstanteLogica('verdadeiro') Experado: tipo[CONSTANTE_LOGICA], valor[verdadeiro], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("verdadeiro123")
	if !(token.tipo == "" && token.valor == "" && r == "verdadeiro123") {
		t.Errorf("extraiConstanteLogica('verdadeiro123') Experado: tipo[], valor[], resto[verdadeiro123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("verdadeiroabcde")
	if !(token.tipo == "" && token.valor == "" && r == "verdadeiroabcde") {
		t.Errorf("extraiConstanteLogica('verdadeiroabcde') Experado: tipo[], valor[], resto[verdadeiroabcde] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("verdadeiroAbcde")
	if !(token.tipo == "" && token.valor == "" && r == "verdadeiroAbcde") {
		t.Errorf("extraiConstanteLogica('verdadeiroAbcde') Experado: tipo[], valor[], resto[verdadeiroAbcde] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("verdadeiro#")
	if !(token.tipo == "CONSTANTE_LOGICA" && token.valor == "verdadeiro" && r == "#") {
		t.Errorf("extraiConstanteLogica('verdadeiro#') Experado: tipo[CONSTANTE_LOGICA], valor[verdadeiro], resto[#] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("verdadeiro ")
	if !(token.tipo == "CONSTANTE_LOGICA" && token.valor == "verdadeiro" && r == " ") {
		t.Errorf("extraiConstanteLogica('verdadeiro ') Experado: tipo[CONSTANTE_LOGICA], valor[verdadeiro], resto[ ] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("123verdadeiro")
	if !(token.tipo == "" && token.valor == "" && r == "123verdadeiro") {
		t.Errorf("extraiConstanteLogica('123verdadeiro') Experado: tipo[], valor[], resto[123verdadeiro] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}
}

func TestExtraiConstanteCaractere(t *testing.T) {
	token, r := extraiConstanteCaractere("")
	if !(token.tipo == "" && token.valor == "" && r == "") {
		t.Errorf("extraiConstanteCaractere('') Experado: tipo[], valor[], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteCaractere("abc123")
	if !(token.tipo == "" && token.valor == "" && r == "abc123") {
		t.Errorf("extraiConstanteCaractere('abc123') Experado: tipo[], valor[], resto[abc123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	texto := `abc: "123`
	token, r = extraiConstanteCaractere(texto)
	if !(token.tipo == "" && token.valor == "" && r == texto) {
		t.Errorf("extraiConstanteCaractere('%v') Experado: tipo[], valor[], resto[%v] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", texto, texto, token.tipo, token.valor, r)
	}

	texto = `"abc: 123`
	token, r = extraiConstanteCaractere(texto)
	if !(token.tipo == "" && token.valor == "" && r == texto) {
		t.Errorf("extraiConstanteCaractere('%v') Experado: tipo[], valor[], resto[%v] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", texto, texto, token.tipo, token.valor, r)
	}

	texto = `"abc: 123"`
	token, r = extraiConstanteCaractere(texto)
	if !(token.tipo == "CONSTANTE_CARACTERE" && token.valor == texto && r == "") {
		t.Errorf("extraiConstanteCaractere('%v') Experado: tipo[CONSTANTE_CARACTERE], valor[%v], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", texto, texto, token.tipo, token.valor, r)
	}

	texto = `"abc: \" 123"`
	token, r = extraiConstanteCaractere(texto)
	if !(token.tipo == "CONSTANTE_CARACTERE" && token.valor == texto && r == "") {
		t.Errorf("extraiConstanteCaractere('%v') Experado: tipo[CONSTANTE_CARACTERE], valor[%v], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", texto, texto, token.tipo, token.valor, r)
	}

	texto = `"abc: \" 123""`
	token, r = extraiConstanteCaractere(texto)
	if !(token.tipo == "CONSTANTE_CARACTERE" && token.valor == `"abc: \" 123"` && r == `"`) {
		t.Errorf("extraiConstanteCaractere('%v') Experado: tipo[CONSTANTE_CARACTERE], valor[\"abc: \\\" 123\"], resto[\"] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", texto, token.tipo, token.valor, r)
	}
}

func TestExtraiVariavel(t *testing.T) {
	token, r := extraiVariavel("")
	if !(token.tipo == "" && token.valor == "" && r == "") {
		t.Errorf("extraiVariavel('') Experado: tipo[], valor[], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiVariavel("123")
	if !(token.tipo == "" && token.valor == "" && r == "123") {
		t.Errorf("extraiVariavel('123') Experado: tipo[], valor[], resto[123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiVariavel("A123")
	if !(token.tipo == "VARIAVEL" && token.valor == "A123" && r == "") {
		t.Errorf("extraiVariavel('A123') Experado: tipo[VARIAVEL], valor[A123], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiVariavel("A123+45")
	if !(token.tipo == "VARIAVEL" && token.valor == "A123" && r == "+45") {
		t.Errorf("extraiVariavel('A123+45') Experado: tipo[VARIAVEL], valor[A123], resto[+45] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiVariavel("a123+45")
	if !(token.tipo == "" && token.valor == "" && r == "a123+45") {
		t.Errorf("extraiVariavel('a123+45') Experado: tipo[], valor[], resto[a123+45] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiVariavel("123+45")
	if !(token.tipo == "" && token.valor == "" && r == "123+45") {
		t.Errorf("extraiVariavel('123+45') Experado: tipo[], valor[], resto[123+45] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiVariavel("verdadeiro")
	if !(token.tipo == "" && token.valor == "" && r == "verdadeiro") {
		t.Errorf("extraiVariavel('verdadeiro') Experado: tipo[], valor[], resto[verdadeiro] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}
}

func TestExtraiTipoVariavel(t *testing.T) {
	token, r := extraiTipoVariavel("")
	if !(token.tipo == "" && token.valor == "" && r == "") {
		t.Errorf("extraiTipoVariavel('') Experado: tipo[], valor[], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("real")
	if !(token.tipo == "TIPO_VARIAVEL" && token.valor == "real" && r == "") {
		t.Errorf("extraiTipoVariavel('real') Experado: tipo[TIPO_VARIAVEL], valor[real], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("realabc")
	if !(token.tipo == "" && token.valor == "" && r == "realabc") {
		t.Errorf("extraiTipoVariavel('realabc') Experado: tipo[], valor[], resto[realabc] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("real123")
	if !(token.tipo == "" && token.valor == "" && r == "real123") {
		t.Errorf("extraiTipoVariavel('real123') Experado: tipo[], valor[], resto[real123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("real: x")
	if !(token.tipo == "TIPO_VARIAVEL" && token.valor == "real" && r == ": x") {
		t.Errorf("extraiTipoVariavel('real: x') Experado: tipo[TIPO_VARIAVEL], valor[real], resto[: x] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("inteiro")
	if !(token.tipo == "TIPO_VARIAVEL" && token.valor == "inteiro" && r == "") {
		t.Errorf("extraiTipoVariavel('inteiro') Experado: tipo[TIPO_VARIAVEL], valor[inteiro], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("inteiroabc")
	if !(token.tipo == "" && token.valor == "" && r == "inteiroabc") {
		t.Errorf("extraiTipoVariavel('inteiroabc') Experado: tipo[], valor[], resto[inteiroabc] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("inteiro123")
	if !(token.tipo == "" && token.valor == "" && r == "inteiro123") {
		t.Errorf("extraiTipoVariavel('inteiro123') Experado: tipo[], valor[], resto[inteiro123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("inteiro: x")
	if !(token.tipo == "TIPO_VARIAVEL" && token.valor == "inteiro" && r == ": x") {
		t.Errorf("extraiTipoVariavel('inteiro: x') Experado: tipo[TIPO_VARIAVEL], valor[inteiro], resto[: x] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("lógico")
	if !(token.tipo == "TIPO_VARIAVEL" && token.valor == "lógico" && r == "") {
		fmt.Printf("--->%v<---", len("lógico"))
		t.Errorf("extraiTipoVariavel('lógico') Experado: tipo[TIPO_VARIAVEL], valor[lógico], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("lógicoabc")
	if !(token.tipo == "" && token.valor == "" && r == "lógicoabc") {
		t.Errorf("extraiTipoVariavel('lógicoabc') Experado: tipo[], valor[], resto[lógicoabc] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("lógico123")
	if !(token.tipo == "" && token.valor == "" && r == "lógico123") {
		t.Errorf("extraiTipoVariavel('lógico123') Experado: tipo[], valor[], resto[lógico123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("lógico: x")
	if !(token.tipo == "TIPO_VARIAVEL" && token.valor == "lógico" && r == ": x") {
		t.Errorf("extraiTipoVariavel('lógico: x') Experado: tipo[TIPO_VARIAVEL], valor[lógico], resto[: x] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("caractere")
	if !(token.tipo == "TIPO_VARIAVEL" && token.valor == "caractere" && r == "") {
		t.Errorf("extraiTipoVariavel('caractere') Experado: tipo[TIPO_VARIAVEL], valor[caractere], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("caractereabc")
	if !(token.tipo == "" && token.valor == "" && r == "caractereabc") {
		t.Errorf("extraiTipoVariavel('caractereabc') Experado: tipo[], valor[], resto[caractereabc] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("caractere123")
	if !(token.tipo == "" && token.valor == "" && r == "caractere123") {
		t.Errorf("extraiTipoVariavel('caractere123') Experado: tipo[], valor[], resto[caractere123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiTipoVariavel("caractere: x")
	if !(token.tipo == "TIPO_VARIAVEL" && token.valor == "caractere" && r == ": x") {
		t.Errorf("extraiTipoVariavel('caractere: x') Experado: tipo[TIPO_VARIAVEL], valor[caractere], resto[: x] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}
}

func TestExtraiDoisPontos(t *testing.T) {
	token, r := extraiDoisPontos("")
	if !(token.tipo == "" && token.valor == "" && r == "") {
		t.Errorf("extraiDoisPontos('') Experado: tipo[], valor[], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiDoisPontos(":")
	if !(token.tipo == "DOIS_PONTOS" && token.valor == ":" && r == "") {
		t.Errorf("extraiDoisPontos(':') Experado: tipo[DOIS_PONTOS], valor[:], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiDoisPontos(":123")
	if !(token.tipo == "DOIS_PONTOS" && token.valor == ":" && r == "123") {
		t.Errorf("extraiDoisPontos(':123') Experado: tipo[DOIS_PONTOS], valor[:], resto[123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiDoisPontos(":abc")
	if !(token.tipo == "DOIS_PONTOS" && token.valor == ":" && r == "abc") {
		t.Errorf("extraiDoisPontos(':abc') Experado: tipo[DOIS_PONTOS], valor[:], resto[abc] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiDoisPontos(":#")
	if !(token.tipo == "DOIS_PONTOS" && token.valor == ":" && r == "#") {
		t.Errorf("extraiDoisPontos(':#') Experado: tipo[DOIS_PONTOS], valor[:], resto[#] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiDoisPontos("123:")
	if !(token.tipo == "" && token.valor == "" && r == "123:") {
		t.Errorf("extraiDoisPontos('123:') Experado: tipo[], valor[], resto[123:] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiDoisPontos("abc:")
	if !(token.tipo == "" && token.valor == "" && r == "abc:") {
		t.Errorf("extraiDoisPontos('abc:') Experado: tipo[], valor[], resto[abc:] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiDoisPontos("#")
	if !(token.tipo == "" && token.valor == "" && r == "#") {
		t.Errorf("extraiDoisPontos('#') Experado: tipo[], valor[], resto[#] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}
}

func TestExtraiPontoEVirgula(t *testing.T) {
	token, r := extraiPontoEVirgula("")
	if !(token.tipo == "" && token.valor == "" && r == "") {
		t.Errorf("extraiPontoEVirgula('') Experado: tipo[], valor[], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiPontoEVirgula(";")
	if !(token.tipo == "PONTO_E_VIRGULA" && token.valor == ";" && r == "") {
		t.Errorf("extraiPontoEVirgula(';') Experado: tipo[PONTO_E_VIRGULA], valor[;], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiPontoEVirgula(";123")
	if !(token.tipo == "PONTO_E_VIRGULA" && token.valor == ";" && r == "123") {
		t.Errorf("extraiPontoEVirgula(';123') Experado: tipo[PONTO_E_VIRGULA], valor[;], resto[123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiPontoEVirgula(";abc")
	if !(token.tipo == "PONTO_E_VIRGULA" && token.valor == ";" && r == "abc") {
		t.Errorf("extraiPontoEVirgula(';abc') Experado: tipo[PONTO_E_VIRGULA], valor[;], resto[abc] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiPontoEVirgula(";#")
	if !(token.tipo == "PONTO_E_VIRGULA" && token.valor == ";" && r == "#") {
		t.Errorf("extraiPontoEVirgula(';#') Experado: tipo[PONTO_E_VIRGULA], valor[;], resto[#] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiPontoEVirgula("123;")
	if !(token.tipo == "" && token.valor == "" && r == "123;") {
		t.Errorf("extraiPontoEVirgula('123;') Experado: tipo[], valor[], resto[123;] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiPontoEVirgula("abc;")
	if !(token.tipo == "" && token.valor == "" && r == "abc;") {
		t.Errorf("extraiPontoEVirgula('abc;') Experado: tipo[], valor[], resto[abc;] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiPontoEVirgula("#")
	if !(token.tipo == "" && token.valor == "" && r == "#") {
		t.Errorf("extraiPontoEVirgula('#') Experado: tipo[], valor[], resto[#] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}
}

func TestExtraiVirgula(t *testing.T) {
	token, r := extraiVirgula("")
	if !(token.tipo == "" && token.valor == "" && r == "") {
		t.Errorf("extraiVirgula('') Experado: tipo[], valor[], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiVirgula(",")
	if !(token.tipo == "VIRGULA" && token.valor == "," && r == "") {
		t.Errorf("extraiVirgula(',') Experado: tipo[VIRGULA], valor[,], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiVirgula(",123")
	if !(token.tipo == "VIRGULA" && token.valor == "," && r == "123") {
		t.Errorf("extraiVirgula(',123') Experado: tipo[VIRGULA], valor[,], resto[123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiVirgula(",abc")
	if !(token.tipo == "VIRGULA" && token.valor == "," && r == "abc") {
		t.Errorf("extraiVirgula(',abc') Experado: tipo[VIRGULA], valor[,], resto[abc] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiVirgula(",#")
	if !(token.tipo == "VIRGULA" && token.valor == "," && r == "#") {
		t.Errorf("extraiVirgula(',#') Experado: tipo[VIRGULA], valor[,], resto[#] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiVirgula("123,")
	if !(token.tipo == "" && token.valor == "" && r == "123,") {
		t.Errorf("extraiVirgula('123,') Experado: tipo[], valor[], resto[123,] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiVirgula("abc,")
	if !(token.tipo == "" && token.valor == "" && r == "abc,") {
		t.Errorf("extraiVirgula('abc,') Experado: tipo[], valor[], resto[abc,] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiVirgula("#")
	if !(token.tipo == "" && token.valor == "" && r == "#") {
		t.Errorf("extraiVirgula('#') Experado: tipo[], valor[], resto[#] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}
}

func TestExtraiEspaco(t *testing.T) {
	token, r := extraiEspaco("")
	if !(token.tipo == "" && token.valor == "" && r == "") {
		t.Errorf("extraiEspaco('') Experado: tipo[], valor[], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiEspaco(" ")
	if !(token.tipo == "ESPACO" && token.valor == " " && r == "") {
		t.Errorf("extraiEspaco(' ') Experado: tipo[ESPACO], valor[ ], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiEspaco(" 123")
	if !(token.tipo == "ESPACO" && token.valor == " " && r == "123") {
		t.Errorf("extraiEspaco(' 123') Experado: tipo[ESPACO], valor[ ], resto[123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiEspaco(" abc")
	if !(token.tipo == "ESPACO" && token.valor == " " && r == "abc") {
		t.Errorf("extraiEspaco(' abc') Experado: tipo[ESPACO], valor[ ], resto[abc] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiEspaco(" #")
	if !(token.tipo == "ESPACO" && token.valor == " " && r == "#") {
		t.Errorf("extraiEspaco(' #') Experado: tipo[ESPACO], valor[ ], resto[#] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiEspaco("123 ")
	if !(token.tipo == "" && token.valor == "" && r == "123 ") {
		t.Errorf("extraiEspaco('123 ') Experado: tipo[], valor[], resto[123 ] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiEspaco("abc ")
	if !(token.tipo == "" && token.valor == "" && r == "abc ") {
		t.Errorf("extraiEspaco('abc ') Experado: tipo[], valor[], resto[abc ] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiEspaco("#")
	if !(token.tipo == "" && token.valor == "" && r == "#") {
		t.Errorf("extraiEspaco('#') Experado: tipo[], valor[], resto[#] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}
}
