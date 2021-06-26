package kappas

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Kappa interface {
	villagers.Villager
}

type kappa struct {
	name string
}

func (k kappa) Animal() string {
	return animals.Kappa.Name()
}

func (k kappa) Name() string {
	return k.name
}

var (
	_ villagers.Villager = kappa{}
)
