package gyroid

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Gyroid is an Animal Crossing animal type.
type Gyroid struct{}

func (v Gyroid) Alternate() {
	// TBC
}

func (v Gyroid) Icon() {
	// TBC
}

func (v Gyroid) Id() string {
	return "gyroid"
}

func (v Gyroid) Fictional() bool {
	return true
}

func (v Gyroid) Languages() {
	// TBC
}

func (v Gyroid) Relatives() {
	// TBC
}

func (v Gyroid) Special() bool {
	return true
}

var (
	_ a.Animal = (Gyroid{})
)
