package hamster

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Hamster is an Animal Crossing animal type.
type Hamster struct{}

func (v Hamster) Alternate() {
	// TBC
}

func (v Hamster) Icon() {
	// TBC
}

func (v Hamster) Id() string {
	return "hamster"
}

func (v Hamster) Fictional() bool {
	return false
}

func (v Hamster) Languages() {
	// TBC
}

func (v Hamster) Relatives() {
	// TBC
}

func (v Hamster) Special() bool {
	return false
}

var (
	_ a.Animal = (Hamster{})
)
