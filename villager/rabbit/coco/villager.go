package coco

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "coco"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Coco.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Coco,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
