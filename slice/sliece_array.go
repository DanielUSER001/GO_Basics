package main

import (
	"fmt"
)

func Main4() {

	array_enteros := [5]int{1, 2, 3, 4, 5}

	slice := []int{array_enteros[0], array_enteros[1], array_enteros[1], array_enteros[2]}

	fmt.Println(slice)

}
