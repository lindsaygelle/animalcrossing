package eagle

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Eagle is an Animal Crossing animal type.
type Eagle struct{}

func (v Eagle) Alternate() {
	// TBC
}

func (v Eagle) Icon() {
	// TBC
}

func (v Eagle) Id() string {
	return "eagle"
}

func (v Eagle) Fictional() bool {
	return false
}

func (v Eagle) Languages() {
	// TBC
}

func (v Eagle) Relatives() {
	// TBC
}

func (v Eagle) Special() bool {
	return false
}

var (
	_ a.Animal = (Eagle{})
)
