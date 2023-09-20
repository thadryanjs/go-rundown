package functions

import (
    "fmt"
)

// takes a string called "name" and returns a string
func Greet(name string) string {

    // use the %s like back in the day lol
    return fmt.Sprintf("Hello, %s!", name)
}

// note that you must provide a type!
