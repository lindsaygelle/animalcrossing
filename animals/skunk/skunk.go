package skunk

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Skunk is an Animal Crossing animal type.
type Skunk struct{}

func (v Skunk) Alternate() {
	// TBC
}

func (v Skunk) Icon() {
	// TBC
}

func (v Skunk) Id() string {
	return "skunk"
}

func (v Skunk) Fictional() bool {
	return false
}

func (v Skunk) Languages() {
	// TBC
}

func (v Skunk) Relatives() {
	// TBC
}

func (v Skunk) Special() bool {
	return true
}

var (
	_ a.Animal = (Skunk{})
)
