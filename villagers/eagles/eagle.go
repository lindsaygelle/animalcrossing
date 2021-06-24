package eagles

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Eagle interface {
	villagers.Villager
}

type eagle struct {
	name string
}

func (e eagle) Animal() string {
	return animals.Eagle.Name()
}

func (e eagle) Name() string {
	return e.name
}

var (
	_ villagers.Villager = eagle{}
)
