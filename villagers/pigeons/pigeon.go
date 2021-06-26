package pigeons

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Pigeon interface {
	villagers.Villager
}

type pigeon struct {
	name string
}

func (p pigeon) Animal() string {
	return animals.Pigeon.Name()
}

func (p pigeon) Name() string {
	return p.name
}

var (
	_ villagers.Villager = pigeon{}
)
