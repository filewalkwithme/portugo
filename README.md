Portugo
=======

Interpretador de código Portugol escrito em [Go](http://www.golang.org)

O objetivo primário é tornar possível a execução dos códigos de exemplo do livro [Lógica de Programação](http://www.skoob.com.br/livro/176433-logica_de_programacao)

Características atualmente suportadas:
- Declaração de variáveis dos tipos INTEIRO e REAL.
- Funções de entrada e saída de dados: LEIA e ESCREVA.
- Resolução de expressões lógica e matemáticas.

Exemplo:

    $ go run portugo.go -arq=olamundo

Arquivo olamundo

    início
      caractere: msg;
      msg <- "oi";
      escreva("Olá Mundo! ", msg);
    fim.

Saída

    TOKENS:
    ------>
    [{INICIO início} {TIPOVAR caractere} {: :} {v msg} {; ;} {v msg} {<- <-} {STRING oi} {; ;} {ESCREVA escreva} {( (} {STRING Olá Mundo! } {, ,} {v msg} {) )} {; ;} {FIM fim} {PONTO .}]


    ÁRVORE SINTÁTICA:
    ------>
    P[]
      INICIO[início]
      V[V]
        V1[V1]
          D[D]
	    D1[D1]
	      TIPOVAR[caractere]
	      :[:]
	    D2[D2]
	      v[msg]
	      D3[D3]
	        _[_]
          ;[;]
          V1[V1]
	    _[_]
      AC[AC]
        AC1[AC1]
          A[A]
	    v[msg]
	    <-[<-]
	    L[L]
	      L1[L1]
	        L3[L3]
	          L5[L5]
	            R1[R1]
	              STRING[oi]
	            R2[R2]
	              _[_]
	          L6[L6]
	            _[_]
	        L4[L4]
	          _[_]
	      L2[L2]
	        _[_]
          ;[;]
          AC1[AC1]
	    A[A]
	      ESCREVA[escreva]
	      ([(]
	      ESCV1[ESCV1]
	        STRING[Olá Mundo! ]
	      ESCV2[ESCV2]
	        ,[,]
	        ESCV1[ESCV1]
	          L[L]
	            L1[L1]
	              L3[L3]
	                L5[L5]
	                  R1[R1]
	                    M1[M1]
	                      M3[M3]
	                        M5[M5]
	                          M7[M7]
	                            M9[M9]
	                              v[msg]
	                            M10[M10]
	                              _[_]
	                          M8[M8]
	                            _[_]
	                        M6[M6]
	                          _[_]
	                      M4[M4]
	                        _[_]
	                    M2[M2]
	                      _[_]
	                  R2[R2]
	                    _[_]
	                L6[L6]
	                  _[_]
	              L4[L4]
	                _[_]
	            L2[L2]
	              _[_]
	        ESCV2[ESCV2]
	          _[_]
	      )[)]
	    ;[;]
	    AC1[AC1]
	      _[_]
      FIM[fim]
      PONTO[.]


    AST - ÁRVORE SINTÁTICA ABSTRATA:
    ------>
    P[]
      D[D]
        TIPOVAR[caractere]
          v[msg]
      A[A]
        <-[<-]
          v[msg]
          STRING[oi]
      A[A]
        ESCREVA[escreva]
          STRING[Olá Mundo! ]
          v[msg]


    -----------inicio do programa-------
    Olá Mundo! oi
    -----------fim do programa----------
