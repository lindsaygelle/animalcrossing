package deer

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/deer/bam"
	"github.com/lindsaygelle/animalcrossing/villager/deer/beau"
	"github.com/lindsaygelle/animalcrossing/villager/deer/bruce"
	"github.com/lindsaygelle/animalcrossing/villager/deer/chelsea"
	"github.com/lindsaygelle/animalcrossing/villager/deer/deirdre"
	"github.com/lindsaygelle/animalcrossing/villager/deer/diana"
	"github.com/lindsaygelle/animalcrossing/villager/deer/erik"
	"github.com/lindsaygelle/animalcrossing/villager/deer/fauna"
	"github.com/lindsaygelle/animalcrossing/villager/deer/fuchsia"
	"github.com/lindsaygelle/animalcrossing/villager/deer/lopez"
	"github.com/lindsaygelle/animalcrossing/villager/deer/zell"
)

var (
	Villagers = villager.NewVillagers(
		bam.Villager,
		beau.Villager,
		bruce.Villager,
		chelsea.Villager,
		deirdre.Villager,
		diana.Villager,
		erik.Villager,
		fauna.Villager,
		fuchsia.Villager,
		lopez.Villager,
		zell.Villager)
)
