# portugo

#### O que é
Portugo é um interpretador de linguagem portugol escrito em Go.
Os objetivos principais deste projeto são:
- Fornecer uma ferramenta capaz de executar códigos escritos em linguagem Portugol e assim auxiliar no aprendizado de algoritmos e lógica de programação
- Implementar um interpretador utilizando a linguagem Go

#### Funcionalidades atualmente suportadas:
- Extração de tokens

#### Uso

    # percorre o arquivo e mostra todos os tokens extraídos e classificados por tipo
    ./portugo -tokens nome_aquivo.txt
    ---------------------------------  
    inteiro		TIPO_VARIAVEL
    :				  DOIS_PONTOS
              ESPACO
    X				  VARIAVEL
    ;				  PONTO_E_VIRGULA
