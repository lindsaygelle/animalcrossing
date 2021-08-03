package violet

import (
	"github.com/lindsaygelle/animalcrossing/animal/gorilla"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "violet"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Violet.
	Villager = villager.Villager{
		Animal:  gorilla.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
