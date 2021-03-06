package madamrosa

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = ""
)

const (
	gender string = "female"
)

const (
	id string = "madamrosa"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for MadamRosa.
	Villager = villager.Villager{
		Animal:      bird.Animal,
		Birthday:    birthday,
		Code:        code,
		Gender:      gender,
		Id:          id,
		Key:         villager.MadamRosa,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
