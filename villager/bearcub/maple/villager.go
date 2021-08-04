package maple

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "maple"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Maple.
	Villager = villager.Villager{
		Animal:      bearcub.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Maple,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
