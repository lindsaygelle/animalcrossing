package otters

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Otter interface {
	villagers.Villager
}

type otter struct {
	name string
}

func (o otter) Animal() string {
	return animals.Otter.Name()
}

func (o otter) Name() string {
	return o.name
}

var (
	_ villagers.Villager = otter{}
)
