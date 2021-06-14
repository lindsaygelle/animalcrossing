package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTigerClass(t *testing.T) {
	if ok := Tiger.Class() == mammalia; !ok {
		t.Fatal("Tiger.Class() != mammalia")
	}
}

func TestTigerConservation(t *testing.T) {
	if ok := Tiger.Conservation() == endangered; !ok {
		t.Fatal("Tiger.Conservation() != endangered")
	}
}

func TestTigerDomain(t *testing.T) {
	if ok := Tiger.Domain() == eukarya; !ok {
		t.Fatal("Tiger.Domain() != eukarya")
	}
}

func TestTigerFamily(t *testing.T) {
	if ok := Tiger.Family() == felidae; !ok {
		t.Fatal("Tiger.Family() != felidae")
	}
}

func TestTigerFictional(t *testing.T) {
	if ok := Tiger.Fictional() == (!fictional); !ok {
		t.Fatal("Tiger.Fictional() != false")
	}
}

func TestTigerGenus(t *testing.T) {
	if ok := Tiger.Genus() == panthera; !ok {
		t.Fatal("Tiger.Genus() != panthera")
	}
}

func TestTigerKingdom(t *testing.T) {
	if ok := Tiger.Kingdom() == animalia; !ok {
		t.Fatal("Tiger.Kingdom() != animalia")
	}
}

func TestTigerName(t *testing.T) {
	if ok := Tiger.Name() != na; !ok {
		t.Fatalf("Tiger.Name != %s", animals.Tiger.Name())
	}
}

func TestTigerOrder(t *testing.T) {
	if ok := Tiger.Order() == carnivora; !ok {
		t.Fatal("Tiger.Order() != carnivora")
	}
}

func TestTigerPhylum(t *testing.T) {
	if ok := Tiger.Phylum() == chordata; !ok {
		t.Fatal("Tiger.Phylum() != chordata")
	}
}

func TestTigerSpecies(t *testing.T) {
	if ok := Tiger.Species() == pTigris; !ok {
		t.Fatal("Tiger.Species() != pTigris")
	}
}
