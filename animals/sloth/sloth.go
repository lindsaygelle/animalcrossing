package sloth

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Sloth is an Animal Crossing animal type.
type Sloth struct{}

func (v Sloth) Alternate() {
	// TBC
}

func (v Sloth) Icon() {
	// TBC
}

func (v Sloth) Id() string {
	return "sloth"
}

func (v Sloth) Fictional() bool {
	return false
}

func (v Sloth) Languages() {
	// TBC
}

func (v Sloth) Relatives() {
	// TBC
}

func (v Sloth) Special() bool {
	return true
}

var (
	_ a.Animal = (Sloth{})
)
