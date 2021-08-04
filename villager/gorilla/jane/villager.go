package jane

import (
	"github.com/lindsaygelle/animalcrossing/animal/gorilla"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "jane"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Jane.
	Villager = villager.Villager{
		Animal:      gorilla.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Jane,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
