package goat

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Goat is an Animal Crossing animal type.
type Goat struct{}

func (v Goat) Alternate() {
	// TBC
}

func (v Goat) Icon() {
	// TBC
}

func (v Goat) Id() string {
	return "goat"
}

func (v Goat) Fictional() bool {
	return false
}

func (v Goat) Languages() {
	// TBC
}

func (v Goat) Relatives() {
	// TBC
}

func (v Goat) Special() bool {
	return false
}

var (
	_ a.Animal = (Goat{})
)
