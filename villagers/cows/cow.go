package cows

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Cow interface {
	villagers.Villager
}

type cow struct {
	name string
}

func (c cow) Animal() string {
	return animals.Cow.Name()
}

func (c cow) Name() string {
	return c.name
}

var (
	_ villagers.Villager = cow{}
)
