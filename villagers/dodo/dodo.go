
package dodo

import (
	a "github.com/lindsaygelle/animalcrossing/animals/dodo"
	s "github.com/lindsaygelle/animalcrossing/species/dodo"
)

// Dodo is an Animal Crossing villager type.
type Dodo struct{}

func (v Dodo) Animal() string {
	return (a.Dodo{}).Id()
}

func (v Dodo) Species() string {
	return (s.Dodo{}).Id()
}
