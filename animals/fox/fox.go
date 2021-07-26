package fox

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Fox is an Animal Crossing animal type.
type Fox struct{}

func (v Fox) Alternate() {
	// TBC
}

func (v Fox) Icon() {
	// TBC
}

func (v Fox) Id() string {
	return "fox"
}

func (v Fox) Fictional() bool {
	return false
}

func (v Fox) Languages() {
	// TBC
}

func (v Fox) Relatives() {
	// TBC
}

func (v Fox) Special() bool {
	return true
}

var (
	_ a.Animal = (Fox{})
)
