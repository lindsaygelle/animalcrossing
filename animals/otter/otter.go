package otter

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Otter is an Animal Crossing animal type.
type Otter struct{}

func (v Otter) Alternate() {
	// TBC
}

func (v Otter) Icon() {
	// TBC
}

func (v Otter) Id() string {
	return "otter"
}

func (v Otter) Fictional() bool {
	return false
}

func (v Otter) Languages() {
	// TBC
}

func (v Otter) Relatives() {
	// TBC
}

func (v Otter) Special() bool {
	return true
}

var (
	_ a.Animal = (Otter{})
)
