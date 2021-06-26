package penguins

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Penguin interface {
	villagers.Villager
}

type penguin struct {
	name string
}

func (p penguin) Animal() string {
	return animals.Penguin.Name()
}

func (p penguin) Name() string {
	return p.name
}

var (
	_ villagers.Villager = penguin{}
)
