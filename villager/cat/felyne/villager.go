package felyne

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat22"
)

const (
	gender string = "male"
)

const (
	id string = "felyne"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Felyne.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Felyne,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
