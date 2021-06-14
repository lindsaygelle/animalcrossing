package species

import "testing"

func TestBullClass(t *testing.T) {
	if ok := Bull.Class() == mammalia; !ok {
		t.Fatal("Bull.Class() != mammalia")
	}
}

func TestBullConservation(t *testing.T) {
	if ok := Bull.Conservation() == domesticated; !ok {
		t.Fatal("Bull.Conservation() != domesticated")
	}
}

func TestBullDomain(t *testing.T) {
	if ok := Bull.Domain() == eukarya; !ok {
		t.Fatal("Bull.Domain() != eukarya")
	}
}

func TestBullFamily(t *testing.T) {
	if ok := Bull.Family() == bovidae; !ok {
		t.Fatal("Bull.Family() != bovidae")
	}
}

func TestBullFictional(t *testing.T) {
	if ok := Bull.Fictional() == (!fictional); !ok {
		t.Fatal("Bull.Fictional() != false")
	}
}

func TestBullGenus(t *testing.T) {
	if ok := Bull.Genus() == bos; !ok {
		t.Fatal("Bull.Genus() != bos")
	}
}

func TestBullKingdom(t *testing.T) {
	if ok := Bull.Kingdom() == animalia; !ok {
		t.Fatal("Bull.Kingdom() != animalia")
	}
}

func TestBullOrder(t *testing.T) {
	if ok := Bull.Order() == artiodactyla; !ok {
		t.Fatal("Bull.Order() != artiodactyla")
	}
}

func TestBullPhylum(t *testing.T) {
	if ok := Bull.Phylum() == chordata; !ok {
		t.Fatal("Bull.Phylum() != chordata")
	}
}

func TestBullSpecies(t *testing.T) {
	if ok := Bull.Species() == bTaurus; !ok {
		t.Fatal("Bull.Species() != bTaurus")
	}
}
