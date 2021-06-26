package horses

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Horse interface {
	villagers.Villager
}

type horse struct {
	name string
}

func (h horse) Animal() string {
	return animals.Horse.Name()
}

func (h horse) Name() string {
	return h.name
}

var (
	_ villagers.Villager = horse{}
)
