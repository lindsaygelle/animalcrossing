package camels

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Camel interface {
	villagers.Villager
}

type camel struct {
	name string
}

func (c camel) Animal() string {
	return animals.Camel.Name()
}

func (c camel) Name() string {
	return c.name
}

var (
	_ villagers.Villager = camel{}
)
