
package goat

import (
	a "github.com/lindsaygelle/animalcrossing/animals/goat"
	s "github.com/lindsaygelle/animalcrossing/species/goat"
)

// Goat is an Animal Crossing villager type.
type Goat struct{}

func (v Goat) Animal() string {
	return (a.Goat{}).Id()
}

func (v Goat) Species() string {
	return (s.Goat{}).Id()
}
