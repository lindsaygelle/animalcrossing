package monkey

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Monkey is an Animal Crossing animal type.
type Monkey struct{}

func (v Monkey) Alternate() {
	// TBC
}

func (v Monkey) Icon() {
	// TBC
}

func (v Monkey) Id() string {
	return "monkey"
}

func (v Monkey) Fictional() bool {
	return false
}

func (v Monkey) Languages() {
	// TBC
}

func (v Monkey) Relatives() {
	// TBC
}

func (v Monkey) Special() bool {
	return false
}

var (
	_ a.Animal = (Monkey{})
)
