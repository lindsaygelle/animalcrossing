package skunks

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Skunk interface {
	villagers.Villager
}

type skunk struct {
	name string
}

func (s skunk) Animal() string {
	return animals.Skunk.Name()
}

func (s skunk) Name() string {
	return s.name
}

var (
	_ villagers.Villager = skunk{}
)
