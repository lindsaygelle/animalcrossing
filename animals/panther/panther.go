package panther

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Panther is an Animal Crossing animal type.
type Panther struct{}

func (v Panther) Alternate() {
	// TBC
}

func (v Panther) Icon() {
	// TBC
}

func (v Panther) Id() string {
	return "panther"
}

func (v Panther) Fictional() bool {
	return false
}

func (v Panther) Languages() {
	// TBC
}

func (v Panther) Relatives() {
	// TBC
}

func (v Panther) Special() bool {
	return true
}

var (
	_ a.Animal = (Panther{})
)
