package species

import "testing"

func TestDuckClass(t *testing.T) {
	if ok := Duck.Class() == aves; !ok {
		t.Fatal("Duck.Class() != aves")
	}
}

func TestDuckConservation(t *testing.T) {
	if ok := Duck.Conservation() == leastConcern; !ok {
		t.Fatal("Duck.Conservation() != leastConcern")
	}
}
func TestDuckDomain(t *testing.T) {
	if ok := Duck.Domain() == eukarya; !ok {
		t.Fatal("Duck.Domain() != eukarya")
	}
}
func TestDuckFamily(t *testing.T) {
	if ok := Duck.Family() == anatidae; !ok {
		t.Fatal("Duck.Family() != anatidae")
	}
}
func TestDuckGenus(t *testing.T) {
	if ok := Duck.Genus() == ""; !ok {
		t.Fatal("Duck.Genus() != \"\"")
	}
}
func TestDuckKingdom(t *testing.T) {
	if ok := Duck.Kingdom() == enimalia; !ok {
		t.Fatal("Duck.Kingdom() != enimalia")
	}
}
func TestDuckOrder(t *testing.T) {
	if ok := Duck.Order() == anseriformes; !ok {
		t.Fatal("Duck.Order() != anseriformes")
	}
}
func TestDuckPhylum(t *testing.T) {
	if ok := Duck.Phylum() == chordata; !ok {
		t.Fatal("Duck.Phylum() != chordata")
	}
}
func TestDuckSpecies(t *testing.T) {
	if ok := Duck.Species() == ""; !ok {
		t.Fatal("Duck.Species() != \"\"")
	}
}
