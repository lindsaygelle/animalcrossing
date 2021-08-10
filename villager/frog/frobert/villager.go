package frobert

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "flg02"
)

const (
	gender string = "male"
)

const (
	id string = "frobert"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Frobert.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Frobert,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
