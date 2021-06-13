package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChameleonClass(t *testing.T) {
	if ok := Chameleon.Class() == reptilia; !ok {
		t.Fatal("Chameleon.Class() != reptilia")
	}
}

func TestChameleonConservation(t *testing.T) {
	if ok := Chameleon.Conservation() == leastConcern; !ok {
		t.Fatal("Chameleon.Conservation() != leastConcern")
	}
}

func TestChameleonDomain(t *testing.T) {
	if ok := Chameleon.Domain() == eukarya; !ok {
		t.Fatal("Chameleon.Domain() != eukarya")
	}
}

func TestChameleonFamily(t *testing.T) {
	if ok := Chameleon.Family() == chamaeleonidae; !ok {
		t.Fatal("Chameleon.Family() != chamaeleonidae")
	}
}

func TestChameleonGenus(t *testing.T) {
	if ok := Chameleon.Genus() == na; !ok {
		t.Fatal("Chameleon.Genus() != na")
	}
}

func TestChameleonKingdom(t *testing.T) {
	if ok := Chameleon.Kingdom() == animalia; !ok {
		t.Fatal("Chameleon.Kingdom() != animalia")
	}
}

func TestChameleonName(t *testing.T) {
	if ok := Chameleon.Name() != na; !ok {
		t.Fatalf("Chameleon.Name != %s", animals.Chameleon.Name())
	}
}

func TestChameleonOrder(t *testing.T) {
	if ok := Chameleon.Order() == squamata; !ok {
		t.Fatal("Chameleon.Order() != squamata")
	}
}

func TestChameleonPhylum(t *testing.T) {
	if ok := Chameleon.Phylum() == chordata; !ok {
		t.Fatal("Chameleon.Phylum() != chordata")
	}
}

func TestChameleonSpecies(t *testing.T) {
	if ok := Chameleon.Species() == na; !ok {
		t.Fatal("Chameleon.Species() != na")
	}
}
