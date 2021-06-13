package species

import "testing"

func TestBovineClass(t *testing.T) {
	if ok := Bovine.Class() == mammalia; !ok {
		t.Fatal("Bovine.Class() != mammalia")
	}
}

func TestBovineConservation(t *testing.T) {
	if ok := Bovine.Conservation() == domesticated; !ok {
		t.Fatal("Bovine.Conservation() != domesticated")
	}
}
func TestBovineDomain(t *testing.T) {
	if ok := Bovine.Domain() == eukarya; !ok {
		t.Fatal("Bovine.Domain() != eukarya")
	}
}
func TestBovineFamily(t *testing.T) {
	if ok := Bovine.Family() == bovidae; !ok {
		t.Fatal("Bovine.Family() != bovidae")
	}
}
func TestBovineGenus(t *testing.T) {
	if ok := Bovine.Genus() == bos; !ok {
		t.Fatal("Bovine.Genus() != bos")
	}
}
func TestBovineKingdom(t *testing.T) {
	if ok := Bovine.Kingdom() == animalia; !ok {
		t.Fatal("Bovine.Kingdom() != animalia")
	}
}
func TestBovineOrder(t *testing.T) {
	if ok := Bovine.Order() == artiodactyla; !ok {
		t.Fatal("Bovine.Order() != artiodactyla")
	}
}
func TestBovinePhylum(t *testing.T) {
	if ok := Bovine.Phylum() == chordata; !ok {
		t.Fatal("Bovine.Phylum() != chordata")
	}
}
func TestBovineSpecies(t *testing.T) {
	if ok := Bovine.Species() == bTaurus; !ok {
		t.Fatal("Bovine.Species() != bTaurus")
	}
}
