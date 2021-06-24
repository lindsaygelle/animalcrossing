package giraffes

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Giraffe interface {
	villagers.Villager
}

type giraffe struct {
	name string
}

func (g giraffe) Animal() string {
	return animals.Giraffe.Name()
}

func (g giraffe) Name() string {
	return g.name
}

var (
	_ villagers.Villager = giraffe{}
)
