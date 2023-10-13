package main
// ^ we need to specify the package name at the top of the file


import (
    "fmt"
    "strings"
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
    fmt.Println("\n/* variables and printing, etc.")

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

    /* Conditionals */
    fmt.Println("\n/* Conditionals")

    // These are all pretty predictable

    number := 5
    if number > 5 {
        fmt.Println("The number is greater than 5")
    } else if number < 5 {
        fmt.Println("The number is less than 5")
    } else {
        fmt.Println("The number is 5")
    }


    // Go also has switch statement, which can be cleaner than if/else statements
    // when you have a lot of conditions. You provide a default case, which is like
    // the 'else' in an if/else statement.

    // I find these very nice and read them as "in case of, x, do this",

    switch number {
        case 5:
            fmt.Println("The number is 5")
        case 6:
            fmt.Println("The number is 6")
        default:
            fmt.Println("The number is neither 5 nor 6")
    }


    /* Manipulating strings */
    fmt.Println("\n/* Manipulating strings")


    fmt.Println("Hello, World!")

    // make a list of names
    namesList := []string{"Marty", "Doc", "Einstein"}

    // iterate over the list and call the function on each name
    for i := 0; i < len(namesList); i++ {
        fmt.Println(functions.Greet(namesList[i]))
    }

    // print that we're done
    fmt.Println("Done!")

    // Strings are immutable in Go, so you can't change them once they are created.
    // You make new strings based on the old ones.

    myName := "Thadryan"

    // How long is the string?
    fmt.Println(len(myName))

    // To iterate over a string, you can use a for loop in the stule of C, C++, Java, etc.
    for i := 0; i < len(myName); i++ {
        fmt.Println(string(myName[i]))
    }


    // Indexing contains some suprises for those used to Python, etc:
    // You will get the byte value of the character, not the character itself
    fmt.Println(myName[0])
    // to get the character, you need to cast it to a string
    fmt.Println("First letter is:", string(myName[0]))
    // You can also use a different print function that will do this for you
    fmt.Println("First letter is:", fmt.Sprintf("%c", myName[0]))

    // Make the string all uppercase
    fmt.Println(strings.ToUpper(myName))
    // Make it lowercase
    fmt.Println(strings.ToLower(myName))
    // Replace a substring
    fmt.Println(strings.Replace(myName, "y", "i", 1))

    // The 1 at the end of the replace function is the number of times to replace.
    // If you want to replace all instances, you can use -1.
    fmt.Println(strings.Replace(myName, "a", "!", -1))
    // This is a little confusing - you'd expect something all=True
    // Kind of obnoxious, but there is probably some performance reason for it or something


    /* Arrays & Slices */
    fmt.Println("\n/* Arrays & Slices")

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


    /* Maps */
    fmt.Println("\n/* Maps")

    // Maps/hashs/dictionary structures are like are called maps in go and behave
    // in predictable way. Note you need to specify the type of the key and the value.

    // declare a map of type string to string
    namesMap := map[string]string{
        "Kerry": "Bishe",
        "Lee": "Pace",
        "Mackenzie": "Davis",
    }

    // look up a value in the map
    fmt.Println(namesMap["Kerry"])

    // add a value to the map
    namesMap["Scoot"] = "McNairy"
    fmt.Println(namesMap["Scoot"])

    // delete a value from the map
    delete(namesMap, "Scoot")
    fmt.Println(namesMap["Scoot"])


    /* Pointers for people in a hurry */

    // This is a hasty overview of pointers because we need them, not an in-depth explanation.

    // for a better explanation of pointers: https://dev.to/ottonicasio/explain-c-pointers-like-im-five-36m8
    // Copilot suggested this video: https://www.youtube.com/watch?v=CF9S4QZuV30 when I was writing about

    // Essentially pointer refers to the memory address of the variable, not the value of the variable.
    // I think of this as the address of that house. If you only know the address of the house, you can
    // get to it, but don't have access to the contents without looking it up.

    // We use the & operator to get the address of a variable and the * operator to it.

    // It's sometimes put as & "points" to the variable. I think of it as a windy arrow to remember it.

    // poiners

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

    fmt.Println("\n/* Structs")

}
