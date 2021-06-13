package species

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"testing"
)

func TestGorillaClass(t *testing.T) {
	if ok := Gorilla.Class() == mammalia; !ok {
		t.Fatal("Gorilla.Class() != mammalia")
	}
}

func TestGorillaConservation(t *testing.T) {
	if ok := Gorilla.Conservation() == endangered; !ok {
		t.Fatal("Gorilla.Conservation() != endangered")
	}
}

func TestGorillaDomain(t *testing.T) {
	if ok := Gorilla.Domain() == eukarya; !ok {
		t.Fatal("Gorilla.Domain() != eukarya")
	}
}

func TestGorillaFamily(t *testing.T) {
	if ok := Gorilla.Family() == hominidae; !ok {
		t.Fatal("Gorilla.Family() != hominidae")
	}
}

func TestGorillaGenus(t *testing.T) {
	if ok := Gorilla.Genus() == gorilla; !ok {
		t.Fatal("Gorilla.Genus() != gorilla")
	}
}

func TestGorillaKingdom(t *testing.T) {
	if ok := Gorilla.Kingdom() == animalia; !ok {
		t.Fatal("Gorilla.Kingdom() != animalia")
	}
}

func TestGorillaName(t *testing.T) {
	if ok := Gorilla.Name() != na; !ok {
		t.Fatalf("Gorilla.Name != %s", animals.Gorilla.Name())
	}
}

func TestGorillaOrder(t *testing.T) {
	if ok := Gorilla.Order() == primates; !ok {
		t.Fatal("Gorilla.Order() != primates")
	}
}

func TestGorillaPhylum(t *testing.T) {
	if ok := Gorilla.Phylum() == chordata; !ok {
		t.Fatal("Gorilla.Phylum() != chordata")
	}
}

func TestGorillaSpecies(t *testing.T) {
	if ok := Gorilla.Species() == na; !ok {
		t.Fatal("Gorilla.Species() != na")
	}
}
