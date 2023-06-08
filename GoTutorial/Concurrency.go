package main

import "fmt"


//A go routine is a function that is capable of running concurrently with other functions.
func f(n int){
	for i:=0; i<10; i++{
		fmt.Println(n, ":", i)
	}
}

func main(){
	go f(0)
	var input string
	fmt.Scanln(&input)
}

//This program consists of two goroutines. 
//The first goroutine is implicit and is the main function itself. 
//The second goroutine is created when we call go f(0). 
//Normally when we invoke a function our program will execute all the statements in a function 
//and then return to the next line following the invocation. 
//With a goroutine we return immediately to the next line and don't wait for the function to complete. 
//This is why the call to the Scanln function has been included; without it the program 
//would exit before being given the opportunity to print all the numbers.