package main

import "fmt"

func main() {
	// var x [5]int
	// x[4] = 100
	// fmt.Println(x) // 0 0 0 0 100

	//var x [5]float64
	// x[0] = 98
	// x[1] = 81
	// x[2] = 77
	// x[3] = 87
	// x[4] = 43
	x := [5]float64{89, 78, 98, 91, 100}

	var total float64 = 0
	for _, value := range x { // a single _ tells compiler that we don't need this
		total += value
	}
	fmt.Println(total / float64(len(x)))

}
