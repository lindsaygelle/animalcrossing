package becky

import (
	"github.com/lindsaygelle/animalcrossing/animal/chicken"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "chn09"
)

const (
	gender string = "female"
)

const (
	id string = "becky"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Becky.
	Villager = villager.Villager{
		Animal:      chicken.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Becky,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
