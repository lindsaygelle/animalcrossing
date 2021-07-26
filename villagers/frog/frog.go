
package frog

import (
	a "github.com/lindsaygelle/animalcrossing/animals/frog"
	s "github.com/lindsaygelle/animalcrossing/species/frog"
)

// Frog is an Animal Crossing villager type.
type Frog struct{}

func (v Frog) Animal() string {
	return (a.Frog{}).Id()
}

func (v Frog) Species() string {
	return (s.Frog{}).Id()
}
