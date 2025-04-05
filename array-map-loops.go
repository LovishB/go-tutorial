package main

import "fmt"

func printArrays() {
	var arr [3]int32 // length of array is 3 and type is int32
	arr[0] = 123     // first element and the rest are default 0
	fmt.Printf("Array first element: %v\n", arr[0])
	fmt.Printf("Array complete: %v\n", arr[0:3])

	//array are continuous in memory and fixed in size
	fmt.Printf("Array length: %v\n", len(arr))
	fmt.Printf("Array capacity: %v\n", cap(arr))
	fmt.Printf("memory location of first element: %v\n", &arr[0])

	//initializing array
	var arr2 = []int32{1, 2, 3}
	fmt.Printf("Array2: %v\n", arr2)

	//appending to array
	var arr3 = append(arr2, 7)
	fmt.Printf("Array3: %v\n", arr3)

	// interating over array
	for i, value := range arr3 {
		fmt.Printf(("Index: %v, Value: %v\n"), i, value)
	}
}

func printMaps() {
	var myMap = make(map[string]int32) // defining empty map
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	fmt.Printf("Map: %v\n", myMap)

	// initializing map
	var myMap2 = map[string]int32{"one": 1, "two": 2, "three": 3}
	fmt.Printf("Map2: %v\n", myMap2)

	// fetching value from map
	var value, ok = myMap2["one"]
	fmt.Printf("Value: %v, Ok: %v\n", value, ok)

	var value2, ok2 = myMap2["four"]
	fmt.Printf("Value2: %v, Ok2: %v\n", value2, ok2)

	//deleting from map
	delete(myMap2, "one")
	fmt.Printf("Map2 after delete: %v\n", myMap2)

	//iterating over map
	for key, value := range myMap2 {
		fmt.Printf("Key: %v, Value %v\n", key, value)
	}
}

func printLoops() {
	// for loop
	fmt.Println("For Loop")
	for i := 0; i < 5; i++ {
		fmt.Printf("%v,", i)
	}
	fmt.Println()

	// for as while loop
	fmt.Println("For as While Loop")
	i := 0
	for i < 5 {
		fmt.Printf("%v,", i)
		i++
	}
	fmt.Println()

	// for as infinite loop
	fmt.Println("For as Infinite Loop")
	i = 0
	for {
		if i > 4 {
			break
		}
		fmt.Printf("%v,", i)
		i++
	}
	fmt.Println()

}
