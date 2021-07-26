
package fox

import (
	a "github.com/lindsaygelle/animalcrossing/animals/fox"
	s "github.com/lindsaygelle/animalcrossing/species/fox"
)

// Fox is an Animal Crossing villager type.
type Fox struct{}

func (v Fox) Animal() string {
	return (a.Fox{}).Id()
}

func (v Fox) Species() string {
	return (s.Fox{}).Id()
}
