package main

import (
	"fmt"
)

func Main6() {

	slice1 := []int{1, 2,3}
	slice2 := []int{2, 4,6}

	combinacion := append(slice1, slice2...)
	

	fmt.Println(combinacion)

}
