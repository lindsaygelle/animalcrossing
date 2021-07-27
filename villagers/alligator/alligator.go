package alligator

import (
	a "github.com/lindsaygelle/animalcrossing/animals/alligator"
	s "github.com/lindsaygelle/animalcrossing/species/alligator"
)

// Alligator is a Animal Crossing villager type.
type Alligator struct{}

func (v Alligator) Animal() string {
	return (a.Alligator{}).Id()
}

func (v Alligator) Code() string {
	return "crd"
}

func (v Alligator) Species() string {
	return (s.Alligator{}).Id()
}
