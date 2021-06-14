package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTapirClass(t *testing.T) {
	if ok := Tapir.Class() == mammalia; !ok {
		t.Fatal("Tapir.Class() != mammalia")
	}
}

func TestTapirConservation(t *testing.T) {
	if ok := Tapir.Conservation() == unknown; !ok {
		t.Fatal("Tapir.Conservation() != unknown")
	}
}

func TestTapirDomain(t *testing.T) {
	if ok := Tapir.Domain() == eukarya; !ok {
		t.Fatal("Tapir.Domain() != eukarya")
	}
}

func TestTapirFamily(t *testing.T) {
	if ok := Tapir.Family() == tapiridae; !ok {
		t.Fatal("Tapir.Family() != tapiridae")
	}
}

func TestTapirFictional(t *testing.T) {
	if ok := Tapir.Fictional() == (!fictional); !ok {
		t.Fatal("Tapir.Fictional() != false")
	}
}

func TestTapirGenus(t *testing.T) {
	if ok := Tapir.Genus() == tapirus; !ok {
		t.Fatal("Tapir.Genus() != tapirus")
	}
}

func TestTapirKingdom(t *testing.T) {
	if ok := Tapir.Kingdom() == animalia; !ok {
		t.Fatal("Tapir.Kingdom() != animalia")
	}
}

func TestTapirName(t *testing.T) {
	if ok := Tapir.Name() != na; !ok {
		t.Fatalf("Tapir.Name != %s", animals.Tapir.Name())
	}
}

func TestTapirOrder(t *testing.T) {
	if ok := Tapir.Order() == perissodactyla; !ok {
		t.Fatal("Tapir.Order() != perissodactyla")
	}
}

func TestTapirPhylum(t *testing.T) {
	if ok := Tapir.Phylum() == chordata; !ok {
		t.Fatal("Tapir.Phylum() != chordata")
	}
}

func TestTapirSpecies(t *testing.T) {
	if ok := Tapir.Species() == na; !ok {
		t.Fatal("Tapir.Species() != na")
	}
}
