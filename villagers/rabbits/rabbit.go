package rabbits

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Rabbit interface {
	villagers.Villager
}

type rabbit struct {
	name string
}

func (r rabbit) Animal() string {
	return animals.Rabbit.Name()
}

func (r rabbit) Name() string {
	return r.name
}

var (
	_ villagers.Villager = rabbit{}
)
