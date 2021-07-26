
package furseal

import (
	a "github.com/lindsaygelle/animalcrossing/animals/furseal"
	s "github.com/lindsaygelle/animalcrossing/species/furseal"
)

// Furseal is an Animal Crossing villager type.
type Furseal struct{}

func (v Furseal) Animal() string {
	return (a.Furseal{}).Id()
}

func (v Furseal) Species() string {
	return (s.Furseal{}).Id()
}
