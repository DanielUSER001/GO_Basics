package main

import "fmt"

func main() {
	mapa := make(map[string]string)

	mapa["ciudad_1"] = "Morelos"
	mapa["ciudad_2"] = "San Luis"
	mapa["ciudad_3"] = "Ciudad de MX"

	// checar si existe una ciuidd

	_, existe := mapa["ciudad_4"]
	if !existe {
		fmt.Println("no exites  ese valor")

	}

	


	mapa2 := map[int]string{ 1:"America", 2:"Europa"}
	mapa2[3] = "Asia"

	

	// eliminar una clave valor con delete
	delete(mapa, "ciudad_1")

	fmt.Println("el valor de ciudad 1 es :",mapa["ciudad_1"])

}
