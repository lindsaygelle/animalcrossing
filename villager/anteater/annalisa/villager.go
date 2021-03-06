package annalisa

import (
	"github.com/lindsaygelle/animalcrossing/animal/anteater"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ant08"
)

const (
	gender string = "female"
)

const (
	id string = "annalisa"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Annalisa.
	Villager = villager.Villager{
		Animal:      anteater.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Annalisa,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
