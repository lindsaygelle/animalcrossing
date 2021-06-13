package species

import "testing"

func TestKappaClass(t *testing.T) {
	if ok := Kappa.Class() == ""; !ok {
		t.Fatal("Kappa.Class() != \"\"")
	}
}

func TestKappaConservation(t *testing.T) {
	if ok := Kappa.Conservation() == basedOnFolklore; !ok {
		t.Fatal("Kappa.Conservation() != basedOnFolklore")
	}
}
func TestKappaDomain(t *testing.T) {
	if ok := Kappa.Domain() == eukarya; !ok {
		t.Fatal("Kappa.Domain() != eukarya")
	}
}
func TestKappaFamily(t *testing.T) {
	if ok := Kappa.Family() == ""; !ok {
		t.Fatal("Kappa.Family() != \"\"")
	}
}
func TestKappaGenus(t *testing.T) {
	if ok := Kappa.Genus() == ""; !ok {
		t.Fatal("Kappa.Genus() != \"\"")
	}
}
func TestKappaKingdom(t *testing.T) {
	if ok := Kappa.Kingdom() == ""; !ok {
		t.Fatal("Kappa.Kingdom() != \"\"")
	}
}
func TestKappaOrder(t *testing.T) {
	if ok := Kappa.Order() == ""; !ok {
		t.Fatal("Kappa.Order() != \"\"")
	}
}
func TestKappaPhylum(t *testing.T) {
	if ok := Kappa.Phylum() == ""; !ok {
		t.Fatal("Kappa.Phylum() != \"\"")
	}
}
func TestKappaSpecies(t *testing.T) {
	if ok := Kappa.Species() == ""; !ok {
		t.Fatal("Kappa.Species() != \"\"")
	}
}
