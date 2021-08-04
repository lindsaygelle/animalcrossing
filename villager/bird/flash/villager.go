package flash

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "flash"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Flash.
	Villager = villager.Villager{
		Animal:      bird.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Flash,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
