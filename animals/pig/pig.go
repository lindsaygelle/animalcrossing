package pig

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Pig is an Animal Crossing animal type.
type Pig struct{}

func (v Pig) Alternate() {
	// TBC
}

func (v Pig) Icon() {
	// TBC
}

func (v Pig) Id() string {
	return "pig"
}

func (v Pig) Fictional() bool {
	return false
}

func (v Pig) Languages() {
	// TBC
}

func (v Pig) Relatives() {
	// TBC
}

func (v Pig) Special() bool {
	return false
}

var (
	_ a.Animal = (Pig{})
)
