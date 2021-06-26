package ostriches

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Ostrich interface {
	villagers.Villager
}

type ostrich struct {
	name string
}

func (o ostrich) Animal() string {
	return animals.Ostrich.Name()
}

func (o ostrich) Name() string {
	return o.name
}

var (
	_ villagers.Villager = ostrich{}
)
