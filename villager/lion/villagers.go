package lion

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/lion/aziz"
	"github.com/lindsaygelle/animalcrossing/villager/lion/bud"
	"github.com/lindsaygelle/animalcrossing/villager/lion/elvis"
	"github.com/lindsaygelle/animalcrossing/villager/lion/jubei"
	"github.com/lindsaygelle/animalcrossing/villager/lion/leopold"
	"github.com/lindsaygelle/animalcrossing/villager/lion/lionel"
	"github.com/lindsaygelle/animalcrossing/villager/lion/mott"
	"github.com/lindsaygelle/animalcrossing/villager/lion/rex"
	"github.com/lindsaygelle/animalcrossing/villager/lion/rory"
)

var (
	Villagers = villager.NewVillagers(
		aziz.Villager,
		bud.Villager,
		elvis.Villager,
		jubei.Villager,
		leopold.Villager,
		lionel.Villager,
		mott.Villager,
		rex.Villager,
		rory.Villager)
)
