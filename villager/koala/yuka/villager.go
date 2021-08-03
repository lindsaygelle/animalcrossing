package yuka

import (
	"github.com/lindsaygelle/animalcrossing/animal/koala"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "yuka"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Yuka.
	Villager = villager.Villager{
		Animal:  koala.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
