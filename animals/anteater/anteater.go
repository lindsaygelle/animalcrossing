package anteater

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Anteater is an Animal Crossing animal type.
type Anteater struct{}

func (v Anteater) Alternate() {
	// TBC
}

func (v Anteater) Icon() {
	// TBC
}

func (v Anteater) Id() string {
	return "anteater"
}

func (v Anteater) Fictional() bool {
	return false
}

func (v Anteater) Languages() {
	// TBC
}

func (v Anteater) Relatives() {
	// TBC
}

func (v Anteater) Special() bool {
	return false
}

var (
	_ a.Animal = (Anteater{})
)
