package chicken

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Chicken is an Animal Crossing animal type.
type Chicken struct{}

func (v Chicken) Alternate() {
	// TBC
}

func (v Chicken) Icon() {
	// TBC
}

func (v Chicken) Id() string {
	return "chicken"
}

func (v Chicken) Fictional() bool {
	return false
}

func (v Chicken) Languages() {
	// TBC
}

func (v Chicken) Relatives() {
	// TBC
}

func (v Chicken) Special() bool {
	return false
}

var (
	_ a.Animal = (Chicken{})
)
