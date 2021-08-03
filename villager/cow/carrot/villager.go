package carrot

import (
	"github.com/lindsaygelle/animalcrossing/animal/cow"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "carrot"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Carrot.
	Villager = villager.Villager{
		Animal:  cow.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
