package rod

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "mus05"
)

const (
	gender string = "male"
)

const (
	id string = "rod"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rod.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rod,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
