package anteaters

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Anteater interface {
	villagers.Villager
}

type anteater struct {
	name string
}

func (a anteater) Animal() string {
	return animals.Anteater.Name()
}

func (a anteater) Name() string {
	return a.name
}

var (
	_ villagers.Villager = anteater{}
)
