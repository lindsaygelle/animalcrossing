package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFurSealClass(t *testing.T) {
	if ok := FurSeal.Class() == mammalia; !ok {
		t.Fatal("FurSeal.Class() != mammalia")
	}
}

func TestFurSealConservation(t *testing.T) {
	if ok := FurSeal.Conservation() == unknown; !ok {
		t.Fatal("FurSeal.Conservation() != unknown")
	}
}

func TestFurSealDomain(t *testing.T) {
	if ok := FurSeal.Domain() == eukarya; !ok {
		t.Fatal("FurSeal.Domain() != eukarya")
	}
}

func TestFurSealFamily(t *testing.T) {
	if ok := FurSeal.Family() == otariidae; !ok {
		t.Fatal("FurSeal.Family() != otariidae")
	}
}

func TestFurSealFictional(t *testing.T) {
	if ok := FurSeal.Fictional() == (!fictional); !ok {
		t.Fatal("FurSeal.Fictional() != false")
	}
}

func TestFurSealGenus(t *testing.T) {
	if ok := FurSeal.Genus() == na; !ok {
		t.Fatal("FurSeal.Genus() != na")
	}
}

func TestFurSealKingdom(t *testing.T) {
	if ok := FurSeal.Kingdom() == animalia; !ok {
		t.Fatal("FurSeal.Kingdom() != animalia")
	}
}

func TestFurSealName(t *testing.T) {
	if ok := FurSeal.Name() != na; !ok {
		t.Fatalf("FurSeal.Name != %s", animals.FurSeal.Name())
	}
}

func TestFurSealOrder(t *testing.T) {
	if ok := FurSeal.Order() == carnivora; !ok {
		t.Fatal("FurSeal.Order() != carnivora")
	}
}

func TestFurSealPhylum(t *testing.T) {
	if ok := FurSeal.Phylum() == chordata; !ok {
		t.Fatal("FurSeal.Phylum() != chordata")
	}
}

func TestFurSealSpecies(t *testing.T) {
	if ok := FurSeal.Species() == na; !ok {
		t.Fatal("FurSeal.Species() != na")
	}
}
