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

// Plant is a struct
type Plant struct {
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
	animal := Animal{"ほげほげ"}
	dog := Dog{Animal{"タロ"}}
	cat := Cat{Animal{"タマ"}}
	doWalk(animal)
	doWalk(dog)
	doWalk(cat)

	// it will fail
	// plant := Plant{"ふがふが"}
	// doWalk(plant)

	japanese := Japanese{Animal{"鈴木一郎"}}
	doWalk(japanese)
	japanese.Animal.Walk()
}
