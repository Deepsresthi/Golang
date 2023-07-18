package math

import (
	"testing"
	"fmt"	
)

// func TestAdd(t *testing.T) {

// 	got := Add(4, 6)
// 	want := 10

// 	if got != want {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}
// }

type addTest struct{
	arg1, arg2, expected int
}

//arg1 and arg2 are argumenst, and the expected stands for the 'result we expect' 

var addTests = []addTest{
	{2, 3, 5},
	{2, 8, 10},
	{5, 7, 12},
	{3, 13, 16},
}

func TestAdd(t *testing.T){
	
	for _, test := range addTests{
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expecetd %q", output, test.expected)
		}
	}
} //for testing purposes

func BenchmarkAdd(b *testing.B) {
	for i:=0; i<b.N; i++{
		Add(4, 6)
	}
}

func ExampleAdd() {
	fmt.Println(Add(4,6))
	//output 10
}