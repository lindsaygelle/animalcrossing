package kangaroo

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Kangaroo is an Animal Crossing animal type.
type Kangaroo struct{}

func (v Kangaroo) Alternate() {
	// TBC
}

func (v Kangaroo) Icon() {
	// TBC
}

func (v Kangaroo) Id() string {
	return "kangaroo"
}

func (v Kangaroo) Fictional() bool {
	return false
}

func (v Kangaroo) Languages() {
	// TBC
}

func (v Kangaroo) Relatives() {
	// TBC
}

func (v Kangaroo) Special() bool {
	return false
}

var (
	_ a.Animal = (Kangaroo{})
)
