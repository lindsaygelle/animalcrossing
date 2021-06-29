package families

import "testing"

func TestLaridae(t *testing.T) {
	var s string = "Laridae"
	if ok := laridae == s; !ok {
		t.Fatalf("laridae != %s", s)
	}
}
