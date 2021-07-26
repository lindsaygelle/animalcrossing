package octopus

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Octopus is an Animal Crossing animal type.
type Octopus struct{}

func (v Octopus) Alternate() {
	// TBC
}

func (v Octopus) Icon() {
	// TBC
}

func (v Octopus) Id() string {
	return "octopus"
}

func (v Octopus) Fictional() bool {
	return false
}

func (v Octopus) Languages() {
	// TBC
}

func (v Octopus) Relatives() {
	// TBC
}

func (v Octopus) Special() bool {
	return false
}

var (
	_ a.Animal = (Octopus{})
)
