package main

//var x map[string] int

import "fmt"

func main() {
	// var x map[string]int
	// x["key"] = 10
	// fmt.Println(x)
	//Up till now we have only seen compile-time errors.
	//This is an example of a runtime error.
	//As the name would imply, runtime errors happen when you run the program,
	//while compile-time errors happen when you try to compile the program.

	//x := make(map[string]int)
	//"make" in Golang is a keyword that is used to initialize only slices, maps, and channels.
	//Unlike the new keywords, it does not return the pointer to the item.
	//It initializes the data in memory with zero value.
	// x["key"] = 10
	// fmt.Println(x["key"])

	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
