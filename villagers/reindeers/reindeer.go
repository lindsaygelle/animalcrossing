package reindeers

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Reindeer interface {
	villagers.Villager
}

type reindeer struct {
	name string
}

func (r reindeer) Animal() string {
	return animals.Reindeer.Name()
}

func (r reindeer) Name() string {
	return r.name
}

var (
	_ villagers.Villager = reindeer{}
)
