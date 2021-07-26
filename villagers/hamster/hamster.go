
package hamster

import (
	a "github.com/lindsaygelle/animalcrossing/animals/hamster"
	s "github.com/lindsaygelle/animalcrossing/species/hamster"
)

// Hamster is an Animal Crossing villager type.
type Hamster struct{}

func (v Hamster) Animal() string {
	return (a.Hamster{}).Id()
}

func (v Hamster) Species() string {
	return (s.Hamster{}).Id()
}
