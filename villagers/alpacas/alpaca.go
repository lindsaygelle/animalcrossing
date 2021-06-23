package alpacas

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type Alpaca interface {
	villagers.Villager
}

type alpaca struct {
	name string
}

func (a alpaca) Animal() string {
	return animals.Alpaca.Name()
}

func (a alpaca) Name() string {
	return a.name
}

var (
	_ villagers.Villager = alpaca{}
)
