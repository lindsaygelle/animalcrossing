
package anteater

import (
	a "github.com/lindsaygelle/animalcrossing/animals/anteater"
	s "github.com/lindsaygelle/animalcrossing/species/anteater"
)

// Anteater is an Animal Crossing villager type.
type Anteater struct{}

func (v Anteater) Animal() string {
	return (a.Anteater{}).Id()
}

func (v Anteater) Species() string {
	return (s.Anteater{}).Id()
}
