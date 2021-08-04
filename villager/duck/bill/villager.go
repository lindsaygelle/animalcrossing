package bill

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
	id string = "bill"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Bill.
	Villager = villager.Villager{
		Animal:      duck.Animal,
		Birthday:    birthday,
		Code:        code,
		Gender:      gender,
		Id:          id,
		Key:         villager.Bill,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)