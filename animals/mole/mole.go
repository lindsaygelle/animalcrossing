package mole

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Mole is an Animal Crossing animal type.
type Mole struct{}

func (v Mole) Alternate() {
	// TBC
}

func (v Mole) Icon() {
	// TBC
}

func (v Mole) Id() string {
	return "mole"
}

func (v Mole) Fictional() bool {
	return false
}

func (v Mole) Languages() {
	// TBC
}

func (v Mole) Relatives() {
	// TBC
}

func (v Mole) Special() bool {
	return true
}

var (
	_ a.Animal = (Mole{})
)
