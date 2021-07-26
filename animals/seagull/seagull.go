package seagull

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Seagull is an Animal Crossing animal type.
type Seagull struct{}

func (v Seagull) Alternate() {
	// TBC
}

func (v Seagull) Icon() {
	// TBC
}

func (v Seagull) Id() string {
	return "seagull"
}

func (v Seagull) Fictional() bool {
	return false
}

func (v Seagull) Languages() {
	// TBC
}

func (v Seagull) Relatives() {
	// TBC
}

func (v Seagull) Special() bool {
	return true
}

var (
	_ a.Animal = (Seagull{})
)
