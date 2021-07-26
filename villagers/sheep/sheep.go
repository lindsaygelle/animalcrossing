
package sheep

import (
	a "github.com/lindsaygelle/animalcrossing/animals/sheep"
	s "github.com/lindsaygelle/animalcrossing/species/sheep"
)

// Sheep is an Animal Crossing villager type.
type Sheep struct{}

func (v Sheep) Animal() string {
	return (a.Sheep{}).Id()
}

func (v Sheep) Species() string {
	return (s.Sheep{}).Id()
}
