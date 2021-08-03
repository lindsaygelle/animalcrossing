package olivia

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "olivia"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Olivia.
	Villager = villager.Villager{
		Animal:  cat.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
