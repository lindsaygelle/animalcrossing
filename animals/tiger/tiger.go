package tiger

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Tiger is an Animal Crossing animal type.
type Tiger struct{}

func (v Tiger) Alternate() {
	// TBC
}

func (v Tiger) Icon() {
	// TBC
}

func (v Tiger) Id() string {
	return "tiger"
}

func (v Tiger) Fictional() bool {
	return false
}

func (v Tiger) Languages() {
	// TBC
}

func (v Tiger) Relatives() {
	// TBC
}

func (v Tiger) Special() bool {
	return false
}

var (
	_ a.Animal = (Tiger{})
)
