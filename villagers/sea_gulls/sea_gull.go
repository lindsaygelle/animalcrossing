package sea_gulls

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type SeaGull interface {
	villagers.Villager
}

type seaGull struct {
	name string
}

func (s seaGull) Animal() string {
	return animals.SeaGull.Name()
}

func (s seaGull) Name() string {
	return s.name
}

var (
	_ villagers.Villager = seaGull{}
)
