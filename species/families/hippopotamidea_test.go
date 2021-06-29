package families

import "testing"

func TestFamilyHippopotamidea(t *testing.T) {
	var s string = "Hippopotamidea"
	if ok := hippopotamidea == s; !ok {
		t.Fatalf("hippopotamidea != %s", s)
	}
}
