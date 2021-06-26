package hamsters

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Hamster interface {
	villagers.Villager
}

type hamster struct {
	name string
}

func (h hamster) Animal() string {
	return animals.Hamster.Name()
}

func (h hamster) Name() string {
	return h.name
}

var (
	_ villagers.Villager = hamster{}
)
