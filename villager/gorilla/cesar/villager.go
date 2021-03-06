package cesar

import (
	"github.com/lindsaygelle/animalcrossing/animal/gorilla"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "gor00"
)

const (
	gender string = "male"
)

const (
	id string = "cesar"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cesar.
	Villager = villager.Villager{
		Animal:      gorilla.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Cesar,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
