package main

import "fmt"

func main() {
	// var x string
	// x = "first"
	// fmt.Println(x)
	// x = "second"
	// fmt.Println(x)

	var x string
	x = "first "
	fmt.Println(x)
	x = x + "second" // first second
	fmt.Println(x)

	var y string = "hello"
	var z string = "world"
	fmt.Println(y == z) // false
}

func f(){
	const w string = "Hello World" // doesn't let change the value of the variables
	fmt.Println(w)
	//w = "will show undefined"
}
