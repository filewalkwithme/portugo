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
		t.Errorf("extraiConstanteInteira('abc123') Experado: tipo[CONSTANTE_INTEIRA], valor[], resto[abc123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}
}

func TestExtraiConstanteReal(t *testing.T) {
	token, r := extraiConstanteReal("")
	if !(token.tipo == "" && token.valor == "" && r == "") {
		t.Errorf("extraiConstanteReal('') Experado tipo[], valor[], resto[] --> Obtido: b[%v], v[%v], r[%v]. \n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteReal("123")
	if !(token.tipo == "" && token.valor == "123" && r == "123") {
		t.Errorf("extraiConstanteReal('123') Experado tipo[], valor[123], resto[123] --> Obtido: b[%v], v[%v], r[%v]\n", token.tipo, token.valor, r)
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
		t.Errorf("extraiConstanteLogica('') Experado: tipo[false], valor[], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica(" ")
	if !(token.tipo == "" && token.valor == "" && r == " ") {
		t.Errorf("extraiConstanteLogica(' ') Experado: tipo[false], valor[], resto[ ] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("falso")
	if !(token.tipo == "CONSTANTE_LOGICA" && token.valor == "falso" && r == "") {
		t.Errorf("extraiConstanteLogica('falso') Experado: tipo[false], valor[falso], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("falso123")
	if !(token.tipo == "" && token.valor == "" && r == "falso123") {
		t.Errorf("extraiConstanteLogica('falso123') Experado: tipo[false], valor[], resto[falso123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("falsoabcde")
	if !(token.tipo == "" && token.valor == "" && r == "falsoabcde") {
		t.Errorf("extraiConstanteLogica('falsoabcde') Experado: tipo[false], valor[], resto[falsoabcde] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("falsoAbcde")
	if !(token.tipo == "" && token.valor == "" && r == "falsoAbcde") {
		t.Errorf("extraiConstanteLogica('falsoAbcde') Experado: tipo[false], valor[], resto[falsoAbcde] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("falso#")
	if !(token.tipo == "CONSTANTE_LOGICA" && token.valor == "falso" && r == "#") {
		t.Errorf("extraiConstanteLogica('falso#') Experado: tipo[true], valor[falso], resto[#] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("falso ")
	if !(token.tipo == "CONSTANTE_LOGICA" && token.valor == "falso" && r == " ") {
		t.Errorf("extraiConstanteLogica('falso ') Experado: tipo[true], valor[falso], resto[ ] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("123falso")
	if !(token.tipo == "" && token.valor == "" && r == "123falso") {
		t.Errorf("extraiConstanteLogica('123falso') Experado: tipo[false], valor[], resto[123falso] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("verdadeiro")
	if !(token.tipo == "CONSTANTE_LOGICA" && token.valor == "verdadeiro" && r == "") {
		t.Errorf("extraiConstanteLogica('verdadeiro') Experado: b2[true], valor[verdadeiro], resto[] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("verdadeiro123")
	if !(token.tipo == "" && token.valor == "" && r == "verdadeiro123") {
		t.Errorf("extraiConstanteLogica('verdadeiro123') Experado: tipo[false], valor[], resto[verdadeiro123] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("verdadeiroabcde")
	if !(token.tipo == "" && token.valor == "" && r == "verdadeiroabcde") {
		t.Errorf("extraiConstanteLogica('verdadeiroabcde') Experado: tipo[false], valor[], resto[verdadeiroabcde] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("verdadeiroAbcde")
	if !(token.tipo == "" && token.valor == "" && r == "verdadeiroAbcde") {
		t.Errorf("extraiConstanteLogica('verdadeiroAbcde') Experado: tipo[false], valor[], resto[verdadeiroAbcde] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("verdadeiro#")
	if !(token.tipo == "CONSTANTE_LOGICA" && token.valor == "verdadeiro" && r == "#") {
		t.Errorf("extraiConstanteLogica('verdadeiro#') Experado: tipo[true], valor[verdadeiro], resto[#] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("verdadeiro ")
	if !(token.tipo == "CONSTANTE_LOGICA" && token.valor == "verdadeiro" && r == " ") {
		t.Errorf("extraiConstanteLogica('verdadeiro ') Experado: tipo[true], valor[verdadeiro], resto[ ] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}

	token, r = extraiConstanteLogica("123verdadeiro")
	if !(token.tipo == "" && token.valor == "" && r == "123verdadeiro") {
		t.Errorf("extraiConstanteLogica('123verdadeiro') Experado: tipo[false], valor[], resto[123verdadeiro] --> Obtido: tipo[%v], valor[%v], resto[%v]\n", token.tipo, token.valor, r)
	}
}

func TestExtraiConstanteCaractere(t *testing.T) {
	b, v, r := extraiConstanteCaractere("")
	if !(b == false && v == "" && r == "") {
		t.Errorf("extraiConstanteCaractere('') Experado: b[false], v[], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteCaractere("abc123")
	if !(b == false && v == "" && r == "abc123") {
		t.Errorf("extraiConstanteCaractere('abc123') Experado: b[false], v[], r[abc123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	texto := `abc: "123`
	b, v, r = extraiConstanteCaractere(texto)
	if !(b == false && v == "" && r == texto) {
		t.Errorf("extraiConstanteCaractere('%v') Experado: b[false], v[], r[%v] --> Obtido: b[%v], v[%v], r[%v]\n", texto, texto, b, v, r)
	}

	texto = `"abc: 123`
	b, v, r = extraiConstanteCaractere(texto)
	if !(b == false && v == "" && r == texto) {
		t.Errorf("extraiConstanteCaractere('%v') Experado: b[false], v[], r[%v] --> Obtido: b[%v], v[%v], r[%v]\n", texto, texto, b, v, r)
	}

	texto = `"abc: 123"`
	b, v, r = extraiConstanteCaractere(texto)
	if !(b == true && v == texto && r == "") {
		t.Errorf("extraiConstanteCaractere('%v') Experado: b[true], v[%v], r[] --> Obtido: b[%v], v[%v], r[%v]\n", texto, texto, b, v, r)
	}

	texto = `"abc: \" 123"`
	b, v, r = extraiConstanteCaractere(texto)
	if !(b == true && v == texto && r == "") {
		t.Errorf("extraiConstanteCaractere('%v') Experado: b[true], v[%v], r[] --> Obtido: b[%v], v[%v], r[%v]\n", texto, texto, b, v, r)
	}

	texto = `"abc: \" 123""`
	b, v, r = extraiConstanteCaractere(texto)
	if !(b == true && v == `"abc: \" 123"` && r == `"`) {
		t.Errorf("extraiConstanteCaractere('%v') Experado: b[false], v[\"abc: \\\" 123\"], r[\"] --> Obtido: b[%v], v[%v], r[%v]\n", texto, b, v, r)
	}
}

func TestExtraiVariavel(t *testing.T) {
	b, v, r := extraiVariavel("")
	if !(b == false && v == "" && r == "") {
		t.Errorf("extraiVariavel('') Experado: b[false], v[], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVariavel("123")
	if !(b == false && v == "" && r == "123") {
		t.Errorf("extraiVariavel('123') Experado: b[false], v[], r[123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVariavel("A123")
	if !(b == true && v == "A123" && r == "") {
		t.Errorf("extraiVariavel('A123') Experado: b[true], v[A123], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVariavel("A123+45")
	if !(b == true && v == "A123" && r == "+45") {
		t.Errorf("extraiVariavel('A123+45') Experado: b[true], v[A123], r[+45] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVariavel("a123+45")
	if !(b == false && v == "" && r == "a123+45") {
		t.Errorf("extraiVariavel('a123+45') Experado: b[false], v[], r[a123+45] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVariavel("123+45")
	if !(b == false && v == "" && r == "123+45") {
		t.Errorf("extraiVariavel('123+45') Experado: b[false], v[], r[123+45] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVariavel("verdadeiro")
	if !(b == false && v == "" && r == "verdadeiro") {
		t.Errorf("extraiVariavel('verdadeiro') Experado: b[false], v[], r[verdadeiro] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}
}

func TestExtraiTipoVariavel(t *testing.T) {
	b, v, r := extraiTipoVariavel("")
	if !(b == false && v == "" && r == "") {
		t.Errorf("extraiTipoVariavel('') Experado: b[false], v[], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("real")
	if !(b == true && v == "real" && r == "") {
		t.Errorf("extraiTipoVariavel('real') Experado: b[true], v[real], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("realabc")
	if !(b == false && v == "" && r == "realabc") {
		t.Errorf("extraiTipoVariavel('realabc') Experado: b[false], v[], r[realabc] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("real123")
	if !(b == false && v == "" && r == "real123") {
		t.Errorf("extraiTipoVariavel('real123') Experado: b[false], v[], r[real123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("real: x")
	if !(b == true && v == "real" && r == ": x") {
		t.Errorf("extraiTipoVariavel('real: x') Experado: b[true], v[real], r[: x] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("inteiro")
	if !(b == true && v == "inteiro" && r == "") {
		t.Errorf("extraiTipoVariavel('inteiro') Experado: b[true], v[inteiro], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("inteiroabc")
	if !(b == false && v == "" && r == "inteiroabc") {
		t.Errorf("extraiTipoVariavel('inteiroabc') Experado: b[false], v[], r[inteiroabc] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("inteiro123")
	if !(b == false && v == "" && r == "inteiro123") {
		t.Errorf("extraiTipoVariavel('inteiro123') Experado: b[false], v[], r[inteiro123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("inteiro: x")
	if !(b == true && v == "inteiro" && r == ": x") {
		t.Errorf("extraiTipoVariavel('inteiro: x') Experado: b[true], v[inteiro], r[: x] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("lógico")
	if !(b == true && v == "lógico" && r == "") {
		fmt.Printf("--->%v<---", len("lógico"))
		t.Errorf("extraiTipoVariavel('lógico') Experado: b[true], v[lógico], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("lógicoabc")
	if !(b == false && v == "" && r == "lógicoabc") {
		t.Errorf("extraiTipoVariavel('lógicoabc') Experado: b[false], v[], r[lógicoabc] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("lógico123")
	if !(b == false && v == "" && r == "lógico123") {
		t.Errorf("extraiTipoVariavel('lógico123') Experado: b[false], v[], r[lógico123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("lógico: x")
	if !(b == true && v == "lógico" && r == ": x") {
		t.Errorf("extraiTipoVariavel('lógico: x') Experado: b[true], v[lógico], r[: x] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("caractere")
	if !(b == true && v == "caractere" && r == "") {
		t.Errorf("extraiTipoVariavel('caractere') Experado: b[true], v[caractere], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("caractereabc")
	if !(b == false && v == "" && r == "caractereabc") {
		t.Errorf("extraiTipoVariavel('caractereabc') Experado: b[false], v[], r[caractereabc] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("caractere123")
	if !(b == false && v == "" && r == "caractere123") {
		t.Errorf("extraiTipoVariavel('caractere123') Experado: b[false], v[], r[caractere123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiTipoVariavel("caractere: x")
	if !(b == true && v == "caractere" && r == ": x") {
		t.Errorf("extraiTipoVariavel('caractere: x') Experado: b[true], v[caractere], r[: x] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}
}

func TestExtraiDoisPontos(t *testing.T) {
	b, v, r := extraiDoisPontos("")
	if !(b == false && v == "" && r == "") {
		t.Errorf("extraiDoisPontos('') Experado: b[false], v[], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiDoisPontos(":")
	if !(b == true && v == ":" && r == "") {
		t.Errorf("extraiDoisPontos(':') Experado: b[true], v[:], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiDoisPontos(":123")
	if !(b == true && v == ":" && r == "123") {
		t.Errorf("extraiDoisPontos(':123') Experado: b[true], v[:], r[123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiDoisPontos(":abc")
	if !(b == true && v == ":" && r == "abc") {
		t.Errorf("extraiDoisPontos(':abc') Experado: b[true], v[:], r[abc] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiDoisPontos(":#")
	if !(b == true && v == ":" && r == "#") {
		t.Errorf("extraiDoisPontos(':#') Experado: b[true], v[:], r[#] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiDoisPontos("123:")
	if !(b == false && v == "" && r == "123:") {
		t.Errorf("extraiDoisPontos('123:') Experado: b[false], v[], r[123:] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiDoisPontos("abc:")
	if !(b == false && v == "" && r == "abc:") {
		t.Errorf("extraiDoisPontos('abc:') Experado: b[false], v[], r[abc:] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiDoisPontos("#")
	if !(b == false && v == "" && r == "#") {
		t.Errorf("extraiDoisPontos('#') Experado: b[false], v[], r[#] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}
}

func TestExtraiPontoEVirgula(t *testing.T) {
	b, v, r := extraiPontoEVirgula("")
	if !(b == false && v == "" && r == "") {
		t.Errorf("extraiPontoEVirgula('') Experado: b[false], v[], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiPontoEVirgula(";")
	if !(b == true && v == ";" && r == "") {
		t.Errorf("extraiPontoEVirgula(';') Experado: b[true], v[;], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiPontoEVirgula(";123")
	if !(b == true && v == ";" && r == "123") {
		t.Errorf("extraiPontoEVirgula(';123') Experado: b[true], v[;], r[123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiPontoEVirgula(";abc")
	if !(b == true && v == ";" && r == "abc") {
		t.Errorf("extraiPontoEVirgula(';abc') Experado: b[true], v[;], r[abc] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiPontoEVirgula(";#")
	if !(b == true && v == ";" && r == "#") {
		t.Errorf("extraiPontoEVirgula(';#') Experado: b[true], v[;], r[#] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiPontoEVirgula("123;")
	if !(b == false && v == "" && r == "123;") {
		t.Errorf("extraiPontoEVirgula('123;') Experado: b[false], v[], r[123;] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiPontoEVirgula("abc;")
	if !(b == false && v == "" && r == "abc;") {
		t.Errorf("extraiPontoEVirgula('abc;') Experado: b[false], v[], r[abc;] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiPontoEVirgula("#")
	if !(b == false && v == "" && r == "#") {
		t.Errorf("extraiPontoEVirgula('#') Experado: b[false], v[], r[#] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}
}

func TestExtraiVirgula(t *testing.T) {
	b, v, r := extraiVirgula("")
	if !(b == false && v == "" && r == "") {
		t.Errorf("extraiVirgula('') Experado: b[false], v[], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVirgula(",")
	if !(b == true && v == "," && r == "") {
		t.Errorf("extraiVirgula(',') Experado: b[true], v[,], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVirgula(",123")
	if !(b == true && v == "," && r == "123") {
		t.Errorf("extraiVirgula(',123') Experado: b[true], v[,], r[123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVirgula(",abc")
	if !(b == true && v == "," && r == "abc") {
		t.Errorf("extraiVirgula(',abc') Experado: b[true], v[,], r[abc] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVirgula(",#")
	if !(b == true && v == "," && r == "#") {
		t.Errorf("extraiVirgula(',#') Experado: b[true], v[,], r[#] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVirgula("123,")
	if !(b == false && v == "" && r == "123,") {
		t.Errorf("extraiVirgula('123,') Experado: b[false], v[], r[123,] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVirgula("abc,")
	if !(b == false && v == "" && r == "abc,") {
		t.Errorf("extraiVirgula('abc,') Experado: b[false], v[], r[abc,] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVirgula("#")
	if !(b == false && v == "" && r == "#") {
		t.Errorf("extraiVirgula('#') Experado: b[false], v[], r[#] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}
}
