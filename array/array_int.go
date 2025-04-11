package main

import "fmt"

func Main2() {
	var numeros [4]int
	numeros[0] = 1
	numeros[1] = 12
	numeros[2] = 20
	numeros[3] = 69
	var suma int
	for _,numero := range numeros {
		suma += numero
	}
	fmt.Printf("La suma total es : %v", suma)

}