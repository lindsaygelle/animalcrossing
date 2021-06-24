package boars

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Boar interface {
	villagers.Villager
}

type boar struct {
	name string
}

func (b boar) Animal() string {
	return animals.Boar.Name()
}

func (b boar) Name() string {
	return b.name
}

var (
	_ villagers.Villager = boar{}
)
