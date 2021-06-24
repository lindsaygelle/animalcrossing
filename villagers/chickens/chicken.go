package chickens

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Chicken interface {
	villagers.Villager
}

type chicken struct {
	name string
}

func (c chicken) Animal() string {
	return animals.Chicken.Name()
}

func (c chicken) Name() string {
	return c.name
}

var (
	_ villagers.Villager = chicken{}
)
