package lions

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Lion interface {
	villagers.Villager
}

type lion struct {
	name string
}

func (l lion) Animal() string {
	return animals.Lion.Name()
}

func (l lion) Name() string {
	return l.name
}

var (
	_ villagers.Villager = lion{}
)
