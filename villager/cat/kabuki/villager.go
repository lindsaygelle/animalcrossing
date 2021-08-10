package kabuki

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat09"
)

const (
	gender string = "male"
)

const (
	id string = "kabuki"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Kabuki.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Kabuki,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
