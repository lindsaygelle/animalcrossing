package species

import (
	"strings"
	"testing"
)

func TestPhylumAll(t *testing.T) {
	var (
		c = [...]string{
			chordata,
			mollusca}
	)
	for i := 0; i < len(c); i++ {
		var s string = c[i]
		if ok := string(s[0]) == strings.ToUpper(string(s[0])); !ok {
			t.Fatalf("%s does not start with a capital letter", s)
		}
	}
}

func TestPhylumChordata(t *testing.T) {
	var s string = "Chordata"
	if ok := chordata == s; !ok {
		t.Fatalf("chordata != %s", s)
	}
}

func TestPhylumMollusca(t *testing.T) {
	var s string = "Mollusca"
	if ok := mollusca == s; !ok {
		t.Fatalf("mollusca != %s", s)
	}
}
