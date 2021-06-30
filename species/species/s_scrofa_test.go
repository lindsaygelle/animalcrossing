package species

import "testing"

func TestSScrofa(t *testing.T) {
	var s string = "S. scrofa"
	if ok := sScrofa == s; !ok {
		t.Fatalf("sScrofa != %s", s)
	}
}
