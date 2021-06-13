package species

import "testing"

func TestMonkeyClass(t *testing.T) {
	if ok := Monkey.Class() == mammalia; !ok {
		t.Fatal("Monkey.Class() != mammalia")
	}
}

func TestMonkeyConservation(t *testing.T) {
	if ok := Monkey.Conservation() == endangered; !ok {
		t.Fatal("Monkey.Conservation() != endangered")
	}
}
func TestMonkeyDomain(t *testing.T) {
	if ok := Monkey.Domain() == eukarya; !ok {
		t.Fatal("Monkey.Domain() != eukarya")
	}
}
func TestMonkeyFamily(t *testing.T) {
	if ok := Monkey.Family() == ""; !ok {
		t.Fatal("Monkey.Family() != \"\"")
	}
}
func TestMonkeyGenus(t *testing.T) {
	if ok := Monkey.Genus() == ""; !ok {
		t.Fatal("Monkey.Genus() != \"\"")
	}
}
func TestMonkeyKingdom(t *testing.T) {
	if ok := Monkey.Kingdom() == animalia; !ok {
		t.Fatal("Monkey.Kingdom() != animalia")
	}
}
func TestMonkeyOrder(t *testing.T) {
	if ok := Monkey.Order() == primates; !ok {
		t.Fatal("Monkey.Order() != primates")
	}
}
func TestMonkeyPhylum(t *testing.T) {
	if ok := Monkey.Phylum() == chordata; !ok {
		t.Fatal("Monkey.Phylum() != chordata")
	}
}
func TestMonkeySpecies(t *testing.T) {
	if ok := Monkey.Species() == ""; !ok {
		t.Fatal("Monkey.Species() != \"\"")
	}
}
