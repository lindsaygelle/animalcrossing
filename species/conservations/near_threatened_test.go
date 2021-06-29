package conservations

import "testing"

func TestNearThreatened(t *testing.T) {
	var s string = "Near Threatened"
	if ok := nearThreatened == s; !ok {
		t.Fatalf("nearThreatened != %s", s)
	}
}
