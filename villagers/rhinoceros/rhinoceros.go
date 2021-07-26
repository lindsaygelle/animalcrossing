
package rhinoceros

import (
	a "github.com/lindsaygelle/animalcrossing/animals/rhinoceros"
	s "github.com/lindsaygelle/animalcrossing/species/rhinoceros"
)

// Rhinoceros is an Animal Crossing villager type.
type Rhinoceros struct{}

func (v Rhinoceros) Animal() string {
	return (a.Rhinoceros{}).Id()
}

func (v Rhinoceros) Species() string {
	return (s.Rhinoceros{}).Id()
}
