- Diferente do Array não tem um número de elementos especificados (Mesmo esquema da List do Java)
- Declarado com: slice := \[]int{1, 2, 3, 4}

Quando criamos um novo slice a partir de um slice já existente e o slice que será utilizado para dizer quais elementos conterão no segundo slice ainda tem cap, o corte dele com append irá modificalo, então estamos dando um re-slice no slice 
```go
primeiroSlice := []int{1, 2, 3, 4, 5}

fmt.Println(primeiroSlice) // [1 2 3 4 5]

segundoSlice := append(primeiroSlice[:2], primeiroSlice[4:]...)

fmt.Println(segundoSlice) // [1 2 5]

fmt.Println(primeiroSlice) // 1 2 5 4 5

```
Para evitar isso podemos utilizar um for loop ou modificar diretamente o primeiro slice que criamos, para n ter esse problema em N ponto.

#### Slices multidimensionais 
```go
ss := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
```

## Range

- Atravessa toda a extensão do slice e nos mostra oque tem lá
- Recebe 2 parametros: 
	- 1º indice do lugar que estamos passando
	- 2º valor que está no indice que estamos passando
- Se não precisarmos usar os 2 parans podemos omitilos com \_

