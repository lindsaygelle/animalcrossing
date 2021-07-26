
package octopus

import (
	a "github.com/lindsaygelle/animalcrossing/animals/octopus"
	s "github.com/lindsaygelle/animalcrossing/species/octopus"
)

// Octopus is an Animal Crossing villager type.
type Octopus struct{}

func (v Octopus) Animal() string {
	return (a.Octopus{}).Id()
}

func (v Octopus) Species() string {
	return (s.Octopus{}).Id()
}
