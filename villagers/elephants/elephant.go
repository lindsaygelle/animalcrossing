package elephants

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Elephant interface {
	villagers.Villager
}

type elephant struct {
	name string
}

func (e elephant) Animal() string {
	return animals.Elephant.Name()
}

func (e elephant) Name() string {
	return e.name
}

var (
	_ villagers.Villager = elephant{}
)
