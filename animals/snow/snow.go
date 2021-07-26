package snow

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Snow is an Animal Crossing animal type.
type Snow struct{}

func (v Snow) Alternate() {
	// TBC
}

func (v Snow) Icon() {
	// TBC
}

func (v Snow) Id() string {
	return "snow"
}

func (v Snow) Fictional() bool {
	return true
}

func (v Snow) Languages() {
	// TBC
}

func (v Snow) Relatives() {
	// TBC
}

func (v Snow) Special() bool {
	return true
}

var (
	_ a.Animal = (Snow{})
)
