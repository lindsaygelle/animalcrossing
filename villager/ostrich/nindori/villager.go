package nindori

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "nindori"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Nindori.
	Villager = villager.Villager{
		Animal:  ostrich.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
