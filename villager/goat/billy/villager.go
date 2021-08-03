package billy

import (
	"github.com/lindsaygelle/animalcrossing/animal/goat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "billy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Billy.
	Villager = villager.Villager{
		Animal:  goat.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
