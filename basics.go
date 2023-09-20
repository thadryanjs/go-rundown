package main

// ^ we need to specify the package name at the top of the file

import (
    "fmt"
    // importing local modules from this directory
    "basics/lib/submod1"
    "basics/lib/submod2"
    "basics/lib/functions"
    "basics/lib/players"
    "basics/lib/animals"
)

/* you can also import like this:

```go
import "fmt"
import "basics/lib/submod1"
import "basics/lib/submod2"
```

...but it's considered good form to use the other way.

Notice that we don't need a , on each line, which is nice. This makes it easier to
comment out lines when testing.
*/


// As you may have noticed, single-line comments are like this.

// multiline comments are like this:

/* <- start here

muiltiline comments!

end here -> */


// The main function is where the program starts
func main() {

    /* variables and printing, etc. */

    // In keeping with tradition
    fmt.Println("Hello, World!")

    // Make sure imported our local modules ok
    submod1.TestSubmod1()
    submod2.TestSubmod2()

    /* You will notice the capitalization of the function name
    this is because it's a "public" function, that can be accessed
    from outside the package. This is opposed to using the "export"
    keyword like in other languages. It kind of makes me nuts, but
    I will probably get used to it.
    */

    // Delcare a variable
    testName := "Marty McFly"

    /* The "walrus" operator is := and it is used to declare and assign
    a variable at the same time. It is only available inside functions
    */

    // the other way we would declare a variable is like this:
    var testAge int = 17
    // This would be read as "the variable testAge is of type int and is 17"

    // The basic way to print a variable:
    fmt.Println(testAge)

    // This is a test to make sure I wrote and imported the functions module correctly
    fmt.Println(functions.Greet(testName))


    /* Arrays & Slices */

    // An array of type string. Note the you do need to specify the size and it can't grow
    names := [3]string{"Marty", "Doc", "Einstein"}

    // for loops look like a cleaner version of the idiom used in C, C++, and Java, etc.

    // iterate over and call the function on each name
    for i := 0; i < len(names); i++ {
        fmt.Println(functions.Greet(names[i]))
    }

    // If you haven't seen something like this before, it's making a variable i, setting it to 0,
    // comparing it to the length of the array, comparing it to the length of the array, and incrementing
    // that variable by 1 each time the loop runs (this is the i++ part). When i is not lower than the
    // length of the array, the loop stops.

    // slices are like arrays, but you don't need to specify the size and they can grow
    dynamicList := []string{"Here", "There", "Everywhere"}
    // see the contents of the slice
    fmt.Println(dynamicList)
    // add an item to the slice
    dynamicList = append(dynamicList, "Somewhere")
    // see the contents of the slice again
    fmt.Println(dynamicList)
    // I rememver this as "append to list, this item" because I'm used to Python
    // where append is a method of the array, ie, array.append("item")


    /* Structs and Pointers */

    // Go uses "structs" to organize data. They are simple collections of "fields".
    // The can be declared like this:
    type Person struct {
        Name string
        Age int
    }
    // note that the types are specified after the field names

    // NOTE:
    // I'm going to use the struct from the players module for the rest of this section

    // create a struct (the source is in lib/player/player.go)
    somePlayer := players.Player{Name: "Marty", Age: 17}

    // print the struct
    fmt.Println(somePlayer)

    /* Go uses methods on types instead of classes with methods, but before we do that let's review pointers because I don't use them often enough to remember how they work all the time. Essentially the pointer refers to the memory address of the variable, not the value of it. It's sometimes called a "reference" because it refers to the variable. I remeber this by thinking of the variable as a house and the pointer as the address of that house. If you only know the address of the house, you can get to it, but don't have access to the contents without looking it up. This is called "dereferencing" the pointer because it's not just a reference anymore, it's the actual value.

    We use the & operator to get the address of a variable and the * operator to dereference it.

    It's sometimes put as & "points" to the variable. I think of it as a windy arrow to remember it.

    */

    // create a variable
    x := 5
    // y "points" to x
    y := &x
    // print the value of x "through" the pointer
    fmt.Println(*y)
    // if we had just printed y, it would have printed the address of x, like this:
    fmt.Println(y)

    // I remember these operators as & being a windy arrow that points
    // to the value and * being a portal you get to the value "through"

    // change the value of x "through" the pointers
    *y = 10

    // print the value of x to see it's differnt
    fmt.Println(x)

    // methods on types

    // The reason this is relevant tp structs is because we need to pass a pointer to the struct,
    // rather than the whole thing, to the methods on the struct. Go doesn't support nested functions,
    // and this program is running in "main", so I can't define them here. The source code is in
    // lib/players/players.go if you want to mess with it, and it looks like this:

    /*
    package players

    import (
        "fmt"
    )

    // define a structure
    type Player struct {
        Name string
        Age int
    }

    // define a method on the structure
    func (p *Player) GetName() string {
        return p.Name
    }

    // define another method on the structure
    func(p* Player) IntroduceSelf() {
        fmt.Println("Hi, I'm " + p.Name)
    }

    */

    // Note the instead of being in a class, Go "knows" this is a method for Player because
    // of the type of the pointer we're passing to it

    // "void" function
    somePlayer.IntroduceSelf()
    // function that returns a string
    fmt.Println(somePlayer.GetName())

    /* Interfaces */

    // Interfaces are a way to create different types that have different versions of the same behavior.
    // For instance, consider a dog, cat, and a duck. All of them make noise, but they all do so differently.
    // We can create an interface that animals make noise, and write a fuction that does the noise specific
    // to that animal. You can think of them all as "NoiseMakers" that have a different version of "MakeNoise"
    // The source code is in lib/animals/animals.go and looks like this:

    // make animals (they empty structs because I didn't need them to have a name or age)
    dog := animals.Dog{}
    cat := animals.Cat{}
    duck := animals.Duck{}

    // have them do their thing:
    dog.MakeNoise()
    cat.MakeNoise()
    duck.MakeNoise()

    // Why is this useful? We could just write a DogNoise(), CatNoise(), and DuckNoise() function, no?
    // This is true, but it becomes handy when we want to iterate over a list of animals and have them
    // make their noise. This way we dont have to have a different function for each animal in our loop.
    // we just tell them to "MakeNoise" and the do it in their own way.

    // make a list of animals
    animalList := []animals.Animal{&dog, &cat, &duck}
    for i := 0; i < len(animalList); i++ {
        animalList[i].MakeNoise()
    }
    // notice that we pass the pointer to the address here, because the functions in the interfaces
    // are expecting a pointer to the address of the struct, not the struct itself.
    // The source code is in lib/animals/animals.go and looks like this:

    /*

    package animals

    import (
        "fmt"
    )

    type Animal interface {
        MakeNoise()
    }


    type Dog struct {}

    func (d *Dog) MakeNoise() {
        fmt.Println("Woof!")
    }


    type Cat struct {}

    func (c *Cat) MakeNoise() {
        fmt.Println("Meow!")
    }


    type Duck struct {}

    func (d *Duck) MakeNoise() {
        fmt.Println("Quack!")
    }

    */

    // We specify the structs and them make methods that take pointers and make them do
    // whatever it is we want them to do.







}
