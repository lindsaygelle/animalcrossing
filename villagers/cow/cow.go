
package cow

import (
	a "github.com/lindsaygelle/animalcrossing/animals/cow"
	s "github.com/lindsaygelle/animalcrossing/species/cow"
)

// Cow is an Animal Crossing villager type.
type Cow struct{}

func (v Cow) Animal() string {
	return (a.Cow{}).Id()
}

func (v Cow) Species() string {
	return (s.Cow{}).Id()
}
