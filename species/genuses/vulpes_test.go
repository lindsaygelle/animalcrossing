package genuses

import "testing"

func TestVulpes(t *testing.T) {
	var s string = "Vulpes"
	if ok := vulpes == s; !ok {
		t.Fatalf("vulpes != %s", s)
	}
}
