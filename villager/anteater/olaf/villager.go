package olaf

import (
	"github.com/lindsaygelle/animalcrossing/animal/anteater"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ant09"
)

const (
	gender string = "male"
)

const (
	id string = "olaf"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Olaf.
	Villager = villager.Villager{
		Animal:      anteater.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Olaf,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
