package species

import "testing"

func TestHedgehogClass(t *testing.T) {
	if ok := Hedgehog.Class() == mammalia; !ok {
		t.Fatal("Hedgehog.Class() != mammalia")
	}
}

func TestHedgehogConservation(t *testing.T) {
	if ok := Hedgehog.Conservation() == leastConcern; !ok {
		t.Fatal("Hedgehog.Conservation() != leastConcern")
	}
}
func TestHedgehogDomain(t *testing.T) {
	if ok := Hedgehog.Domain() == eukarya; !ok {
		t.Fatal("Hedgehog.Domain() != eukarya")
	}
}
func TestHedgehogFamily(t *testing.T) {
	if ok := Hedgehog.Family() == erinaceidae; !ok {
		t.Fatal("Hedgehog.Family() != erinaceidae")
	}
}
func TestHedgehogGenus(t *testing.T) {
	if ok := Hedgehog.Genus() == ""; !ok {
		t.Fatal("Hedgehog.Genus() != \"\"")
	}
}
func TestHedgehogKingdom(t *testing.T) {
	if ok := Hedgehog.Kingdom() == animalia; !ok {
		t.Fatal("Hedgehog.Kingdom() != animalia")
	}
}
func TestHedgehogOrder(t *testing.T) {
	if ok := Hedgehog.Order() == eulipotyphla; !ok {
		t.Fatal("Hedgehog.Order() != eulipotyphla")
	}
}
func TestHedgehogPhylum(t *testing.T) {
	if ok := Hedgehog.Phylum() == chordata; !ok {
		t.Fatal("Hedgehog.Phylum() != chordata")
	}
}
func TestHedgehogSpecies(t *testing.T) {
	if ok := Hedgehog.Species() == ""; !ok {
		t.Fatal("Hedgehog.Species() != \"\"")
	}
}
