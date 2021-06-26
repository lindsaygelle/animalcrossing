package turkeys

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Turkey interface {
	villagers.Villager
}

type turkey struct {
	name string
}

func (t turkey) Animal() string {
	return animals.Turkey.Name()
}

func (t turkey) Name() string {
	return t.name
}

var (
	_ villagers.Villager = turkey{}
)
