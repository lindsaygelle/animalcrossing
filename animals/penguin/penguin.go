package penguin

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Penguin is an Animal Crossing animal type.
type Penguin struct{}

func (v Penguin) Alternate() {
	// TBC
}

func (v Penguin) Icon() {
	// TBC
}

func (v Penguin) Id() string {
	return "penguin"
}

func (v Penguin) Fictional() bool {
	return false
}

func (v Penguin) Languages() {
	// TBC
}

func (v Penguin) Relatives() {
	// TBC
}

func (v Penguin) Special() bool {
	return false
}

var (
	_ a.Animal = (Penguin{})
)
