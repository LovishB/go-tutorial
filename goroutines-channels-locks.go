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
var dataChannel chan string

/*
Simple goroutine example
*/
func dbCallsWithGoroutine() {
	var timeStart = time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1)   // Add a counter
		go dbCall() // This will run the dbCall function in a goroutine
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Execution Time: ", time.Since(timeStart))
}

func dbCall() {
	// Simulating a database call
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond) // 2sec delay
	fmt.Println("DB", dbData)
	wg.Done() // Decrement the counter when the goroutine completes
}

/*
Goroutine with write lock
This is used to protect the data from being accessed by multiple goroutines at the same time
*/
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

func dbCallWithWriteLock(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("Fetched DB data[%d]: %s\n", i, dbData[i])

	mutex.Lock()                               // Create a mutex lock to safeguard write access
	resultData = append(resultData, dbData[i]) // writing to resultData (here the order will not be preserved)
	mutex.Unlock()                             // Unlock the mutex after writing

	wg.Done() // Decrement the counter when the goroutine completes
}

/*
Channels are used to communicate between goroutines
They hold data and help to synchronize go routines

Two phases Channel Approach: 1. Wait for all goroutines to finish, 2. Collect all results from the channel
*/
func dbCallsWithGoroutineAndChannels() {
	var timeStart = time.Now()

	// Create a channel to receive data
	dataChannel = make(chan string, len(dbData)) // Buffered channel

	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCallWithChannels(i)
	}
	wg.Wait()

	// Collect all results from the channel
	channelResults := []string{}
	for i := 0; i < len(dbData); i++ {
		channelResults = append(channelResults, <-dataChannel) // Read from the channel one by one
	}

	close(dataChannel) // Close the channel after all goroutines are done

	fmt.Println("Execution Time: ", time.Since(timeStart))
	fmt.Println("ResultData ", channelResults)
}

func dbCallWithChannels(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("Fetched DB data[%d]: %s\n", i, dbData[i])

	// Send the data to the channel
	dataChannel <- dbData[i]

	wg.Done()
}

/*
One phase Approach: values are sent to the channel as they are read simultaneously
*/
func dbCallsWithGoroutineAndChannelsSimultaneousReads() {
	var timeStart = time.Now()

	// Create a channel to receive data
	dataChannel = make(chan string, len(dbData)) // Buffered channel

	for i := 0; i < len(dbData); i++ {
		go dbCallWithChannelsSimultaneousReads(i)
	}

	// Collection from the channel is done simultaneously
	channelResults := []string{}
	for i := 0; i < len(dbData); i++ {
		channelResults = append(channelResults, <-dataChannel)
	}

	close(dataChannel)

	fmt.Println("Execution Time: ", time.Since(timeStart))
	fmt.Println("ResultData ", channelResults)
}

func dbCallWithChannelsSimultaneousReads(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("Fetched DB data[%d]: %s\n", i, dbData[i])

	// Send the data to the channel
	dataChannel <- dbData[i]
}
