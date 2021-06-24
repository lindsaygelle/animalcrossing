package dodos

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Dodo interface {
	villagers.Villager
}

type dodo struct {
	name string
}

func (d dodo) Animal() string {
	return animals.Dodo.Name()
}

func (d dodo) Name() string {
	return d.name
}

var (
	_ villagers.Villager = dodo{}
)
