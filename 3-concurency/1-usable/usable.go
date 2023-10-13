// very similar to: https://go.dev/tour/concurrency/2
package main

import (
    "fmt"
    "time"
)

// We use a channel to communicate with the main function.
// This will tell it to wait for us until the routines are
// done, solving the problem from the simple example

// This function will count to a number with a specified
// time delay between steps, and total up numbers to pass
// to the channel

// by giving them different totals to count to and different
// delays to wait between steps, we can see that they are
// out of sync, and that the main function will wait for them

// pass how high to count, how long to wait, what channel
func countSlowly(n int, sleepTime int, c chan int) {
    var total = 0
    for i := 0; i <= n; i++ {
        total += i
        fmt.Println(i, "of", n)
        // convert sleepTime to a time.Duration
        time.Sleep(time.Duration(sleepTime) * time.Millisecond)
    }
    // add the sum we calculated to the channel with <-
    c <- total
}

func main() {

    // note that channels are typed, so these only work with ints
    ch1 := make(chan int)
    ch2 := make(chan int)

    // start them with different totals and time delays
    go countSlowly(10, 250, ch1)
    go countSlowly(11, 450, ch2)

    // I know, it's ugly lol, but it works
    // the <- is used to assign from a channel
    total1 := <- ch1
    total2 := <- ch2

    // we will now get the results without having to
    // artificially delay the overall program
    fmt.Println("Total1:", total1)
    fmt.Println("Total2:", total2)

}
