
package alpaca

import (
	a "github.com/lindsaygelle/animalcrossing/animals/alpaca"
	s "github.com/lindsaygelle/animalcrossing/species/alpaca"
)

// Alpaca is an Animal Crossing villager type.
type Alpaca struct{}

func (v Alpaca) Animal() string {
	return (a.Alpaca{}).Id()
}

func (v Alpaca) Species() string {
	return (s.Alpaca{}).Id()
}
