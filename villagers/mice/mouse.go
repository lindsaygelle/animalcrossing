package mice

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Mouse interface {
	villagers.Villager
}

type mouse struct {
	name string
}

func (m mouse) Animal() string {
	return animals.Mouse.Name()
}

func (m mouse) Name() string {
	return m.name
}

var (
	_ villagers.Villager = mouse{}
)
