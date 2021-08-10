package pashmina

import (
	"github.com/lindsaygelle/animalcrossing/animal/goat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "goa08"
)

const (
	gender string = "female"
)

const (
	id string = "pashmina"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Pashmina.
	Villager = villager.Villager{
		Animal:      goat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Pashmina,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
