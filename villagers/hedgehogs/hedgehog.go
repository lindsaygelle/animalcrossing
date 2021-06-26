package hedgehogs

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Hedgehog interface {
	villagers.Villager
}

type hedgehog struct {
	name string
}

func (h hedgehog) Animal() string {
	return animals.Hedgehog.Name()
}

func (h hedgehog) Name() string {
	return h.name
}

var (
	_ villagers.Villager = hedgehog{}
)
