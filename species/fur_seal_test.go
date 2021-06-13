package species

import "testing"

func TestFurSealClass(t *testing.T) {
	if ok := FurSeal.Class() == mammalia; !ok {
		t.Fatal("FurSeal.Class() != mammalia")
	}
}

func TestFurSealConservation(t *testing.T) {
	if ok := FurSeal.Conservation() == unknown; !ok {
		t.Fatal("FurSeal.Conservation() != unknown")
	}
}
func TestFurSealDomain(t *testing.T) {
	if ok := FurSeal.Domain() == eukarya; !ok {
		t.Fatal("FurSeal.Domain() != eukarya")
	}
}
func TestFurSealFamily(t *testing.T) {
	if ok := FurSeal.Family() == otariidae; !ok {
		t.Fatal("FurSeal.Family() != otariidae")
	}
}
func TestFurSealGenus(t *testing.T) {
	if ok := FurSeal.Genus() == ""; !ok {
		t.Fatal("FurSeal.Genus() != \"\"")
	}
}
func TestFurSealKingdom(t *testing.T) {
	if ok := FurSeal.Kingdom() == animalia; !ok {
		t.Fatal("FurSeal.Kingdom() != animalia")
	}
}
func TestFurSealOrder(t *testing.T) {
	if ok := FurSeal.Order() == carnivora; !ok {
		t.Fatal("FurSeal.Order() != carnivora")
	}
}
func TestFurSealPhylum(t *testing.T) {
	if ok := FurSeal.Phylum() == chordata; !ok {
		t.Fatal("FurSeal.Phylum() != chordata")
	}
}
func TestFurSealSpecies(t *testing.T) {
	if ok := FurSeal.Species() == ""; !ok {
		t.Fatal("FurSeal.Species() != \"\"")
	}
}
