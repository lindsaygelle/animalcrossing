package tortoises

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Tortoise interface {
	villagers.Villager
}

type tortoise struct {
	name string
}

func (t tortoise) Animal() string {
	return animals.Tortoise.Name()
}

func (t tortoise) Name() string {
	return t.name
}

var (
	_ villagers.Villager = tortoise{}
)
