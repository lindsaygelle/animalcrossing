package species

import "testing"

func TestCowClass(t *testing.T) {
	if ok := Cow.Class() == mammalia; !ok {
		t.Fatal("Cow.Class() != mammalia")
	}
}

func TestCowConservation(t *testing.T) {
	if ok := Cow.Conservation() == domesticated; !ok {
		t.Fatal("Cow.Conservation() != domesticated")
	}
}
func TestCowDomain(t *testing.T) {
	if ok := Cow.Domain() == eukarya; !ok {
		t.Fatal("Cow.Domain() != eukarya")
	}
}
func TestCowFamily(t *testing.T) {
	if ok := Cow.Family() == bovidae; !ok {
		t.Fatal("Cow.Family() != bovidae")
	}
}
func TestCowGenus(t *testing.T) {
	if ok := Cow.Genus() == bos; !ok {
		t.Fatal("Cow.Genus() != bos")
	}
}
func TestCowKingdom(t *testing.T) {
	if ok := Cow.Kingdom() == animalia; !ok {
		t.Fatal("Cow.Kingdom() != animalia")
	}
}
func TestCowOrder(t *testing.T) {
	if ok := Cow.Order() == artiodactyla; !ok {
		t.Fatal("Cow.Order() != artiodactyla")
	}
}
func TestCowPhylum(t *testing.T) {
	if ok := Cow.Phylum() == chordata; !ok {
		t.Fatal("Cow.Phylum() != chordata")
	}
}
func TestCowSpecies(t *testing.T) {
	if ok := Cow.Species() == bTaurus; !ok {
		t.Fatal("Cow.Species() != bTaurus")
	}
}
