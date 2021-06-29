package conservations

import "testing"

func TestVulnerable(t *testing.T) {
	var s string = "Vulnerable"
	if ok := vulnerable == s; !ok {
		t.Fatalf("vulnerable != %s", s)
	}
}
