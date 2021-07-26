
package bull

import (
	a "github.com/lindsaygelle/animalcrossing/animals/bull"
	s "github.com/lindsaygelle/animalcrossing/species/bull"
)

// Bull is an Animal Crossing villager type.
type Bull struct{}

func (v Bull) Animal() string {
	return (a.Bull{}).Id()
}

func (v Bull) Species() string {
	return (s.Bull{}).Id()
}
