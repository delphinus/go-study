package main

import (
	"fmt"
)

// Animal is a struct
type Animal struct {
	Name string
}

// Walker is an interface
type Walker interface {
	Walk()
}

// Walk will walk
func (a Animal) Walk() {
	fmt.Printf("%s said: Let's walk!\n", a.Name)
}

// Plant is a struct
type Plant struct {
	Name string
}

func doWalk(walker Walker) {
	walker.Walk()
}

func main() {
	// it will fail
	// plant := Plant{"ふがふが"}
	// doWalk(plant)
}
