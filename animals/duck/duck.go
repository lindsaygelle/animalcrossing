package duck

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Duck is an Animal Crossing animal type.
type Duck struct{}

func (v Duck) Alternate() {
	// TBC
}

func (v Duck) Icon() {
	// TBC
}

func (v Duck) Id() string {
	return "duck"
}

func (v Duck) Fictional() bool {
	return false
}

func (v Duck) Languages() {
	// TBC
}

func (v Duck) Relatives() {
	// TBC
}

func (v Duck) Special() bool {
	return false
}

var (
	_ a.Animal = (Duck{})
)
