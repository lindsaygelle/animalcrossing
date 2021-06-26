package squirrels

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Squirrel interface {
	villagers.Villager
}

type squirrel struct {
	name string
}

func (s squirrel) Animal() string {
	return animals.Squirrel.Name()
}

func (s squirrel) Name() string {
	return s.name
}

var (
	_ villagers.Villager = squirrel{}
)
