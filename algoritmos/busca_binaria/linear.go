package main

import "fmt"

func main() {
	numbers := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println(search(numbers, 30))
}

func search(arr []int, element int) int {
	// Percorre o array
	for i := 0; i < len(arr); i++ {
		// Caso o elemento desejado seja o presente no indice do array o retorna
		if arr[i] == element {
			return i
		}
	}
	// Caso o valor não seja encontrado retorna -1
	return -1
}
