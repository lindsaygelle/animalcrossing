package cranston

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ost06"
)

const (
	gender string = "male"
)

const (
	id string = "cranston"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cranston.
	Villager = villager.Villager{
		Animal:      ostrich.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Cranston,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
