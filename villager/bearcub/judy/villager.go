package judy

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "judy"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Judy.
	Villager = villager.Villager{
		Animal:      bearcub.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Judy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
