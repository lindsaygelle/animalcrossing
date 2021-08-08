package ostrich

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/ostrich/blanche"
	"github.com/lindsaygelle/animalcrossing/villager/ostrich/cranston"
	"github.com/lindsaygelle/animalcrossing/villager/ostrich/flora"
	"github.com/lindsaygelle/animalcrossing/villager/ostrich/gladys"
	"github.com/lindsaygelle/animalcrossing/villager/ostrich/julia"
	"github.com/lindsaygelle/animalcrossing/villager/ostrich/nindori"
	"github.com/lindsaygelle/animalcrossing/villager/ostrich/phil"
	"github.com/lindsaygelle/animalcrossing/villager/ostrich/phoebe"
	"github.com/lindsaygelle/animalcrossing/villager/ostrich/queenie"
	"github.com/lindsaygelle/animalcrossing/villager/ostrich/rio"
	"github.com/lindsaygelle/animalcrossing/villager/ostrich/sandy"
	"github.com/lindsaygelle/animalcrossing/villager/ostrich/sprocket"
)

var (
	Villagers = villager.NewVillagers(
		blanche.Villager,
		cranston.Villager,
		flora.Villager,
		gladys.Villager,
		julia.Villager,
		nindori.Villager,
		phil.Villager,
		phoebe.Villager,
		queenie.Villager,
		rio.Villager,
		sandy.Villager,
		sprocket.Villager)
)
