
package hippopotamus

import (
	a "github.com/lindsaygelle/animalcrossing/animals/hippopotamus"
	s "github.com/lindsaygelle/animalcrossing/species/hippopotamus"
)

// Hippopotamus is an Animal Crossing villager type.
type Hippopotamus struct{}

func (v Hippopotamus) Animal() string {
	return (a.Hippopotamus{}).Id()
}

func (v Hippopotamus) Species() string {
	return (s.Hippopotamus{}).Id()
}
