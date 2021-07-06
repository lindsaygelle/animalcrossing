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

type bear struct{}

func (b bear) Animal() string       { return animals.Bear.Name() }
func (b bear) Class() string        { return classes.Mammalia.Name() }
func (b bear) Conservation() string { return conservations.LeastConcern.Name() }
func (b bear) Domain() string       { return domains.Eukarya.Name() }
func (b bear) Family() string       { return families.Ursidae.Name() }
func (b bear) Fictional() bool      { return false }
func (b bear) Genus() string        { return genuses.Ursus.Name() }
func (b bear) Kingdom() string      { return kingdoms.Animalia.Name() }
func (b bear) Order() string        { return orders.Carnivora.Name() }
func (b bear) Phylum() string       { return phylums.Chordata.Name() }
func (b bear) Species() string      { return "" }

var (
	Bear Species = bear{}
)
