
package skunk

import (
	a "github.com/lindsaygelle/animalcrossing/animals/skunk"
	s "github.com/lindsaygelle/animalcrossing/species/skunk"
)

// Skunk is an Animal Crossing villager type.
type Skunk struct{}

func (v Skunk) Animal() string {
	return (a.Skunk{}).Id()
}

func (v Skunk) Species() string {
	return (s.Skunk{}).Id()
}
