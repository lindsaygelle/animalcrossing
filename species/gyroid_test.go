package species

import "testing"

func TestGyroidClass(t *testing.T) {
	if ok := Gyroid.Class() == ""; !ok {
		t.Fatal("Gyroid.Class() != \"\"")
	}
}

func TestGyroidConservation(t *testing.T) {
	if ok := Gyroid.Conservation() == unknown; !ok {
		t.Fatal("Gyroid.Conservation() != unknown")
	}
}
func TestGyroidDomain(t *testing.T) {
	if ok := Gyroid.Domain() == eukarya; !ok {
		t.Fatal("Gyroid.Domain() != eukarya")
	}
}
func TestGyroidFamily(t *testing.T) {
	if ok := Gyroid.Family() == ""; !ok {
		t.Fatal("Gyroid.Family() != \"\"")
	}
}
func TestGyroidGenus(t *testing.T) {
	if ok := Gyroid.Genus() == ""; !ok {
		t.Fatal("Gyroid.Genus() != \"\"")
	}
}
func TestGyroidKingdom(t *testing.T) {
	if ok := Gyroid.Kingdom() == ""; !ok {
		t.Fatal("Gyroid.Kingdom() != \"\"")
	}
}
func TestGyroidOrder(t *testing.T) {
	if ok := Gyroid.Order() == ""; !ok {
		t.Fatal("Gyroid.Order() != \"\"")
	}
}
func TestGyroidPhylum(t *testing.T) {
	if ok := Gyroid.Phylum() == ""; !ok {
		t.Fatal("Gyroid.Phylum() != \"\"")
	}
}
func TestGyroidSpecies(t *testing.T) {
	if ok := Gyroid.Species() == ""; !ok {
		t.Fatal("Gyroid.Species() != \"\"")
	}
}
