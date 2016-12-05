package main

import (
	"fmt"
)

// Walker is an interface
type Walker interface {
	Walk()
}

// Animal is a struct
type Animal struct {
	Name string
}

// Walk will walk
func (a Animal) Walk() {
	fmt.Printf("%s said: Let's walk!\n", a.Name)
}

// Dog is a struct
type Dog struct {
	Animal
}

// Cat is a struct
type Cat struct {
	Animal
}

// Japanese is a struct
type Japanese struct {
	Animal
}

// Walk will walk
func (j Japanese) Walk() {
	fmt.Printf("%s は言った「歩こう！」\n", j.Name)
}

func main() {
	dog := Dog{Animal{"タロ"}}
	cat := Cat{Animal{"タマ"}}
	doWalk(dog)
	doWalk(cat)

	japanese := Japanese{Animal{"鈴木一郎"}}
	doWalk(japanese)
	japanese.Animal.Walk()
}

func doWalk(walker Walker) {
	walker.Walk()
}
