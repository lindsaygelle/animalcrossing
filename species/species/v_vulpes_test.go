package species

import "testing"

func TestVVulpes(t *testing.T) {
	var s string = "V. vulpes"
	if ok := vVulpes == s; !ok {
		t.Fatalf("vVulpes != %s", s)
	}
}
