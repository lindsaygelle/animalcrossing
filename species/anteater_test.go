package species

import "testing"

func TestAnteaterClass(t *testing.T) {
	if ok := Anteater.Class() == mammalia; !ok {
		t.Fatal("Anteater.Class() != mammalia")
	}
}
func TestAnteaterConservation(t *testing.T) {
	if ok := Anteater.Conservation() == vulnerable; !ok {
		t.Fatal("Anteater.Conservation() != vulnerable")
	}
}
func TestAnteaterDomain(t *testing.T) {
	if ok := Anteater.Domain() == eukarya; !ok {
		t.Fatal("Anteater.Domain() != eukarya")
	}
}
func TestAnteaterFamily(t *testing.T) {
	if ok := Anteater.Family() == ""; !ok {
		t.Fatal("Anteater.Family() != \"\"")
	}
}
func TestAnteaterGenus(t *testing.T) {
	if ok := Anteater.Genus() == ""; !ok {
		t.Fatal("Anteater.Genus() != \"\"")
	}
}
func TestAnteaterKingdom(t *testing.T) {
	if ok := Anteater.Kingdom() == animalia; !ok {
		t.Fatal("Anteater.Kingdom() != animalia")
	}
}

func TestAnteaterOrder(t *testing.T) {
	if ok := Anteater.Order() == pilosa; !ok {
		t.Fatal("Anteater.Order != pilosa")
	}
}

func TestAnteaterPhylum(t *testing.T) {
	if ok := Anteater.Phylum() == chordata; !ok {
		t.Fatal("Anteater.Phylum() != chordata")
	}
}

func TestAnteaterSpecies(t *testing.T) {
	if ok := Anteater.Species() == ""; !ok {
		t.Fatal("Anteater.Species() != \"\"")
	}
}
