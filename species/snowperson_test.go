package species

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSnowpersonClass(t *testing.T) {
	if ok := Snowperson.Class() == na; !ok {
		t.Fatal("Snowperson.Class() != na")
	}
}

func TestSnowpersonConservation(t *testing.T) {
	if ok := Snowperson.Conservation() == na; !ok {
		t.Fatal("Snowperson.Conservation() != na")
	}
}

func TestSnowpersonDomain(t *testing.T) {
	if ok := Snowperson.Domain() == na; !ok {
		t.Fatal("Snowperson.Domain() != na")
	}
}

func TestSnowpersonFamily(t *testing.T) {
	if ok := Snowperson.Family() == na; !ok {
		t.Fatal("Snowperson.Family() != na")
	}
}

func TestSnowpersonFictional(t *testing.T) {
	if ok := Snowperson.Fictional() == (fictional); !ok {
		t.Fatal("Snowperson.Fictional() != true")
	}
}

func TestSnowpersonGenus(t *testing.T) {
	if ok := Snowperson.Genus() == na; !ok {
		t.Fatal("Snowperson.Genus() != na")
	}
}

func TestSnowpersonKingdom(t *testing.T) {
	if ok := Snowperson.Kingdom() == na; !ok {
		t.Fatal("Snowperson.Kingdom() != na")
	}
}

func TestSnowpersonName(t *testing.T) {
	if ok := Snowperson.Name() == animals.Snowperson.Name(); !ok {
		t.Fatalf("Snowperson.Name != %s", animals.Snowperson.Name())
	}
}

func TestSnowpersonOrder(t *testing.T) {
	if ok := Snowperson.Order() == na; !ok {
		t.Fatal("Snowperson.Order() != na")
	}
}

func TestSnowpersonPhylum(t *testing.T) {
	if ok := Snowperson.Phylum() == na; !ok {
		t.Fatal("Snowperson.Phylum() != na")
	}
}

func TestSnowpersonSpecies(t *testing.T) {
	if ok := Snowperson.Species() == na; !ok {
		t.Fatal("Snowperson.Species() != na")
	}
}
