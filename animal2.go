package main

import (
	"fmt"
)

// Animal is a struct
type Animal struct {
	Name string
}

// Dog is a struct
type Dog struct {
	Animal
}

// Cat is a struct
type Cat struct {
	Animal
}

// Walker is an interface
type Walker interface {
	Walk()
}

// Walk will walk
func (a Animal) Walk() {
	fmt.Printf("%s said: Let's walk!\n", a.Name)
}

func doWalk(walker Walker) {
	walker.Walk()
}

func main() {
	dog := Dog{Animal{"タロ"}}
	cat := Cat{Animal{"タマ"}}
	doWalk(dog)
	doWalk(cat)
}
