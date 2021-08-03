package groucho

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "groucho"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Groucho.
	Villager = villager.Villager{
		Animal:  bear.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
