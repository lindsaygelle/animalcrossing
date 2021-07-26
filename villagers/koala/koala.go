
package koala

import (
	a "github.com/lindsaygelle/animalcrossing/animals/koala"
	s "github.com/lindsaygelle/animalcrossing/species/koala"
)

// Koala is an Animal Crossing villager type.
type Koala struct{}

func (v Koala) Animal() string {
	return (a.Koala{}).Id()
}

func (v Koala) Species() string {
	return (s.Koala{}).Id()
}
