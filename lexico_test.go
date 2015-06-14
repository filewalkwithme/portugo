package main

import "testing"

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

func TestExtraiConstanteInteira(t *testing.T) {
	b, v, r := extraiConstanteInteira("")
	if !(b == false && v == "" && r == "") {
		t.Errorf("extraiConstanteInteira('') Experado: b[false], v[], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteInteira("123")
	if !(b == true && v == "123" && r == "") {
		t.Errorf("extraiConstanteInteira('123') Experado: b[true], v[123], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteInteira("123abc")
	if !(b == true && v == "123" && r == "abc") {
		t.Errorf("extraiConstanteInteira('123abc') Experado: b[true], v[123], r[abc] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteInteira("123.45")
	if !(b == false && v == "" && r == "123.45") {
		t.Errorf("extraiConstanteInteira('123.45') Experado: b[false], v[], r[123.45] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteInteira("abc123")
	if !(b == false && v == "" && r == "abc123") {
		t.Errorf("extraiConstanteInteira('abc123') Experado: b[false], v[], r[abc123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}
}

func TestExtraiConstanteReal(t *testing.T) {
	b, v, r := extraiConstanteReal("")
	if !(b == false && v == "" && r == "") {
		t.Errorf("extraiConstanteReal('') Experado b[false], v[], r[] --> Obtido: b[%v], v[%v], r[%v]. \n", b, v, r)
	}

	b, v, r = extraiConstanteReal("123")
	if !(b == false && v == "123" && r == "123") {
		t.Errorf("extraiConstanteReal('123') Experado b[false], v[123], r[123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteReal("123abc")
	if !(b == false && v == "123" && r == "123abc") {
		t.Errorf("extraiConstanteReal('123abc') Experado: b[false], v[123], r[123abc] --> Obtido: b[%v], v[%v], r[%v] \n", b, v, r)
	}

	b, v, r = extraiConstanteReal("123.45")
	if !(b == true && v == "123.45" && r == "") {
		t.Errorf("extraiConstanteReal('123.45') Experado: b[true], v[123.45], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteReal("123.45abc")
	if !(b == true && v == "123.45" && r == "abc") {
		t.Errorf("extraiConstanteReal('123.45abc') Experado: b[true], v[123.45], r[abc] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteReal("123.45.67")
	if !(b == false && v == "" && r == "123.45.67") {
		t.Errorf("extraiConstanteReal('123.45.67') Experado: b[false], v[], r[123.45.67] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteReal("123.45.abc")
	if !(b == false && v == "" && r == "123.45.abc") {
		t.Errorf("extraiConstanteReal('123.45.abc') Experado: b[false], v[], r[123.45.abc] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteReal("abc123")
	if !(b == false && v == "" && r == "abc123") {
		t.Errorf("extraiConstanteReal('abc123') Experado: b[false], v[], r[abc123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}
}

func TestExtraiConstanteLogica(t *testing.T) {
	b, v, r := extraiConstanteLogica("")
	if !(b == false && v == "" && r == "") {
		t.Errorf("extraiConstanteLogica('') Experado: b[false], v[], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica(" ")
	if !(b == false && v == "" && r == " ") {
		t.Errorf("extraiConstanteLogica(' ') Experado: b[false], v[], r[ ] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("falso")
	if !(b == true && v == "falso" && r == "") {
		t.Errorf("extraiConstanteLogica('falso') Experado: b[false], v[falso], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("falso123")
	if !(b == false && v == "" && r == "falso123") {
		t.Errorf("extraiConstanteLogica('falso123') Experado: b[false], v[], r[falso123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("falsoabcde")
	if !(b == false && v == "" && r == "falsoabcde") {
		t.Errorf("extraiConstanteLogica('falsoabcde') Experado: b[false], v[], r[falsoabcde] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("falsoAbcde")
	if !(b == false && v == "" && r == "falsoAbcde") {
		t.Errorf("extraiConstanteLogica('falsoAbcde') Experado: b[false], v[], r[falsoAbcde] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("falso#")
	if !(b == true && v == "falso" && r == "#") {
		t.Errorf("extraiConstanteLogica('falso#') Experado: b[true], v[falso], r[#] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("falso ")
	if !(b == true && v == "falso" && r == " ") {
		t.Errorf("extraiConstanteLogica('falso ') Experado: b[true], v[falso], r[ ] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("123falso")
	if !(b == false && v == "" && r == "123falso") {
		t.Errorf("extraiConstanteLogica('123falso') Experado: b[false], v[], r[123falso] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("verdadeiro")
	if !(b == true && v == "verdadeiro" && r == "") {
		t.Errorf("extraiConstanteLogica('verdadeiro') Experado: b2[true], v[verdadeiro], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("verdadeiro123")
	if !(b == false && v == "" && r == "verdadeiro123") {
		t.Errorf("extraiConstanteLogica('verdadeiro123') Experado: b[false], v[], r[verdadeiro123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("verdadeiroabcde")
	if !(b == false && v == "" && r == "verdadeiroabcde") {
		t.Errorf("extraiConstanteLogica('verdadeiroabcde') Experado: b[false], v[], r[verdadeiroabcde] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("verdadeiroAbcde")
	if !(b == false && v == "" && r == "verdadeiroAbcde") {
		t.Errorf("extraiConstanteLogica('verdadeiroAbcde') Experado: b[false], v[], r[verdadeiroAbcde] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("verdadeiro#")
	if !(b == true && v == "verdadeiro" && r == "#") {
		t.Errorf("extraiConstanteLogica('verdadeiro#') Experado: b[true], v[verdadeiro], r[#] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("verdadeiro ")
	if !(b == true && v == "verdadeiro" && r == " ") {
		t.Errorf("extraiConstanteLogica('verdadeiro ') Experado: b[true], v[verdadeiro], r[ ] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiConstanteLogica("123verdadeiro")
	if !(b == false && v == "" && r == "123verdadeiro") {
		t.Errorf("extraiConstanteLogica('123verdadeiro') Experado: b[false], v[], r[123verdadeiro] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
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
		t.Errorf("extraiVariavel('') Experado: b[false], v[], r[123] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVariavel("a123")
	if !(b == true && v == "a123" && r == "") {
		t.Errorf("extraiVariavel('') Experado: b[true], v[a123], r[] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVariavel("a123+45")
	if !(b == true && v == "a123" && r == "+45") {
		t.Errorf("extraiVariavel('') Experado: b[true], v[a123], r[+45] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}

	b, v, r = extraiVariavel("123+45")
	if !(b == false && v == "" && r == "123+45") {
		t.Errorf("extraiVariavel('') Experado: b[false], v[], r[123+45] --> Obtido: b[%v], v[%v], r[%v]\n", b, v, r)
	}
}
