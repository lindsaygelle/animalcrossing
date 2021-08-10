package gonzo

import (
	"github.com/lindsaygelle/animalcrossing/animal/koala"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "kal04"
)

const (
	gender string = "male"
)

const (
	id string = "gonzo"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gonzo.
	Villager = villager.Villager{
		Animal:      koala.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Gonzo,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
