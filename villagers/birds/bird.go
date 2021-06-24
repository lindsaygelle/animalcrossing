package birds

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Bird interface {
	villagers.Villager
}

type bird struct {
	name string
}

func (b bird) Animal() string {
	return animals.Bird.Name()
}

func (b bird) Name() string {
	return b.name
}

var (
	_ villagers.Villager = bird{}
)
