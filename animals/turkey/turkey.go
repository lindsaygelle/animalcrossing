package turkey

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Turkey is an Animal Crossing animal type.
type Turkey struct{}

func (v Turkey) Alternate() {
	// TBC
}

func (v Turkey) Icon() {
	// TBC
}

func (v Turkey) Id() string {
	return "turkey"
}

func (v Turkey) Fictional() bool {
	return false
}

func (v Turkey) Languages() {
	// TBC
}

func (v Turkey) Relatives() {
	// TBC
}

func (v Turkey) Special() bool {
	return true
}

var (
	_ a.Animal = (Turkey{})
)
