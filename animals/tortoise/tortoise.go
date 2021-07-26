package tortoise

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Tortoise is an Animal Crossing animal type.
type Tortoise struct{}

func (v Tortoise) Alternate() {
	// TBC
}

func (v Tortoise) Icon() {
	// TBC
}

func (v Tortoise) Id() string {
	return "tortoise"
}

func (v Tortoise) Fictional() bool {
	return false
}

func (v Tortoise) Languages() {
	// TBC
}

func (v Tortoise) Relatives() {
	// TBC
}

func (v Tortoise) Special() bool {
	return true
}

var (
	_ a.Animal = (Tortoise{})
)
