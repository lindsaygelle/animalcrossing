package species

import "testing"

func TestGullClass(t *testing.T) {
	if ok := Gull.Class() == aaves; !ok {
		t.Fatal("Gull.Class() != aaves")
	}
}

func TestGullConservation(t *testing.T) {
	if ok := Gull.Conservation() == leastConcern; !ok {
		t.Fatal("Gull.Conservation() != leastConcern")
	}
}
func TestGullDomain(t *testing.T) {
	if ok := Gull.Domain() == eukarya; !ok {
		t.Fatal("Gull.Domain() != eukarya")
	}
}
func TestGullFamily(t *testing.T) {
	if ok := Gull.Family() == laridae; !ok {
		t.Fatal("Gull.Family() != laridae")
	}
}
func TestGullGenus(t *testing.T) {
	if ok := Gull.Genus() == ""; !ok {
		t.Fatal("Gull.Genus() != \"\"")
	}
}
func TestGullKingdom(t *testing.T) {
	if ok := Gull.Kingdom() == animalia; !ok {
		t.Fatal("Gull.Kingdom() != animalia")
	}
}
func TestGullOrder(t *testing.T) {
	if ok := Gull.Order() == charadriiformes; !ok {
		t.Fatal("Gull.Order() != charadriiformes")
	}
}
func TestGullPhylum(t *testing.T) {
	if ok := Gull.Phylum() == chordata; !ok {
		t.Fatal("Gull.Phylum() != chordata")
	}
}
func TestGullSpecies(t *testing.T) {
	if ok := Gull.Species() == ""; !ok {
		t.Fatal("Gull.Species() != \"\"")
	}
}
