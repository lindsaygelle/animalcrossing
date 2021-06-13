package species

import "testing"

func TestChamelonClass(t *testing.T) {
	if ok := Chamelon.Class() == reptilia; !ok {
		t.Fatal("Chamelon.Class() != reptilia")
	}
}

func TestChamelonConservation(t *testing.T) {
	if ok := Chamelon.Conservation() == leastConcern; !ok {
		t.Fatal("Chamelon.Conservation() != leastConcern")
	}
}
func TestChamelonDomain(t *testing.T) {
	if ok := Chamelon.Domain() == eukarya; !ok {
		t.Fatal("Chamelon.Domain() != eukarya")
	}
}
func TestChamelonFamily(t *testing.T) {
	if ok := Chamelon.Family() == chamaeleonidae; !ok {
		t.Fatal("Chamelon.Family() != chamaeleonidae")
	}
}
func TestChamelonGenus(t *testing.T) {
	if ok := Chamelon.Genus() == ""; !ok {
		t.Fatal("Chamelon.Genus() != \"\"")
	}
}
func TestChamelonKingdom(t *testing.T) {
	if ok := Chamelon.Kingdom() == animalia; !ok {
		t.Fatal("Chamelon.Kingdom() != animalia")
	}
}
func TestChamelonOrder(t *testing.T) {
	if ok := Chamelon.Order() == squamata; !ok {
		t.Fatal("Chamelon.Order() != squamata")
	}
}
func TestChamelonPhylum(t *testing.T) {
	if ok := Chamelon.Phylum() == chordata; !ok {
		t.Fatal("Chamelon.Phylum() != chordata")
	}
}
func TestChamelonSpecies(t *testing.T) {
	if ok := Chamelon.Species() == ""; !ok {
		t.Fatal("Chamelon.Species() != \"\"")
	}
}
