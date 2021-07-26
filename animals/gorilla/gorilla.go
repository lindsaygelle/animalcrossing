package gorilla

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Gorilla is an Animal Crossing animal type.
type Gorilla struct{}

func (v Gorilla) Alternate() {
	// TBC
}

func (v Gorilla) Icon() {
	// TBC
}

func (v Gorilla) Id() string {
	return "gorilla"
}

func (v Gorilla) Fictional() bool {
	return false
}

func (v Gorilla) Languages() {
	// TBC
}

func (v Gorilla) Relatives() {
	// TBC
}

func (v Gorilla) Special() bool {
	return false
}

var (
	_ a.Animal = (Gorilla{})
)
