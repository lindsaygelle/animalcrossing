
package boar

import (
	a "github.com/lindsaygelle/animalcrossing/animals/boar"
	s "github.com/lindsaygelle/animalcrossing/species/boar"
)

// Boar is an Animal Crossing villager type.
type Boar struct{}

func (v Boar) Animal() string {
	return (a.Boar{}).Id()
}

func (v Boar) Species() string {
	return (s.Boar{}).Id()
}
