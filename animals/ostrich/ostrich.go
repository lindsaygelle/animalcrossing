package ostrich

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Ostrich is an Animal Crossing animal type.
type Ostrich struct{}

func (v Ostrich) Alternate() {
	// TBC
}

func (v Ostrich) Icon() {
	// TBC
}

func (v Ostrich) Id() string {
	return "ostrich"
}

func (v Ostrich) Fictional() bool {
	return false
}

func (v Ostrich) Languages() {
	// TBC
}

func (v Ostrich) Relatives() {
	// TBC
}

func (v Ostrich) Special() bool {
	return false
}

var (
	_ a.Animal = (Ostrich{})
)
