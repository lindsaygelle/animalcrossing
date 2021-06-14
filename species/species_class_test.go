package species

import (
	"strings"
	"testing"
)

func TestClassAll(t *testing.T) {
	var (
		c = [...]string{
			aaves,
			ammalia,
			amphibia,
			aves,
			cephalopoda,
			mammalia,
			reptilia,
			sauropsida}
	)
	for i := 0; i < len(c); i++ {
		var s string = c[i]
		if ok := string(s[0]) == strings.ToUpper(string(s[0])); !ok {
			t.Fatalf("%s does not start with a capital letter", s)
		}
	}
}

func TestClassAaves(t *testing.T) {
	var s string = "Aaves"
	if ok := aaves == s; !ok {
		t.Fatalf("aaves != %s", s)
	}
}

func TestClassAmmalia(t *testing.T) {
	var s string = "Ammalia"
	if ok := ammalia == s; !ok {
		t.Fatalf("ammalia != %s", s)
	}
}

func TestClassAmphibia(t *testing.T) {
	var s string = "Amphibia"
	if ok := amphibia == s; !ok {
		t.Fatalf("amphibia != %s", s)
	}
}

func TestClassAves(t *testing.T) {
	var s string = "Aves"
	if ok := aves == s; !ok {
		t.Fatalf("aves != %s", s)
	}
}

func TestClassCephalopoda(t *testing.T) {
	var s string = "Cephalopoda"
	if ok := cephalopoda == s; !ok {
		t.Fatalf("cephalopoda != %s", s)
	}
}

func TestClassMammalia(t *testing.T) {
	var s string = "Mammalia"
	if ok := mammalia == s; !ok {
		t.Fatalf("mammalia != %s", s)
	}
}

func TestClassReptilia(t *testing.T) {
	var s string = "Reptilia"
	if ok := reptilia == s; !ok {
		t.Fatalf("reptilia != %s", s)
	}
}

func TestClassSauropsida(t *testing.T) {
	var s string = "Sauropsida"
	if ok := sauropsida == s; !ok {
		t.Fatalf("sauropsida != %s", s)
	}
}
