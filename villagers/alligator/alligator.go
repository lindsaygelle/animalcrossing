package alligator

import (
	v "github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator/alli"
)

// Alligator is a Animal Crossing villager type.
type Alligator interface {
	v.Villager
}

var (
	Alligators = map[string]Alligator{
		(alli.Alli{}.Id()): (alli.Alli{})}
)
