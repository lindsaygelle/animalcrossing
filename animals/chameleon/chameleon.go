package chameleon

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Chameleon is an Animal Crossing animal type.
type Chameleon struct{}

func (v Chameleon) Alternate() {
	// TBC
}

func (v Chameleon) Icon() {
	// TBC
}

func (v Chameleon) Id() string {
	return "chameleon"
}

func (v Chameleon) Fictional() bool {
	return false
}

func (v Chameleon) Languages() {
	// TBC
}

func (v Chameleon) Relatives() {
	// TBC
}

func (v Chameleon) Special() bool {
	return true
}

var (
	_ a.Animal = (Chameleon{})
)
