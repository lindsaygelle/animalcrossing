package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTurkeyClass(t *testing.T) {
	if ok := Turkey.Class() == aves; !ok {
		t.Fatal("Turkey.Class() != aves")
	}
}

func TestTurkeyConservation(t *testing.T) {
	if ok := Turkey.Conservation() == unknown; !ok {
		t.Fatal("Turkey.Conservation() != unknown")
	}
}

func TestTurkeyDomain(t *testing.T) {
	if ok := Turkey.Domain() == eukarya; !ok {
		t.Fatal("Turkey.Domain() != eukarya")
	}
}

func TestTurkeyFamily(t *testing.T) {
	if ok := Turkey.Family() == phasianidae; !ok {
		t.Fatal("Turkey.Family() != phasianidae")
	}
}

func TestTurkeyGenus(t *testing.T) {
	if ok := Turkey.Genus() == meleagris; !ok {
		t.Fatal("Turkey.Genus() != meleagris")
	}
}

func TestTurkeyKingdom(t *testing.T) {
	if ok := Turkey.Kingdom() == animalia; !ok {
		t.Fatal("Turkey.Kingdom() != animalia")
	}
}

func TestTurkeyName(t *testing.T) {
	if ok := Turkey.Name() != na; !ok {
		t.Fatalf("Turkey.Name != %s", animals.Turkey.Name())
	}
}

func TestTurkeyOrder(t *testing.T) {
	if ok := Turkey.Order() == galliformes; !ok {
		t.Fatal("Turkey.Order() != galliformes")
	}
}

func TestTurkeyPhylum(t *testing.T) {
	if ok := Turkey.Phylum() == chordata; !ok {
		t.Fatal("Turkey.Phylum() != chordata")
	}
}

func TestTurkeySpecies(t *testing.T) {
	if ok := Turkey.Species() == na; !ok {
		t.Fatal("Turkey.Species() != na")
	}
}
