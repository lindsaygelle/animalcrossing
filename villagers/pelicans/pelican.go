package pelicans

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Pelican interface {
	villagers.Villager
}

type pelican struct {
	name string
}

func (p pelican) Animal() string {
	return animals.Pelican.Name()
}

func (p pelican) Name() string {
	return p.name
}

var (
	_ villagers.Villager = pelican{}
)
