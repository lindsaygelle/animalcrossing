
package deer

import (
	a "github.com/lindsaygelle/animalcrossing/animals/deer"
	s "github.com/lindsaygelle/animalcrossing/species/deer"
)

// Deer is an Animal Crossing villager type.
type Deer struct{}

func (v Deer) Animal() string {
	return (a.Deer{}).Id()
}

func (v Deer) Species() string {
	return (s.Deer{}).Id()
}
