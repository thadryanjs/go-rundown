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
