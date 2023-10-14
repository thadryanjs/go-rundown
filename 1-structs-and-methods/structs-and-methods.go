package main

import (
    "fmt"
)

/* Structs and Methods */

// Go uses "structs" to organize data. They are simple collections of "fields".
// The can be declared like this:
type Player struct {
    Name string
    Age int
}


// Structs have methods associated with them ("methods on types")
// The reason we revied pointers in the last module is because we need to pass a pointer to the struct

// We define a method on the structure
func (p *Player) GetName() string {
    return p.Name
}

// We define another method on the structure
func(p* Player) IntroduceSelf() {
    fmt.Println("Hi, I'm " + p.Name)
}

// Note the instead of being in a class with a built int method, Go "knows" this is a method
// for Player because of the type of the pointer we're passing to it

func main() {
    // create a new player
    somePlayer := Player{"Rebecca", 34}
    // "void" function
    somePlayer.IntroduceSelf()
    // function that returns a string
    fmt.Println(somePlayer.GetName())
}
