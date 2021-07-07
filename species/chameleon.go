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

type chameleon struct{}

func (c chameleon) Animal() string       { return animals.Chameleon.Name() }
func (c chameleon) Class() string        { return classes.Reptilia.Name() }
func (c chameleon) Conservation() string { return conservations.LeastConcern.Name() }
func (c chameleon) Domain() string       { return domains.Eukarya.Name() }
func (c chameleon) Family() string       { return families.Chamaeleonidae.Name() }
func (c chameleon) Fictional() bool      { return false }
func (c chameleon) Genus() string        { return genuses.Felis.Name() }
func (c chameleon) Kingdom() string      { return kingdoms.Animalia.Name() }
func (c chameleon) Order() string        { return orders.Squamata.Name() }
func (c chameleon) Phylum() string       { return phylums.Chordata.Name() }
func (c chameleon) Species() string      { return "" }

var (
	Chameleon Species = chameleon{}
)
