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
