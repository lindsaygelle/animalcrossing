
package frillneckedlizard

import (
	a "github.com/lindsaygelle/animalcrossing/animals/frillneckedlizard"
	s "github.com/lindsaygelle/animalcrossing/species/frillneckedlizard"
)

// Frillneckedlizard is an Animal Crossing villager type.
type Frillneckedlizard struct{}

func (v Frillneckedlizard) Animal() string {
	return (a.Frillneckedlizard{}).Id()
}

func (v Frillneckedlizard) Species() string {
	return (s.Frillneckedlizard{}).Id()
}
