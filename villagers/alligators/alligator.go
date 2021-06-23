package alligators

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Alligator interface {
	villagers.Villager
}

type alligator struct {
	name string
}

func (a alligator) Animal() string {
	return animals.Alligator.Name()
}

func (a alligator) Name() string {
	return a.name
}

var (
	_ villagers.Villager = alligator{}
)
