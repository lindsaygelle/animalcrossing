package tank

import (
	"github.com/lindsaygelle/animalcrossing/animal/rhinoceros"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "tank"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tank.
	Villager = villager.Villager{
		Animal:      rhinoceros.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Tank,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
