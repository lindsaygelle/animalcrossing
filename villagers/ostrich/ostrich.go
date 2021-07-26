
package ostrich

import (
	a "github.com/lindsaygelle/animalcrossing/animals/ostrich"
	s "github.com/lindsaygelle/animalcrossing/species/ostrich"
)

// Ostrich is an Animal Crossing villager type.
type Ostrich struct{}

func (v Ostrich) Animal() string {
	return (a.Ostrich{}).Id()
}

func (v Ostrich) Species() string {
	return (s.Ostrich{}).Id()
}
