package agnes

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pig17"
)

const (
	gender string = "female"
)

const (
	id string = "agnes"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Agnes.
	Villager = villager.Villager{
		Animal:      pig.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Agnes,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
