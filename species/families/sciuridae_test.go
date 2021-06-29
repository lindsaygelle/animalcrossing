package families

import "testing"

func TestSciuridae(t *testing.T) {
	var s string = "Sciuridae"
	if ok := sciuridae == s; !ok {
		t.Fatalf("sciuridae != %s", s)
	}
}
