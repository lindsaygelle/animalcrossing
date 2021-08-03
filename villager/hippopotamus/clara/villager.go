package clara

import (
	"github.com/lindsaygelle/animalcrossing/animal/hippopotamus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "clara"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Clara.
	Villager = villager.Villager{
		Animal:  hippopotamus.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
