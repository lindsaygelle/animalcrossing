package cow

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Cow is an Animal Crossing animal type.
type Cow struct{}

func (v Cow) Alternate() {
	// TBC
}

func (v Cow) Icon() {
	// TBC
}

func (v Cow) Id() string {
	return "cow"
}

func (v Cow) Fictional() bool {
	return false
}

func (v Cow) Languages() {
	// TBC
}

func (v Cow) Relatives() {
	// TBC
}

func (v Cow) Special() bool {
	return false
}

var (
	_ a.Animal = (Cow{})
)
