
package raccoon

import (
	a "github.com/lindsaygelle/animalcrossing/animals/raccoon"
	s "github.com/lindsaygelle/animalcrossing/species/raccoon"
)

// Raccoon is an Animal Crossing villager type.
type Raccoon struct{}

func (v Raccoon) Animal() string {
	return (a.Raccoon{}).Id()
}

func (v Raccoon) Species() string {
	return (s.Raccoon{}).Id()
}
