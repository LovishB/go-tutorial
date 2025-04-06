package main

import (
	"fmt"
	"sync"
	"time"
)

// Goroutines are lightweight threads (2kbs and dynamic size change)
// They are managed by the Go runtime and not the OS scheduler
// Goroutines -> Go's runtime scheduler → Multiplexed in OS threads → CPU cores

// A CPU core can run only one thread at a time(parallism is running threads parallel in multi core)
// But multiple threads can be run on a single core when the core is free (conurrency)
// Goroutines can run in thousands on a single thread concurrently (due to it's lightweight nature)

var dbData = []string{"data1", "data2", "data3", "data4", "data5"}
var resultData = []string{} // Slice to store the results
var wg = sync.WaitGroup{}   // they are counters to wait for all goroutines to finish
var mutex = sync.Mutex{}    // Mutex to protect access to resultData

func dbCall() {
	// Simulating a database call
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond) // 2sec delay
	fmt.Println("DB", dbData)
	wg.Done() // Decrement the counter when the goroutine completes
}

func dbCallsWithGoroutine() {
	var timeStart = time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1)   // Add a counter
		go dbCall() // This will run the dbCall function in a goroutine
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Execution Time: ", time.Since(timeStart))
}

func dbCallWithWriteLock(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("DB", dbData)

	mutex.Lock()                               // Create a mutex lock to safeguard write access
	resultData = append(resultData, dbData[i]) // writing to resultData (here the order will not be preserved)
	mutex.Unlock()                             // Unlock the mutex after writing

	wg.Done() // Decrement the counter when the goroutine completes
}

func dbCallsWithGoroutineAndLocks() {
	var timeStart = time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCallWithWriteLock(i)
	}
	wg.Wait()
	fmt.Println("Execution Time: ", time.Since(timeStart))
	fmt.Println("ResultData ", resultData)
}
