
package hedgehog

import (
	a "github.com/lindsaygelle/animalcrossing/animals/hedgehog"
	s "github.com/lindsaygelle/animalcrossing/species/hedgehog"
)

// Hedgehog is an Animal Crossing villager type.
type Hedgehog struct{}

func (v Hedgehog) Animal() string {
	return (a.Hedgehog{}).Id()
}

func (v Hedgehog) Species() string {
	return (s.Hedgehog{}).Id()
}
