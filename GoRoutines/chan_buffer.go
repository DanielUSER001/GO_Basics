package main

import "fmt"

func Main2() {
	ch := make(chan int , 2)
	go func() {
		ch <- 50
		ch <- 100
	}()
	// Esperamos a recibir por un valor
	Numero1 := <-ch
	fmt.Printf("Recibimos un valor de %v y vamos por el siguiente", Numero1)
	// Esperamos a recibir por el siguiente  valor
	Numero2 := <-ch
	fmt.Printf("Recibimos el ultimo valor que esperavamos %v ", Numero2)
	
}
