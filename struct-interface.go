package main

import "fmt"

// Structs can be considered as classes in Go
type User struct {
	Name string
	Age  int
}

// Assigned a method to the User struct
func (user User) isAdult() bool {
	return user.Age >= 18
}

func printStructures() {
	var user User = User{Name: "John", Age: 30}
	fmt.Println("User Name:", user)
	fmt.Println("Is Adult:", user.isAdult())
}

// Interfaces are used to implement polymorphism in Go
// All structs with same method signature can be used as an interface
type Area interface {
	calculateArea() float64
}

type Circle struct {
	Radius float64
}

func (circle Circle) calculateArea() float64 {
	return 3.14 * circle.Radius * circle.Radius
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (rectangle Rectangle) calculateArea() float64 {
	return rectangle.Length * rectangle.Width
}

func printInterfaces() {
	// Initialize the structs
	var circle Circle = Circle{Radius: 5}
	var rectangle Rectangle = Rectangle{Length: 10, Width: 5}

	var area = []Area{circle, rectangle}
	for _, a := range area {
		fmt.Printf("Area: %v\n", a.calculateArea())
	}
}
