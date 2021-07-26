
package beaver

import (
	a "github.com/lindsaygelle/animalcrossing/animals/beaver"
	s "github.com/lindsaygelle/animalcrossing/species/beaver"
)

// Beaver is an Animal Crossing villager type.
type Beaver struct{}

func (v Beaver) Animal() string {
	return (a.Beaver{}).Id()
}

func (v Beaver) Species() string {
	return (s.Beaver{}).Id()
}
