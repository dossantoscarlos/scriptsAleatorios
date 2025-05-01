# ASCII Happy Day

Este é um programa simples em Go que imprime mensagens artísticas em ASCII no terminal, celebrando um "Happy Day" de forma divertida.

## Como funciona

O programa exibe, em blocos separados, as palavras "HAPPY", "DAY", "TO" e "WORK" utilizando arte em ASCII.

## Como executar

1. **Pré-requisitos:**  
   - Ter o [Go](https://golang.org/dl/) instalado na sua máquina.

2. **Clone ou baixe este repositório.**

3. **Execute o programa:**

   No terminal, navegue até a pasta onde está o arquivo `main.go` e rode:

   ```sh
   go run main.go
   ```

   Ou, para compilar e executar:

   ```sh
   go build -o happyday
   ./happyday
   ```

## Exemplo de saída

```
=======================================

 00000  0000   0       00000 00000 
 0      0      0  	 0      0        
 00000  0000   0         0     0   
 0      0      0    0    0    0    
 0      0000   000000  00000 00000 

=======================================

0000   00000   000  
0   0    0    0   0 
0    0   0    00000 
0   0    0    0   0 
0000   00000  0   0 

=======================================

0000      00  
0   0   0    0
0    0  0    0
0   0   0    0
0000\t  00  

=======================================

0000000  0000      000   00000     000    0        0    0    00000   
   0     0   0    0   0  0    0   0   0   0        0    0   0     0  
   0     0000     00000  00000    00000   0        000000   0     0  
   0     0   0    0   0  0    0   0   0   0    0   0    0   0     0  
   0     0    0   0   0  00000    0   0   000000   0    0    00000   

=======================================
```

## Licença

Este projeto é livre para uso e modificação. 