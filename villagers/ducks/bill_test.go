package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBillAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Bill.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bill.Animal(), s)
	}
}

func TestBillName(t *testing.T) {
	if ok := Bill.Name() == bill; !ok {
		t.Fatalf("%s != %s", Bill.Name(), bill)
	}
}
