package main

/*
Generics are a way to write functions and data structures that can work with multiple data type.
*/
func sum[T int32 | int64 | float32 | float64](numbers []T) T {
	var total T
	for _, num := range numbers {
		total += num
	}
	return total
}

func printGenerics() {
	println(sum([]int32{1, 2, 3, 4, 5}))
	println(sum([]float64{1.5, 2.5, 3.5, 4.5, 5.5}))
}
