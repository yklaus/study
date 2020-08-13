package pointer

import "fmt"

func SampleCopyValue() {
	a := 2
	b := a
	fmt.Println(a, b)
	fmt.Println(&a, &b)
	fmt.Println("-------------------------------------------")
	a = 10
	fmt.Println(a, b)
	fmt.Println("-------------------------------------------")
}

func SampleCopyMemoryAddress() {
	a := 2
	b := &a
	fmt.Println(a, b)
	fmt.Println(&a, b)
	fmt.Println(a, *b)
	fmt.Println("-------------------------------------------")
	a = 10
	fmt.Println(a, *b)
	fmt.Println("-------------------------------------------")
	*b = 20
	fmt.Println(a, *b)
}
