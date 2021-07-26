package dog

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Dog is an Animal Crossing animal type.
type Dog struct{}

func (v Dog) Alternate() {
	// TBC
}

func (v Dog) Icon() {
	// TBC
}

func (v Dog) Id() string {
	return "dog"
}

func (v Dog) Fictional() bool {
	return false
}

func (v Dog) Languages() {
	// TBC
}

func (v Dog) Relatives() {
	// TBC
}

func (v Dog) Special() bool {
	return false
}

var (
	_ a.Animal = (Dog{})
)
