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

type anteater struct{}

func (a anteater) Animal() string       { return animals.Anteater.Name() }
func (a anteater) Class() string        { return classes.Mammalia.Name() }
func (a anteater) Conservation() string { return conservations.Vulnerable.Name() }
func (a anteater) Domain() string       { return domains.Eukarya.Name() }
func (a anteater) Family() string       { return families.Myrmecophagidae.Name() }
func (a anteater) Fictional() bool      { return false }
func (a anteater) Genus() string        { return genuses.Myrmecophaga.Name() }
func (a anteater) Kingdom() string      { return kingdoms.Animalia.Name() }
func (a anteater) Order() string        { return orders.Pilosa.Name() }
func (a anteater) Phylum() string       { return phylums.Chordata.Name() }
func (a anteater) Species() string      { return "" }

var (
	Anteater Anteaterer = anteater{}
)
