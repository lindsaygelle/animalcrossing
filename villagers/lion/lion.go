
package lion

import (
	a "github.com/lindsaygelle/animalcrossing/animals/lion"
	s "github.com/lindsaygelle/animalcrossing/species/lion"
)

// Lion is an Animal Crossing villager type.
type Lion struct{}

func (v Lion) Animal() string {
	return (a.Lion{}).Id()
}

func (v Lion) Species() string {
	return (s.Lion{}).Id()
}
