package walrus

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Walrus is an Animal Crossing animal type.
type Walrus struct{}

func (v Walrus) Alternate() {
	// TBC
}

func (v Walrus) Icon() {
	// TBC
}

func (v Walrus) Id() string {
	return "walrus"
}

func (v Walrus) Fictional() bool {
	return false
}

func (v Walrus) Languages() {
	// TBC
}

func (v Walrus) Relatives() {
	// TBC
}

func (v Walrus) Special() bool {
	return true
}

var (
	_ a.Animal = (Walrus{})
)
