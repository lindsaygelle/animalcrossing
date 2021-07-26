
package panther

import (
	a "github.com/lindsaygelle/animalcrossing/animals/panther"
	s "github.com/lindsaygelle/animalcrossing/species/panther"
)

// Panther is an Animal Crossing villager type.
type Panther struct{}

func (v Panther) Animal() string {
	return (a.Panther{}).Id()
}

func (v Panther) Species() string {
	return (s.Panther{}).Id()
}
