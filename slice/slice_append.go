package main

import (
	"fmt"
)

func Main5() {

	slice := []int{1, 2}
	fmt.Println(slice)
	slice = append(slice, 3)

	fmt.Println(slice)

}
