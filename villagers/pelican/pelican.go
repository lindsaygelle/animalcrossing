
package pelican

import (
	a "github.com/lindsaygelle/animalcrossing/animals/pelican"
	s "github.com/lindsaygelle/animalcrossing/species/pelican"
)

// Pelican is an Animal Crossing villager type.
type Pelican struct{}

func (v Pelican) Animal() string {
	return (a.Pelican{}).Id()
}

func (v Pelican) Species() string {
	return (s.Pelican{}).Id()
}
