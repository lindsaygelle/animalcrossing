package bearcub

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Bearcub is an Animal Crossing animal type.
type Bearcub struct{}

func (v Bearcub) Alternate() {
	// TBC
}

func (v Bearcub) Icon() {
	// TBC
}

func (v Bearcub) Id() string {
	return "bearcub"
}

func (v Bearcub) Fictional() bool {
	return false
}

func (v Bearcub) Languages() {
	// TBC
}

func (v Bearcub) Relatives() {
	// TBC
}

func (v Bearcub) Special() bool {
	return false
}

var (
	_ a.Animal = (Bearcub{})
)
