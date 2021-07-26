
package cat

import (
	a "github.com/lindsaygelle/animalcrossing/animals/cat"
	s "github.com/lindsaygelle/animalcrossing/species/cat"
)

// Cat is an Animal Crossing villager type.
type Cat struct{}

func (v Cat) Animal() string {
	return (a.Cat{}).Id()
}

func (v Cat) Species() string {
	return (s.Cat{}).Id()
}
