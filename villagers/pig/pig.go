
package pig

import (
	a "github.com/lindsaygelle/animalcrossing/animals/pig"
	s "github.com/lindsaygelle/animalcrossing/species/pig"
)

// Pig is an Animal Crossing villager type.
type Pig struct{}

func (v Pig) Animal() string {
	return (a.Pig{}).Id()
}

func (v Pig) Species() string {
	return (s.Pig{}).Id()
}
