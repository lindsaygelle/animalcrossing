
package tapir

import (
	a "github.com/lindsaygelle/animalcrossing/animals/tapir"
	s "github.com/lindsaygelle/animalcrossing/species/tapir"
)

// Tapir is an Animal Crossing villager type.
type Tapir struct{}

func (v Tapir) Animal() string {
	return (a.Tapir{}).Id()
}

func (v Tapir) Species() string {
	return (s.Tapir{}).Id()
}
