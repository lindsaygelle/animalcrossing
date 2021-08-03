package pinky

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "pinky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Pinky.
	Villager = villager.Villager{
		Animal:  bear.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
