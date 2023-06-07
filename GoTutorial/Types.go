package main

import "fmt"

//Integer type of Go are: uint8, uint16, uint32, uint64, int8, int16, int32, and int64
//Nan- Not a Number

func main() {
	fmt.Println("1 + 1 =", 1+1)
	//output 2

	//strings:
	fmt.Println(len("Hello World")) // 11
	fmt.Println("Hello World"[1])   //101->represented as bytes
	fmt.Println("Hello " + "World ")//Hello World

	//Booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

