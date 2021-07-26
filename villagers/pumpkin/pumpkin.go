
package pumpkin

import (
	a "github.com/lindsaygelle/animalcrossing/animals/pumpkin"
	s "github.com/lindsaygelle/animalcrossing/species/pumpkin"
)

// Pumpkin is an Animal Crossing villager type.
type Pumpkin struct{}

func (v Pumpkin) Animal() string {
	return (a.Pumpkin{}).Id()
}

func (v Pumpkin) Species() string {
	return (s.Pumpkin{}).Id()
}
