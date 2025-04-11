package main

import "fmt"

func Main() {
	var nombres [3]string
	nombres[0] = "Daniel"
	nombres[1] = "Juan"
	nombres[2] = "luis"

	for key, value := range nombres {
		fmt.Printf("Nombre[%v]  tiene el valor de %v ",key, value)
	}

}