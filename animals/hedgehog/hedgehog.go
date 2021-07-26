package hedgehog

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Hedgehog is an Animal Crossing animal type.
type Hedgehog struct{}

func (v Hedgehog) Alternate() {
	// TBC
}

func (v Hedgehog) Icon() {
	// TBC
}

func (v Hedgehog) Id() string {
	return "hedgehog"
}

func (v Hedgehog) Fictional() bool {
	return false
}

func (v Hedgehog) Languages() {
	// TBC
}

func (v Hedgehog) Relatives() {
	// TBC
}

func (v Hedgehog) Special() bool {
	return true
}

var (
	_ a.Animal = (Hedgehog{})
)
