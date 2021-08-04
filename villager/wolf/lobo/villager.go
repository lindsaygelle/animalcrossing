package lobo

import (
	"github.com/lindsaygelle/animalcrossing/animal/wolf"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "lobo"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Lobo.
	Villager = villager.Villager{
		Animal:      wolf.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Lobo,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
