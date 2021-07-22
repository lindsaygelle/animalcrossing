package main

import (
	"fmt"

	"github.com/lindsaygelle/animalcrossing/animals/cat"
)

func main() {
	for _, t := range (cat.Cat{}).Translations() {
		fmt.Println(t.Name(), t.Gender(), t.Value())
	}
}
