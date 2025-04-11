package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)

	fmt.Printf("El largo del slice es: %v y su capacida es de: %v", len(slice), cap(slice))
}