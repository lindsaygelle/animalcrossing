package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHippoClass(t *testing.T) {
	if ok := Hippo.Class() == mammalia; !ok {
		t.Fatal("Hippo.Class() != mammalia")
	}
}

func TestHippoConservation(t *testing.T) {
	if ok := Hippo.Conservation() == vulnerable; !ok {
		t.Fatal("Hippo.Conservation() != vulnerable")
	}
}

func TestHippoDomain(t *testing.T) {
	if ok := Hippo.Domain() == eukarya; !ok {
		t.Fatal("Hippo.Domain() != eukarya")
	}
}

func TestHippoFamily(t *testing.T) {
	if ok := Hippo.Family() == hippopotamidea; !ok {
		t.Fatal("Hippo.Family() != hippopotamidea")
	}
}

func TestHippoFictional(t *testing.T) {
	if ok := Hippo.Fictional() == (!fictional); !ok {
		t.Fatal("Hippo.Fictional() != false")
	}
}

func TestHippoGenus(t *testing.T) {
	if ok := Hippo.Genus() == hippopotamus; !ok {
		t.Fatal("Hippo.Genus() != hippopotamus")
	}
}

func TestHippoKingdom(t *testing.T) {
	if ok := Hippo.Kingdom() == animalia; !ok {
		t.Fatal("Hippo.Kingdom() != animalia")
	}
}

func TestHippoName(t *testing.T) {
	if ok := Hippo.Name() != na; !ok {
		t.Fatalf("Hippo.Name != %s", animals.Hippo.Name())
	}
}

func TestHippoOrder(t *testing.T) {
	if ok := Hippo.Order() == artiodactyla; !ok {
		t.Fatal("Hippo.Order() != artiodactyla")
	}
}

func TestHippoPhylum(t *testing.T) {
	if ok := Hippo.Phylum() == chordata; !ok {
		t.Fatal("Hippo.Phylum() != chordata")
	}
}

func TestHippoSpecies(t *testing.T) {
	if ok := Hippo.Species() == na; !ok {
		t.Fatal("Hippo.Species() != na")
	}
}
