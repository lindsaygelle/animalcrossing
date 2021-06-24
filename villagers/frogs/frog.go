package frogs

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Frog interface {
	villagers.Villager
}

type frog struct {
	name string
}

func (f frog) Animal() string {
	return animals.Frog.Name()
}

func (f frog) Name() string {
	return f.name
}

var (
	_ villagers.Villager = frog{}
)
