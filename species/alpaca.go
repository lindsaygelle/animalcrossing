package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/species/classes"
	"github.com/lindsaygelle/animalcrossing/species/conservations"
	"github.com/lindsaygelle/animalcrossing/species/domains"
	"github.com/lindsaygelle/animalcrossing/species/families"
	"github.com/lindsaygelle/animalcrossing/species/genuses"
	"github.com/lindsaygelle/animalcrossing/species/kingdoms"
	"github.com/lindsaygelle/animalcrossing/species/orders"
	"github.com/lindsaygelle/animalcrossing/species/phylums"
)

type alpaca struct{}

func (a alpaca) Animal() string       { return animals.Alpaca.Name() }
func (a alpaca) Class() string        { return classes.Mammalia.Name() }
func (a alpaca) Conservation() string { return conservations.Domesticated.Name() }
func (a alpaca) Domain() string       { return domains.Eukarya.Name() }
func (a alpaca) Family() string       { return families.Camelidae.Name() }
func (a alpaca) Fictional() bool      { return false }
func (a alpaca) Genus() string        { return genuses.Vicugna.Name() }
func (a alpaca) Kingdom() string      { return kingdoms.Animalia.Name() }
func (a alpaca) Order() string        { return orders.Artiodactyla.Name() }
func (a alpaca) Phylum() string       { return phylums.Chordata.Name() }
func (a alpaca) Species() string      { return "" }

var (
	Alpaca Species = alpaca{}
)
