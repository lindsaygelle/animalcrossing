package axolotls

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Axolotl interface {
	villagers.Villager
}

type axolotl struct {
	name string
}

func (a axolotl) Animal() string {
	return animals.Axolotl.Name()
}

func (a axolotl) Name() string {
	return a.name
}

var (
	_ villagers.Villager = axolotl{}
)
