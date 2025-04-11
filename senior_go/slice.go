package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println("-----------------")
	fmt.Println(removeNewSlice(slice, 2))
	fmt.Println("-----------------")
	fmt.Println(removeLoop(slice, 3))
	fmt.Println("-----------------")
	fmt.Println(removeCopy(slice, 4))
	fmt.Println("-----------------")
	fmt.Println(removeAppend(slice, 1))
	fmt.Println("-----------------")
	fmt.Println(removeWithSlecePackage(slice, 4))
	fmt.Println("-----------------")

}
func removeLoop(slice []int, index int) []int {

	result := make([]int, 0, len(slice)-1)

	for key, value := range slice {
		if index == key {
			continue
		}
		result = append(result, value)
	}
	return result
}


func removeNewSlice(slice []int, index int) []int {
	result := make([]int, len(slice)-1)

	copy(result[:index], slice[:index])
	copy(result[index:], slice[index+1:])
	return result
}

func removeCopy(slice []int, index int) []int {

	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}


func removeAppend(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)

}

func removeWithSlecePackage(slice []int, index int) []int {
	return slices.Delete(slice, index, index+1)

}
