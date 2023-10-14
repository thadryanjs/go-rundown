// Very similar to this: https://go.dev/tour/concurrency/4
package main

import (
    "fmt"
)

// Technically these are "bytes" not "chars" but we're
// using this to get the intuition for the process

// this takes a list of bytes and puts them in a channel
func putCharsInChannel(chars []byte, c chan byte) {
    // iterate over the list of chars and add them
    for _, char := range chars {
        c <- char
    }
    // close the channel because it's at the end of a range
    close(c)
}

func main() {

    // make a chanell of bytes
    c := make(chan byte, 5)

    // make a list of bytes
    chars := []byte{'H', 'e', 'r', 'e', '!'}

    // put the chars in the channel
    go putCharsInChannel(chars, c)

    // iterate over the channel to see if it worked
    for i := range c {
        // print them as characters
        fmt.Printf("%c\n", i)
    }

}

