package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKoalaClass(t *testing.T) {
	if ok := Koala.Class() == mammalia; !ok {
		t.Fatal("Koala.Class() != mammalia")
	}
}

func TestKoalaConservation(t *testing.T) {
	if ok := Koala.Conservation() == vulnerable; !ok {
		t.Fatal("Koala.Conservation() != vulnerable")
	}
}

func TestKoalaDomain(t *testing.T) {
	if ok := Koala.Domain() == eukarya; !ok {
		t.Fatal("Koala.Domain() != eukarya")
	}
}

func TestKoalaFamily(t *testing.T) {
	if ok := Koala.Family() == phascolarctidae; !ok {
		t.Fatal("Koala.Family() != phascolarctidae")
	}
}

func TestKoalaFictional(t *testing.T) {
	if ok := Koala.Fictional() == (!fictional); !ok {
		t.Fatal("Koala.Fictional() != false")
	}
}

func TestKoalaGenus(t *testing.T) {
	if ok := Koala.Genus() == phascolarctos; !ok {
		t.Fatal("Koala.Genus() != phascolarctos")
	}
}

func TestKoalaKingdom(t *testing.T) {
	if ok := Koala.Kingdom() == animalia; !ok {
		t.Fatal("Koala.Kingdom() != animalia")
	}
}

func TestKoalaName(t *testing.T) {
	if ok := Koala.Name() != na; !ok {
		t.Fatalf("Koala.Name != %s", animals.Koala.Name())
	}
}

func TestKoalaOrder(t *testing.T) {
	if ok := Koala.Order() == diprotodontia; !ok {
		t.Fatal("Koala.Order() != diprotodontia")
	}
}

func TestKoalaPhylum(t *testing.T) {
	if ok := Koala.Phylum() == chordata; !ok {
		t.Fatal("Koala.Phylum() != chordata")
	}
}

func TestKoalaSpecies(t *testing.T) {
	if ok := Koala.Species() == pCinereus; !ok {
		t.Fatal("Koala.Species() != pCinereus")
	}
}
