package octopuses

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Octopus interface {
	villagers.Villager
}

type octopus struct {
	name string
}

func (o octopus) Animal() string {
	return animals.Octopus.Name()
}

func (o octopus) Name() string {
	return o.name
}

var (
	_ villagers.Villager = octopus{}
)
