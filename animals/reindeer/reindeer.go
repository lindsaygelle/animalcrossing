package reindeer

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Reindeer is an Animal Crossing animal type.
type Reindeer struct{}

func (v Reindeer) Alternate() {
	// TBC
}

func (v Reindeer) Icon() {
	// TBC
}

func (v Reindeer) Id() string {
	return "reindeer"
}

func (v Reindeer) Fictional() bool {
	return false
}

func (v Reindeer) Languages() {
	// TBC
}

func (v Reindeer) Relatives() {
	// TBC
}

func (v Reindeer) Special() bool {
	return true
}

var (
	_ a.Animal = (Reindeer{})
)

