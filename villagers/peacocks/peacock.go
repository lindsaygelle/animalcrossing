package peacocks

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Peacock interface {
	villagers.Villager
}

type peacock struct {
	name string
}

func (p peacock) Animal() string {
	return animals.Peacock.Name()
}

func (p peacock) Name() string {
	return p.name
}

var (
	_ villagers.Villager = peacock{}
)
