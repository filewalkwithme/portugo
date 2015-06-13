package main

import (
	"fmt"
	"testing"
)

func TestExtraiInteiro(t *testing.T) {
	b1, v1, r1 := extraiInteiro("")

	b2, v2, r2 := extraiInteiro("123")
	fmt.Printf("%v, %v, %v\n", b2, v2, r2)

	b3, v3, r3 := extraiInteiro("123abc")
	fmt.Printf("%v, %v, %v\n", b3, v3, r3)

	b4, v4, r4 := extraiInteiro("123.45")
	fmt.Printf("%v, %v, %v\n", b4, v4, r4)

	b5, v5, r5 := extraiInteiro("abc123")

	if !(b1 == false && v1 == "" && r1 == "") {
		t.Errorf("b1: [%v], v1: [%v], r1: [%v]\n Expected: b1: [false], v1: [], r1: []\n", b1, v1, r1)
	}

	if !(b2 == true && v2 == "123" && r2 == "") {
		t.Errorf("b2: [%v], v2: [%v], r2: [%v]\n Expected: b2: [true], v2: [123], r2: []\n", b2, v2, r2)
	}

	if !(b3 == true && v3 == "123" && r3 == "abc") {
		t.Errorf("b3: [%v], v3: [%v], r3: [%v]\n Expected: b3: [true], v3: [123], r3: [abc]\n", b3, v3, r3)
	}

	if !(b4 == false && v4 == "" && r4 == "123.45") {
		t.Errorf("b4: [%v], v4: [%v], r4: [%v]\n Expected: b4: [false], v4: [], r4: [123.45]\n", b4, v4, r4)
	}

	if !(b5 == false && v5 == "" && r5 == "abc123") {
		t.Errorf("b5: [%v], v5: [%v], r5: [%v]\n Expected: b5: [false], v5: [], r5: [abc123]\n", b5, v5, r5)
	}
}

func TestExtraiReal(t *testing.T) {
	b1, v1, r1 := extraiInteiro("")

	b2, v2, r2 := extraiInteiro("123")
	fmt.Printf("%v, %v, %v\n", b2, v2, r2)

	b3, v3, r3 := extraiInteiro("123abc")
	fmt.Printf("%v, %v, %v\n", b3, v3, r3)

	b4, v4, r4 := extraiInteiro("123.45")
	fmt.Printf("%v, %v, %v\n", b4, v4, r4)

	b5, v5, r5 := extraiInteiro("abc123")

	if !(b1 == false && v1 == "" && r1 == "") {
		t.Errorf("b1: [%v], v1: [%v], r1: [%v]\n Expected: b1: [false], v1: [], r1: []\n", b1, v1, r1)
	}

	if !(b2 == true && v2 == "123" && r2 == "") {
		t.Errorf("b2: [%v], v2: [%v], r2: [%v]\n Expected: b2: [true], v2: [123], r2: []\n", b2, v2, r2)
	}

	if !(b3 == true && v3 == "123" && r3 == "abc") {
		t.Errorf("b3: [%v], v3: [%v], r3: [%v]\n Expected: b3: [true], v3: [123], r3: [abc]\n", b3, v3, r3)
	}

	if !(b4 == false && v4 == "" && r4 == "123.45") {
		t.Errorf("b4: [%v], v4: [%v], r4: [%v]\n Expected: b4: [false], v4: [], r4: [123.45]\n", b4, v4, r4)
	}

	if !(b5 == false && v5 == "" && r5 == "abc123") {
		t.Errorf("b5: [%v], v5: [%v], r5: [%v]\n Expected: b5: [false], v5: [], r5: [abc123]\n", b5, v5, r5)
	}

}
