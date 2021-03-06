package broffina

import (
	"github.com/lindsaygelle/animalcrossing/animal/chicken"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "chn12"
)

const (
	gender string = "female"
)

const (
	id string = "broffina"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Broffina.
	Villager = villager.Villager{
		Animal:      chicken.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Broffina,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
