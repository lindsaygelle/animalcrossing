package twirp

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "twirp"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Twirp.
	Villager = villager.Villager{
		Animal:  bird.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
