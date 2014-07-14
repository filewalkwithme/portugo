Portugo
=======

Interpretador de código Portugol escrito em [Go](http://www.golang.org)

O objetivo primário é tornar possível a execução dos códigos de exemplo do livro [Lógica de Programação](http://www.skoob.com.br/livro/176433-logica_de_programacao)

Características atualmente suportadas:
- Declaração de variáveis dos tipos INTEIRO e REAL.
- Funções de entrada e saída de dados: LEIA e ESCREVA.
- Resolução de expressões lógica e matemáticas.

Exemplo:

    $ go run portugo.go -arq=texte2

Arquivo texte2

    início
      escreva(rnd(20));
      escreva(rnd(20));
      escreva(rnd(20));
    fim.

Saída   

    P[]
      V[V]
      A[A]
        ESCREVA[escreva]
          FUNCMAT[rnd]
            INTEIRO[20]
      A[A]
        ESCREVA[escreva]
          FUNCMAT[rnd]
            INTEIRO[20]
      A[A]
        ESCREVA[escreva]
          FUNCMAT[rnd]
            INTEIRO[20]


    -----------inicio do programa-------
    2
    3
    0
    -----------fim do programa----------

