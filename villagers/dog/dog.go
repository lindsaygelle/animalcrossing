
package dog

import (
	a "github.com/lindsaygelle/animalcrossing/animals/dog"
	s "github.com/lindsaygelle/animalcrossing/species/dog"
)

// Dog is an Animal Crossing villager type.
type Dog struct{}

func (v Dog) Animal() string {
	return (a.Dog{}).Id()
}

func (v Dog) Species() string {
	return (s.Dog{}).Id()
}
