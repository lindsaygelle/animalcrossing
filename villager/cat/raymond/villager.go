package raymond

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat23"
)

const (
	gender string = "male"
)

const (
	id string = "raymond"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Raymond.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Raymond,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
