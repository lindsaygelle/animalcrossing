package kyle

import (
	"github.com/lindsaygelle/animalcrossing/animal/wolf"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "wol10"
)

const (
	gender string = "male"
)

const (
	id string = "kyle"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Kyle.
	Villager = villager.Villager{
		Animal:      wolf.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Kyle,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
