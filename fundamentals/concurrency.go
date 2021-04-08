package main

import (
	"fmt"
	"time"
)

// Function to be executed concurrently.
func f() {
    fmt.Println("f function")
}

// This function uses channels to transmit back
// information to the calling thread.
// Notice that every channel has a datatype.
// These channels seem to work as POSIX queues.
func strlen(s string, c chan int) {
	c <- len(s)
}

func main() {
    go f()  // Execute f

    // We need to keep the main thread alive until f() has finished
    // in order to see its results.
    time.Sleep(1 * time.Second)
    fmt.Println("main function")


	c := make(chan int)
	// Two executions
	go strlen("Salutations", c) // First execution
	go strlen("World", c) // Second execution

	// Channel information must be stored into
	// local variable, of the same type as the
	// transmitted back from the channel.
	// After two executions we can fetch two results.
	// Seems fairly deterministic.
	x := <-c // Result from first execution.
	y := <-c // Result from second execution.

	fmt.Println(x, y, x+y)
}

