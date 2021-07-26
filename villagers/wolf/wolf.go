
package wolf

import (
	a "github.com/lindsaygelle/animalcrossing/animals/wolf"
	s "github.com/lindsaygelle/animalcrossing/species/wolf"
)

// Wolf is an Animal Crossing villager type.
type Wolf struct{}

func (v Wolf) Animal() string {
	return (a.Wolf{}).Id()
}

func (v Wolf) Species() string {
	return (s.Wolf{}).Id()
}
