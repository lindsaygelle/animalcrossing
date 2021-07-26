
package rabbit

import (
	a "github.com/lindsaygelle/animalcrossing/animals/rabbit"
	s "github.com/lindsaygelle/animalcrossing/species/rabbit"
)

// Rabbit is an Animal Crossing villager type.
type Rabbit struct{}

func (v Rabbit) Animal() string {
	return (a.Rabbit{}).Id()
}

func (v Rabbit) Species() string {
	return (s.Rabbit{}).Id()
}
