# Mergesort

Assim como o Quicksort o Mergesort é um algoritmo de ordenação que utiliza do paradigma [divide and conquer](https://github.com/GuilhermehChaves/google-skills/tree/master/algoritmos/divide_and_conquer),
que consiste em ramificar a entrada várias vezes para assim separar os problemas em vários outros menores.

### Merge

Este algoritmo divide o array em duas partes e utiliza recursão para executar a função em ambas as partes,
tendo no final um array para cada elemento com uma única posição, em seguida compara os valores entre um e outro
e verifica qual é menor, o menor valor é colocado à esquerda e em seguida realiza o merge destas partes do array
ordenado, formando agora grupos cada vez maiores e realizando o mesmo processo de comparação até o fim onde o array estará ordenado
por completo.

Para visualizar melhor como funciona vejamos a imagem a seguir:

![001_mergesort](https://user-images.githubusercontent.com/48635609/91215813-638b5480-e6eb-11ea-82ca-fa414d5a5d2f.png)

Observando a imagem podemos notar as divisões que ocorrem no array e os momentos em que são feitos o merge (junção).

![002_mergesort](https://user-images.githubusercontent.com/48635609/91216255-0643d300-e6ec-11ea-82ba-2b8af5fe5b42.gif)

Com o gif podemos entender mais claramente o processo, o array começa de cima para baixo realizando a separação dos elementos até que tenha
vários arrays com apenas uma posição referente a cada elemento do array principal, ao chegar ao final o array "sobe" novamente, desta vez 
realizando as comparações dos valores entre as partes do array as colocando em ordem e em seguida realiza o merge, ordenando os sub arrays até que chegue 
ao seu tamanho original e a ordenação seja realizada por completo.

Abaixo podemos ver uma representação em código do algoritmo `mergesort`

```go
package main

import (
	"fmt"
)

func main() {
	array := []int{5, 3, 2, 1, 4}
	fmt.Println("\n--- Desordenado --- \n\n", array)
	fmt.Println("\n--- Ordenado ---\n\n", mergeSort(array), "\n")
}

func mergeSort(a []int) []int {
	var num = len(a)

	if num == 1 {
		return a
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)

	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = a[i]
		} else {
			right[i-middle] = a[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

```

`OBS: No arquivo "mergesort.go" o código acima está com comentários explicando cada linha`

### Referências

https://www.golangprograms.com/golang-program-for-implementation-of-mergesort.html

https://www.geeksforgeeks.org/merge-sort/

### Materiais complementares

- [MERGE SORT | Algoritmos #7](https://www.youtube.com/watch?v=5prE6Mz8Vh0) - vídeo

- [Algoritmo de ordenação - MergeSort Video Aula](https://www.youtube.com/watch?v=PKCMMSXQyJE) - vídeo
