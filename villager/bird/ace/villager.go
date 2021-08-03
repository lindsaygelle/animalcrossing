package ace

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "ace"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ace.
	Villager = villager.Villager{
		Animal:  bird.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
