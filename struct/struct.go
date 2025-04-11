package main

import "fmt"

type Libro struct {
	Titulo  string
	Autor   string
	Paginas int
}

type Persona struct {
	Direccion Direccion
	Nombre string
}
type Direccion struct {
	Calle  string
	Ciudad string
}



func Main() {
	slice_libros := make([]Libro, 2)
	
	
	persona_1 := Persona{
		Direccion: Direccion{Calle: "los azules", Ciudad: "Morelos"},
		Nombre: "Chefson",
	}

	persona_1.ImprimirNombre()

	libro_coran := Libro{
		Titulo:  "El coran",
		Autor:   "Enoc",
		Paginas: 657,
	}
	libro_mateo := Libro{
		Titulo:  "Mateo IV",
		Autor:   "Dios",
		Paginas: 1237,
	}

	slice_libros[0] = libro_coran
	slice_libros[1] = libro_mateo

}

func (p *Persona) ImprimirNombre() {
	fmt.Println("el nombre de la persona es ", p.Nombre)
}
