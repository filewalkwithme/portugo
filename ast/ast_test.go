package ast

import (
	core "github.com/maiconio/portugo/core"
	"testing"
)

func TestRemoveNodeFilho(t *testing.T) {
	pai := core.Node{nil, nil, "pai", 0, 0, core.Token{"", ""}}
	filho0 := core.Node{nil, nil, "filho0", 0, 0, core.Token{"", ""}}
	filho1 := core.Node{nil, nil, "filho1", 0, 0, core.Token{"", ""}}
	filho2 := core.Node{nil, nil, "filho2", 0, 0, core.Token{"", ""}}
	filho3 := core.Node{nil, nil, "filho3", 0, 0, core.Token{"", ""}}

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho0)
	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho1)
	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho2)
	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho3)

	f0 := len(pai.Filhos)

	RemoveNodeFilho(&pai, 1)

	f1 := len(pai.Filhos)

	if f1 != (f0 - 1) {
		t.Error("Número de filhos depois da remoção é diferente do número de filhos original menos 1")
	}

	if pai.Filhos[0].Valor != "filho0" || pai.Filhos[1].Valor != "filho2" || pai.Filhos[2].Valor != "filho3" || len(pai.Filhos) != 3 {
		t.Error("Remoção do filho1 falhou")
	}
}

func TestAdicionaNodeFilho(t *testing.T) {

	pai := core.Node{nil, nil, "pai", 0, 0, core.Token{"", ""}}
	filho0 := core.Node{nil, nil, "filho0", 0, 0, core.Token{"", ""}}
	filho1 := core.Node{nil, nil, "filho1", 0, 0, core.Token{"", ""}}
	filho2 := core.Node{nil, nil, "filho2", 0, 0, core.Token{"", ""}}

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho0)
	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho1)
	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho2)

	f1 := len(pai.Filhos)

	if f1 != 3 {
		t.Error("Número de filhos depois de 3 inclusões é diferente do esperado")
	}

	if pai.Filhos[0].Valor != "filho0" {
		t.Error("Inclusão do filho0 falhou")
	}

	if pai.Filhos[1].Valor != "filho1" {
		t.Error("Inclusão do filho1 falhou")
	}

	if pai.Filhos[2].Valor != "filho2" {
		t.Error("Inclusão do filho2 falhou")
	}

}

func TestRetiraNodesVazios(t *testing.T) {
	pai := core.Node{nil, nil, "pai", 0, 0, core.Token{"", ""}}

	deletar := map[string]bool{
		"INICIO": true,
		"FIM":    true,
		"PONTO":  true,
		",":      true,
		";":      true,
		":":      true,
		"(":      true,
		")":      true,
		"ESCV2":  true,
		"AC1":    true,
		"V1":     true,
		"R2":     true,
		"M10":    true,
		"M8":     true,
		"M6":     true,
		"M4":     true,
		"M2":     true,
		"L2":     true,
		"L4":     true,
		"L6":     true,
		"D3":     true,
	}

	for chave, _ := range deletar {
		filho := core.Node{nil, nil, chave, 0, 0, core.Token{"", ""}}

		AdicionaNodeFilho(&pai, len(pai.Filhos), &filho)
	}

	naoDeletar := map[string]bool{
		"R1": true,
		"M9": true,
		"M7": true,
		"M5": true,
		"M3": true,
		"M1": true,
		"L1": true,
		"L3": true,
	}

	for chave, _ := range naoDeletar {
		filho := core.Node{nil, nil, chave, 0, 0, core.Token{"", ""}}

		AdicionaNodeFilho(&pai, len(pai.Filhos), &filho)
	}

	for RetiraNodesVazios(&pai, 0) > 0 {
	}

	f1 := len(pai.Filhos)

	if f1 != len(naoDeletar) {
		t.Error("Número de filhos é diferente do esperado")
	}
}

func TestPromoveNodeSimples(t *testing.T) {
}

func TestComandoOperador(t *testing.T) {
}

func TestComandoAtribuicao(t *testing.T) {
}

func TestComandoEscreva(t *testing.T) {
}

func TestComandoLeia(t *testing.T) {
}

func TestComandoFuncaoMatematica(t *testing.T) {
}

func TestComandoDeclaraVariavel(t *testing.T) {
}

func TestPromoveAcoes(t *testing.T) {
}

func TestConfiguraAST(t *testing.T) {
}
