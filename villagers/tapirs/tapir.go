package tapirs

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Tapir interface {
	villagers.Villager
}

type tapir struct {
	name string
}

func (t tapir) Animal() string {
	return animals.Tapir.Name()
}

func (t tapir) Name() string {
	return t.name
}

var (
	_ villagers.Villager = tapir{}
)
