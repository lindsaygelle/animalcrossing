package wolves

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Wolf interface {
	villagers.Villager
}

type wolf struct {
	name string
}

func (w wolf) Animal() string {
	return animals.Wolf.Name()
}

func (w wolf) Name() string {
	return w.name
}

var (
	_ villagers.Villager = wolf{}
)
