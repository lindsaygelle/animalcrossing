
package giraffe

import (
	a "github.com/lindsaygelle/animalcrossing/animals/giraffe"
	s "github.com/lindsaygelle/animalcrossing/species/giraffe"
)

// Giraffe is an Animal Crossing villager type.
type Giraffe struct{}

func (v Giraffe) Animal() string {
	return (a.Giraffe{}).Id()
}

func (v Giraffe) Species() string {
	return (s.Giraffe{}).Id()
}
