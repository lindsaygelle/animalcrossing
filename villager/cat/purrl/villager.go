package purrl

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat07"
)

const (
	gender string = "female"
)

const (
	id string = "purrl"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Purrl.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Purrl,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
