
package walrus

import (
	a "github.com/lindsaygelle/animalcrossing/animals/walrus"
	s "github.com/lindsaygelle/animalcrossing/species/walrus"
)

// Walrus is an Animal Crossing villager type.
type Walrus struct{}

func (v Walrus) Animal() string {
	return (a.Walrus{}).Id()
}

func (v Walrus) Species() string {
	return (s.Walrus{}).Id()
}
