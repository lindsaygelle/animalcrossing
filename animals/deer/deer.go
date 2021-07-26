package deer

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Deer is an Animal Crossing animal type.
type Deer struct{}

func (v Deer) Alternate() {
	// TBC
}

func (v Deer) Icon() {
	// TBC
}

func (v Deer) Id() string {
	return "deer"
}

func (v Deer) Fictional() bool {
	return false
}

func (v Deer) Languages() {
	// TBC
}

func (v Deer) Relatives() {
	// TBC
}

func (v Deer) Special() bool {
	return false
}

var (
	_ a.Animal = (Deer{})
)
