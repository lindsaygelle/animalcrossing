package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAlligatorClass(t *testing.T) {
	if ok := Alligator.Class() == reptilia; !ok {
		t.Fatal("Alligator.Class() != reptilia")
	}
}

func TestAlligatorConservation(t *testing.T) {
	if ok := Alligator.Conservation() == vulnerable; !ok {
		t.Fatal("Alligator.Conservation() != vulnerable")
	}
}

func TestAlligatorDomain(t *testing.T) {
	if ok := Alligator.Domain() == eukarya; !ok {
		t.Fatal("Alligator.Domain() != eukarya")
	}
}

func TestAlligatorFamily(t *testing.T) {
	if ok := Alligator.Family() == alligatoridae; !ok {
		t.Fatal("Alligator.Family() != alligatoridae")
	}
}

func TestAlligatorFictional(t *testing.T) {
	if ok := Alligator.Fictional() == (!fictional); !ok {
		t.Fatal("Alligator.Fictional() != false")
	}
}

func TestAlligatorGenus(t *testing.T) {
	if ok := Alligator.Genus() == alligator; !ok {
		t.Fatal("Alligator.Genus() != alligator")
	}
}

func TestAlligatorKingdom(t *testing.T) {
	if ok := Alligator.Kingdom() == animalia; !ok {
		t.Fatal("Alligator.Kingdom() != animalia")
	}
}

func TestAlligatorOrder(t *testing.T) {
	if ok := Alligator.Order() == crocodylia; !ok {
		t.Fatal("Alligator.Order != crocodylia")
	}
}

func TestAlligatorName(t *testing.T) {
	if ok := Alligator.Name() != na; !ok {
		t.Fatalf("Alliagtor.Name != %s", animals.Alligator.Name())
	}
}

func TestAlligatorPhylum(t *testing.T) {
	if ok := Alligator.Phylum() == chordata; !ok {
		t.Fatal("Alligator.Phylum() != chordata")
	}
}

func TestAlligatorSpecies(t *testing.T) {
	if ok := Alligator.Species() == na; !ok {
		t.Fatal("Alligator.Species() != na")
	}
}
