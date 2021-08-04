package spike

import (
	"github.com/lindsaygelle/animalcrossing/animal/rhinoceros"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "spike"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Spike.
	Villager = villager.Villager{
		Animal:      rhinoceros.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Spike,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
