package species

import "testing"

func TestAlligatorClass(t *testing.T) {
	if ok := Alligator.Class() == reptilia; !ok {
		t.Fatal("Alligator.Class() != reptilia")
	}
}
func TestAlligatorConservation(t *testing.T) {
	if ok := Alligator.Conservation() == vulnerable; !ok {
		t.Fatal("Alligator.Conservation() != vulnerable")
	}
}
func TestAlligatorDomain(t *testing.T) {
	if ok := Alligator.Domain() == eukarya; !ok {
		t.Fatal("Alligator.Domain() != eukarya")
	}
}
func TestAlligatorFamily(t *testing.T) {
	if ok := Alligator.Family() == alligatoridae; !ok {
		t.Fatal("Alligator.Family() != alligatoridae")
	}
}
func TestAlligatorGenus(t *testing.T) {
	if ok := Alligator.Genus() == alligator; !ok {
		t.Fatal("Alligator.Genus() != alligator")
	}
}
func TestAlligatorKingdom(t *testing.T) {
	if ok := Alligator.Kingdom() == animalia; !ok {
		t.Fatal("Alligator.Kingdom() != animalia")
	}
}
func TestAlligatorPhylum(t *testing.T) {
	if ok := Alligator.Phylum() == chordata; !ok {
	}
}
