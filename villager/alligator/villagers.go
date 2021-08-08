package alligator

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/alligator/alfonso"
	"github.com/lindsaygelle/animalcrossing/villager/alligator/alli"
	"github.com/lindsaygelle/animalcrossing/villager/alligator/boots"
	"github.com/lindsaygelle/animalcrossing/villager/alligator/del"
	"github.com/lindsaygelle/animalcrossing/villager/alligator/drago"
	"github.com/lindsaygelle/animalcrossing/villager/alligator/gayle"
	"github.com/lindsaygelle/animalcrossing/villager/alligator/liz"
	"github.com/lindsaygelle/animalcrossing/villager/alligator/pironkon"
	"github.com/lindsaygelle/animalcrossing/villager/alligator/sly"
)

var (
	Villagers = villager.NewVillagers(
		alfonso.Villager,
		alli.Villager,
		boots.Villager,
		del.Villager,
		drago.Villager,
		gayle.Villager,
		liz.Villager,
		pironkon.Villager,
		sly.Villager)
)
