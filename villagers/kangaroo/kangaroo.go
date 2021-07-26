
package kangaroo

import (
	a "github.com/lindsaygelle/animalcrossing/animals/kangaroo"
	s "github.com/lindsaygelle/animalcrossing/species/kangaroo"
)

// Kangaroo is an Animal Crossing villager type.
type Kangaroo struct{}

func (v Kangaroo) Animal() string {
	return (a.Kangaroo{}).Id()
}

func (v Kangaroo) Species() string {
	return (s.Kangaroo{}).Id()
}
