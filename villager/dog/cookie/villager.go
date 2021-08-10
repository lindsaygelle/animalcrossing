package cookie

import (
	"github.com/lindsaygelle/animalcrossing/animal/dog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "dog08"
)

const (
	gender string = "female"
)

const (
	id string = "cookie"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cookie.
	Villager = villager.Villager{
		Animal:      dog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Cookie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
