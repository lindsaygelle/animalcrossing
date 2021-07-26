
package human

import (
	a "github.com/lindsaygelle/animalcrossing/animals/human"
	s "github.com/lindsaygelle/animalcrossing/species/human"
)

// Human is an Animal Crossing villager type.
type Human struct{}

func (v Human) Animal() string {
	return (a.Human{}).Id()
}

func (v Human) Species() string {
	return (s.Human{}).Id()
}
