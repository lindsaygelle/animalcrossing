package gorillas

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Gorilla interface {
	villagers.Villager
}

type gorilla struct {
	name string
}

func (g gorilla) Animal() string {
	return animals.Gorilla.Name()
}

func (g gorilla) Name() string {
	return g.name
}

var (
	_ villagers.Villager = gorilla{}
)
