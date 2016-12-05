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

// Japanese is a struct
type Japanese struct {
	Animal
}

// Walk will walk
func (j Japanese) Walk() {
	fmt.Printf("%s は言った「歩こう！」\n", j.Name)
}

func doWalk(walker Walker) {
	walker.Walk()
}

func main() {
	japanese := Japanese{Animal{"鈴木一郎"}}
	doWalk(japanese)
	// japanese.Walk()
	doWalk(japanese.Animal)
	// japanese.Animal.Walk()
}
