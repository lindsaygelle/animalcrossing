package cats

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Cat interface {
	villagers.Villager
}

type cat struct {
	name string
}

func (c cat) Animal() string {
	return animals.Cat.Name()
}

func (c cat) Name() string {
	return c.name
}

var (
	_ villagers.Villager = cat{}
)
