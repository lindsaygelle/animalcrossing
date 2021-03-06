package puddles

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "flg06"
)

const (
	gender string = "female"
)

const (
	id string = "puddles"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Puddles.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Puddles,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
