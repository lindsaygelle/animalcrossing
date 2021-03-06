package maddie

import (
	"github.com/lindsaygelle/animalcrossing/animal/dog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "dog09"
)

const (
	gender string = "female"
)

const (
	id string = "maddie"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Maddie.
	Villager = villager.Villager{
		Animal:      dog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Maddie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
