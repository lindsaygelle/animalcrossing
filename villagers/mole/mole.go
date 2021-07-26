
package mole

import (
	a "github.com/lindsaygelle/animalcrossing/animals/mole"
	s "github.com/lindsaygelle/animalcrossing/species/mole"
)

// Mole is an Animal Crossing villager type.
type Mole struct{}

func (v Mole) Animal() string {
	return (a.Mole{}).Id()
}

func (v Mole) Species() string {
	return (s.Mole{}).Id()
}
