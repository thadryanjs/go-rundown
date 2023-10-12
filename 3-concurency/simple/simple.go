package main

import (
    "fmt"
    "time"
)

// count to ten with time between each number
func firstTenNumbers() {
    count := 0
    for count < 10 {
        fmt.Println(count)
        count++
        time.Sleep(500 * time.Millisecond)
    }
}

// print the first ten letters of the alphabet with time between each
func firstTenLetters() {
    letters := "ABCDEFGHIJ"
    for _, letter := range letters {
        fmt.Printf("\t%c\n", letter)
        time.Sleep(750 * time.Millisecond)
    }
}

func main() {

    // start the routines one after the other
    go firstTenLetters()
    go firstTenNumbers()

    // Wait for the goroutines to finish
    time.Sleep(8 * time.Second)
    // Note:
    // This is required becuase the main routine will not wait for the
    // go routines to finish before exiting! The next example will show
    // how to use channels to wait for the go routines to finish.
    // This isn't much use if you have to know excactly how long the
    // go routines will take to finish afterall
}
