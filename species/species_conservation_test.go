package species

import (
	"strings"
	"testing"
)

func TestConservationAll(t *testing.T) {

	var (
		c = [...]string{
			basedOnFolklore,
			criticallyEndangered,
			domesticated,
			endangered,
			extinct,
			leastConcern,
			nearThreatened,
			unknown,
			vulnerable}
	)

	for i := 0; i < len(c); i++ {
		var s string = c[i]
		if ok := string(s[0]) == strings.ToUpper(string(s[0])); !ok {
			t.Fatalf("%s does not start with a capital letter", s)
		}
	}
}

func TestConservationCriticallyEndangered(t *testing.T) {
	var s string = "Critically Endangered"
	if ok := criticallyEndangered == s; !ok {
		t.Fatalf("criticallyEndangered != %s", s)
	}
}

func TestConservationDomesticated(t *testing.T) {
	var s string = "Domesticated"
	if ok := domesticated == s; !ok {
		t.Fatalf("domesticated != %s", s)
	}
}

func TestConservationEndangered(t *testing.T) {
	var s string = "Endangered"
	if ok := endangered == s; !ok {
		t.Fatalf("endangered != %s", s)
	}
}

func TestConservationExtinct(t *testing.T) {
	var s string = "Extinct"
	if ok := extinct == s; !ok {
		t.Fatalf("extinct != %s", s)
	}
}

func TestConservationLeastConcern(t *testing.T) {
	var s string = "Least Concern"
	if ok := leastConcern == s; !ok {
		t.Fatalf("leastConcern != %s", s)
	}
}

func TestConservationNearThreatened(t *testing.T) {
	var s string = "Near Threatened"
	if ok := nearThreatened == s; !ok {
		t.Fatalf("nearThreatened != %s", s)
	}
}

func TestConservationVulnerable(t *testing.T) {
	var s string = "Vulnerable"
	if ok := vulnerable == s; !ok {
		t.Fatalf("vulnerable != %s", s)
	}
}
