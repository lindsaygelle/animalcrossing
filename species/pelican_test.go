package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPelicanClass(t *testing.T) {
	if ok := Pelican.Class() == aves; !ok {
		t.Fatal("Pelican.Class() != aves")
	}
}

func TestPelicanConservation(t *testing.T) {
	if ok := Pelican.Conservation() == unknown; !ok {
		t.Fatal("Pelican.Conservation() != unknown")
	}
}

func TestPelicanDomain(t *testing.T) {
	if ok := Pelican.Domain() == eukarya; !ok {
		t.Fatal("Pelican.Domain() != eukarya")
	}
}

func TestPelicanFamily(t *testing.T) {
	if ok := Pelican.Family() == pelecanidae; !ok {
		t.Fatal("Pelican.Family() != pelecanidae")
	}
}

func TestPelicanFictional(t *testing.T) {
	if ok := Pelican.Fictional() == (!fictional); !ok {
		t.Fatal("Pelican.Fictional() != false")
	}
}

func TestPelicanGenus(t *testing.T) {
	if ok := Pelican.Genus() == pelecanus; !ok {
		t.Fatal("Pelican.Genus() != pelecanus")
	}
}

func TestPelicanKingdom(t *testing.T) {
	if ok := Pelican.Kingdom() == animalia; !ok {
		t.Fatal("Pelican.Kingdom() != animalia")
	}
}

func TestPelicanName(t *testing.T) {
	if ok := Pelican.Name() != na; !ok {
		t.Fatalf("Pelican.Name != %s", animals.Pelican.Name())
	}
}

func TestPelicanOrder(t *testing.T) {
	if ok := Pelican.Order() == pelecaniformes; !ok {
		t.Fatal("Pelican.Order() != pelecaniformes")
	}
}

func TestPelicanPhylum(t *testing.T) {
	if ok := Pelican.Phylum() == chordata; !ok {
		t.Fatal("Pelican.Phylum() != chordata")
	}
}

func TestPelicanSpecies(t *testing.T) {
	if ok := Pelican.Species() == na; !ok {
		t.Fatal("Pelican.Species() != na")
	}
}
