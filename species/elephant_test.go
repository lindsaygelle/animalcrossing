package species

import "testing"

func TestElephantClass(t *testing.T) {
	if ok := Elephant.Class() == mammalia; !ok {
		t.Fatal("Elephant.Class() != mammalia")
	}
}

func TestElephantConservation(t *testing.T) {
	if ok := Elephant.Conservation() == endangered; !ok {
		t.Fatal("Elephant.Conservation() != endangered")
	}
}
func TestElephantDomain(t *testing.T) {
	if ok := Elephant.Domain() == eukarya; !ok {
		t.Fatal("Elephant.Domain() != eukarya")
	}
}
func TestElephantFamily(t *testing.T) {
	if ok := Elephant.Family() == elephantidae; !ok {
		t.Fatal("Elephant.Family() != elephantidae")
	}
}
func TestElephantGenus(t *testing.T) {
	if ok := Elephant.Genus() == ""; !ok {
		t.Fatal("Elephant.Genus() != \"\"")
	}
}
func TestElephantKingdom(t *testing.T) {
	if ok := Elephant.Kingdom() == animalia; !ok {
		t.Fatal("Elephant.Kingdom() != animalia")
	}
}
func TestElephantOrder(t *testing.T) {
	if ok := Elephant.Order() == proboscidea; !ok {
		t.Fatal("Elephant.Order() != proboscidea")
	}
}
func TestElephantPhylum(t *testing.T) {
	if ok := Elephant.Phylum() == chordata; !ok {
		t.Fatal("Elephant.Phylum() != chordata")
	}
}
func TestElephantSpecies(t *testing.T) {
	if ok := Elephant.Species() == bTaurus; !ok {
		t.Fatal("Elephant.Species() != \"\"")
	}
}
