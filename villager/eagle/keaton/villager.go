package keaton

import (
	"github.com/lindsaygelle/animalcrossing/animal/eagle"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "keaton"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Keaton.
	Villager = villager.Villager{
		Animal:  eagle.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
