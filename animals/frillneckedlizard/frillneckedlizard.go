package frillneckedlizard

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Frillneckedlizard is an Animal Crossing animal type.
type Frillneckedlizard struct{}

func (v Frillneckedlizard) Alternate() {
	// TBC
}

func (v Frillneckedlizard) Icon() {
	// TBC
}

func (v Frillneckedlizard) Id() string {
	return "frillneckedlizard"
}

func (v Frillneckedlizard) Fictional() bool {
	return false
}

func (v Frillneckedlizard) Languages() {
	// TBC
}

func (v Frillneckedlizard) Relatives() {
	// TBC
}

func (v Frillneckedlizard) Special() bool {
	return true
}

var (
	_ a.Animal = (Frillneckedlizard{})
)
