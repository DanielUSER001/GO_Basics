package main

import "testing"




func BenchmarkRemoveLoop(b *testing.B) {
	a := []int{1,2,3,4,5,6}
	index := 2
	for i := 0; i < b.N; i++ {
		removeLoop(a, index)
	}

}

func BenchmarkRemoveNewSlice(b *testing.B) {

	a := []int{1,2,3,4,5,6}
	index := 2
	for i := 0; i < b.N; i++ {
		removeNewSlice(a, index)
	}
}
func BenchmarkRemoveCopy(b *testing.B) {
	a := []int{1,2,3,4,5,6}
	index := 2
	for i := 0; i < b.N; i++ {
		removeCopy(a, index)
	}

}
func BenchmarkRemoveAppend(b *testing.B) {

	a := []int{1,2,3,4,5,6}
	index := 2
	for i := 0; i < b.N; i++ {
		removeAppend(a, index)
	}
}
func BenchmarkRemoveWithSlecePackage(b *testing.B) {

	a := []int{1,2,3,4,5,6}
	index := 2
	for i := 0; i < b.N; i++ {
		removeWithSlecePackage(a, index)
	}
}
