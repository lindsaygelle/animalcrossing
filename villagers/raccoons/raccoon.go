package raccoons

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Raccoon interface {
	villagers.Villager
}

type raccoon struct {
	name string
}

func (r raccoon) Animal() string {
	return animals.Raccoon.Name()
}

func (r raccoon) Name() string {
	return r.name
}

var (
	_ villagers.Villager = raccoon{}
)
