package main

import "fmt"

//A function is an independent section of code that maps zero or more input parameters
// to zero or more output parameters. Functions (also known as procedures or subroutines)
//are often represented as a black box: (the black box represents the function)

// func average(xs []float64) float64{
// 	total:=0.0
// 	for _, v := range xs {
// 		total+=v
// 	}
// 	return total/float64(len(xs))
// }

// func main(){
// 	xs:=[]float64{43,46,66, 56,44}
// 	fmt.Println(average(xs))
// }

//Multiple input values
func add(args ...int) int { // a way to give 0 or more parameters
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
func main() {
	// fmt.Println(add(1, 2, 3))
	xs := []int{1,2,3}
	fmt.Println(add(xs...))
}
