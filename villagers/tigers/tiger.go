package tigers

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Tiger interface {
	villagers.Villager
}

type tiger struct {
	name string
}

func (t tiger) Animal() string {
	return animals.Tiger.Name()
}

func (t tiger) Name() string {
	return t.name
}

var (
	_ villagers.Villager = tiger{}
)
