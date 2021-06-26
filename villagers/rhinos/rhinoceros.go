package rhinoceroses

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Rhinoceros interface {
	villagers.Villager
}

type rhinoceros struct {
	name string
}

func (r rhinoceros) Animal() string {
	return animals.Rhinoceros.Name()
}

func (r rhinoceros) Name() string {
	return r.name
}

var (
	_ villagers.Villager = rhinoceros{}
)
