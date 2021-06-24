package dogs

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Dog interface {
	villagers.Villager
}

type dog struct {
	name string
}

func (d dog) Animal() string {
	return animals.Dog.Name()
}

func (d dog) Name() string {
	return d.name
}

var (
	_ villagers.Villager = dog{}
)
