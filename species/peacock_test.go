package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPeacockClass(t *testing.T) {
	if ok := Peacock.Class() == aves; !ok {
		t.Fatal("Peacock.Class() != aves")
	}
}

func TestPeacockConservation(t *testing.T) {
	if ok := Peacock.Conservation() == unknown; !ok {
		t.Fatal("Peacock.Conservation() != unknown")
	}
}

func TestPeacockDomain(t *testing.T) {
	if ok := Peacock.Domain() == eukarya; !ok {
		t.Fatal("Peacock.Domain() != eukarya")
	}
}

func TestPeacockFamily(t *testing.T) {
	if ok := Peacock.Family() == phasianidae; !ok {
		t.Fatal("Peacock.Family() != phasianidae")
	}
}

func TestPeacockGenus(t *testing.T) {
	if ok := Peacock.Genus() == pavo; !ok {
		t.Fatal("Peacock.Genus() != pavo")
	}
}

func TestPeacockKingdom(t *testing.T) {
	if ok := Peacock.Kingdom() == animalia; !ok {
		t.Fatal("Peacock.Kingdom() != animalia")
	}
}

func TestPeacockName(t *testing.T) {
	if ok := Peacock.Name() != na; !ok {
		t.Fatalf("Peacock.Name != %s", animals.Peacock.Name())
	}
}

func TestPeacockOrder(t *testing.T) {
	if ok := Peacock.Order() == galliformes; !ok {
		t.Fatal("Peacock.Order() != galliformes")
	}
}

func TestPeacockPhylum(t *testing.T) {
	if ok := Peacock.Phylum() == chordata; !ok {
		t.Fatal("Peacock.Phylum() != chordata")
	}
}

func TestPeacockSpecies(t *testing.T) {
	if ok := Peacock.Species() == na; !ok {
		t.Fatal("Peacock.Species() != na")
	}
}
