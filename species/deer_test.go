package species

import "testing"

func TestDeerClass(t *testing.T) {
	if ok := Deer.Class() == mammalia; !ok {
		t.Fatal("Deer.Class() != mammalia")
	}
}

func TestDeerConservation(t *testing.T) {
	if ok := Deer.Conservation() == leastConcern; !ok {
		t.Fatal("Deer.Conservation() != leastConcern")
	}
}
func TestDeerDomain(t *testing.T) {
	if ok := Deer.Domain() == eukarya; !ok {
		t.Fatal("Deer.Domain() != eukarya")
	}
}
func TestDeerFamily(t *testing.T) {
	if ok := Deer.Family() == cervidae; !ok {
		t.Fatal("Deer.Family() != cervidae")
	}
}
func TestDeerGenus(t *testing.T) {
	if ok := Deer.Genus() == ""; !ok {
		t.Fatal("Deer.Genus() != \"\"")
	}
}
func TestDeerKingdom(t *testing.T) {
	if ok := Deer.Kingdom() == animalia; !ok {
		t.Fatal("Deer.Kingdom() != animalia")
	}
}
func TestDeerOrder(t *testing.T) {
	if ok := Deer.Order() == artiodactyla; !ok {
		t.Fatal("Deer.Order() != artiodactyla")
	}
}
func TestDeerPhylum(t *testing.T) {
	if ok := Deer.Phylum() == chordata; !ok {
		t.Fatal("Deer.Phylum() != chordata")
	}
}
func TestDeerSpecies(t *testing.T) {
	if ok := Deer.Species() == ""; !ok {
		t.Fatal("Deer.Species() != \"\"")
	}
}