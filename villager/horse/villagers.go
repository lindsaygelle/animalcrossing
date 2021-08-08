package horse

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/horse/annalise"
	"github.com/lindsaygelle/animalcrossing/villager/horse/buck"
	"github.com/lindsaygelle/animalcrossing/villager/horse/cleo"
	"github.com/lindsaygelle/animalcrossing/villager/horse/clyde"
	"github.com/lindsaygelle/animalcrossing/villager/horse/colton"
	"github.com/lindsaygelle/animalcrossing/villager/horse/ed"
	"github.com/lindsaygelle/animalcrossing/villager/horse/elmer"
	"github.com/lindsaygelle/animalcrossing/villager/horse/epona"
	"github.com/lindsaygelle/animalcrossing/villager/horse/filly"
	"github.com/lindsaygelle/animalcrossing/villager/horse/julian"
	"github.com/lindsaygelle/animalcrossing/villager/horse/papi"
	"github.com/lindsaygelle/animalcrossing/villager/horse/peaches"
	"github.com/lindsaygelle/animalcrossing/villager/horse/reneigh"
	"github.com/lindsaygelle/animalcrossing/villager/horse/roscoe"
	"github.com/lindsaygelle/animalcrossing/villager/horse/savannah"
	"github.com/lindsaygelle/animalcrossing/villager/horse/victoria"
	"github.com/lindsaygelle/animalcrossing/villager/horse/winnie"
)

var (
	Villagers = villager.NewVillagers(
		annalise.Villager,
		buck.Villager,
		cleo.Villager,
		clyde.Villager,
		colton.Villager,
		ed.Villager,
		elmer.Villager,
		epona.Villager,
		filly.Villager,
		julian.Villager,
		papi.Villager,
		peaches.Villager,
		reneigh.Villager,
		roscoe.Villager,
		savannah.Villager,
		victoria.Villager,
		winnie.Villager)
)
