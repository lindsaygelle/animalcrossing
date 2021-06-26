package pumpkins

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Pumpkin interface {
	villagers.Villager
}

type pumpkin struct {
	name string
}

func (p pumpkin) Animal() string {
	return animals.Pumpkin.Name()
}

func (p pumpkin) Name() string {
	return p.name
}

var (
	_ villagers.Villager = pumpkin{}
)
