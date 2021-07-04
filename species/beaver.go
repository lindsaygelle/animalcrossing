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

type beaver struct{}

func (b beaver) Animal() string       { return animals.Beaver.Name() }
func (b beaver) Class() string        { return classes.Mammalia.Name() }
func (b beaver) Conservation() string { return conservations.LeastConcern.Name() }
func (b beaver) Domain() string       { return domains.Eukarya.Name() }
func (b beaver) Family() string       { return families.Castoridae.Name() }
func (b beaver) Fictional() bool      { return false }
func (b beaver) Genus() string        { return genuses.Castor.Name() }
func (b beaver) Kingdom() string      { return kingdoms.Animalia.Name() }
func (b beaver) Order() string        { return orders.Rodentia.Name() }
func (b beaver) Phylum() string       { return phylums.Chordata.Name() }
func (b beaver) Species() string      { return "" }

var (
	Beaver Beaverer = beaver{}
)
