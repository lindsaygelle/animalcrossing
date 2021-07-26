package cat

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Cat is an Animal Crossing animal type.
type Cat struct{}

func (v Cat) Alternate() {
	// TBC
}

func (v Cat) Icon() {
	// TBC
}

func (v Cat) Id() string {
	return "cat"
}

func (v Cat) Fictional() bool {
	return false
}

func (v Cat) Languages() {
	// TBC
}

func (v Cat) Relatives() {
	// TBC
}

func (v Cat) Special() bool {
	return false
}

var (
	_ a.Animal = (Cat{})
)
