package rocket

import (
	"github.com/lindsaygelle/animalcrossing/animal/gorilla"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "gor09"
)

const (
	gender string = "female"
)

const (
	id string = "rocket"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rocket.
	Villager = villager.Villager{
		Animal:      gorilla.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rocket,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
