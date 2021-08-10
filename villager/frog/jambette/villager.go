package jambette

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "flg13"
)

const (
	gender string = "female"
)

const (
	id string = "jambette"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Jambette.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Jambette,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
