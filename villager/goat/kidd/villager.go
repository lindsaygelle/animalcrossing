package kidd

import (
	"github.com/lindsaygelle/animalcrossing/animal/goat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "kidd"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Kidd.
	Villager = villager.Villager{
		Animal:  goat.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
