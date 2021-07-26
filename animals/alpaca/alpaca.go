package alpaca

import (
	a "github.com/lindsaygelle/animalcrossing/animals"
)

// Alpaca is an Animal Crossing animal type.
type Alpaca struct{}

func (v Alpaca) Alternate() {
	// TBC
}

func (v Alpaca) Icon() {
	// TBC
}

func (v Alpaca) Id() string {
	return "alpaca"
}

func (v Alpaca) Fictional() bool {
	return false
}

func (v Alpaca) Languages() {
	// TBC
}

func (v Alpaca) Relatives() {
	// TBC
}

func (v Alpaca) Special() bool {
	return true
}

var (
	_ a.Animal = (Alpaca{})
)
