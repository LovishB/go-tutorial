package main

import "fmt"

func printPointers() {
	// A pointer is a variable that stores the memory address of another variable.
	var a int32 = 58
	var p *int32 = &a // (& operator means address of) it will have the first memory address of 4 bytes memory allocated to a(int32)
	fmt.Println("a:", a)
	fmt.Println("p:", p)
	fmt.Println(*p) // (* operator means value at address) dereferencing the pointer

	// We can chnage the value of a using the pointer
	*p = 100
	fmt.Println("a:", a)

	// always check nil pointer before dereferencing otherwise program will panic and crash
	var p2 *int // a nil pointer i.e. it does not point to any variable
	if p2 == nil {
		fmt.Println("p2 is nil")
	}

	// pinter to a pointer
	var x int = 10
	var pX *int = &x    // pX is a pointer to x i.e. it will have the address of x
	var ppX **int = &pX // ppX is a pointer to i.e. pX it will have the address of pX
	fmt.Println("x:", x)
	fmt.Println("pX:", pX)
	fmt.Println("ppX:", ppX)
	fmt.Println("value at ppX:", **ppX) // dereferencing twice to get the value of x

	**ppX = 100
	fmt.Println("x:", x) // x is changed to 100
}
