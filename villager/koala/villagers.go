package koala

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/koala/alice"
	"github.com/lindsaygelle/animalcrossing/villager/koala/canberra"
	"github.com/lindsaygelle/animalcrossing/villager/koala/eugene"
	"github.com/lindsaygelle/animalcrossing/villager/koala/faith"
	"github.com/lindsaygelle/animalcrossing/villager/koala/gonzo"
	"github.com/lindsaygelle/animalcrossing/villager/koala/huggy"
	"github.com/lindsaygelle/animalcrossing/villager/koala/lyman"
	"github.com/lindsaygelle/animalcrossing/villager/koala/melba"
	"github.com/lindsaygelle/animalcrossing/villager/koala/ozzie"
	"github.com/lindsaygelle/animalcrossing/villager/koala/sydney"
	"github.com/lindsaygelle/animalcrossing/villager/koala/yuka"
)

var (
	Villagers = villager.NewVillagers(
		alice.Villager,
		canberra.Villager,
		eugene.Villager,
		faith.Villager,
		gonzo.Villager,
		huggy.Villager,
		lyman.Villager,
		melba.Villager,
		ozzie.Villager,
		sydney.Villager,
		yuka.Villager)
)
