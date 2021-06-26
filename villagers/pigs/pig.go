package pigs

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Pig interface {
	villagers.Villager
}

type pig struct {
	name string
}

func (p pig) Animal() string {
	return animals.Pig.Name()
}

func (p pig) Name() string {
	return p.name
}

var (
	_ villagers.Villager = pig{}
)
