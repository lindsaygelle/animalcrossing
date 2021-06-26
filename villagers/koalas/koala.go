package koalas

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Koala interface {
	villagers.Villager
}

type koala struct {
	name string
}

func (k koala) Animal() string {
	return animals.Koala.Name()
}

func (k koala) Name() string {
	return k.name
}

var (
	_ villagers.Villager = koala{}
)
