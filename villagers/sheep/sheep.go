package sheep

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Sheep interface {
	villagers.Villager
}

type sheep struct {
	name string
}

func (s sheep) Animal() string {
	return animals.Sheep.Name()
}

func (s sheep) Name() string {
	return s.name
}

var (
	_ villagers.Villager = sheep{}
)
