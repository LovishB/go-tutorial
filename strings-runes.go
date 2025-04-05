package main

import (
	"fmt"
	"strings"
)

func printStrings() {
	// String are immutable
	str := "Hello, World!"
	// str[0] = 'h' // Error: cannot assign to str[0]
	str = "Hello, Go!" // This is allowed, as it creates a new string
	println(str)

	// Strings are UTF-32 encoded(32-bit ASCII)
	str2 := "Hello, 世界" // Contains both ASCII and non-ASCII characters

	for i, r := range str2 {
		println(i, r) // Prints the index and the UTF-32 code point
	}

	for i, r := range str2 {
		println(i, string(r)) // Prints the index and the rune
	}

	for _, r := range str2 {
		fmt.Printf("%c\n", r) // Prints the rune
	}

	// String builder is much faster than string concatenation
	var sb strings.Builder
	sb.WriteString("Hello")
	sb.WriteString(" ")
	sb.WriteString("Go")
	fmt.Println(sb.String()) // Prints the concatenated string
}
