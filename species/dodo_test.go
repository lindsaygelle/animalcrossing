package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDodoClass(t *testing.T) {
	if ok := Dodo.Class() == aves; !ok {
		t.Fatal("Dodo.Class() != aves")
	}
}

func TestDodoConservation(t *testing.T) {
	if ok := Dodo.Conservation() == extinct; !ok {
		t.Fatal("Dodo.Conservation() != extinct")
	}
}

func TestDodoDomain(t *testing.T) {
	if ok := Dodo.Domain() == eukarya; !ok {
		t.Fatal("Dodo.Domain() != eukarya")
	}
}

func TestDodoFamily(t *testing.T) {
	if ok := Dodo.Family() == columbidae; !ok {
		t.Fatal("Dodo.Family() != columbidae")
	}
}

func TestDodoFictional(t *testing.T) {
	if ok := Dodo.Fictional() == (!fictional); !ok {
		t.Fatal("Dodo.Fictional() != false")
	}
}

func TestDodoGenus(t *testing.T) {
	if ok := Dodo.Genus() == raphus; !ok {
		t.Fatal("Dodo.Genus() != raphus")
	}
}

func TestDodoKingdom(t *testing.T) {
	if ok := Dodo.Kingdom() == animalia; !ok {
		t.Fatal("Dodo.Kingdom() != animalia")
	}
}

func TestDodoName(t *testing.T) {
	if ok := Dodo.Name() != na; !ok {
		t.Fatalf("Dodo.Name != %s", animals.Dodo.Name())
	}
}

func TestDodoOrder(t *testing.T) {
	if ok := Dodo.Order() == columbiformes; !ok {
		t.Fatal("Dodo.Order() != columbiformes")
	}
}

func TestDodoPhylum(t *testing.T) {
	if ok := Dodo.Phylum() == chordata; !ok {
		t.Fatal("Dodo.Phylum() != chordata")
	}
}

func TestDodoSpecies(t *testing.T) {
	if ok := Dodo.Species() == rCucullatus; !ok {
		t.Fatal("Dodo.Species() != rCucullatus")
	}
}
