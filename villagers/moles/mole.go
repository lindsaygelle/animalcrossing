package moles

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Mole interface {
	villagers.Villager
}

type mole struct {
	name string
}

func (m mole) Animal() string {
	return animals.Mole.Name()
}

func (m mole) Name() string {
	return m.name
}

var (
	_ villagers.Villager = mole{}
)
