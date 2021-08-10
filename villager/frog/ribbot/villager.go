package ribbot

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "flg01"
)

const (
	gender string = "male"
)

const (
	id string = "ribbot"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ribbot.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Ribbot,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
