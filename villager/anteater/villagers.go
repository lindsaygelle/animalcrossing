package anteater

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/anteater/anabelle"
	"github.com/lindsaygelle/animalcrossing/villager/anteater/annalisa"
	"github.com/lindsaygelle/animalcrossing/villager/anteater/antonio"
	"github.com/lindsaygelle/animalcrossing/villager/anteater/cyrano"
	"github.com/lindsaygelle/animalcrossing/villager/anteater/lulu"
	"github.com/lindsaygelle/animalcrossing/villager/anteater/nosegay"
	"github.com/lindsaygelle/animalcrossing/villager/anteater/olaf"
	"github.com/lindsaygelle/animalcrossing/villager/anteater/pango"
	"github.com/lindsaygelle/animalcrossing/villager/anteater/snooty"
	"github.com/lindsaygelle/animalcrossing/villager/anteater/zoe"
)

var (
	Villagers = villager.NewVillagers(
		anabelle.Villager,
		annalisa.Villager,
		antonio.Villager,
		cyrano.Villager,
		lulu.Villager,
		nosegay.Villager,
		olaf.Villager,
		pango.Villager,
		snooty.Villager,
		zoe.Villager)
)
