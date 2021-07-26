package raccoon

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Raccoon is an Animal Crossing animal type.
type Raccoon struct{}

func (v Raccoon) Alternate() {
	// TBC
}

func (v Raccoon) Icon() {
	// TBC
}

func (v Raccoon) Id() string {
	return "raccoon"
}

func (v Raccoon) Fictional() bool {
	return false
}

func (v Raccoon) Languages() {
	// TBC
}

func (v Raccoon) Relatives() {
	// TBC
}

func (v Raccoon) Special() bool {
	return true
}

var (
	_ a.Animal = (Raccoon{})
)
