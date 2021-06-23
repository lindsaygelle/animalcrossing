package bears

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Bear interface {
	villagers.Villager
}

type bear struct {
	name string
}

func (b bear) Animal() string {
	return animals.Bear.Name()
}

func (b bear) Name() string {
	return b.name
}
