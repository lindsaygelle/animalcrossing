package monkeys

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Monkey interface {
	villagers.Villager
}

type monkey struct {
	name string
}

func (m monkey) Animal() string {
	return animals.Monkey.Name()
}

func (m monkey) Name() string {
	return m.name
}

var (
	_ villagers.Villager = monkey{}
)
