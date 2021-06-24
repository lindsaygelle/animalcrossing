package foxes

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Fox interface {
	villagers.Villager
}

type fox struct {
	name string
}

func (f fox) Animal() string {
	return animals.Fox.Name()
}

func (f fox) Name() string {
	return f.name
}

var (
	_ villagers.Villager = fox{}
)
