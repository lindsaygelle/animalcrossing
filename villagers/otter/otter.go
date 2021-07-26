
package otter

import (
	a "github.com/lindsaygelle/animalcrossing/animals/otter"
	s "github.com/lindsaygelle/animalcrossing/species/otter"
)

// Otter is an Animal Crossing villager type.
type Otter struct{}

func (v Otter) Animal() string {
	return (a.Otter{}).Id()
}

func (v Otter) Species() string {
	return (s.Otter{}).Id()
}
