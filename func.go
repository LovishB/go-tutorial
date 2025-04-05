package main

import (
	"errors"
	"fmt"
)

func printValue(val string) {
	println(val)
}

func intDivision(a int32, b int8) (int32, int32, error) {
	var err error //nil error by default
	if b == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int32 = a / int32(b)
	var remainder int32 = a % int32(b)
	return result, remainder, err
}

func checkDivision(a int32, b int8) (int32, int32, error) {
	var result, remainder, err = intDivision(a, b)

	switch err {
	case nil:
		fmt.Println("Division successful")
	default:
		fmt.Println("Division failed")
	}
	return result, remainder, err
}
