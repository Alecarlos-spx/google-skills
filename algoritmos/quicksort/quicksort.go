package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := []int{10, 20, 7, 4, 90, 50, 70, 22, 44, 55, 33}

	fmt.Println("\n --- Desordenado ---\n\n", array, "\n")
	quicksort(array)
	fmt.Println("\n --- Ordenado ---\n\n", array, "\n")
}

func quicksort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	// Indice do primeiro e último elemento
	var start, end int = 0, len(a) - 1

	// Setando número aleatório ao pivot
	pivot := rand.Int() % len(a)

	fmt.Println("\n --- Trocando ----\n\n", a[pivot], " <---> ", a[end], "\n")
	// Realizando a troca do último elemento com o pivot
	a[pivot], a[end] = a[end], a[pivot]
	fmt.Println("\n --- Ordenando ---\n\n", a, "\n")

	for i, _ := range a {
		fmt.Println("\n", i, "i \n")
		// Se o elemento na posição i for menor que o elemento na posição final
		if a[i] < a[end] {
			// Troca as posição do primeiro elemento com o elemento na posição i
			fmt.Println("\n --- Trocando ---\n\n", a[start], " <---> ", a[i], "\n")
			a[start], a[i] = a[i], a[start]
			fmt.Println("\n --- Ordenando ---\n\n", a, "\n")
			// Somando + 1 no start
			start++
		}
	}

	fmt.Println("\n --- Trocando ---\n\n", a[start], " <---> ", a[end], "\n")
	a[start], a[end] = a[end], a[start]
	fmt.Println("\n --- Ordenando ---\n\n", a, "\n", end)

	// Chama a função recursivamente, pegando da posição 0 até o start
	quicksort(a[:start])
	// Chama a função recursivamente, começando do start + 1 até o final
	quicksort(a[start+1:])

	return a
}
