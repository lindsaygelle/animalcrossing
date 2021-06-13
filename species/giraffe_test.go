package species

import "testing"

func TestGiraffeClass(t *testing.T) {
	if ok := Giraffe.Class() == mammalia; !ok {
		t.Fatal("Giraffe.Class() != mammalia")
	}
}

func TestGiraffeConservation(t *testing.T) {
	if ok := Giraffe.Conservation() == vulnerable; !ok {
		t.Fatal("Giraffe.Conservation() != vulnerable")
	}
}
func TestGiraffeDomain(t *testing.T) {
	if ok := Giraffe.Domain() == eukarya; !ok {
		t.Fatal("Giraffe.Domain() != eukarya")
	}
}
func TestGiraffeFamily(t *testing.T) {
	if ok := Giraffe.Family() == giraffidae; !ok {
		t.Fatal("Giraffe.Family() != giraffidae")
	}
}
func TestGiraffeGenus(t *testing.T) {
	if ok := Giraffe.Genus() == giraffa; !ok {
		t.Fatal("Giraffe.Genus() != giraffa")
	}
}
func TestGiraffeKingdom(t *testing.T) {
	if ok := Giraffe.Kingdom() == animalia; !ok {
		t.Fatal("Giraffe.Kingdom() != animalia")
	}
}
func TestGiraffeOrder(t *testing.T) {
	if ok := Giraffe.Order() == artiodactyla; !ok {
		t.Fatal("Giraffe.Order() != artiodactyla")
	}
}
func TestGiraffePhylum(t *testing.T) {
	if ok := Giraffe.Phylum() == chordata; !ok {
		t.Fatal("Giraffe.Phylum() != chordata")
	}
}
func TestGiraffeSpecies(t *testing.T) {
	if ok := Giraffe.Species() == gCamelopardalis; !ok {
		t.Fatal("Giraffe.Species() != gCamelopardalis")
	}
}
