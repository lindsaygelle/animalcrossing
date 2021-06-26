package sloths

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Sloth interface {
	villagers.Villager
}

type sloth struct {
	name string
}

func (s sloth) Animal() string {
	return animals.Sloth.Name()
}

func (s sloth) Name() string {
	return s.name
}

var (
	_ villagers.Villager = sloth{}
)
