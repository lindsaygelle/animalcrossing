
package gyroid

import (
	a "github.com/lindsaygelle/animalcrossing/animals/gyroid"
	s "github.com/lindsaygelle/animalcrossing/species/gyroid"
)

// Gyroid is an Animal Crossing villager type.
type Gyroid struct{}

func (v Gyroid) Animal() string {
	return (a.Gyroid{}).Id()
}

func (v Gyroid) Species() string {
	return (s.Gyroid{}).Id()
}
