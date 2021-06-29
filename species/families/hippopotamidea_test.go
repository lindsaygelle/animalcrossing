package families

import "testing"

func TestHippopotamidea(t *testing.T) {
	var s string = "Hippopotamidea"
	if ok := hippopotamidea == s; !ok {
		t.Fatalf("hippopotamidea != %s", s)
	}
}
