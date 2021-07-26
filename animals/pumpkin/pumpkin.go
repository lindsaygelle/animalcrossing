package pumpkin

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Pumpkin is an Animal Crossing animal type.
type Pumpkin struct{}

func (v Pumpkin) Alternate() {
	// TBC
}

func (v Pumpkin) Icon() {
	// TBC
}

func (v Pumpkin) Id() string {
	return "pumpkin"
}

func (v Pumpkin) Fictional() bool {
	return true
}

func (v Pumpkin) Languages() {
	// TBC
}

func (v Pumpkin) Relatives() {
	// TBC
}

func (v Pumpkin) Special() bool {
	return true
}

var (
	_ a.Animal = (Pumpkin{})
)
