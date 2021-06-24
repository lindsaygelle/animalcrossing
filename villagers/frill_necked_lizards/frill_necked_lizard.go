package frill_necked_lizards

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/villagers"
)

type FrillNeckedLizard interface {
	villagers.Villager
}

type frillNeckedLizard struct {
	name string
}

func (f frillNeckedLizard) Animal() string {
	return animals.FrillNeckedLizard.Name()
}

func (f frillNeckedLizard) Name() string {
	return f.name
}

var (
	_ villagers.Villager = frillNeckedLizard{}
)
