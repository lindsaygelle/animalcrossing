package gorilla

import (
	"github.com/lindsaygelle/animalcrossing/villager"
	"github.com/lindsaygelle/animalcrossing/villager/gorilla/al"
	"github.com/lindsaygelle/animalcrossing/villager/gorilla/boone"
	"github.com/lindsaygelle/animalcrossing/villager/gorilla/boyd"
	"github.com/lindsaygelle/animalcrossing/villager/gorilla/cesar"
	"github.com/lindsaygelle/animalcrossing/villager/gorilla/hans"
	"github.com/lindsaygelle/animalcrossing/villager/gorilla/jane"
	"github.com/lindsaygelle/animalcrossing/villager/gorilla/louie"
	"github.com/lindsaygelle/animalcrossing/villager/gorilla/peewee"
	"github.com/lindsaygelle/animalcrossing/villager/gorilla/rilla"
	"github.com/lindsaygelle/animalcrossing/villager/gorilla/rocket"
	"github.com/lindsaygelle/animalcrossing/villager/gorilla/violet"
	"github.com/lindsaygelle/animalcrossing/villager/gorilla/yodel"
)

var (
	Villagers = villager.NewVillagers(
		al.Villager,
		boone.Villager,
		boyd.Villager,
		cesar.Villager,
		hans.Villager,
		jane.Villager,
		louie.Villager,
		peewee.Villager,
		rilla.Villager,
		rocket.Villager,
		violet.Villager,
		yodel.Villager)
)
