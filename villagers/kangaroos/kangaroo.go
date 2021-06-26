package kangaroos

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Kangaroo interface {
	villagers.Villager
}

type kangaroo struct {
	name string
}

func (k kangaroo) Animal() string {
	return animals.Kangaroo.Name()
}

func (k kangaroo) Name() string {
	return k.name
}

var (
	_ villagers.Villager = kangaroo{}
)
