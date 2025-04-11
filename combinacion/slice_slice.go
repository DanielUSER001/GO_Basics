package main

import "fmt"


func Main3()  {
	slice_slice := make([][]int , 3)
    slice := []int{ 1,2,3,4}

	slice_slice = append(slice_slice, slice )

	fmt.Print(slice_slice)
}