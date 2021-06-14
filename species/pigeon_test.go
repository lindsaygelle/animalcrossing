package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPigeonClass(t *testing.T) {
	if ok := Pigeon.Class() == aves; !ok {
		t.Fatal("Pigeon.Class() != aves")
	}
}

func TestPigeonConservation(t *testing.T) {
	if ok := Pigeon.Conservation() == unknown; !ok {
		t.Fatal("Pigeon.Conservation() != unknown")
	}
}

func TestPigeonDomain(t *testing.T) {
	if ok := Pigeon.Domain() == eukarya; !ok {
		t.Fatal("Pigeon.Domain() != eukarya")
	}
}

func TestPigeonFamily(t *testing.T) {
	if ok := Pigeon.Family() == columbidae; !ok {
		t.Fatal("Pigeon.Family() != columbidae")
	}
}

func TestPigeonFictional(t *testing.T) {
	if ok := Pigeon.Fictional() == (!fictional); !ok {
		t.Fatal("Pigeon.Fictional() != false")
	}
}

func TestPigeonGenus(t *testing.T) {
	if ok := Pigeon.Genus() == na; !ok {
		t.Fatal("Pigeon.Genus() != na")
	}
}

func TestPigeonKingdom(t *testing.T) {
	if ok := Pigeon.Kingdom() == animalia; !ok {
		t.Fatal("Pigeon.Kingdom() != animalia")
	}
}

func TestPigeonName(t *testing.T) {
	if ok := Pigeon.Name() != na; !ok {
		t.Fatalf("Pigeon.Name != %s", animals.Pigeon.Name())
	}
}

func TestPigeonOrder(t *testing.T) {
	if ok := Pigeon.Order() == calumbiformes; !ok {
		t.Fatal("Pigeon.Order() != calumbiformes")
	}
}

func TestPigeonPhylum(t *testing.T) {
	if ok := Pigeon.Phylum() == chordata; !ok {
		t.Fatal("Pigeon.Phylum() != chordata")
	}
}

func TestPigeonSpecies(t *testing.T) {
	if ok := Pigeon.Species() == na; !ok {
		t.Fatal("Pigeon.Species() != na")
	}
}
