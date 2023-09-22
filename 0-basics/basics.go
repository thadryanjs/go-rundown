package main

// ^ we need to specify the package name at the top of the file

import (
    "fmt"
    // importing local modules from this directory
    "basics/lib/submod1"
    "basics/lib/submod2"
    "basics/lib/functions"
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
}
