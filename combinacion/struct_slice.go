package main

import "fmt"

type Estudiante struct {
	Notas []int
}

func Main() {
	notas := []int{1, 2, 3, 4, 5}

	Estudiante_1 := Estudiante{
		Notas: notas,
	}

	fmt.Println(Estudiante_1)
}
