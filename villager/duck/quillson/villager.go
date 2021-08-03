package quillson

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "quillson"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Quillson.
	Villager = villager.Villager{
		Animal:  duck.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
