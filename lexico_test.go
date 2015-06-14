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

func TestExtraiInteiro(t *testing.T) {
	b1, v1, r1 := extraiInteiro("")
	b2, v2, r2 := extraiInteiro("123")
	b3, v3, r3 := extraiInteiro("123abc")
	b4, v4, r4 := extraiInteiro("123.45")
	b5, v5, r5 := extraiInteiro("abc123")

	if !(b1 == false && v1 == "" && r1 == "") {
		t.Errorf("b1: [%v], v1: [%v], r1: [%v]\n Experado: b1: [false], v1: [], r1: []\n", b1, v1, r1)
	}

	if !(b2 == true && v2 == "123" && r2 == "") {
		t.Errorf("b2: [%v], v2: [%v], r2: [%v]\n Experado: b2: [true], v2: [123], r2: []\n", b2, v2, r2)
	}

	if !(b3 == true && v3 == "123" && r3 == "abc") {
		t.Errorf("b3: [%v], v3: [%v], r3: [%v]\n Experado: b3: [true], v3: [123], r3: [abc]\n", b3, v3, r3)
	}

	if !(b4 == false && v4 == "" && r4 == "123.45") {
		t.Errorf("b4: [%v], v4: [%v], r4: [%v]\n Experado: b4: [false], v4: [], r4: [123.45]\n", b4, v4, r4)
	}

	if !(b5 == false && v5 == "" && r5 == "abc123") {
		t.Errorf("b5: [%v], v5: [%v], r5: [%v]\n Experado: b5: [false], v5: [], r5: [abc123]\n", b5, v5, r5)
	}
}

func TestExtraiReal(t *testing.T) {
	b1, v1, r1 := extraiReal("")
	b2, v2, r2 := extraiReal("123")
	b3, v3, r3 := extraiReal("123abc")
	b4, v4, r4 := extraiReal("123.45")
	b5, v5, r5 := extraiReal("123.45abc")
	b6, v6, r6 := extraiReal("123.45.67")
	b7, v7, r7 := extraiReal("123.45.abc")
	b8, v8, r8 := extraiReal("abc123")

	if !(b1 == false && v1 == "" && r1 == "") {
		t.Errorf("b1: [%v], v1: [%v], r1: [%v]\n Experado: b1: [false], v1: [], r1: []\n", b1, v1, r1)
	}

	if !(b2 == false && v2 == "123" && r2 == "123") {
		t.Errorf("b2: [%v], v2: [%v], r2: [%v]\n Experado: b2: [false], v2: [123], r2: [123]\n", b2, v2, r2)
	}

	if !(b3 == false && v3 == "123" && r3 == "123abc") {
		t.Errorf("b3: [%v], v3: [%v], r3: [%v]\n Experado: b3: [false], v3: [123], r3: [123abc]\n", b3, v3, r3)
	}

	if !(b4 == true && v4 == "123.45" && r4 == "") {
		t.Errorf("b4: [%v], v4: [%v], r4: [%v]\n Experado: b4: [true], v4: [123.45], r4: []\n", b4, v4, r4)
	}

	if !(b5 == true && v5 == "123.45" && r5 == "abc") {
		t.Errorf("b5: [%v], v5: [%v], r5: [%v]\n Experado: b5: [true], v5: [123.45], r5: [abc]\n", b5, v5, r5)
	}

	if !(b6 == false && v6 == "" && r6 == "123.45.67") {
		t.Errorf("b6: [%v], v6: [%v], r6: [%v]\n Experado: b6: [false], v6: [], r6: [123.45.67]\n", b6, v6, r6)
	}

	if !(b7 == false && v7 == "" && r7 == "123.45.abc") {
		t.Errorf("b7: [%v], v7: [%v], r7: [%v]\n Experado: b7: [false], v7: [], r7: [123.45.abc]\n", b7, v7, r7)
	}

	if !(b8 == false && v8 == "" && r8 == "abc123") {
		t.Errorf("b8: [%v], v8: [%v], r8: [%v]\n Experado: b8: [false], v8: [], r8: [abc123]\n", b8, v8, r8)
	}
}

func TestExtraiLogico(t *testing.T) {
	b, v, r := extraiLogico("")
	if !(b == false && v == "" && r == "") {
		t.Errorf("extraiLogico('') Experado: b[false], v[], r[] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico(" ")
	if !(b == false && v == "" && r == " ") {
		t.Errorf("extraiLogico(' ') Experado: b2[false], v[], r[ ] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("falso")
	if !(b == true && v == "falso" && r == "") {
		t.Errorf("extraiLogico('falso') Experado: b2[false], v[falso], r[] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("falso123")
	if !(b == false && v == "" && r == "falso123") {
		t.Errorf("extraiLogico('falso123') Experado: b[false], v[], r[falso123] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("falsoabcde")
	if !(b == false && v == "" && r == "falsoabcde") {
		t.Errorf("extraiLogico('falsoabcde') Experado: b[false], v[], r[falsoabcde] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("falsoAbcde")
	if !(b == false && v == "" && r == "falsoAbcde") {
		t.Errorf("extraiLogico('falsoAbcde') Experado: b[false], v[], r[falsoAbcde] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("falso#")
	if !(b == true && v == "falso" && r == "#") {
		t.Errorf("extraiLogico('falso#') Experado: b[true], v[falso], r[#] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("falso ")
	if !(b == true && v == "falso" && r == " ") {
		t.Errorf("extraiLogico('falso ') Experado: b[true], v[falso], r[ ] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("123falso")
	if !(b == false && v == "" && r == "123falso") {
		t.Errorf("extraiLogico('123falso') Experado: b[false], v[], r[123falso] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("verdadeiro")
	if !(b == true && v == "verdadeiro" && r == "") {
		t.Errorf("extraiLogico('verdadeiro') Experado: b2[true], v[verdadeiro], r[] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("verdadeiro123")
	if !(b == false && v == "" && r == "verdadeiro123") {
		t.Errorf("extraiLogico('verdadeiro123') Experado: b[false], v[], r[verdadeiro123] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("verdadeiroabcde")
	if !(b == false && v == "" && r == "verdadeiroabcde") {
		t.Errorf("extraiLogico('verdadeiroabcde') Experado: b[false], v[], r[verdadeiroabcde] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("verdadeiroAbcde")
	if !(b == false && v == "" && r == "verdadeiroAbcde") {
		t.Errorf("extraiLogico('verdadeiroAbcde') Experado: b[false], v[], r[verdadeiroAbcde] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("verdadeiro#")
	if !(b == true && v == "verdadeiro" && r == "#") {
		t.Errorf("extraiLogico('verdadeiro#') Experado: b[true], v[verdadeiro], r[#] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("verdadeiro ")
	if !(b == true && v == "verdadeiro" && r == " ") {
		t.Errorf("extraiLogico('verdadeiro ') Experado: b[true], v[verdadeiro], r[ ] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiLogico("123verdadeiro")
	if !(b == false && v == "" && r == "123verdadeiro") {
		t.Errorf("extraiLogico('123verdadeiro') Experado: b[false], v[], r[123verdadeiro] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}
}

func TestExtraiCaractere(t *testing.T) {
	b, v, r := extraiCaractere("")
	if !(b == false && v == "" && r == "") {
		t.Errorf("extraiCaractere('') Experado: b[false], v[], r[] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	b, v, r = extraiCaractere("abc123")
	if !(b == false && v == "" && r == "abc123") {
		t.Errorf("extraiCaractere('abc123') Experado: b[false], v[], r[abc123] Obtido: b[%v], v[%v], r[%v]\n \n", b, v, r)
	}

	texto := `abc: "123`
	b, v, r = extraiCaractere(texto)
	if !(b == false && v == "" && r == texto) {
		t.Errorf("extraiCaractere('%v') Experado: b[false], v[], r[%v] Obtido: b[%v], v[%v], r[%v]\n \n", texto, texto, b, v, r)
	}

	texto = `"abc: 123`
	b, v, r = extraiCaractere(texto)
	if !(b == false && v == "" && r == texto) {
		t.Errorf("extraiCaractere('%v') Experado: b[false], v[], r[%v] Obtido: b[%v], v[%v], r[%v]\n \n", texto, texto, b, v, r)
	}

	texto = `"abc: 123"`
	b, v, r = extraiCaractere(texto)
	if !(b == true && v == texto && r == "") {
		t.Errorf("extraiCaractere('%v') Experado: b[true], v[%v], r[] Obtido: b[%v], v[%v], r[%v]\n \n", texto, texto, b, v, r)
	}

	texto = `"abc: \" 123"`
	b, v, r = extraiCaractere(texto)
	if !(b == true && v == texto && r == "") {
		t.Errorf("extraiCaractere('%v') Experado: b[true], v[%v], r[] Obtido: b[%v], v[%v], r[%v]\n \n", texto, texto, b, v, r)
	}

	texto = `"abc: \" 123""`
	b, v, r = extraiCaractere(texto)
	if !(b == true && v == `"abc: \" 123"` && r == `"`) {
		t.Errorf("extraiCaractere('%v') Experado: b[false], v[\"abc: \\\" 123\"], r[\"] Obtido: b[%v], v[%v], r[%v]\n \n", texto, b, v, r)
	}
}
