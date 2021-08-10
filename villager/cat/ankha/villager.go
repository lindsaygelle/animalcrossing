package ankha

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat19"
)

const (
	gender string = "female"
)

const (
	id string = "ankha"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ankha.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Ankha,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
