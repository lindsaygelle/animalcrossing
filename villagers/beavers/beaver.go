package beavers

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Beaver interface {
	villagers.Villager
}

type beaver struct {
	name string
}

func (b beaver) Animal() string {
	return animals.Beaver.Name()
}

func (b beaver) Name() string {
	return b.name
}

var (
	_ villagers.Villager = beaver{}
)
