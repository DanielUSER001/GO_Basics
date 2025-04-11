package main

import "fmt"

func main() {
	slice_mapas := make([]map[string]int, 3)

	slice_mapas[0] = map[string]int{"Queso": 999, "Dinamo": 123}

	fmt.Println(slice_mapas)
}
