package ghosts

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Ghost interface {
	villagers.Villager
}

type ghost struct {
	name string
}

func (g ghost) Animal() string {
	return animals.Ghost.Name()
}
func (g ghost) Name() string {
	return g.name
}

var (
	_ villagers.Villager = ghost{}
)
