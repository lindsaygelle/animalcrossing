package alligator

import (
	v "github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator/alli"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator/boots"
)

// Alligator is a Animal Crossing villager type.
type Alligator interface {
	v.Villager
}

var (
	// Alligators is a map of Animal Crossing Alligator villagers.
	Alligators = map[string]Alligator{
		(alli.Alli{}).Id():   alli.Alli{},
		(boots.Boots{}).Id(): boots.Boots{}}
)
