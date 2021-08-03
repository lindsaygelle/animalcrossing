package petunia

import (
	"github.com/lindsaygelle/animalcrossing/animal/cow"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "petunia"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Petunia.
	Villager = villager.Villager{
		Animal:  cow.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
