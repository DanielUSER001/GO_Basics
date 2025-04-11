package main

import "fmt"

func Main() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	// Esperamos a recibir el numero
	NumeroRecibido := <-ch
	fmt.Println(NumeroRecibido)
}
