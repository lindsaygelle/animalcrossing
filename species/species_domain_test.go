package species

import (
	"strings"
	"testing"
)

func TestDomainAll(t *testing.T) {
	var (
		c = [...]string{
			eukarya}
	)

	for i := 0; i < len(c); i++ {
		var s string = c[i]
		if ok := string(s[0]) == strings.ToUpper(string(s[0])); !ok {
			t.Fatalf("%s does not start with a capital letter", s)
		}
	}
}

func TestDomainEukarya(t *testing.T) {
	var s string = "Eukarya"
	if ok := eukarya == s; !ok {
		t.Fatalf("eukarya != %s", s)
	}
}
