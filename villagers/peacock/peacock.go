
package peacock

import (
	a "github.com/lindsaygelle/animalcrossing/animals/peacock"
	s "github.com/lindsaygelle/animalcrossing/species/peacock"
)

// Peacock is an Animal Crossing villager type.
type Peacock struct{}

func (v Peacock) Animal() string {
	return (a.Peacock{}).Id()
}

func (v Peacock) Species() string {
	return (s.Peacock{}).Id()
}
