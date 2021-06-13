package species

import "testing"

func TestBoarClass(t *testing.T) {
	if ok := Boar.Class() == mammalia; !ok {
		t.Fatal("Boar.Class() != mammalia")
	}
}

func TestBoarConservation(t *testing.T) {
	if ok := Boar.Conservation() == leastConcern; !ok {
		t.Fatal("Boar.Conservation() != leastConcern")
	}
}
func TestBoarDomain(t *testing.T) {
	if ok := Boar.Domain() == eukarya; !ok {
		t.Fatal("Boar.Domain() != eukarya")
	}
}
func TestBoarFamily(t *testing.T) {
	if ok := Boar.Family() == suidae; !ok {
		t.Fatal("Boar.Family() != suidae")
	}
}
func TestBoarGenus(t *testing.T) {
	if ok := Boar.Genus() == sus; !ok {
		t.Fatal("Boar.Genus() != sus")
	}
}
func TestBoarKingdom(t *testing.T) {
	if ok := Boar.Kingdom() == animalia; !ok {
		t.Fatal("Boar.Kingdom() != animalia")
	}
}
func TestBoarOrder(t *testing.T) {
	if ok := Boar.Order() == artiodactyla; !ok {
		t.Fatal("Boar.Order() != artiodactyla")
	}
}
func TestBoarPhylum(t *testing.T) {
	if ok := Boar.Phylum() == chordata; !ok {
		t.Fatal("Boar.Phylum() != chordata")
	}
}
func TestBoarSpecies(t *testing.T) {
	if ok := Boar.Species() == sScrofa; !ok {
		t.Fatal("Boar.Species() != sScrofa")
	}
}
