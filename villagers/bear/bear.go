
package bear

import (
	a "github.com/lindsaygelle/animalcrossing/animals/bear"
	s "github.com/lindsaygelle/animalcrossing/species/bear"
)

// Bear is an Animal Crossing villager type.
type Bear struct{}

func (v Bear) Animal() string {
	return (a.Bear{}).Id()
}

func (v Bear) Species() string {
	return (s.Bear{}).Id()
}
