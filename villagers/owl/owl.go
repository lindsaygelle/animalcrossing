
package owl

import (
	a "github.com/lindsaygelle/animalcrossing/animals/owl"
	s "github.com/lindsaygelle/animalcrossing/species/owl"
)

// Owl is an Animal Crossing villager type.
type Owl struct{}

func (v Owl) Animal() string {
	return (a.Owl{}).Id()
}

func (v Owl) Species() string {
	return (s.Owl{}).Id()
}
