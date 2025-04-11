package main

import "fmt"

func Main2() {
	type Empleado struct {
		Nombre  string
		Salario float64
	}

	var Array_empleado [3]Empleado = [3]Empleado{
		{Nombre: "Daniel", Salario: 1500.00},
		{Nombre: "daniel", Salario: 1500.00},
	}

	fmt.Println(Array_empleado)

}
