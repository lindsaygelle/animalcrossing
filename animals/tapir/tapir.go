package tapir

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Tapir is an Animal Crossing animal type.
type Tapir struct{}

func (v Tapir) Alternate() {
	// TBC
}

func (v Tapir) Icon() {
	// TBC
}

func (v Tapir) Id() string {
	return "tapir"
}

func (v Tapir) Fictional() bool {
	return false
}

func (v Tapir) Languages() {
	// TBC
}

func (v Tapir) Relatives() {
	// TBC
}

func (v Tapir) Special() bool {
	return true
}

var (
	_ a.Animal = (Tapir{})
)
