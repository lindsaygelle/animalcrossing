
package pigeon

import (
	a "github.com/lindsaygelle/animalcrossing/animals/pigeon"
	s "github.com/lindsaygelle/animalcrossing/species/pigeon"
)

// Pigeon is an Animal Crossing villager type.
type Pigeon struct{}

func (v Pigeon) Animal() string {
	return (a.Pigeon{}).Id()
}

func (v Pigeon) Species() string {
	return (s.Pigeon{}).Id()
}
