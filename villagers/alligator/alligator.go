package alligator

import (
	v "github.com/lindsaygelle/animalcrossing/villagers"

	"github.com/lindsaygelle/animalcrossing/villagers/alligator/alfonso"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator/alli"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator/boots"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator/del"
)

// Alligator is a Animal Crossing villager type.
type Alligator interface {
	v.Villager
}

var (
	// Alligators is a map of Animal Crossing Alligator villagers.
	Alligators = map[string]Alligator{
		(alfonso.Alfonso{}).Id(): alfonso.Alfonso{},
		(alli.Alli{}).Id():       alli.Alli{},
		(boots.Boots{}).Id():     boots.Boots{},
		(del.Del{}).Id():         del.Del{}}
)
