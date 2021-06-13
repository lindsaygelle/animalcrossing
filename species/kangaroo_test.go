package species

import "testing"

func TestKangarooClass(t *testing.T) {
	if ok := Kangaroo.Class() == mammalia; !ok {
		t.Fatal("Kangaroo.Class() != mammalia")
	}
}

func TestKangarooConservation(t *testing.T) {
	if ok := Kangaroo.Conservation() == leastConcern; !ok {
		t.Fatal("Kangaroo.Conservation() != leastConcern")
	}
}
func TestKangarooDomain(t *testing.T) {
	if ok := Kangaroo.Domain() == eukarya; !ok {
		t.Fatal("Kangaroo.Domain() != eukarya")
	}
}
func TestKangarooFamily(t *testing.T) {
	if ok := Kangaroo.Family() == macropodidae; !ok {
		t.Fatal("Kangaroo.Family() != macropodidae")
	}
}
func TestKangarooGenus(t *testing.T) {
	if ok := Kangaroo.Genus() == macropus; !ok {
		t.Fatal("Kangaroo.Genus() != macropus")
	}
}
func TestKangarooKingdom(t *testing.T) {
	if ok := Kangaroo.Kingdom() == animalia; !ok {
		t.Fatal("Kangaroo.Kingdom() != animalia")
	}
}
func TestKangarooOrder(t *testing.T) {
	if ok := Kangaroo.Order() == diprotodontia; !ok {
		t.Fatal("Kangaroo.Order() != diprotodontia")
	}
}
func TestKangarooPhylum(t *testing.T) {
	if ok := Kangaroo.Phylum() == chordata; !ok {
		t.Fatal("Kangaroo.Phylum() != chordata")
	}
}
func TestKangarooSpecies(t *testing.T) {
	if ok := Kangaroo.Species() == ""; !ok {
		t.Fatal("Kangaroo.Species() != \"\"")
	}
}
