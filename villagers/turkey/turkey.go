
package turkey

import (
	a "github.com/lindsaygelle/animalcrossing/animals/turkey"
	s "github.com/lindsaygelle/animalcrossing/species/turkey"
)

// Turkey is an Animal Crossing villager type.
type Turkey struct{}

func (v Turkey) Animal() string {
	return (a.Turkey{}).Id()
}

func (v Turkey) Species() string {
	return (s.Turkey{}).Id()
}
