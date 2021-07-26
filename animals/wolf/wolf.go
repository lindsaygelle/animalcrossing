package wolf

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Wolf is an Animal Crossing animal type.
type Wolf struct{}

func (v Wolf) Alternate() {
	// TBC
}

func (v Wolf) Icon() {
	// TBC
}

func (v Wolf) Id() string {
	return "wolf"
}

func (v Wolf) Fictional() bool {
	return false
}

func (v Wolf) Languages() {
	// TBC
}

func (v Wolf) Relatives() {
	// TBC
}

func (v Wolf) Special() bool {
	return false
}

var (
	_ a.Animal = (Wolf{})
)
