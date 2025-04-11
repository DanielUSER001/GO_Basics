package main

import "fmt"


func Main3()  {
 
	array_boolean := [5]bool{true, false, true, false, true}

	for _, boolean := range array_boolean {
		fmt.Printf("El array de boolean contien un %v", boolean)
	}
	
}