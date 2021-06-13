package species

import "testing"

func TestOwlClass(t *testing.T) {
	if ok := Owl.Class() == aves; !ok {
		t.Fatal("Owl.Class() != aves")
	}
}

func TestOwlConservation(t *testing.T) {
	if ok := Owl.Conservation() == unknown; !ok {
		t.Fatal("Owl.Conservation() != unknown")
	}
}
func TestOwlDomain(t *testing.T) {
	if ok := Owl.Domain() == eukarya; !ok {
		t.Fatal("Owl.Domain() != eukarya")
	}
}
func TestOwlFamily(t *testing.T) {
	if ok := Owl.Family() == ""; !ok {
		t.Fatal("Owl.Family() != \"\"")
	}
}
func TestOwlGenus(t *testing.T) {
	if ok := Owl.Genus() == ""; !ok {
		t.Fatal("Owl.Genus() != \"\"")
	}
}
func TestOwlKingdom(t *testing.T) {
	if ok := Owl.Kingdom() == animalia; !ok {
		t.Fatal("Owl.Kingdom() != animalia")
	}
}
func TestOwlOrder(t *testing.T) {
	if ok := Owl.Order() == strigiformes; !ok {
		t.Fatal("Owl.Order() != strigiformes")
	}
}
func TestOwlPhylum(t *testing.T) {
	if ok := Owl.Phylum() == chordata; !ok {
		t.Fatal("Owl.Phylum() != chordata")
	}
}
func TestOwlSpecies(t *testing.T) {
	if ok := Owl.Species() == ""; !ok {
		t.Fatal("Owl.Species() != \"\"")
	}
}
