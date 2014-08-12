package lex

import (
	"fmt"
	"testing"
)

func TestCarregaTokens(t *testing.T) {
	r1 :=
		`[]core.Token{core.Token{Tipo:"INTEIRO", Id:"1"}, core.Token{Tipo:"REAL", Id:"1.5"}, core.Token{Tipo:"STRING", Id:"abc"}, core.Token{Tipo:"LOGICO", Id:"falso"}, core.Token{Tipo:"TIPOVAR", Id:"inteiro"}, core.Token{Tipo:":", Id:":"}, core.Token{Tipo:"TIPOVAR", Id:"real"}, core.Token{Tipo:":", Id:":"}, core.Token{Tipo:"TIPOVAR", Id:"caractere"}, core.Token{Tipo:":", Id:":"}, core.Token{Tipo:"TIPOVAR", Id:"lógico"}, core.Token{Tipo:":", Id:":"}, core.Token{Tipo:":", Id:":"}, core.Token{Tipo:"PONTO", Id:"."}, core.Token{Tipo:"(", Id:"("}, core.Token{Tipo:")", Id:")"}, core.Token{Tipo:",", Id:","}, core.Token{Tipo:";", Id:";"}, core.Token{Tipo:"DIV", Id:"div"}, core.Token{Tipo:"MOD", Id:"mod"}, core.Token{Tipo:"FUNCMAT", Id:"sen"}, core.Token{Tipo:"FUNCMAT", Id:"cos"}, core.Token{Tipo:"FUNCMAT", Id:"tg"}, core.Token{Tipo:"FUNCMAT", Id:"arctg"}, core.Token{Tipo:"FUNCMAT", Id:"arccos"}, core.Token{Tipo:"FUNCMAT", Id:"arcsen"}, core.Token{Tipo:"FUNCMAT", Id:"abs"}, core.Token{Tipo:"FUNCMAT", Id:"int"}, core.Token{Tipo:"FUNCMAT", Id:"frac"}, core.Token{Tipo:"FUNCMAT", Id:"ard"}, core.Token{Tipo:"FUNCMAT", Id:"rnd"}, core.Token{Tipo:"OP.LOGICO.E", Id:"e"}, core.Token{Tipo:"OP.LOGICO.OU", Id:"ou"}, core.Token{Tipo:"OP.LOGICO.XOU", Id:"xou"}, core.Token{Tipo:"OP.LOGICO.UN", Id:"não"}, core.Token{Tipo:"FUNCMAT", Id:"sinal"}, core.Token{Tipo:"<-", Id:"<-"}, core.Token{Tipo:"OP.RELACIONAL", Id:"<="}, core.Token{Tipo:"OP.RELACIONAL", Id:">="}, core.Token{Tipo:"OP.RELACIONAL", Id:"<>"}, core.Token{Tipo:"OP.RELACIONAL", Id:"< >"}, core.Token{Tipo:"OP.RELACIONAL", Id:"="}, core.Token{Tipo:"LEIA", Id:"leia"}, core.Token{Tipo:"ESCREVA", Id:"escreva"}, core.Token{Tipo:"v", Id:"inicio"}, core.Token{Tipo:"FIM", Id:"fim"}, core.Token{Tipo:"+-", Id:"+"}, core.Token{Tipo:"+-", Id:"+"}, core.Token{Tipo:"+-", Id:"-"}, core.Token{Tipo:"*/", Id:"*"}, core.Token{Tipo:"*/", Id:"/"}, core.Token{Tipo:"**//", Id:"**"}, core.Token{Tipo:"**//", Id:"//"}, core.Token{Tipo:"v", Id:"variavel"}}`

	listaTokens := CarregaTokens("../testes_src/02.ptg")
	strListaTokens := fmt.Sprintf("%#v", listaTokens)
	if r1 != strListaTokens {
		t.Error("Tokens obtidos são diferentes do esperado", len(r1), len(strListaTokens))
	}
}
