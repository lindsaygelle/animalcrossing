
package horse

import (
	a "github.com/lindsaygelle/animalcrossing/animals/horse"
	s "github.com/lindsaygelle/animalcrossing/species/horse"
)

// Horse is an Animal Crossing villager type.
type Horse struct{}

func (v Horse) Animal() string {
	return (a.Horse{}).Id()
}

func (v Horse) Species() string {
	return (s.Horse{}).Id()
}
