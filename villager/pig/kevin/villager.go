package kevin

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pig15"
)

const (
	gender string = "male"
)

const (
	id string = "kevin"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Kevin.
	Villager = villager.Villager{
		Animal:      pig.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Kevin,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
