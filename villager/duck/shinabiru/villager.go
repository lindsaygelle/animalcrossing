package shinabiru

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = ""
)

const (
	gender string = "male"
)

const (
	id string = "shinabiru"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Shinabiru.
	Villager = villager.Villager{
		Animal:      duck.Animal,
		Birthday:    birthday,
		Code:        code,
		Gender:      gender,
		Id:          id,
		Key:         villager.Shinabiru,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
