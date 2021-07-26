package alligator

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Alligator is an Animal Crossing animal type.
type Alligator struct{}

func (v Alligator) Alternate() {
	// TBC
}

func (v Alligator) Icon() {
	// TBC
}

func (v Alligator) Id() string {
	return "alligator"
}

func (v Alligator) Fictional() bool {
	return false
}

func (v Alligator) Languages() {
	// TBC
}

func (v Alligator) Relatives() {
	// TBC
}

func (v Alligator) Special() bool {
	return false
}

var (
	_ a.Animal = (Alligator{})
)
