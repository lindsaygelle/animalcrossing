
package sloth

import (
	a "github.com/lindsaygelle/animalcrossing/animals/sloth"
	s "github.com/lindsaygelle/animalcrossing/species/sloth"
)

// Sloth is an Animal Crossing villager type.
type Sloth struct{}

func (v Sloth) Animal() string {
	return (a.Sloth{}).Id()
}

func (v Sloth) Species() string {
	return (s.Sloth{}).Id()
}
