package pigeon

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Pigeon is an Animal Crossing animal type.
type Pigeon struct{}

func (v Pigeon) Alternate() {
	// TBC
}

func (v Pigeon) Icon() {
	// TBC
}

func (v Pigeon) Id() string {
	return "pigeon"
}

func (v Pigeon) Fictional() bool {
	return false
}

func (v Pigeon) Languages() {
	// TBC
}

func (v Pigeon) Relatives() {
	// TBC
}

func (v Pigeon) Special() bool {
	return true
}

var (
	_ a.Animal = (Pigeon{})
)
