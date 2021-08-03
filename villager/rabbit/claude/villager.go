package claude

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "claude"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Claude.
	Villager = villager.Villager{
		Animal:  rabbit.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
