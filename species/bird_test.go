package species

import "testing"

func TestBirdClass(t *testing.T) {
	if ok := Bird.Class() == aves; !ok {
		t.Fatal("Bird.Class() != aves")
	}
}

func TestBirdConservation(t *testing.T) {
	if ok := Bird.Conservation() == leastConcern; !ok {
		t.Fatal("Bird.Conservation() != leastConcern")
	}
}
func TestBirdDomain(t *testing.T) {
	if ok := Bird.Domain() == eukarya; !ok {
		t.Fatal("Bird.Domain() != eukarya")
	}
}
func TestBirdFamily(t *testing.T) {
	if ok := Bird.Family() == ""; !ok {
		t.Fatal("Bird.Family() != \"\"")
	}
}
func TestBirdGenus(t *testing.T) {
	if ok := Bird.Genus() == ""; !ok {
		t.Fatal("Bird.Genus() != \"\"")
	}
}
func TestBirdKingdom(t *testing.T) {
	if ok := Bird.Kingdom() == animalia; !ok {
		t.Fatal("Bird.Kingdom() != animalia")
	}
}
func TestBirdOrder(t *testing.T) {
	if ok := Bird.Order() == ""; !ok {
		t.Fatal("Bird.Order() != \"\"")
	}
}
func TestBirdPhylum(t *testing.T) {
	if ok := Bird.Phylum() == chordata; !ok {
		t.Fatal("Bird.Phylum() != chordata")
	}
}
func TestBirdSpecies(t *testing.T) {
	if ok := Bird.Species() == ""; !ok {
		t.Fatal("Bird.Species() != \"\"")
	}
}
