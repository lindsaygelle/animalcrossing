package walruses

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Walrus interface {
	villagers.Villager
}

type walrus struct {
	name string
}

func (w walrus) Animal() string {
	return animals.Walrus.Name()
}

func (w walrus) Name() string {
	return w.name
}

var (
	_ villagers.Villager = walrus{}
)
