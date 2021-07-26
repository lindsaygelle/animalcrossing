package koala

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Koala is an Animal Crossing animal type.
type Koala struct{}

func (v Koala) Alternate() {
	// TBC
}

func (v Koala) Icon() {
	// TBC
}

func (v Koala) Id() string {
	return "koala"
}

func (v Koala) Fictional() bool {
	return false
}

func (v Koala) Languages() {
	// TBC
}

func (v Koala) Relatives() {
	// TBC
}

func (v Koala) Special() bool {
	return false
}

var (
	_ a.Animal = (Koala{})
)
