package alli

import (
	a "github.com/lindsaygelle/animalcrossing/animals/alligator"
	s "github.com/lindsaygelle/animalcrossing/species/alligator"
)

// Alli is an Animal Crossing villager.
type Alli struct{}

func (v Alli) Animal() string {
	return (a.Alligator{}).Id()
}

func (v Alli) Gender() string {
	return "female"
}

func (v Alli) Languages() {
	// TBC
}

func (v Alli) Personality() string {
	return "snooty"
}

func (v Alli) Special() bool {
	return false
}

func (v Alli) Species() string {
	return (s.Alligator{}).Id()
}
