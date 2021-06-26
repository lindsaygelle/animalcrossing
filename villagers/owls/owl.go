package owls

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Owl interface {
	villagers.Villager
}

type owl struct {
	name string
}

func (o owl) Animal() string {
	return animals.Owl.Name()
}

func (o owl) Name() string {
	return o.name
}

var (
	_ villagers.Villager = owl{}
)
