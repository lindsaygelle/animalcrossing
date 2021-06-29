package families

import "testing"

func TestFamilySciuridae(t *testing.T) {
	var s string = "Sciuridae"
	if ok := sciuridae == s; !ok {
		t.Fatalf("sciuridae != %s", s)
	}
}
