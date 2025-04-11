package main

import "fmt"


func main()  {
 
	array_float := [3]float32{ 5.0, 10.0, 33.3}
	var divisor int = 0
	var dividiendo float32 = 0

	for _, value := range array_float {
		dividiendo += value
		divisor += 1

	}
	fmt.Printf("El promedio de los numero en el arreglo es :%v", dividiendo/float32(divisor))
}