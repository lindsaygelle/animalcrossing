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

type alligator struct{}

func (a alligator) Animal() string       { return animals.Alligator.Name() }
func (a alligator) Class() string        { return classes.Reptilia.Name() }
func (a alligator) Conservation() string { return conservations.Vulnerable.Name() }
func (a alligator) Domain() string       { return domains.Eukarya.Name() }
func (a alligator) Family() string       { return families.Alligatoridae.Name() }
func (a alligator) Fictional() bool      { return false }
func (a alligator) Genus() string        { return genuses.Alligator.Name() }
func (a alligator) Kingdom() string      { return kingdoms.Animalia.Name() }
func (a alligator) Order() string        { return orders.Crocodylia.Name() }
func (a alligator) Phylum() string       { return phylums.Chordata.Name() }
func (a alligator) Species() string      { return "" }

var (
	Alligator Species = alligator{}
)
