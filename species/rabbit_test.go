package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRabbitClass(t *testing.T) {
	if ok := Rabbit.Class() == mammalia; !ok {
		t.Fatal("Rabbit.Class() != mammalia")
	}
}

func TestRabbitConservation(t *testing.T) {
	if ok := Rabbit.Conservation() == domesticated; !ok {
		t.Fatal("Rabbit.Conservation() != domesticated")
	}
}

func TestRabbitDomain(t *testing.T) {
	if ok := Rabbit.Domain() == eukarya; !ok {
		t.Fatal("Rabbit.Domain() != eukarya")
	}
}

func TestRabbitFamily(t *testing.T) {
	if ok := Rabbit.Family() == leporidae; !ok {
		t.Fatal("Rabbit.Family() != leporidae")
	}
}

func TestRabbitGenus(t *testing.T) {
	if ok := Rabbit.Genus() == na; !ok {
		t.Fatal("Rabbit.Genus() != na")
	}
}

func TestRabbitKingdom(t *testing.T) {
	if ok := Rabbit.Kingdom() == animalia; !ok {
		t.Fatal("Rabbit.Kingdom() != animalia")
	}
}

func TestRabbitName(t *testing.T) {
	if ok := Rabbit.Name() != na; !ok {
		t.Fatalf("Rabbit.Name != %s", animals.Rabbit.Name())
	}
}

func TestRabbitOrder(t *testing.T) {
	if ok := Rabbit.Order() == lagomorpha; !ok {
		t.Fatal("Rabbit.Order() != lagomorpha")
	}
}

func TestRabbitPhylum(t *testing.T) {
	if ok := Rabbit.Phylum() == chordata; !ok {
		t.Fatal("Rabbit.Phylum() != chordata")
	}
}

func TestRabbitSpecies(t *testing.T) {
	if ok := Rabbit.Species() == na; !ok {
		t.Fatal("Rabbit.Species() != na")
	}
}
