package species

import "testing"

func TestFrogClass(t *testing.T) {
	if ok := Frog.Class() == amphibia; !ok {
		t.Fatal("Frog.Class() != amphibia")
	}
}

func TestFrogConservation(t *testing.T) {
	if ok := Frog.Conservation() == vulnerable; !ok {
		t.Fatal("Frog.Conservation() != vulnerable")
	}
}
func TestFrogDomain(t *testing.T) {
	if ok := Frog.Domain() == eukarya; !ok {
		t.Fatal("Frog.Domain() != eukarya")
	}
}
func TestFrogFamily(t *testing.T) {
	if ok := Frog.Family() == ""; !ok {
		t.Fatal("Frog.Family() != \"\"")
	}
}
func TestFrogGenus(t *testing.T) {
	if ok := Frog.Genus() == ""; !ok {
		t.Fatal("Frog.Genus() != \"\"")
	}
}
func TestFrogKingdom(t *testing.T) {
	if ok := Frog.Kingdom() == animalia; !ok {
		t.Fatal("Frog.Kingdom() != animalia")
	}
}
func TestFrogOrder(t *testing.T) {
	if ok := Frog.Order() == anura; !ok {
		t.Fatal("Frog.Order() != anura")
	}
}
func TestFrogPhylum(t *testing.T) {
	if ok := Frog.Phylum() == chordata; !ok {
		t.Fatal("Frog.Phylum() != chordata")
	}
}
func TestFrogSpecies(t *testing.T) {
	if ok := Frog.Species() == ""; !ok {
		t.Fatal("Frog.Species() != \"\"")
	}
}
