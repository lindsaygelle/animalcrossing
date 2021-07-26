package furseal

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Furseal is an Animal Crossing animal type.
type Furseal struct{}

func (v Furseal) Alternate() {
	// TBC
}

func (v Furseal) Icon() {
	// TBC
}

func (v Furseal) Id() string {
	return "furseal"
}

func (v Furseal) Fictional() bool {
	return false
}

func (v Furseal) Languages() {
	// TBC
}

func (v Furseal) Relatives() {
	// TBC
}

func (v Furseal) Special() bool {
	return true
}

var (
	_ a.Animal = (Furseal{})
)
