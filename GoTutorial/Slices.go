package main

//Slice is a segment of an array.
//Like arrays slices are indexable and have a length.
//Unlike arrays the length is allowed to change

import "fmt"

// var x []float64

//x :=  make([]float64, 5)
//This creates a slice that is associated with an underlying float64 array of length 5. Slices are always associated with some array, and although they can never be longer than the array, they can be smaller.

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2) //[1 2 3] [1 2 3 4 5]
}
