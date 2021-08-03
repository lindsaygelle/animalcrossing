package jubei

import (
	"github.com/lindsaygelle/animalcrossing/animal/lion"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "jubei"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Jubei.
	Villager = villager.Villager{
		Animal:  lion.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
