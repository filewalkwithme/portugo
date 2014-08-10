package ast

import (
	core "github.com/maiconio/portugo/core"
	lex "github.com/maiconio/portugo/lex"
	util "github.com/maiconio/portugo/util"
	sintatico "github.com/maiconio/portugo/sintatico"
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
	pai := core.Node{nil, nil, "pai", 0, 0, core.Token{"", ""}}
	filho0 := core.Node{nil, nil, "filho0", 0, 0, core.Token{"", ""}}
	filho0a := core.Node{nil, nil, "filho0a", 0, 0, core.Token{"", ""}}

	filho1 := core.Node{nil, nil, "FUNCMAT", 0, 0, core.Token{"FUNCMAT", "cos"}}
	filho1a := core.Node{nil, nil, "filho1a", 0, 0, core.Token{"INTEIRO", "1"}}

	filho2 := core.Node{nil, nil, "filho2", 0, 0, core.Token{"", ""}}
	filho2a := core.Node{nil, nil, "filho2a", 0, 0, core.Token{"", ""}}
	filho2b := core.Node{nil, nil, "filho2b", 0, 0, core.Token{"", ""}}

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho0)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0a)

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho1)
	AdicionaNodeFilho(&filho1, len(filho1.Filhos), &filho1a)

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho2)
	AdicionaNodeFilho(&filho2, len(filho2.Filhos), &filho2a)
	AdicionaNodeFilho(&filho2, len(filho2.Filhos), &filho2b)

	for PromoveNodeSimples(&pai, 0) > 0 {
	}

	if len(pai.Filhos) != 3 {
		t.Error("Número de filhos do nó pai é diferente do esperado [3]")
	} else {
		if len(pai.Filhos[0].Filhos) != 0 {
			t.Error("Número de filhos do nó filho0 é diferente do esperado [0]")
		}

		if len(pai.Filhos[1].Filhos) != 1 {
			t.Error("Número de filhos do nó filho1 é diferente do esperado [1]")
		}

		if len(pai.Filhos[2].Filhos) != 2 {
			t.Error("Número de filhos do nó filho2 é diferente do esperado [2]")
		}
	}
}

func TestComandoOperador(t *testing.T) {
	pai := core.Node{nil, nil, "pai", 0, 0, core.Token{"", ""}}

	/*

	   				R1[R1]
	                     INTEIRO[1]
	                     M2[M2]
	                       +-[+]
	                       INTEIRO[2]
	*/

	filho0 := core.Node{nil, nil, "R1", 0, 0, core.Token{"R1", "R1"}}

	filho0a := core.Node{nil, nil, "INTEIRO", 0, 0, core.Token{"INTEIRO", "1"}}
	filho0b := core.Node{nil, nil, "M2", 0, 0, core.Token{"M2", "M2"}}

	filho0b_a := core.Node{nil, nil, "+-", 0, 0, core.Token{"+-", "+"}}
	filho0b_b := core.Node{nil, nil, "INTEIRO", 0, 0, core.Token{"INTEIRO", "2"}}

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho0)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0a)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0b)
	AdicionaNodeFilho(&filho0b, len(filho0b.Filhos), &filho0b_a)
	AdicionaNodeFilho(&filho0b, len(filho0b.Filhos), &filho0b_b)
	//--------------------

	/*
       R1[R1]
         <-[<-]
         M2[M2]
           +-[+]
           INTEIRO[2]
	*/

	filho1 := core.Node{nil, nil, "R1", 0, 0, core.Token{"R1", "R1"}}
	filho1a := core.Node{nil, nil, "<-", 0, 0, core.Token{"<-", "<-"}}
	filho1b := core.Node{nil, nil, "M2", 0, 0, core.Token{"M2", "M2"}}

	filho1b_a := core.Node{nil, nil, "+-", 0, 0, core.Token{"+-", "+"}}
	filho1b_b := core.Node{nil, nil, "INTEIRO", 0, 0, core.Token{"INTEIRO", "2"}}

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho1)
	AdicionaNodeFilho(&filho1, len(filho1.Filhos), &filho1a)
	AdicionaNodeFilho(&filho1, len(filho1.Filhos), &filho1b)
	AdicionaNodeFilho(&filho1b, len(filho1b.Filhos), &filho1b_a)
	AdicionaNodeFilho(&filho1b, len(filho1b.Filhos), &filho1b_b)
	//--------------------

	for ComandoOperador(&pai, 0) > 0 {
	}

	if len(pai.Filhos) != 2 {
		t.Error("Número de filhos do nó pai é diferente do esperado [2]")
	} else {
		if len(pai.Filhos[0].Filhos) != 1 {
			t.Error("Número de filhos do nó filho0 é diferente do esperado [1]")
		} else {
			if len(pai.Filhos[0].Filhos[0].Filhos) != 2 {
				t.Error("Número de filhos do nó filho0.0 é diferente do esperado [2]")
			}
		}

		if len(pai.Filhos[1].Filhos) != 2 {
			t.Error("Número de filhos do nó filho1 é diferente do esperado [1]")
		}
	}
}

func TestComandoAtribuicao(t *testing.T) {
	pai := core.Node{nil, nil, "pai", 0, 0, core.Token{"", ""}}

	/*
	   	  A[A]
	        v[c]
	        <-[<-]
	       	INTEIRO[1]
	*/

	filho0 := core.Node{nil, nil, "A", 0, 0, core.Token{"A", "A"}}

	filho0a := core.Node{nil, nil, "v", 0, 0, core.Token{"A", "A"}}
	filho0b := core.Node{nil, nil, "<-", 0, 0, core.Token{"<-", "<-"}}
	filho0c := core.Node{nil, nil, "INTEIRO", 0, 0, core.Token{"INTEIRO", "1"}}

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho0)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0a)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0b)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0c)
	//--------------------

	/*
	   	  A[A]
	        v[c]
	        <-[<-]
	       	INTEIRO[1]
	       	INTEIRO[2]
	*/

	filho1 := core.Node{nil, nil, "A", 0, 0, core.Token{"A", "A"}}

	filho1a := core.Node{nil, nil, "v", 0, 0, core.Token{"A", "A"}}
	filho1b := core.Node{nil, nil, "<-", 0, 0, core.Token{"<-", "<-"}}
	filho1c := core.Node{nil, nil, "INTEIRO", 0, 0, core.Token{"INTEIRO", "1"}}
	filho1d := core.Node{nil, nil, "INTEIRO", 0, 0, core.Token{"INTEIRO", "2"}}

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho1)
	AdicionaNodeFilho(&filho1, len(filho1.Filhos), &filho1a)
	AdicionaNodeFilho(&filho1, len(filho1.Filhos), &filho1b)
	AdicionaNodeFilho(&filho1, len(filho1.Filhos), &filho1c)
	AdicionaNodeFilho(&filho1, len(filho1.Filhos), &filho1d)
	//--------------------

	for ComandoAtribuicao(&pai, 0) > 0 {
	}

	if len(pai.Filhos) != 2 {
		t.Error("Número de filhos do nó pai é diferente do esperado [2]")
	} else {
		if len(pai.Filhos[0].Filhos) != 1 {
			t.Error("Número de filhos do nó filho0 é diferente do esperado [1]")
		} else {
			if len(pai.Filhos[0].Filhos[0].Filhos) != 2 {
				t.Error("Número de filhos do nó filho0.0 é diferente do esperado [2]")
			}
		}

		if len(pai.Filhos[1].Filhos) != 4 {
			t.Error("Número de filhos do nó filho1 é diferente do esperado [4]")
		}
	}

}

func TestComandoEscreva(t *testing.T) {
	pai := core.Node{nil, nil, "pai", 0, 0, core.Token{"", ""}}

/*
  A[A]
    ESCREVA[escreva]
    STRING[Oi]
    STRING[a]
 */
	filho0 := core.Node{nil, nil, "A", 0, 0, core.Token{"A", "A"}}

	filho0a := core.Node{nil, nil, "ESCREVA", 0, 0, core.Token{"ESCREVA", "escreva"}}
	filho0b := core.Node{nil, nil, "STRING", 0, 0, core.Token{"ESCREVA", "Oi"}}
	filho0c := core.Node{nil, nil, "STRING", 0, 0, core.Token{"ESCREVA", "a"}}

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho0)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0a)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0b)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0c)
	
	
	//--------------------
	for ComandoEscreva(&pai, 0) > 0 {
	}

	if len(pai.Filhos) != 1 {
		t.Error("Número de filhos do nó pai é diferente do esperado [1]")
	} else {
		if len(pai.Filhos[0].Filhos) != 1 {
			t.Error("Número de filhos do nó filho0 é diferente do esperado [1]")
		} else {
			if len(pai.Filhos[0].Filhos[0].Filhos) != 2 {
				t.Error("Número de filhos do nó filho0a é diferente do esperado [2]")
			} 
		}
	}
}

func TestComandoLeia(t *testing.T) {
	pai := core.Node{nil, nil, "pai", 0, 0, core.Token{"", ""}}
/*
   A[A]
    LEIA[leia]
    v[a]
    LEIA2[LEIA2]
      v[b]
      LEIA2[LEIA2]
*/

	filho0 := core.Node{nil, nil, "A", 0, 0, core.Token{"A", "A"}}

	filho0a := core.Node{nil, nil, "LEIA", 0, 0, core.Token{"LEIA", "leia"}}
	filho0b := core.Node{nil, nil, "v", 0, 0, core.Token{"v", "a"}}
	filho0c := core.Node{nil, nil, "LEIA2", 0, 0, core.Token{"LEIA2", "leia2"}}

	filho0c_a := core.Node{nil, nil, "v", 0, 0, core.Token{"v", "b"}}
	filho0c_b := core.Node{nil, nil, "LEIA2", 0, 0, core.Token{"LEIA2", "leia2"}}

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho0)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0a)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0b)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0c)

	AdicionaNodeFilho(&filho0c, len(filho0c.Filhos), &filho0c_a)
	AdicionaNodeFilho(&filho0c, len(filho0c.Filhos), &filho0c_b)
	
	
	//--------------------
	for ComandoLeia(&pai, 0) > 0 {
	}

	if len(pai.Filhos) != 1 {
		t.Error("Número de filhos do nó pai é diferente do esperado [1]")
	} else {
		if len(pai.Filhos[0].Filhos) != 1 {
			t.Error("Número de filhos do nó filho0 é diferente do esperado [1]")
		} else {
			if len(pai.Filhos[0].Filhos[0].Filhos) != 2 {
				t.Error("Número de filhos do nó filho0a é diferente do esperado [2]")
			} 
		}
	}
}

func TestComandoFuncaoMatematica(t *testing.T) {
	pai := core.Node{nil, nil, "pai", 0, 0, core.Token{"", ""}}
/*	 
    M9[M9]
      FUNCMAT[sen]
      INTEIRO[1]
*/
	filho0 := core.Node{nil, nil, "M9", 0, 0, core.Token{"M9", "M9"}}

	filho0a := core.Node{nil, nil, "FUNCMAT", 0, 0, core.Token{"FUNCMAT", "sen"}}
	filho0b := core.Node{nil, nil, "INTEIRO", 0, 0, core.Token{"INTEIRO", "1"}}

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho0)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0a)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0b)
		
	//--------------------
	for ComandoFuncaoMatematica(&pai, 0) > 0 {
	}

	if len(pai.Filhos) != 1 {
		t.Error("Número de filhos do nó pai é diferente do esperado [1]")
	} else {
		if len(pai.Filhos[0].Filhos) != 1 {
			t.Error("Número de filhos do nó filho0 é diferente do esperado [1]")
		} else {
			if len(pai.Filhos[0].Filhos[0].Filhos) != 1 {
				t.Error("Número de filhos do nó filho0a é diferente do esperado [1]")
			} 
		}
	}
}

func TestComandoDeclaraVariavel(t *testing.T) {
	pai := core.Node{nil, nil, "pai", 0, 0, core.Token{"", ""}}
/*
  D[D]
    TIPOVAR[real]
    D2[D2]
      v[a]
      D2[D2]
        v[b]
        v[c] 
*/	
	filho0 := core.Node{nil, nil, "D", 0, 0, core.Token{"D", "D"}}

	filho0a := core.Node{nil, nil, "TIPOVAR", 0, 0, core.Token{"TIPOVAR", "real"}}
	filho0b := core.Node{nil, nil, "D2", 0, 0, core.Token{"D2", "D2"}}

	filho0b_a := core.Node{nil, nil, "v", 0, 0, core.Token{"v", "a"}}
	filho0b_b := core.Node{nil, nil, "D2", 0, 0, core.Token{"D2", "D2"}}

	filho0b_b_a := core.Node{nil, nil, "v", 0, 0, core.Token{"v", "b"}}
	filho0b_b_b := core.Node{nil, nil, "v", 0, 0, core.Token{"v", "c"}}

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho0)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0a)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0b)

	AdicionaNodeFilho(&filho0b, len(filho0b.Filhos), &filho0b_a)
	AdicionaNodeFilho(&filho0b, len(filho0b.Filhos), &filho0b_b)

	AdicionaNodeFilho(&filho0b_b, len(filho0b_b.Filhos), &filho0b_b_a)
	AdicionaNodeFilho(&filho0b_b, len(filho0b_b.Filhos), &filho0b_b_b)
		
	//--------------------
	for ComandoDeclaraVariavel(&pai, 0) > 0 {
	}

	if len(pai.Filhos) != 1 {
		t.Error("Número de filhos do nó pai é diferente do esperado [1]")
	} else {
		if len(pai.Filhos[0].Filhos) != 1 {
			t.Error("Número de filhos do nó filho0 é diferente do esperado [1]")
		} else {
			if len(pai.Filhos[0].Filhos[0].Filhos) != 3 {
				t.Error("Número de filhos do nó filho0a é diferente do esperado [3]")
			} 
		}
	}
}

func TestPromoveAcoes(t *testing.T) {
	pai := core.Node{nil, nil, "pai", 0, 0, core.Token{"", ""}}
/*
P[]
  V1[V1]
    D[D]
      TIPOVAR[real]
        v[a]
    D[D]
      TIPOVAR[inteiro]
        v[b]

*/
	filho0 := core.Node{nil, nil, "V1", 0, 0, core.Token{"D", "D"}}

	filho0a := core.Node{nil, nil, "D", 0, 0, core.Token{"D", "D"}}
	filho0a_a := core.Node{nil, nil, "TIPOVAR", 0, 0, core.Token{"TIPOVAR", "real"}}
	filho0a_a_a := core.Node{nil, nil, "v", 0, 0, core.Token{"v", "a"}}

	filho0b := core.Node{nil, nil, "D", 0, 0, core.Token{"D", "D"}}
	filho0b_a := core.Node{nil, nil, "TIPOVAR", 0, 0, core.Token{"TIPOVAR", "real"}}
	filho0b_a_a := core.Node{nil, nil, "v", 0, 0, core.Token{"v", "b"}}

	AdicionaNodeFilho(&pai, len(pai.Filhos), &filho0)
	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0a)
	AdicionaNodeFilho(&filho0a, len(filho0a.Filhos), &filho0a_a)
	AdicionaNodeFilho(&filho0a_a, len(filho0a_a.Filhos), &filho0a_a_a)

	AdicionaNodeFilho(&filho0, len(filho0.Filhos), &filho0b)
	AdicionaNodeFilho(&filho0b, len(filho0b.Filhos), &filho0b_a)
	AdicionaNodeFilho(&filho0b_a, len(filho0b_a.Filhos), &filho0b_a_a)

	//--------------------
	for PromoveAcoes(&pai, 0) > 0 {
	}

	if len(pai.Filhos) != 2 {
		t.Error("Número de filhos do nó pai é diferente do esperado [2]")
	} else {
		if len(pai.Filhos[0].Filhos) != 1 {
			t.Error("Número de filhos do nó filho0 é diferente do esperado [1]")
		} else {
			if len(pai.Filhos[0].Filhos[0].Filhos) != 1 {
				t.Error("Número de filhos do nó filho0a é diferente do esperado [1]")
			} else {
				if len(pai.Filhos[0].Filhos[0].Filhos[0].Filhos) != 0 {
					t.Error("Número de filhos do nó filho0a_a é diferente do esperado [0]")
				} 
			}
		}
	}

}

func TestConfiguraAST(t *testing.T) {
r1 :=
`P[]
  D[D]
    TIPOVAR[real]
      v[a]
  D[D]
    TIPOVAR[inteiro]
      v[b]
      v[c]
  A[A]
    LEIA[leia]
      v[a]
      v[b]
  A[A]
    LEIA[leia]
      v[c]
  A[A]
    <-[<-]
      v[c]
      +-[+]
        v[a]
        v[b]
  A[A]
    ESCREVA[escreva]
      v[a]
      v[b]
  A[A]
    ESCREVA[escreva]
      v[c]
`

		listaTokens := lex.CarregaTokens("../testes_src/01.ptg")
			
		parseTree := core.Node{nil, nil, "P", 0, 0, core.Token{"", ""}}
		sintatico.MontaParsingTree(&parseTree, listaTokens)
		ConfiguraAST(&parseTree)
		if r1 != util.MostraTree(&parseTree) {
			t.Error(len(util.MostraTree(&parseTree)),"---------",len(r1))
		}
		
}
