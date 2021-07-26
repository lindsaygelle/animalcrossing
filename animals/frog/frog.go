package frog

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Frog is an Animal Crossing animal type.
type Frog struct{}

func (v Frog) Alternate() {
	// TBC
}

func (v Frog) Icon() {
	// TBC
}

func (v Frog) Id() string {
	return "frog"
}

func (v Frog) Fictional() bool {
	return false
}

func (v Frog) Languages() {
	// TBC
}

func (v Frog) Relatives() {
	// TBC
}

func (v Frog) Special() bool {
	return false
}

var (
	_ a.Animal = (Frog{})
)
