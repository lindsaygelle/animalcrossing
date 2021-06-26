package panthers

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Panther interface {
	villagers.Villager
}

type panther struct {
	name string
}

func (p panther) Animal() string {
	return animals.Panther.Name()
}

func (p panther) Name() string {
	return p.name
}

var (
	_ villagers.Villager = panther{}
)
