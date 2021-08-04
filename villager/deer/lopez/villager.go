package lopez

import (
	"github.com/lindsaygelle/animalcrossing/animal/deer"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "lopez"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Lopez.
	Villager = villager.Villager{
		Animal:      deer.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Lopez,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
