package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHippopotamusClass(t *testing.T) {
	if ok := Hippopotamus.Class() == mammalia; !ok {
		t.Fatal("Hippopotamus.Class() != mammalia")
	}
}

func TestHippopotamusConservation(t *testing.T) {
	if ok := Hippopotamus.Conservation() == vulnerable; !ok {
		t.Fatal("Hippopotamus.Conservation() != vulnerable")
	}
}

func TestHippopotamusDomain(t *testing.T) {
	if ok := Hippopotamus.Domain() == eukarya; !ok {
		t.Fatal("Hippopotamus.Domain() != eukarya")
	}
}

func TestHippopotamusFamily(t *testing.T) {
	if ok := Hippopotamus.Family() == hippopotamidea; !ok {
		t.Fatal("Hippopotamus.Family() != hippopotamidea")
	}
}

func TestHippopotamusFictional(t *testing.T) {
	if ok := Hippopotamus.Fictional() == (!fictional); !ok {
		t.Fatal("Hippopotamus.Fictional() != false")
	}
}

func TestHippopotamusGenus(t *testing.T) {
	if ok := Hippopotamus.Genus() == hippopotamus; !ok {
		t.Fatal("Hippopotamus.Genus() != hippopotamus")
	}
}

func TestHippopotamusKingdom(t *testing.T) {
	if ok := Hippopotamus.Kingdom() == animalia; !ok {
		t.Fatal("Hippopotamus.Kingdom() != animalia")
	}
}

func TestHippopotamusName(t *testing.T) {
	if ok := Hippopotamus.Name() != na; !ok {
		t.Fatalf("Hippopotamus.Name != %s", animals.Hippopotamus.Name())
	}
}

func TestHippopotamusOrder(t *testing.T) {
	if ok := Hippopotamus.Order() == artiodactyla; !ok {
		t.Fatal("Hippopotamus.Order() != artiodactyla")
	}
}

func TestHippopotamusPhylum(t *testing.T) {
	if ok := Hippopotamus.Phylum() == chordata; !ok {
		t.Fatal("Hippopotamus.Phylum() != chordata")
	}
}

func TestHippopotamusSpecies(t *testing.T) {
	if ok := Hippopotamus.Species() == na; !ok {
		t.Fatal("Hippopotamus.Species() != na")
	}
}
