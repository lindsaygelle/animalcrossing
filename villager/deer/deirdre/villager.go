package deirdre

import (
	"github.com/lindsaygelle/animalcrossing/animal/deer"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "der04"
)

const (
	gender string = "female"
)

const (
	id string = "deirdre"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Deirdre.
	Villager = villager.Villager{
		Animal:      deer.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Deirdre,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
