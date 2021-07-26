
package elephant

import (
	a "github.com/lindsaygelle/animalcrossing/animals/elephant"
	s "github.com/lindsaygelle/animalcrossing/species/elephant"
)

// Elephant is an Animal Crossing villager type.
type Elephant struct{}

func (v Elephant) Animal() string {
	return (a.Elephant{}).Id()
}

func (v Elephant) Species() string {
	return (s.Elephant{}).Id()
}
