
package chameleon

import (
	a "github.com/lindsaygelle/animalcrossing/animals/chameleon"
	s "github.com/lindsaygelle/animalcrossing/species/chameleon"
)

// Chameleon is an Animal Crossing villager type.
type Chameleon struct{}

func (v Chameleon) Animal() string {
	return (a.Chameleon{}).Id()
}

func (v Chameleon) Species() string {
	return (s.Chameleon{}).Id()
}
