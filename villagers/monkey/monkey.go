
package monkey

import (
	a "github.com/lindsaygelle/animalcrossing/animals/monkey"
	s "github.com/lindsaygelle/animalcrossing/species/monkey"
)

// Monkey is an Animal Crossing villager type.
type Monkey struct{}

func (v Monkey) Animal() string {
	return (a.Monkey{}).Id()
}

func (v Monkey) Species() string {
	return (s.Monkey{}).Id()
}
