package deers

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Deer interface {
	villagers.Villager
}

type deer struct {
	name string
}

func (d deer) Animal() string {
	return animals.Deer.Name()
}

func (d deer) Name() string {
	return d.name
}

var (
	_ villagers.Villager = deer{}
)
