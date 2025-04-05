package main //inside the main package

import (
	"fmt"
	"unicode/utf8"
) // fmt package is form go standard library

// Global variables are declared outside of any function
// and can be accessed throughout the package.
var (
	appName    string = "MyGoApp"
	version    string = "1.0.0"
	maxRetries int8   = 3
)

func printVarConstants() {
	var a int8 = 5
	fmt.Println(a)
	var b float32 = 5.5405
	fmt.Println(b)
	fmt.Println(b + float32(a)) // type conversion
	var c string = `Hello
Go`
	fmt.Println(c)

	fmt.Println("Hello" + "Go") // string concatenation
	var f string = "World"
	fmt.Printf("Hello %v Go\n", f)         // string formatting
	fmt.Println(utf8.RuneCountInString(c)) // count the number of char in a string
	fmt.Println(len(c))                    // count the number of bytes in a string

	var d rune = 'A'      // go has rune instead of char
	fmt.Println(d)        //this will print the UTF-32 code point
	fmt.Printf("%c\n", d) // this will print the char

	var e bool = true
	fmt.Println(e)

	//fancy ways of defining variables
	var x float32
	fmt.Println(x) // go assigns default value to the variable in now manual initialization

	var y = 5.5
	fmt.Println(y) // go will figure out the type of the variable on its own

	z := true
	fmt.Println(z) // short hand declaration, only works inside functions

	const pi = 3.14 // constant declaration
	fmt.Println(pi)
}
