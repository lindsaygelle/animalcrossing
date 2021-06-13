package species

import "testing"

func TestMoleClass(t *testing.T) {
	if ok := Mole.Class() == mammalia; !ok {
		t.Fatal("Mole.Class() != mammalia")
	}
}

func TestMoleConservation(t *testing.T) {
	if ok := Mole.Conservation() == leastConcern; !ok {
		t.Fatal("Mole.Conservation() != leastConcern")
	}
}
func TestMoleDomain(t *testing.T) {
	if ok := Mole.Domain() == eukarya; !ok {
		t.Fatal("Mole.Domain() != eukarya")
	}
}
func TestMoleFamily(t *testing.T) {
	if ok := Mole.Family() == talpidae; !ok {
		t.Fatal("Mole.Family() != talpidae")
	}
}
func TestMoleGenus(t *testing.T) {
	if ok := Mole.Genus() == ""; !ok {
		t.Fatal("Mole.Genus() != \"\"")
	}
}
func TestMoleKingdom(t *testing.T) {
	if ok := Mole.Kingdom() == animalia; !ok {
		t.Fatal("Mole.Kingdom() != animalia")
	}
}
func TestMoleOrder(t *testing.T) {
	if ok := Mole.Order() == eulipotyphla; !ok {
		t.Fatal("Mole.Order() != eulipotyphla")
	}
}
func TestMolePhylum(t *testing.T) {
	if ok := Mole.Phylum() == chordata; !ok {
		t.Fatal("Mole.Phylum() != chordata")
	}
}
func TestMoleSpecies(t *testing.T) {
	if ok := Mole.Species() == ""; !ok {
		t.Fatal("Mole.Species() != \"\"")
	}
}
