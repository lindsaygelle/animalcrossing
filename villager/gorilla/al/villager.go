package al

import (
	"github.com/lindsaygelle/animalcrossing/animal/gorilla"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "al"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Al.
	Villager = villager.Villager{
		Animal:  gorilla.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
