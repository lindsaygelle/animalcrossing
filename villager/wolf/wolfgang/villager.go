package wolfgang

import (
	"github.com/lindsaygelle/animalcrossing/animal/wolf"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "wol02"
)

const (
	gender string = "male"
)

const (
	id string = "wolfgang"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Wolfgang.
	Villager = villager.Villager{
		Animal:      wolf.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Wolfgang,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
