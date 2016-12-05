package main

import (
	"fmt"
)

// Animal is a struct
type Animal struct {
	Name string
}

// Walk will walk
func (a Animal) Walk() {
	fmt.Printf("%s said: Let's walk!\n", a.Name)
}

func doWalk(animal Animal) {
	animal.Walk()
}

func main() {
	animal := Animal{"ほげほげ"}
	doWalk(animal)
}
