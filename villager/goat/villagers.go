package goat

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/goat/billy"
	"github.com/lindsaygelle/animalcrossing/villager/goat/chevre"
	"github.com/lindsaygelle/animalcrossing/villager/goat/gruff"
	"github.com/lindsaygelle/animalcrossing/villager/goat/iggy"
	"github.com/lindsaygelle/animalcrossing/villager/goat/kidd"
	"github.com/lindsaygelle/animalcrossing/villager/goat/nan"
	"github.com/lindsaygelle/animalcrossing/villager/goat/pashmina"
	"github.com/lindsaygelle/animalcrossing/villager/goat/sherb"
	"github.com/lindsaygelle/animalcrossing/villager/goat/sven"
	"github.com/lindsaygelle/animalcrossing/villager/goat/velma"
)

var (
	Villagers = villager.NewVillagers(
		billy.Villager,
		chevre.Villager,
		gruff.Villager,
		iggy.Villager,
		kidd.Villager,
		nan.Villager,
		pashmina.Villager,
		sherb.Villager,
		sven.Villager,
		velma.Villager)
)
