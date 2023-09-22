package main

import (
    "fmt"
)

/* Interfaces */

// Interfaces are a way to create different types that have different versions of the same behavior.
// For instance, consider a dog, cat, and a duck. All of them make noise, but they all do so differently.
// We can create an interface that animals make noise, and write a fuction that does the noise specific
// to that animal. You can think of them all as "NoiseMakers" that have a different version of "MakeNoise"
// The source code is in lib/animals/animals.go and looks like this:

// notice that we pass the pointer to the address here, because the functions in the interfaces
// are expecting a pointer to the address of the struct, not the struct itself.

// define an type with an interface
type Animal interface {
    // give it a method
    MakeNoise()
}


// define the dog struct - we leave it empty because we don't need it to have any fields
type Dog struct {}

// now we describe the way a day makes noise
func (d *Dog) MakeNoise() {
    // Make the dog say woof
    fmt.Println("Woof!")
}


// We do the same thing for the other animals
type Cat struct {}
// Same as above, but meowing
func (c *Cat) MakeNoise() {
    fmt.Println("Meow!")
}


// one more time for the duck
type Duck struct {}

// the same as above, but quacking
func (d *Duck) MakeNoise() {
    fmt.Println("Quack!")
}



func main() {

    // make animals
    dog := Dog{}
    cat := Cat{}
    duck := Duck{}

    // have them do their thing:
    dog.MakeNoise()
    cat.MakeNoise()
    duck.MakeNoise()

    // Why is this useful? We could just write a DogNoise(), CatNoise(), and DuckNoise() function, no?
    // This is true, but it becomes handy when we want to iterate over a list of animals and have them
    // make their noise. This way we dont have to have a different function for each animal in our loop.
    // we just tell them to "MakeNoise" and the do it in their own way.

    // make a list of references to animals
    animalList := []Animal{&dog, &cat, &duck}
    // iterate over the list and have them make noise, each in their own way
    for i := 0; i < len(animalList); i++ {
        animalList[i].MakeNoise()
    }

}
