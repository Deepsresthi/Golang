package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	// x := 0
	// increment := func() int {
	// 	x++
	// 	return x
	// }

	// fmt.Println(increment())
	// fmt.Println(increment())

	// One way to use closure is by writing a function which returns
	// another function which – when called – can generate a sequence of numbers.
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
}
