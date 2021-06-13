package species

import "testing"

func TestOtterClass(t *testing.T) {
	if ok := Otter.Class() == mammalia; !ok {
		t.Fatal("Otter.Class() != mammalia")
	}
}

func TestOtterConservation(t *testing.T) {
	if ok := Otter.Conservation() == nearThreatened; !ok {
		t.Fatal("Otter.Conservation() != nearThreatened")
	}
}
func TestOtterDomain(t *testing.T) {
	if ok := Otter.Domain() == eukarya; !ok {
		t.Fatal("Otter.Domain() != eukarya")
	}
}
func TestOtterFamily(t *testing.T) {
	if ok := Otter.Family() == mustelidae; !ok {
		t.Fatal("Otter.Family() != mustelidae")
	}
}
func TestOtterGenus(t *testing.T) {
	if ok := Otter.Genus() == ""; !ok {
		t.Fatal("Otter.Genus() != \"\"")
	}
}
func TestOtterKingdom(t *testing.T) {
	if ok := Otter.Kingdom() == animalia; !ok {
		t.Fatal("Otter.Kingdom() != animalia")
	}
}
func TestOtterOrder(t *testing.T) {
	if ok := Otter.Order() == carnivora; !ok {
		t.Fatal("Otter.Order() != carnivora")
	}
}
func TestOtterPhylum(t *testing.T) {
	if ok := Otter.Phylum() == chordata; !ok {
		t.Fatal("Otter.Phylum() != chordata")
	}
}
func TestOtterSpecies(t *testing.T) {
	if ok := Otter.Species() == ""; !ok {
		t.Fatal("Otter.Species() != \"\"")
	}
}
