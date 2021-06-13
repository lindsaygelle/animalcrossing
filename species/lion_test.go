package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"testing"
)

func TestLionClass(t *testing.T) {
	if ok := Lion.Class() == ammalia; !ok {
		t.Fatal("Lion.Class() != ammalia")
	}
}

func TestLionConservation(t *testing.T) {
	if ok := Lion.Conservation() == vulnerable; !ok {
		t.Fatal("Lion.Conservation() != vulnerable")
	}
}

func TestLionDomain(t *testing.T) {
	if ok := Lion.Domain() == eukarya; !ok {
		t.Fatal("Lion.Domain() != eukarya")
	}
}

func TestLionFamily(t *testing.T) {
	if ok := Lion.Family() == felidae; !ok {
		t.Fatal("Lion.Family() != felidae")
	}
}

func TestLionGenus(t *testing.T) {
	if ok := Lion.Genus() == panthera; !ok {
		t.Fatal("Lion.Genus() != panthera")
	}
}

func TestLionKingdom(t *testing.T) {
	if ok := Lion.Kingdom() == animalia; !ok {
		t.Fatal("Lion.Kingdom() != animalia")
	}
}

func TestLionName(t *testing.T) {
	if ok := Lion.Name() != na; !ok {
		t.Fatalf("Lion.Name != %s", animals.Lion.Name())
	}
}

func TestLionOrder(t *testing.T) {
	if ok := Lion.Order() == carnivora; !ok {
		t.Fatal("Lion.Order() != carnivora")
	}
}

func TestLionPhylum(t *testing.T) {
	if ok := Lion.Phylum() == chordata; !ok {
		t.Fatal("Lion.Phylum() != chordata")
	}
}

func TestLionSpecies(t *testing.T) {
	if ok := Lion.Species() == pLeo; !ok {
		t.Fatal("Lion.Species() != pLeo")
	}
}
