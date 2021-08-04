package clara

import (
	"github.com/lindsaygelle/animalcrossing/animal/hippopotamus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "clara"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Clara.
	Villager = villager.Villager{
		Animal:      hippopotamus.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Clara,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
