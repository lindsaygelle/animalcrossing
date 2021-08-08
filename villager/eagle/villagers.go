package eagle

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/eagle/amelia"
	"github.com/lindsaygelle/animalcrossing/villager/eagle/apollo"
	"github.com/lindsaygelle/animalcrossing/villager/eagle/avery"
	"github.com/lindsaygelle/animalcrossing/villager/eagle/buzz"
	"github.com/lindsaygelle/animalcrossing/villager/eagle/celia"
	"github.com/lindsaygelle/animalcrossing/villager/eagle/frank"
	"github.com/lindsaygelle/animalcrossing/villager/eagle/keaton"
	"github.com/lindsaygelle/animalcrossing/villager/eagle/pierce"
	"github.com/lindsaygelle/animalcrossing/villager/eagle/quetzal"
	"github.com/lindsaygelle/animalcrossing/villager/eagle/sterling"
)

var (
	Villagers = villager.NewVillagers(
		amelia.Villager,
		apollo.Villager,
		avery.Villager,
		buzz.Villager,
		celia.Villager,
		frank.Villager,
		keaton.Villager,
		pierce.Villager,
		quetzal.Villager,
		sterling.Villager)
)
