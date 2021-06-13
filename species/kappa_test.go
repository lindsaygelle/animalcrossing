package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"testing"
)

func TestKappaClass(t *testing.T) {
	if ok := Kappa.Class() == na; !ok {
		t.Fatal("Kappa.Class() != na")
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
	if ok := Kappa.Family() == na; !ok {
		t.Fatal("Kappa.Family() != na")
	}
}

func TestKappaGenus(t *testing.T) {
	if ok := Kappa.Genus() == na; !ok {
		t.Fatal("Kappa.Genus() != na")
	}
}

func TestKappaKingdom(t *testing.T) {
	if ok := Kappa.Kingdom() == na; !ok {
		t.Fatal("Kappa.Kingdom() != na")
	}
}

func TestKappaName(t *testing.T) {
	if ok := Kappa.Name() != na; !ok {
		t.Fatalf("Kappa.Name != %s", animals.Kappa.Name())
	}
}

func TestKappaOrder(t *testing.T) {
	if ok := Kappa.Order() == na; !ok {
		t.Fatal("Kappa.Order() != na")
	}
}

func TestKappaPhylum(t *testing.T) {
	if ok := Kappa.Phylum() == na; !ok {
		t.Fatal("Kappa.Phylum() != na")
	}
}

func TestKappaSpecies(t *testing.T) {
	if ok := Kappa.Species() == na; !ok {
		t.Fatal("Kappa.Species() != na")
	}
}
