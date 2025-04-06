package main // compiler will look for the main package to start the program

import "fmt"

func main() {
	fmt.Printf("Starting the Application %s\n", appName)
	fmt.Printf("Version %s\n", version)
	fmt.Printf("Max Retries %d\n", maxRetries)
	printVarConstants()
	printValue("Ahh function works")
	checkDivision(10, 0)
	checkDivision(10, 2)
	printArrays()
	printMaps()
	printLoops()
	printStrings()
	printStructures()
	printInterfaces()
	printPointers()
	dbCallsWithGoroutine()
	dbCallsWithGoroutineAndLocks()
	dbCallsWithGoroutineAndChannels()
	dbCallsWithGoroutineAndChannelsSimultaneousReads()
	printGenerics()
}
