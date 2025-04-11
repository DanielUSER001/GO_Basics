package main

import "fmt"

func main() {
	// strucuras anonimas
	Persona := struct {
		Nombre     string
		Puntuacion int
	}{
		Nombre: "Daniel",
		Puntuacion: 44,
	}
	fmt.Print(Persona)


	type Punto struct {
		X int
		Y int
	}
	// iniacializacion directa
	punto := Punto{
		X: 5,
		Y: 4,
	}
	fmt.Println(punto)
}
