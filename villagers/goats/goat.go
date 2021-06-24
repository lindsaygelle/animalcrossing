package goats

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Goat interface {
	villagers.Villager
}

type goat struct {
	name string
}

func (g goat) Animal() string {
	return animals.Goat.Name()
}

func (g goat) Name() string {
	return g.name
}

var (
	_ villagers.Villager = goat{}
)
