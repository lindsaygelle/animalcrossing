package lyman

import (
	"github.com/lindsaygelle/animalcrossing/animal/koala"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "kal09"
)

const (
	gender string = "male"
)

const (
	id string = "lyman"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Lyman.
	Villager = villager.Villager{
		Animal:      koala.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Lyman,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
