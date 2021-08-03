package cyrus

import (
	"github.com/lindsaygelle/animalcrossing/animal/alpaca"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "cyrus"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cyrus.
	Villager = villager.Villager{
		Animal:  alpaca.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
