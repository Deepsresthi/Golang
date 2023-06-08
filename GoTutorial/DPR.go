package main

//DEFER, PANIC, AND RECOVER

import "fmt"

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func main() {
	defer second()
	first()
	//output 1st 2nd
}

// This program prints 1st followed by 2nd. Basically defer moves the call to second to the end of the function:

// We can handle a run-time panic with the built-in recover function. 
//recover stops the panic and returns the value that was passed to the call to panic. 
