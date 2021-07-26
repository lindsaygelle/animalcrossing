
package bearcub

import (
	a "github.com/lindsaygelle/animalcrossing/animals/bearcub"
	s "github.com/lindsaygelle/animalcrossing/species/bearcub"
)

// Bearcub is an Animal Crossing villager type.
type Bearcub struct{}

func (v Bearcub) Animal() string {
	return (a.Bearcub{}).Id()
}

func (v Bearcub) Species() string {
	return (s.Bearcub{}).Id()
}
