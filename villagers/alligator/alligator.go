package alligator

import (
	v "github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/alligator/alli"
)

type Alligator interface {
	v.Villager
}

var (
	Alligators = map[string]Alligator{
		"alli": alli.Alli{}}
)
