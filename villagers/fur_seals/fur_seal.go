package fur_seals

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type FurSeal interface {
	villagers.Villager
}

type furSeal struct {
	name string
}

func (f furSeal) Animal() string {
	return animals.FurSeal.Name()
}

func (f furSeal) Name() string {
	return f.name
}

var (
	_ villagers.Villager = furSeal{}
)
