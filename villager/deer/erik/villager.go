package erik

import (
	"github.com/lindsaygelle/animalcrossing/animal/deer"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "erik"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Erik.
	Villager = villager.Villager{
		Animal:      deer.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Erik,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
