package hippopotamuses

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Hippopotamus interface {
	villagers.Villager
}

type hippopotamus struct {
	name string
}

func (h hippopotamus) Animal() string {
	return animals.Hippopotamus.Name()
}

func (h hippopotamus) Name() string {
	return h.name
}
