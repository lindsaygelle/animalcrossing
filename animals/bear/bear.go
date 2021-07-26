package bear

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Bear is an Animal Crossing animal type.
type Bear struct{}

func (v Bear) Alternate() {
	// TBC
}

func (v Bear) Icon() {
	// TBC
}

func (v Bear) Id() string {
	return "bear"
}

func (v Bear) Fictional() bool {
	return false
}

func (v Bear) Languages() {
	// TBC
}

func (v Bear) Relatives() {
	// TBC
}

func (v Bear) Special() bool {
	return false
}

var (
	_ a.Animal = (Bear{})
)
