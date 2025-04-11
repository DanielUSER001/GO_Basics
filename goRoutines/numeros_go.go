package main

import (
	"fmt"
	"time"
)

func Main2() {

	go imprimirNumero(2)
	go imprimirNumero(4)
	go imprimirNumero(6)
	time.Sleep(time.Second)

}

func imprimirNumero(numero int) {
	fmt.Printf("--%v--", numero)
}
