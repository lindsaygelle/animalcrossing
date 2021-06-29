package families

import "testing"

func TestFamilyLaridae(t *testing.T) {
	var s string = "Laridae"
	if ok := laridae == s; !ok {
		t.Fatalf("laridae != %s", s)
	}
}
