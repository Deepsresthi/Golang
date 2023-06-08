package main

import "fmt"

func main() {
	// for loop
	i := 1
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	//also be written as
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//if else loop
	for i:=1; i<=10; i++{
		if i%2==0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}
