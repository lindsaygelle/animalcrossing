package chameleons

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Chameleon interface {
	villagers.Villager
}

type chameleon struct {
	name string
}

func (c chameleon) Animal() string {
	return animals.Chameleon.Name()
}

func (c chameleon) Name() string {
	return c.name
}

var (
	_ villagers.Villager = chameleon{}
)
