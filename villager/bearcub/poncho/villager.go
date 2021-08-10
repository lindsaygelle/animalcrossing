package poncho

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cbr02"
)

const (
	gender string = "male"
)

const (
	id string = "poncho"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Poncho.
	Villager = villager.Villager{
		Animal:      bearcub.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Poncho,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
