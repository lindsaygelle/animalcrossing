package species

import "testing"

func TestBearClass(t *testing.T) {
	if ok := Bear.Class() == mammalia; !ok {
		t.Fatal("Bear.Class() != mammalia")
	}
}

func TestBearConservation(t *testing.T) {
	if ok := Bear.Conservation() == leastConcern; !ok {
		t.Fatal("Bear.Conservation() != leastConcern")
	}
}
func TestBearDomain(t *testing.T) {
	if ok := Bear.Domain() == eukarya; !ok {
		t.Fatal("Bear.Domain() != eukarya")
	}
}
func TestBearFamily(t *testing.T) {
	if ok := Bear.Family() == ursidae; !ok {
		t.Fatal("Bear.Family() != ursidae")
	}
}
func TestBearGenus(t *testing.T) {
	if ok := Bear.Genus() == ""; !ok {
		t.Fatal("Bear.Genus() != \"\"")
	}
}
func TestBearKingdom(t *testing.T) {
	if ok := Bear.Kingdom() == animalia; !ok {
		t.Fatal("Bear.Kingdom() != animalia")
	}
}
func TestBearOrder(t *testing.T) {
	if ok := Bear.Order() == carnivora; !ok {
		t.Fatal("Bear.Order() != carnivora")
	}
}
func TestBearPhylum(t *testing.T) {
	if ok := Bear.Phylum() == chordata; !ok {
		t.Fatal("Bear.Phylum() != chordata")
	}
}
func TestBearSpecies(t *testing.T) {
	if ok := Bear.Species() == ""; !ok {
		t.Fatal("Bear.Species() != \"\"")
	}
}
