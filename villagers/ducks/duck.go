package ducks

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Duck interface {
	villagers.Villager
}

type duck struct {
	name string
}

func (d duck) Animal() string {
	return animals.Duck.Name()
}

func (d duck) Name() string {
	return d.name
}

var (
	_ villagers.Villager = duck{}
)
