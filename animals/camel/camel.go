package camel

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Camel is an Animal Crossing animal type.
type Camel struct{}

func (v Camel) Alternate() {
	// TBC
}

func (v Camel) Icon() {
	// TBC
}

func (v Camel) Id() string {
	return "camel"
}

func (v Camel) Fictional() bool {
	return false
}

func (v Camel) Languages() {
	// TBC
}

func (v Camel) Relatives() {
	// TBC
}

func (v Camel) Special() bool {
	return true
}

var (
	_ a.Animal = (Camel{})
)
