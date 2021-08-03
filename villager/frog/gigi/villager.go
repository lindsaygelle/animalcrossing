package gigi

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "gigi"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gigi.
	Villager = villager.Villager{
		Animal:  frog.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
