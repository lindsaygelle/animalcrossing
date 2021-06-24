package bulls

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Bull interface {
	villagers.Villager
}

type bull struct {
	name string
}

func (b bull) Animal() string {
	return animals.Bull.Name()
}

func (b bull) Name() string {
	return b.name
}

var (
	_ villagers.Villager = bull{}
)
