package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

func main() {
	ch := make(chan Persona)

	go func() {
		ch <- Persona{Nombre: "Daniel", Edad: 23}
	}()

	fmt.Printf("Data recibida de el chanel ch Data: %v",<-ch)

}
